package m3u8

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// PlaylistType represents the type of playlist
type PlaylistType string

// CryptMethod represents the encryption method
type CryptMethod string

const (
	PlaylistTypeVOD   PlaylistType = "VOD"
	PlaylistTypeEvent PlaylistType = "EVENT"

	CryptMethodAES  CryptMethod = "AES-128"
	CryptMethodNONE CryptMethod = "NONE"
)

// regex pattern for extracting `key=value` parameters from a line
var linePattern = regexp.MustCompile(`([a-zA-Z-]+)=("[^"]+"|[^",]+)`)

// ParseResult holds the result of parsing an M3U8 file
type ParseResult struct {
	URL  *url.URL
	M3u8 *M3u8
}

// M3u8 represents an M3U8 playlist
type M3u8 struct {
	Version        int8              // EXT-X-VERSION:version
	MediaSequence  uint64            // Default 0, #EXT-X-MEDIA-SEQUENCE:sequence
	Segments       []*Segment        // TS segments
	MasterPlaylist []*MasterPlaylist // Master playlist entries
	Keys           map[int]*KeyInfo  // Encryption keys
	EndList        bool              // #EXT-X-ENDLIST
	PlaylistType   PlaylistType      // VOD or EVENT
	TargetDuration float64           // #EXT-X-TARGETDURATION:duration
}

// Segment represents a media segment
type Segment struct {
	URI      string  // Segment URI
	KeyIndex int     // Index into Keys map
	Title    string  // #EXTINF: duration,<title>
	Duration float32 // #EXTINF: duration,<title>
	Length   uint64  // #EXT-X-BYTERANGE: length[@offset]
	Offset   uint64  // #EXT-X-BYTERANGE: length[@offset]
}

// MasterPlaylist represents a master playlist entry
type MasterPlaylist struct {
	URI        string
	BandWidth  uint32
	Resolution string
	Codecs     string
	ProgramID  uint32
}

// KeyInfo represents encryption key information
type KeyInfo struct {
	Method              CryptMethod // 'AES-128' or 'NONE'
	URI                 string      // Key URI
	Key                 string      // Decrypted key
	IV                  string      // Initialization vector
	AliyunVoDEncryption bool        // Whether using Aliyun VOD encryption
}

// ParseFromURL parses an M3U8 playlist from a URL
func ParseFromURL(ctx context.Context, link string, headers map[string]string) (*ParseResult, error) {
	u, err := url.Parse(link)
	if err != nil {
		return nil, fmt.Errorf("parse URL failed: %w", err)
	}

	// Create independent HTTP client for this parse operation
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 20,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	if referer, ok := headers["Referer"]; ok {
		req.Header.Set("Referer", referer)
	} else {
		req.Header.Set("Referer", u.String())
	}
	if origin, ok := headers["Origin"]; ok {
		req.Header.Set("Origin", origin)
	} else {
		req.Header.Set("Origin", "https://time.geekbang.org")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request m3u8 URL failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status %d", resp.StatusCode)
	}

	// Parse the M3U8 content
	m3u8, err := parseM3u8(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse m3u8 content failed: %w", err)
	}

	return processM3u8Result(ctx, m3u8, u, headers)
}

// ParseFromBytes parses an M3U8 playlist from bytes
func ParseFromBytes(ctx context.Context, data []byte, baseURL string, headers map[string]string) (*ParseResult, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parse base URL failed: %w", err)
	}

	// Parse the M3U8 content from bytes
	m3u8, err := parseM3u8(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("parse m3u8 content failed: %w", err)
	}

	return processM3u8Result(ctx, m3u8, u, headers)
}

