package version

import "testing"

func TestGet(t *testing.T) {
	v := Get()
	if v.Platform == "" {
		t.Error("expected non-empty platform")
	}
	if v.GoVersion == "" {
		t.Error("expected non-empty go version")
	}
}
