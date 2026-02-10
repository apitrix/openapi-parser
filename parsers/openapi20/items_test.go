package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for items.go - parseItems
// =============================================================================

// --- Basic Items ---

func TestParseItems_Basic(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: tags
          in: query
          type: array
          items:
            type: string
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	items := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value.Items()
	assert.Equal(t, "string", items.Type())
}

// --- Items with Format ---

func TestParseItems_WithFormat(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: ids
          in: query
          type: array
          items:
            type: integer
            format: int64
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	items := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value.Items()
	assert.Equal(t, "integer", items.Type())
	assert.Equal(t, "int64", items.Format())
}

// --- Items with Validation ---

func TestParseItems_WithValidation(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: scores
          in: query
          type: array
          items:
            type: integer
            minimum: 0
            maximum: 100
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	items := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value.Items()
	require.NotNil(t, items.Minimum())
	require.NotNil(t, items.Maximum())
	assert.Equal(t, float64(0), *items.Minimum())
	assert.Equal(t, float64(100), *items.Maximum())
}

// --- Items with Enum ---

func TestParseItems_WithEnum(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: statuses
          in: query
          type: array
          items:
            type: string
            enum:
              - available
              - pending
              - sold
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	items := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value.Items()
	require.Len(t, items.Enum(), 3)
	assert.Equal(t, "available", items.Enum()[0])
}

// --- Nested Items ---

func TestParseItems_NestedItems(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /matrix:
    get:
      parameters:
        - name: matrix
          in: query
          type: array
          items:
            type: array
            items:
              type: integer
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	items := result.Document.Paths().Items()["/matrix"].Get().Parameters()[0].Value.Items()
	assert.Equal(t, "array", items.Type())
	require.NotNil(t, items.Items())
	assert.Equal(t, "integer", items.Items().Type())
}

// --- Items with CollectionFormat ---

func TestParseItems_CollectionFormat(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: ids
          in: query
          type: array
          collectionFormat: pipes
          items:
            type: string
            collectionFormat: csv
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	items := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value.Items()
	assert.Equal(t, "csv", items.CollectionFormat())
}

// --- Items with Default ---

func TestParseItems_Default(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: tags
          in: query
          type: array
          items:
            type: string
            default: "unknown"
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	items := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value.Items()
	assert.Equal(t, "unknown", items.Default())
}

// --- Items Extensions ---

func TestParseItems_Extensions(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: tags
          in: query
          type: array
          items:
            type: string
            x-custom: "value"
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	items := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value.Items()
	assert.Equal(t, "value", items.VendorExtensions["x-custom"])
}
