package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/zkep/my-geektime/libs/m3u8"
)

func main() {
	// Parse command line flags
	url := flag.String("url", "", "M3U8 playlist URL")
	output := flag.String("output", "./downloads", "Output directory")
	concurrency := flag.Int("concurrency", 10, "Number of concurrent downloads")
	key := flag.String("key", "", "Encryption key for Aliyun VOD (optional)")
	timeout := flag.Duration("timeout", 10*time.Minute, "Download timeout")
	flag.Parse()

	if *url == "" {
		fmt.Println("Error: URL is required")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Starting M3U8 downloader...\n")
	fmt.Printf("URL: %s\n", *url)
	fmt.Printf("Output: %s\n", *output)
	fmt.Printf("Concurrency: %d\n", *concurrency)
	fmt.Printf("Timeout: %v\n\n", *timeout)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	// Configure downloader
	config := m3u8.DownloadConfig{
		Output:      *output,
		URL:         *url,
		Key:         *key,
		Concurrency: *concurrency,
		Headers: map[string]string{
			"Referer": "https://time.geekbang.org",
			"Origin":  "https://time.geekbang.org",
		},
	}

	// Create downloader
	downloader, err := m3u8.NewDownloader(ctx, config)
	if err != nil {
		fmt.Printf("Failed to create downloader: %v\n", err)
		os.Exit(1)
	}

	// Start download
	fmt.Println("Downloading...")
	if err := downloader.Start(); err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			fmt.Printf("\nDownload timed out after %v\n", *timeout)
		case errors.Is(err, context.Canceled):
			fmt.Println("\nDownload was canceled")
		default:
			fmt.Printf("\nDownload failed: %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Println("\n✓ Download completed successfully!")
}
