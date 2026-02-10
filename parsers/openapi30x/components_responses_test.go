package openapi30x

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
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  responses:
    NotFound:
      description: "Resource not found"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().Responses(), 1)
	assert.Contains(t, result.Document.Components().Responses(), "NotFound")
}

// --- Multiple Responses ---

func TestParseComponentsResponses_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().Responses(), 4)
}

// --- Empty ---

func TestParseComponentsResponses_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  responses: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Components().Responses())
}

// --- With Content ---

func TestParseComponentsResponses_WithContent(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	resp := result.Document.Components().Responses()["Success"].Value
	assert.NotNil(t, resp.Content()["application/json"])
}

// --- With Headers ---

func TestParseComponentsResponses_WithHeaders(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	resp := result.Document.Components().Responses()["RateLimited"].Value
	assert.Contains(t, resp.Headers(), "X-Rate-Limit-Remaining")
}
