package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for response.go - parseResponse, parseResponses
// =============================================================================

// --- Basic Response ---

func TestParseResponse_Basic(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "Success response"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	resp := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value()
	assert.Equal(t, "Success response", resp.Description())
}

// --- Response with Schema ---

func TestParseResponse_WithSchema(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
          schema:
            type: array
            items:
              type: string
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Schema().Value()
	assert.Equal(t, "array", schema.Type())
}

// --- Response with Headers ---

func TestParseResponse_WithHeaders(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
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
              type: integer
              description: "Rate limit"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	headers := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Headers()
	require.Contains(t, headers, "X-Rate-Limit")
	assert.Equal(t, "integer", headers["X-Rate-Limit"].Type())
}

// --- Response with Examples ---

func TestParseResponse_WithExamples(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
          examples:
            application/json:
              name: "Fluffy"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	examples := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Examples()
	require.Contains(t, examples, "application/json")
}

// --- Default Response ---

func TestParseResponse_Default(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        default:
          description: "Unexpected error"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.Paths().Items()["/pets"].Get().Responses().Default())
	assert.Equal(t, "Unexpected error", result.Document.Paths().Items()["/pets"].Get().Responses().Default().Value().Description())
}

// --- Response $ref ---

func TestParseResponse_Ref(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "404":
          $ref: "#/responses/NotFound"
responses:
  NotFound:
    description: "Not found"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "#/responses/NotFound", result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["404"].Ref)
}

// --- Response Extensions ---

func TestParseResponse_Extensions(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
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
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	ext := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().VendorExtensions
	assert.Equal(t, "value", ext["x-custom"])
}
