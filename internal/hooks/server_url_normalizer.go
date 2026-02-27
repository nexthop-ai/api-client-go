package hooks

import (
	"regexp"
	"strings"
)

var schemeRegex = regexp.MustCompile(`(?i)^https?://`)

// ServerURLNormalizerHook normalizes server URLs by prepending https:// if no scheme is provided.
type ServerURLNormalizerHook struct{}

func (h *ServerURLNormalizerHook) SDKInit(baseURL string, client HTTPClient) (string, HTTPClient) {
	normalized := baseURL
	if !schemeRegex.MatchString(normalized) {
		normalized = "https://" + normalized
	}
	normalized = strings.TrimRight(normalized, "/")
	return normalized, client
}
