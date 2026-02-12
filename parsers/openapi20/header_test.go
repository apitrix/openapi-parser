package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for header.go - parseHeader, parseHeaders
// =============================================================================

// --- Basic Header ---

func TestParseHeader_Basic(t *testing.T) {
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
            X-Request-ID:
              type: string
              description: "Request ID"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	header := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Headers()["X-Request-ID"]
	assert.Equal(t, "string", header.Type())
	assert.Equal(t, "Request ID", header.Description())
}

// --- Header with Format ---

func TestParseHeader_WithFormat(t *testing.T) {
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
              format: int32
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	header := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Headers()["X-Rate-Limit"]
	assert.Equal(t, "integer", header.Type())
	assert.Equal(t, "int32", header.Format())
}

// --- Header with Validation ---

func TestParseHeader_WithValidation(t *testing.T) {
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
            X-Limit:
              type: integer
              minimum: 1
              maximum: 100
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	header := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Headers()["X-Limit"]
	require.NotNil(t, header.Minimum())
	require.NotNil(t, header.Maximum())
	assert.Equal(t, float64(1), *header.Minimum())
	assert.Equal(t, float64(100), *header.Maximum())
}

// --- Array Header ---

func TestParseHeader_Array(t *testing.T) {
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
            X-Tags:
              type: array
              items:
                type: string
              collectionFormat: csv
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	header := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Headers()["X-Tags"]
	assert.Equal(t, "array", header.Type())
	assert.Equal(t, "csv", header.CollectionFormat())
	require.NotNil(t, header.Items())
	assert.Equal(t, "string", header.Items().Type())
}

// --- Header with Enum ---

func TestParseHeader_Enum(t *testing.T) {
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
            X-Status:
              type: string
              enum:
                - active
                - pending
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	header := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Headers()["X-Status"]
	require.Len(t, header.Enum(), 2)
	assert.Equal(t, "active", header.Enum()[0])
}

// --- Header with Default ---

func TestParseHeader_Default(t *testing.T) {
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
            X-Page-Size:
              type: integer
              default: 20
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	header := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Headers()["X-Page-Size"]
	assert.Equal(t, 20, header.Default())
}

// --- Header Extensions ---

func TestParseHeader_Extensions(t *testing.T) {
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
            X-Custom:
              type: string
              x-deprecated: true
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	header := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Headers()["X-Custom"]
	assert.Equal(t, true, header.VendorExtensions["x-deprecated"])
}
