package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for operation_responses.go - parseOperationResponses function
// =============================================================================

// --- Single Response ---

func TestParseOperationResponses_Single(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	codes := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()
	assert.Len(t, codes, 1)
}

// --- Multiple Status Codes ---

func TestParseOperationResponses_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
        "400":
          description: "Bad Request"
        "404":
          description: "Not Found"
        "500":
          description: "Internal Error"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	codes := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()
	assert.Len(t, codes, 4)
}

// --- Default Response ---

func TestParseOperationResponses_Default(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
        default:
          description: "Unexpected error"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	defaultResp := result.Document.Paths().Items()["/pets"].Get().Responses().Default()
	require.NotNil(t, defaultResp)
	assert.Equal(t, "Unexpected error", defaultResp.Value().Description())
}

// --- With Content ---

func TestParseOperationResponses_WithContent(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
          content:
            application/json:
              schema:
                type: array
                items:
                  type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	content := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Content()
	assert.Contains(t, content, "application/json")
}

// --- Reference Response ---

func TestParseOperationResponses_Reference(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "404":
          $ref: '#/components/responses/NotFound'
components:
  responses:
    NotFound:
      description: "Not found"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "#/components/responses/NotFound", result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["404"].Ref)
}
