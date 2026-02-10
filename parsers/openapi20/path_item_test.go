package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for pathitem.go - parsePathItem function
// =============================================================================

// --- Basic Paths ---

func TestParsePathItem_SimplePath(t *testing.T) {
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
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.Paths())
	require.Contains(t, result.Document.Paths().Items(), "/pets")
	require.NotNil(t, result.Document.Paths().Items()["/pets"].Get())
}

func TestParsePathItem_AllMethods(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /resource:
    get:
      responses:
        "200":
          description: "OK"
    put:
      responses:
        "200":
          description: "OK"
    post:
      responses:
        "201":
          description: "Created"
    delete:
      responses:
        "204":
          description: "Deleted"
    options:
      responses:
        "200":
          description: "OK"
    head:
      responses:
        "200":
          description: "OK"
    patch:
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	pathItem := result.Document.Paths().Items()["/resource"]
	assert.NotNil(t, pathItem.Get())
	assert.NotNil(t, pathItem.Put())
	assert.NotNil(t, pathItem.Post())
	assert.NotNil(t, pathItem.Delete())
	assert.NotNil(t, pathItem.Options())
	assert.NotNil(t, pathItem.Head())
	assert.NotNil(t, pathItem.Patch())
}

// --- Path Parameters ---

func TestParsePathItem_Parameters(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets/{petId}:
    parameters:
      - name: petId
        in: path
        required: true
        type: string
    get:
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	pathItem := result.Document.Paths().Items()["/pets/{petId}"]
	require.Len(t, pathItem.Parameters(), 1)
	assert.Equal(t, "petId", pathItem.Parameters()[0].Value.Name())
	assert.Equal(t, "path", pathItem.Parameters()[0].Value.In())
}

// --- Path with $ref ---

func TestParsePathItem_Ref(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    $ref: "#/paths/~1other"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "#/paths/~1other", result.Document.Paths().Items()["/pets"].Ref())
}

// --- Extensions ---

func TestParsePathItem_Extensions(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    x-internal: true
    get:
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.Paths().Items()["/pets"].VendorExtensions)
	assert.Equal(t, true, result.Document.Paths().Items()["/pets"].VendorExtensions["x-internal"])
}
