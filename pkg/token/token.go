package token

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

const (
	tokenRequestURLEnv   = "ACTIONS_ID_TOKEN_REQUEST_URL"
	tokenRequestTokenEnv = "ACTIONS_ID_TOKEN_REQUEST_TOKEN"
)

// FetchToken retrieves a GitHub Actions OIDC token from the runtime.
func FetchToken() (string, error) {
	reqURL := os.Getenv(tokenRequestURLEnv)
	reqToken := os.Getenv(tokenRequestTokenEnv)

	if reqURL == "" || reqToken == "" {
		return "", fmt.Errorf("GitHub Actions OIDC env vars not set")
	}

	// Construct the full URL with audience parameter
	fullURL := reqURL + "&audience=api://AzureADTokenExchange"

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", fmt.Errorf("requesting token: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	return string(body), nil
}
