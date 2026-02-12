package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for components_parameters.go - parseComponentsParameters function
// =============================================================================

// --- Single Parameter ---

func TestParseComponentsParameters_Single(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  parameters:
    LimitParam:
      name: limit
      in: query
      schema:
        type: integer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().Parameters(), 1)
	assert.Contains(t, result.Document.Components().Parameters(), "LimitParam")
}

// --- Multiple Parameters ---

func TestParseComponentsParameters_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  parameters:
    LimitParam:
      name: limit
      in: query
      schema:
        type: integer
    OffsetParam:
      name: offset
      in: query
      schema:
        type: integer
    SortParam:
      name: sort
      in: query
      schema:
        type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().Parameters(), 3)
}

// --- Empty ---

func TestParseComponentsParameters_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  parameters: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Components().Parameters())
}

// --- All Locations ---

func TestParseComponentsParameters_AllLocations(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  parameters:
    PathParam:
      name: id
      in: path
      required: true
      schema:
        type: string
    QueryParam:
      name: filter
      in: query
      schema:
        type: string
    HeaderParam:
      name: X-Request-Id
      in: header
      schema:
        type: string
    CookieParam:
      name: session
      in: cookie
      schema:
        type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().Parameters(), 4)
	assert.Equal(t, "path", result.Document.Components().Parameters()["PathParam"].Value().In())
	assert.Equal(t, "query", result.Document.Components().Parameters()["QueryParam"].Value().In())
	assert.Equal(t, "header", result.Document.Components().Parameters()["HeaderParam"].Value().In())
	assert.Equal(t, "cookie", result.Document.Components().Parameters()["CookieParam"].Value().In())
}
