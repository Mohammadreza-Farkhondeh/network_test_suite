package tests

import (
	"fmt"
	"net/http"
	"time"
)

type WebsiteAccessibilityResult struct {
	Website    string
	Accessible bool
	StatusCode int
	Elapsed    time.Duration
	TestTime   time.Time
}

// CheckWebsiteAccessibility checks if a website is accessible via HTTP.
func CheckWebsiteAccessibility(url string) (WebsiteAccessibilityResult, error) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return WebsiteAccessibilityResult{}, fmt.Errorf("failed to access %s: %w", url, err)
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	return WebsiteAccessibilityResult{
		Website:    url,
		Accessible: resp.StatusCode == http.StatusOK,
		StatusCode: resp.StatusCode,
		Elapsed:    elapsed,
		TestTime:   time.Now(),
	}, nil
}
