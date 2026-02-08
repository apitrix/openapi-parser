package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for components_responses.go - parseComponentsResponses function
// =============================================================================

// --- Single Response ---

func TestParseComponentsResponses_Single(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  responses:
    NotFound:
      description: "Resource not found"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Responses, 1)
	assert.Contains(t, doc.Components.Responses, "NotFound")
}

// --- Multiple Responses ---

func TestParseComponentsResponses_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  responses:
    NotFound:
      description: "Not found"
    BadRequest:
      description: "Bad request"
    Unauthorized:
      description: "Unauthorized"
    InternalError:
      description: "Internal server error"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Responses, 4)
}

// --- Empty ---

func TestParseComponentsResponses_Empty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  responses: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, doc.Components.Responses)
}

// --- With Content ---

func TestParseComponentsResponses_WithContent(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  responses:
    Success:
      description: "Successful response"
      content:
        application/json:
          schema:
            type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	resp := doc.Components.Responses["Success"].Value
	assert.NotNil(t, resp.Content["application/json"])
}

// --- With Headers ---

func TestParseComponentsResponses_WithHeaders(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  responses:
    RateLimited:
      description: "Rate limited"
      headers:
        X-Rate-Limit-Remaining:
          schema:
            type: integer
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	resp := doc.Components.Responses["RateLimited"].Value
	assert.Contains(t, resp.Headers, "X-Rate-Limit-Remaining")
}
