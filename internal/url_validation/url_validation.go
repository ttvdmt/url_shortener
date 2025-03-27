package urlvalidation

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

func IsValid(str_url string) bool {
	parsed_url, err := url.ParseRequestURI(str_url)
	if err != nil {
		return false
	}

	if parsed_url.Host == "" {
		return false
	}

	if !isSafe(*parsed_url) || !isAvailable(str_url) {
		return false
	}

	return true
}

func isAvailable(url string) bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Head(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return true
	}

	return false
}

func isSafe(u url.URL) bool {
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	ip := net.ParseIP(u.Hostname())
	if ip != nil && (ip.IsLoopback() || ip.IsPrivate()) {
		return false
	}

	return true
}
