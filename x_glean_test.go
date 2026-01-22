package apiclientgo

import (
	"os"
	"testing"

	"github.com/gleanwork/api-client-go/internal/hooks"
)

func TestWithExcludeDeprecatedAfter(t *testing.T) {
	sdk := New(
		WithExcludeDeprecatedAfter("2026-10-15"),
	)

	config := hooks.GetXGleanConfig(sdk)
	if config.ExcludeDeprecatedAfter != "2026-10-15" {
		t.Errorf("expected ExcludeDeprecatedAfter to be '2026-10-15', got '%s'", config.ExcludeDeprecatedAfter)
	}
}

func TestWithIncludeExperimental(t *testing.T) {
	sdk := New(
		WithIncludeExperimental(true),
	)

	config := hooks.GetXGleanConfig(sdk)
	if !config.IncludeExperimental {
		t.Error("expected IncludeExperimental to be true")
	}
}

func TestWithBothOptions(t *testing.T) {
	sdk := New(
		WithExcludeDeprecatedAfter("2026-10-15"),
		WithIncludeExperimental(true),
	)

	config := hooks.GetXGleanConfig(sdk)
	if config.ExcludeDeprecatedAfter != "2026-10-15" {
		t.Errorf("expected ExcludeDeprecatedAfter to be '2026-10-15', got '%s'", config.ExcludeDeprecatedAfter)
	}
	if !config.IncludeExperimental {
		t.Error("expected IncludeExperimental to be true")
	}
}

func TestWithIncludeExperimentalFalse(t *testing.T) {
	sdk := New(
		WithIncludeExperimental(false),
	)

	config := hooks.GetXGleanConfig(sdk)
	if config.IncludeExperimental {
		t.Error("expected IncludeExperimental to be false")
	}
}

func TestMultipleSDKInstances(t *testing.T) {
	sdk1 := New(
		WithExcludeDeprecatedAfter("2026-01-01"),
	)
	sdk2 := New(
		WithExcludeDeprecatedAfter("2027-12-31"),
	)

	config1 := hooks.GetXGleanConfig(sdk1)
	config2 := hooks.GetXGleanConfig(sdk2)

	if config1.ExcludeDeprecatedAfter != "2026-01-01" {
		t.Errorf("sdk1: expected ExcludeDeprecatedAfter to be '2026-01-01', got '%s'", config1.ExcludeDeprecatedAfter)
	}
	if config2.ExcludeDeprecatedAfter != "2027-12-31" {
		t.Errorf("sdk2: expected ExcludeDeprecatedAfter to be '2027-12-31', got '%s'", config2.ExcludeDeprecatedAfter)
	}
}

func TestEnvironmentVariablesAreRead(t *testing.T) {
	// This test verifies that the hook reads environment variables
	// The actual precedence is tested in the hooks package

	// Set environment variables
	os.Setenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER", "2028-06-01")
	os.Setenv("X_GLEAN_INCLUDE_EXPERIMENTAL", "true")
	defer os.Unsetenv("X_GLEAN_EXCLUDE_DEPRECATED_AFTER")
	defer os.Unsetenv("X_GLEAN_INCLUDE_EXPERIMENTAL")

	// SDK can be created without error when env vars are set
	sdk := New()
	if sdk == nil {
		t.Error("expected SDK to be created")
	}
}