// processM3u8Result processes the parsed M3U8 result (common logic)
func processM3u8Result(ctx context.Context, m3u8 *M3u8, u *url.URL, headers map[string]string) (*ParseResult, error) {

	// Handle master playlist (multiple quality levels)
	if len(m3u8.MasterPlaylist) > 0 {
		// Use the first variant
		sf := m3u8.MasterPlaylist[0]
		resolvedURL := resolveURL(u, sf.URI)
		return ParseFromURL(ctx, resolvedURL, headers)
	}

	if len(m3u8.Segments) == 0 {
		return nil, fmt.Errorf("no segments found in m3u8")
	}

	result := &ParseResult{
		URL:  u,
		M3u8: m3u8,
	}

	// Fetch and process keys
	for idx, key := range m3u8.Keys {
		switch {
		case key.Method == "" || key.Method == CryptMethodNONE:
			continue
		case key.AliyunVoDEncryption && key.Method == CryptMethodAES:
			// For Aliyun VOD encryption, key should be provided externally
			// This will be handled by the caller
			continue
		case !key.AliyunVoDEncryption && key.Method == CryptMethodAES:
			// Skip fetching key if URI indicates it will be provided externally
			if key.URI == "provided-externally" {
				continue
			}
			// Fetch key from URI
			keyURL := resolveURL(u, key.URI)
			keyData, err := fetchKey(ctx, keyURL, headers)
			if err != nil {
				return nil, fmt.Errorf("fetch key failed: %w", err)
			}
			m3u8.Keys[idx].Key = string(keyData)
		default:
			return nil, fmt.Errorf("unsupported encryption method: %s", key.Method)
		}
	}

	return result, nil
}

// parseM3u8 parses M3U8 content from a reader
func parseM3u8(reader io.Reader) (*M3u8, error) {
	s := bufio.NewScanner(reader)
	var lines []string
	for s.Scan() {
		lines = append(lines, strings.TrimSpace(s.Text()))
	}

	var (
		i        = 0
		count    = len(lines)
		m3u8     = &M3u8{Keys: make(map[int]*KeyInfo)}
		keyIndex = 0
		key      *KeyInfo
		seg      *Segment
		extInf   bool
		extByte  bool
	)

	for ; i < count; i++ {
		line := lines[i]

		// First line must be #EXTM3U
		if i == 0 {
			if line != "#EXTM3U" {
				return nil, fmt.Errorf("invalid m3u8, missing #EXTM3U in line 1")
			}
			continue
		}

		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "#EXT-X-PLAYLIST-TYPE:"):
			if _, err := fmt.Sscanf(line, "#EXT-X-PLAYLIST-TYPE:%s", &m3u8.PlaylistType); err != nil {
				return nil, err
			}
			isValid := m3u8.PlaylistType == "" ||
				m3u8.PlaylistType == PlaylistTypeVOD ||
				m3u8.PlaylistType == PlaylistTypeEvent
			if !isValid {
				return nil, fmt.Errorf("invalid playlist type: %s", m3u8.PlaylistType)
			}
		case strings.HasPrefix(line, "#EXT-X-TARGETDURATION:"):
			if _, err := fmt.Sscanf(line, "#EXT-X-TARGETDURATION:%f", &m3u8.TargetDuration); err != nil {
				return nil, err
			}
		case strings.HasPrefix(line, "#EXT-X-MEDIA-SEQUENCE:"):
			if _, err := fmt.Sscanf(line, "#EXT-X-MEDIA-SEQUENCE:%d", &m3u8.MediaSequence); err != nil {
				return nil, err
			}
		case strings.HasPrefix(line, "#EXT-X-VERSION:"):
			if _, err := fmt.Sscanf(line, "#EXT-X-VERSION:%d", &m3u8.Version); err != nil {
				return nil, err
			}
		case strings.HasPrefix(line, "#EXT-X-STREAM-INF:"):
			// Master playlist entry
			mp, err := parseMasterPlaylist(line)
			if err != nil {
				return nil, err
			}
			i++
			if i < count {
				mp.URI = lines[i]
				if mp.URI == "" || strings.HasPrefix(mp.URI, "#") {
					return nil, fmt.Errorf("invalid stream URI at line %d", i+1)
				}
				m3u8.MasterPlaylist = append(m3u8.MasterPlaylist, mp)
			}
			continue
		case strings.HasPrefix(line, "#EXTINF:"):
			if extInf {
				return nil, fmt.Errorf("duplicate EXTINF at line %d", i+1)
			}
			if seg == nil {
				seg = new(Segment)
			}
			var s string
			if _, err := fmt.Sscanf(line, "#EXTINF:%s", &s); err != nil {
				return nil, err
			}
			if strings.Contains(s, ",") {
				split := strings.Split(s, ",")
				seg.Title = split[1]
				s = split[0]
			}
			df, err := strconv.ParseFloat(s, 32)
			if err != nil {
				return nil, err
			}
			seg.Duration = float32(df)
			seg.KeyIndex = keyIndex
			extInf = true
		case strings.HasPrefix(line, "#EXT-X-BYTERANGE:"):
			if extByte {
				return nil, fmt.Errorf("duplicate EXT-X-BYTERANGE at line %d", i+1)
			}
			if seg == nil {
				seg = new(Segment)
			}
			var b string
			if _, err := fmt.Sscanf(line, "#EXT-X-BYTERANGE:%s", &b); err != nil {
				return nil, err
			}
			if b == "" {
				return nil, fmt.Errorf("invalid EXT-X-BYTERANGE at line %d", i+1)
			}
			if strings.Contains(b, "@") {
				split := strings.Split(b, "@")
				offset, err := strconv.ParseUint(split[1], 10, 64)
				if err != nil {
					return nil, err
				}
				seg.Offset = offset
				b = split[0]
			}
			length, err := strconv.ParseUint(b, 10, 64)
			if err != nil {
				return nil, err
			}
			seg.Length = length
			extByte = true
		case !strings.HasPrefix(line, "#"):
			// Segment URI
			if extInf {
				if seg == nil {
					return nil, fmt.Errorf("invalid segment URI at line %d", i+1)
				}
				seg.URI = line
				extByte = false
				extInf = false
				m3u8.Segments = append(m3u8.Segments, seg)
				seg = nil
				continue
			}
		case strings.HasPrefix(line, "#EXT-X-KEY"):
			// Encryption key
			params := parseLineParameters(line)
			if len(params) == 0 {
				return nil, fmt.Errorf("invalid EXT-X-KEY at line %d", i+1)
			}

			method := CryptMethod(params["METHOD"])
			// Check for Aliyun typo (MEATHOD instead of METHOD)
			if meathod := CryptMethod(params["MEATHOD"]); meathod == CryptMethodAES {
				method = meathod
			}

			key = new(KeyInfo)
			if method == CryptMethodAES {
				// Check if it's Aliyun VOD encryption
				if params["MEATHOD"] == string(CryptMethodAES) {
					key.AliyunVoDEncryption = true
				}
			}

			if method != "" && method != CryptMethodAES && method != CryptMethodNONE {
				return nil, fmt.Errorf("unsupported encryption method: %s", method)
			}

			keyIndex++
			key.Method = method
			key.URI = params["URI"]
			key.IV = params["IV"]
			m3u8.Keys[keyIndex] = key
		case line == "#EXT-X-ENDLIST":
			m3u8.EndList = true
		}
	}

	return m3u8, nil
}

