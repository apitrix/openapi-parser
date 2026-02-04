package openapi30

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
	yaml := `openapi: "3.0.3"
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Headers, 1)
	assert.Contains(t, doc.Components.Headers, "X-Rate-Limit")
}

// --- Multiple Headers ---

func TestParseComponentsHeaders_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Headers, 3)
}

// --- Empty ---

func TestParseComponentsHeaders_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  headers: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, doc.Components.Headers)
}

// --- With Description ---

func TestParseComponentsHeaders_WithDescription(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	header := doc.Components.Headers["X-Rate-Limit"].Value
	assert.Equal(t, "The number of allowed requests per hour", header.Description)
}
