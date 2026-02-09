package openapi20

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for parse.go - Parse functions
// =============================================================================

// --- Basic Parsing ---

func TestParse_ValidMinimalSwagger(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test API"
  version: "1.0.0"
paths: {}
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "2.0", result.Document.Swagger)
	assert.Equal(t, "Test API", result.Document.Info.Title)
	assert.Equal(t, "1.0.0", result.Document.Info.Version)
}

func TestParse_JSON(t *testing.T) {
	// Arrange
	json := `{"swagger": "2.0", "info": {"title": "Test API", "version": "1.0"}, "paths": {}}`

	// Act
	result, err := Parse([]byte(json))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "2.0", result.Document.Swagger)
}

func TestParse_InvalidYAML(t *testing.T) {
	// Arrange
	invalidYAML := `swagger: [invalid`

	// Act
	_, err := Parse([]byte(invalidYAML))

	// Assert
	require.Error(t, err)
}

func TestParse_InvalidVersion(t *testing.T) {
	// Arrange
	yaml := `swagger: "3.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
`

	// Act
	_, err := Parse([]byte(yaml))

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported Swagger version")
}

func TestParse_MissingSwaggerVersion(t *testing.T) {
	// Arrange
	yaml := `info:
  title: "Test"
  version: "1.0"
paths: {}
`

	// Act
	_, err := Parse([]byte(yaml))

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "swagger field is required")
}

// --- Parse ---

func TestParse_DetectsUnknown(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
  unknownField: "value"
paths: {}
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	unknownErrors := filterTestUnknownFieldErrors(result)
	assert.NotEmpty(t, unknownErrors)
	assert.True(t, strings.Contains(unknownErrors[0].Message, "unknownField"))
}

func TestParse_IgnoresExtensions(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
  x-custom: "value"
paths: {}
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	unknownErrors := filterTestUnknownFieldErrors(result)
	assert.Empty(t, unknownErrors)
}

// filterTestUnknownFieldErrors returns only errors with Kind "unknown_field" from a ParseResult.
func filterTestUnknownFieldErrors(result *ParseResult) []*ParseError {
	var filtered []*ParseError
	for _, e := range result.Errors {
		if e.Kind == "unknown_field" {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

// --- ParseReader ---

func TestParseReader_Valid(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	reader := strings.NewReader(yaml)

	// Act
	result, err := ParseReader(reader)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "2.0", result.Document.Swagger)
}
