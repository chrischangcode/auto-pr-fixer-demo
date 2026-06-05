package token

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFetchToken_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("oidc-token-value"))
	}))
	defer server.Close()

	os.Setenv(tokenRequestURLEnv, server.URL+"?param=1")
	os.Setenv(tokenRequestTokenEnv, "test-token")
	defer os.Unsetenv(tokenRequestURLEnv)
	defer os.Unsetenv(tokenRequestTokenEnv)

	tok, err := FetchToken()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok != "oidc-token-value" {
		t.Errorf("expected oidc-token-value, got %s", tok)
	}
}

func TestFetchToken_MissingEnv(t *testing.T) {
	os.Unsetenv(tokenRequestURLEnv)
	os.Unsetenv(tokenRequestTokenEnv)

	_, err := FetchToken()
	if err == nil {
		t.Error("expected error when env vars not set")
	}
}
