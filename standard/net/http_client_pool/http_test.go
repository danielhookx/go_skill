package http_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	_, err = io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
}

func TestBigRespClient(t *testing.T) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8080/bigresp", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	data, err := io.ReadAll(resp.Body)
	t.Log("get big resp:", formatBytes(uint64(len(data))))
	resp.Body.Close()
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
}

func BenchmarkClient(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			client := &http.Client{}

			req, err := http.NewRequest("GET", "http://localhost:8080/hello", nil)
			if err != nil {
				b.Fatalf("Failed to create request: %v", err)
			}

			resp, err := client.Do(req)
			if err != nil {
				b.Fatalf("Failed to send request: %v", err)
			}

			_, err = io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				b.Fatalf("Failed to read response body: %v", err)
			}
		}

	})
}

func BenchmarkClientPerHostPool(b *testing.B) {
	trans := http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
	}
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			client := &http.Client{
				Transport: &trans,
			}

			req, err := http.NewRequest("GET", "http://localhost:8080/hello", nil)
			if err != nil {
				b.Fatalf("Failed to create request: %v", err)
			}

			resp, err := client.Do(req)
			if err != nil {
				b.Fatalf("Failed to send request: %v", err)
			}

			_, err = io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				b.Fatalf("Failed to read response body: %v", err)
			}
		}

	})
}

func BenchmarkClientPerHostPoolBigResp(b *testing.B) {
	trans := http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
	}
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			client := &http.Client{
				Transport: &trans,
			}

			req, err := http.NewRequest("GET", "http://localhost:8080/bigresp", nil)
			if err != nil {
				b.Fatalf("Failed to create request: %v", err)
			}

			resp, err := client.Do(req)
			if err != nil {
				b.Fatalf("Failed to send request: %v", err)
			}

			_, err = io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				b.Fatalf("Failed to read response body: %v", err)
			}
		}

	})
}

func formatBytes(bytes uint64) string {
	const (
		KB = 1 << 10 // 1024
		MB = 1 << 20 // 1024 * 1024
		GB = 1 << 30 // 1024 * 1024 * 1024
		TB = 1 << 40 // 1024 * 1024 * 1024 * 1024
		PB = 1 << 50 // 1024 * 1024 * 1024 * 1024 * 1024
	)

	switch {
	case bytes >= PB:
		return fmt.Sprintf("%.2f PB", float64(bytes)/PB)
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/TB)
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d bytes", bytes)
	}
}
