package hooks

import (
	"net/http"
	"testing"
)

func TestNormalizeServerURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no scheme prepends https",
			input:    "example.glean.com",
			expected: "https://example.glean.com",
		},
		{
			name:     "https preserved",
			input:    "https://example.glean.com",
			expected: "https://example.glean.com",
		},
		{
			name:     "http localhost preserved",
			input:    "http://localhost:8080",
			expected: "http://localhost:8080",
		},
		{
			name:     "trailing slashes stripped",
			input:    "https://example.glean.com///",
			expected: "https://example.glean.com",
		},
		{
			name:     "url with path",
			input:    "https://example.glean.com/api/v1",
			expected: "https://example.glean.com/api/v1",
		},
		{
			name:     "no scheme with trailing slash",
			input:    "example.glean.com/",
			expected: "https://example.glean.com",
		},
		{
			name:     "no scheme with path",
			input:    "example.glean.com/api/v1",
			expected: "https://example.glean.com/api/v1",
		},
	}

	hook := &ServerURLNormalizerHook{}
	client := http.DefaultClient

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, resultClient := hook.SDKInit(tt.input, client)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
			if resultClient != client {
				t.Error("expected client to be unchanged")
			}
		})
	}
}