// parseMasterPlaylist parses a master playlist line
func parseMasterPlaylist(line string) (*MasterPlaylist, error) {
	params := parseLineParameters(line)
	if len(params) == 0 {
		return nil, fmt.Errorf("empty master playlist parameters")
	}

	mp := new(MasterPlaylist)
	for k, v := range params {
		switch k {
		case "BANDWIDTH":
			val, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, err
			}
			mp.BandWidth = uint32(val)
		case "RESOLUTION":
			mp.Resolution = v
		case "PROGRAM-ID":
			val, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, err
			}
			mp.ProgramID = uint32(val)
		case "CODECS":
			mp.Codecs = v
		}
	}
	return mp, nil
}

// parseLineParameters extracts key=value parameters from a line
func parseLineParameters(line string) map[string]string {
	r := linePattern.FindAllStringSubmatch(line, -1)
	params := make(map[string]string)
	for _, arr := range r {
		params[arr[1]] = strings.Trim(arr[2], "\"")
	}
	return params
}

// resolveURL resolves a relative URL against a base URL
func resolveURL(base *url.URL, ref string) string {
	if strings.HasPrefix(ref, "http://") || strings.HasPrefix(ref, "https://") || strings.HasPrefix(ref, "file://") {
		return ref
	}

	resolved := base.ResolveReference(&url.URL{Path: ref})
	return resolved.String()
}

// fetchKey fetches an encryption key from a URL
func fetchKey(ctx context.Context, keyURL string, headers map[string]string) ([]byte, error) {
	// Handle file:// protocol
	if strings.HasPrefix(keyURL, "file://") {
		filePath := strings.TrimPrefix(keyURL, "file://")
		// Check if file exists first
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return nil, fmt.Errorf("key file does not exist: %s", filePath)
		}
		data, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("read key file failed (path: %s): %w", filePath, err)
		}
		return data, nil
	}

	// Create independent HTTP client for this key fetch operation
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 20,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, keyURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	if referer, ok := headers["Referer"]; ok {
		req.Header.Set("Referer", referer)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("key fetch failed with status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
