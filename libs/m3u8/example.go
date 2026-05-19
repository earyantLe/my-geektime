package m3u8

import (
	"context"
	"fmt"
	"time"
)

// Example usage of the M3U8 downloader
func Example() {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Configure the downloader
	config := DownloadConfig{
		Output:      "/path/to/output",
		URL:         "https://example.com/playlist.m3u8",
		Key:         "", // Optional encryption key for Aliyun VOD
		Concurrency: 10, // Number of concurrent downloads
		Headers: map[string]string{
			"Referer": "https://time.geekbang.org",
			"Origin":  "https://time.geekbang.org",
		},
	}

	// Create downloader
	downloader, err := NewDownloader(ctx, config)
	if err != nil {
		fmt.Printf("Failed to create downloader: %v\n", err)
		return
	}

	// Start download
	if err := downloader.Start(); err != nil {
		fmt.Printf("Download failed: %v\n", err)
		return
	}

	fmt.Println("Download completed successfully!")
}
