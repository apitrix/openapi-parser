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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Parameters, 1)
	assert.Contains(t, doc.Components.Parameters, "LimitParam")
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Parameters, 3)
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, doc.Components.Parameters)
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Parameters, 4)
	assert.Equal(t, "path", doc.Components.Parameters["PathParam"].Value.In)
	assert.Equal(t, "query", doc.Components.Parameters["QueryParam"].Value.In)
	assert.Equal(t, "header", doc.Components.Parameters["HeaderParam"].Value.In)
	assert.Equal(t, "cookie", doc.Components.Parameters["CookieParam"].Value.In)
}
