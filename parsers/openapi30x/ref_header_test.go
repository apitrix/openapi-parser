package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for ref_header.go - header reference parsing
// =============================================================================

// --- Basic Reference ---

func TestParseRefHeader_Basic(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
          headers:
            X-Rate-Limit:
              $ref: '#/components/headers/RateLimit'
components:
  headers:
    RateLimit:
      schema:
        type: integer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers["X-Rate-Limit"]
	assert.Equal(t, "#/components/headers/RateLimit", ref.Ref)
}

// --- Multiple References ---

func TestParseRefHeader_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
          headers:
            X-Rate-Limit:
              $ref: '#/components/headers/RateLimit'
            X-Request-Id:
              $ref: '#/components/headers/RequestId'
components:
  headers:
    RateLimit:
      schema:
        type: integer
    RequestId:
      schema:
        type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	headers := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers
	assert.Equal(t, "#/components/headers/RateLimit", headers["X-Rate-Limit"].Ref)
	assert.Equal(t, "#/components/headers/RequestId", headers["X-Request-Id"].Ref)
}

// --- Mixed Inline and Reference ---

func TestParseRefHeader_Mixed(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
          headers:
            X-Rate-Limit:
              $ref: '#/components/headers/RateLimit'
            X-Custom:
              description: "Custom header"
              schema:
                type: string
components:
  headers:
    RateLimit:
      schema:
        type: integer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	headers := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers
	assert.Equal(t, "#/components/headers/RateLimit", headers["X-Rate-Limit"].Ref)
	assert.Equal(t, "Custom header", headers["X-Custom"].Value.Description)
}
