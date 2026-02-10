package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for response.go - parseResponse function
// =============================================================================

// --- Basic Response ---

func TestParseResponse_Basic(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "Successful response"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	resp := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value
	assert.Equal(t, "Successful response", resp.Description())
}

// --- Multiple Status Codes ---

func TestParseResponse_MultipleStatusCodes(t *testing.T) {
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
        "201":
          description: "Created"
        "400":
          description: "Bad Request"
        "401":
          description: "Unauthorized"
        "403":
          description: "Forbidden"
        "404":
          description: "Not Found"
        "500":
          description: "Internal Server Error"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	codes := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()
	assert.Len(t, codes, 7)
}

// --- Default Response ---

func TestParseResponse_Default(t *testing.T) {
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
        default:
          description: "Unexpected error"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	defaultResp := result.Document.Paths().Items()["/pets"].Get().Responses().Default()
	require.NotNil(t, defaultResp)
	assert.Equal(t, "Unexpected error", defaultResp.Value.Description())
}

// --- Content Types ---

func TestParseResponse_MultipleContentTypes(t *testing.T) {
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
          content:
            application/json:
              schema:
                type: object
            application/xml:
              schema:
                type: object
            text/plain:
              schema:
                type: string
            text/html:
              schema:
                type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	content := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value.Content()
	assert.Len(t, content, 4)
}

// --- Headers ---

func TestParseResponse_Headers(t *testing.T) {
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
            X-Rate-Limit-Limit:
              description: "Request limit"
              schema:
                type: integer
            X-Rate-Limit-Remaining:
              description: "Remaining requests"
              schema:
                type: integer
            X-Rate-Limit-Reset:
              description: "Reset time"
              schema:
                type: integer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	headers := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value.Headers()
	assert.Len(t, headers, 3)
}

// --- Links ---

func TestParseResponse_Links(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /users/{id}:
    get:
      responses:
        "200":
          description: "OK"
          links:
            GetUserPets:
              operationId: getUserPets
              parameters:
                userId: '$response.body#/id'
            GetUserOrders:
              operationId: getUserOrders
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	links := result.Document.Paths().Items()["/users/{id}"].Get().Responses().Codes()["200"].Value.Links()
	assert.Len(t, links, 2)
}

// --- Reference ---

func TestParseResponse_Reference(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "404":
          $ref: '#/components/responses/NotFound'
        "500":
          $ref: '#/components/responses/InternalError'
components:
  responses:
    NotFound:
      description: "Not found"
    InternalError:
      description: "Internal error"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	resp404 := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["404"]
	assert.Equal(t, "#/components/responses/NotFound", resp404.Ref)
}

// --- Extensions ---

func TestParseResponse_Extensions(t *testing.T) {
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
          x-custom: "value"
          x-internal: true
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	resp := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value
	require.NotNil(t, resp.VendorExtensions)
	assert.Equal(t, "value", resp.VendorExtensions["x-custom"])
}

// --- Complete Response ---

func TestParseResponse_Complete(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "List of pets"
          headers:
            X-Total-Count:
              schema:
                type: integer
          content:
            application/json:
              schema:
                type: array
                items:
                  type: object
          links:
            GetPetById:
              operationId: getPetById
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	resp := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value
	assert.NotEmpty(t, resp.Description())
	assert.NotEmpty(t, resp.Headers())
	assert.NotEmpty(t, resp.Content())
	assert.NotEmpty(t, resp.Links())
}

// --- Wildcard Status Codes ---

func TestParseResponse_WildcardCodes(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "2XX":
          description: "Success"
        "4XX":
          description: "Client error"
        "5XX":
          description: "Server error"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	codes := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()
	assert.Contains(t, codes, "2XX")
	assert.Contains(t, codes, "4XX")
	assert.Contains(t, codes, "5XX")
}
