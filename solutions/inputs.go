package solutions

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var cacheDir = "aoc_cache"
var session = ""

func init() {
	if envCacheDir := os.Getenv("AOC_CACHE_DIR"); envCacheDir != "" {
		cacheDir = envCacheDir
	}
	if envSession, ok := os.LookupEnv("AOC_SESSION"); ok {
		session = envSession
	}
}

func getInput(year, day int) ([]byte, error) {
	path := filepath.Join(cacheDir, strconv.Itoa(year), strconv.Itoa(day)+".txt")

	// already cached to disk, return from file
	s, err := os.Stat(path)
	if err == nil && !s.IsDir() {
		f, err := os.Open(path)
		defer func() { _ = f.Close() }()
		if err == nil {
			data, err := io.ReadAll(f)
			if err == nil {
				return data, nil
			}

			fmt.Println("Failed to read cached file:", err)
			fmt.Println("Fetching fresh input...")
		} else {
			fmt.Println("Failed to open cached file:", err)
			fmt.Println("Fetching fresh input...")
		}
	}

	data, err := fetchInput(year, day)
	if err != nil {
		return nil, fmt.Errorf("fetching input; %w", err)
	}

	if bytes.Equal(data, []byte("Puzzle inputs differ by user.  Please log in to get your puzzle input.")) {
		return nil, errors.New("failed to authenticate")
	}

	if err := writeToFile(path, data); err != nil {
		fmt.Println("Failed to write to cache file:", err)
	}

	return data, nil
}

func fetchInput(year int, day int) ([]byte, error) {
	if session == "" {
		return nil, fmt.Errorf("session not set")
	}

	u, err := url.Parse(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day))
	if err != nil {
		return nil, fmt.Errorf("parsing url; %w", err)
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "github.com/aivarasbaranauskas/aoc")
	req.Header.Set("Cookie", "session="+session)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	defer func() { _ = resp.Body.Close() }()
	if err != nil {
		return nil, fmt.Errorf("making request; %w", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body; %w", err)
	}

	return bytes.Trim(data, "\n"), nil
}

func writeToFile(path string, data []byte) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
		return fmt.Errorf("creating directory for cache; %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("creating cache file; %w", err)
	}

	_, err = f.Write(data)
	if err != nil {
		_ = f.Close()
		return fmt.Errorf("writing to cache file; %w", err)
	}

	return f.Close()
}
