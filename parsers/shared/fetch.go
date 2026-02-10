package shared

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// defaultFetchClient is the HTTP client used by FetchURL.
var defaultFetchClient = &http.Client{Timeout: 30 * time.Second}

// FetchURL fetches a document from an HTTP/HTTPS URL and returns the raw bytes
// along with the base URL (directory) for resolving relative $ref references.
func FetchURL(rawURL string) (data []byte, basePath string, err error) {
	resp, err := defaultFetchClient.Get(rawURL)
	if err != nil {
		return nil, "", fmt.Errorf("failed to fetch URL %q: %w", rawURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("failed to fetch URL %q: HTTP %d", rawURL, resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read response body from %q: %w", rawURL, err)
	}

	basePath = urlBase(rawURL)
	return data, basePath, nil
}

// urlBase extracts the base URL (everything up to and including the last '/').
// e.g. "https://example.com/specs/openapi.yaml" → "https://example.com/specs/"
func urlBase(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		// Fallback: trim after last '/'
		if i := strings.LastIndex(rawURL, "/"); i >= 0 {
			return rawURL[:i+1]
		}
		return rawURL
	}

	if i := strings.LastIndex(parsed.Path, "/"); i >= 0 {
		parsed.Path = parsed.Path[:i+1]
	}
	return parsed.String()
}
