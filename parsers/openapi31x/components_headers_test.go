package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for components_headers.go - parseComponentsHeaders function
// =============================================================================

// --- Single Header ---

func TestParseComponentsHeaders_Single(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  headers:
    X-Rate-Limit:
      schema:
        type: integer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components.Headers, 1)
	assert.Contains(t, result.Document.Components.Headers, "X-Rate-Limit")
}

// --- Multiple Headers ---

func TestParseComponentsHeaders_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  headers:
    X-Rate-Limit:
      schema:
        type: integer
    X-Request-Id:
      schema:
        type: string
    X-Custom:
      schema:
        type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components.Headers, 3)
}

// --- Empty ---

func TestParseComponentsHeaders_Empty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  headers: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Components.Headers)
}

// --- With Description ---

func TestParseComponentsHeaders_WithDescription(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  headers:
    X-Rate-Limit:
      description: "The number of allowed requests per hour"
      schema:
        type: integer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	header := result.Document.Components.Headers["X-Rate-Limit"].Value
	assert.Equal(t, "The number of allowed requests per hour", header.Description)
}
