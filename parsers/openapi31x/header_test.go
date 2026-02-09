package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for header.go - parseHeader function
// =============================================================================

// --- Basic Header ---

func TestParseHeader_Basic(t *testing.T) {
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
          headers:
            X-Rate-Limit:
              schema:
                type: integer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	header := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers["X-Rate-Limit"].Value
	require.NotNil(t, header)
	assert.NotNil(t, header.Schema)
}

// --- With Description ---

func TestParseHeader_WithDescription(t *testing.T) {
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
          headers:
            X-Rate-Limit:
              description: "Number of requests allowed per hour"
              schema:
                type: integer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	header := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers["X-Rate-Limit"].Value
	assert.Equal(t, "Number of requests allowed per hour", header.Description)
}

// --- Multiple Headers ---

func TestParseHeader_Multiple(t *testing.T) {
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
          headers:
            X-Rate-Limit-Limit:
              schema:
                type: integer
            X-Rate-Limit-Remaining:
              schema:
                type: integer
            X-Rate-Limit-Reset:
              schema:
                type: integer
            X-Request-Id:
              schema:
                type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	headers := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers
	assert.Len(t, headers, 4)
}

// --- Required and Deprecated ---

func TestParseHeader_RequiredDeprecated(t *testing.T) {
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
          headers:
            X-Required:
              required: true
              schema:
                type: string
            X-Deprecated:
              deprecated: true
              schema:
                type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	headers := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers
	assert.True(t, headers["X-Required"].Value.Required)
	assert.True(t, headers["X-Deprecated"].Value.Deprecated)
}

// --- Style and Explode ---

func TestParseHeader_StyleExplode(t *testing.T) {
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
          headers:
            X-Custom:
              style: simple
              explode: false
              schema:
                type: array
                items:
                  type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	header := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers["X-Custom"].Value
	assert.Equal(t, "simple", header.Style)
	require.NotNil(t, header.Explode)
	assert.False(t, *header.Explode)
}

// --- Content ---

func TestParseHeader_Content(t *testing.T) {
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
          headers:
            X-Custom:
              content:
                application/json:
                  schema:
                    type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	header := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers["X-Custom"].Value
	require.NotNil(t, header.Content)
	assert.Contains(t, header.Content, "application/json")
}

// --- Examples ---

func TestParseHeader_Examples(t *testing.T) {
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
          headers:
            X-Custom:
              schema:
                type: string
              examples:
                ex1:
                  value: "example1"
                ex2:
                  value: "example2"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	header := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers["X-Custom"].Value
	assert.Len(t, header.Examples, 2)
}

// --- Reference ---

func TestParseHeader_Reference(t *testing.T) {
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
	headerRef := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers["X-Rate-Limit"]
	assert.Equal(t, "#/components/headers/RateLimit", headerRef.Ref)
}

// --- Extensions ---

func TestParseHeader_Extensions(t *testing.T) {
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
          headers:
            X-Custom:
              x-internal: true
              schema:
                type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	header := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers["X-Custom"].Value
	require.NotNil(t, header.VendorExtensions)
	assert.Equal(t, true, header.VendorExtensions["x-internal"])
}
