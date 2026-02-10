package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for parameter.go - parseParameter function
// =============================================================================

// --- Query Parameters ---

func TestParseParameter_QueryParameter(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: limit
          in: query
          description: "Maximum results"
          type: integer
          required: false
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	param := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value
	assert.Equal(t, "limit", param.Name())
	assert.Equal(t, "query", param.In())
	assert.Equal(t, "integer", param.Type())
	assert.Equal(t, "Maximum results", param.Description())
}

// --- Path Parameters ---

func TestParseParameter_PathParameter(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets/{petId}:
    get:
      parameters:
        - name: petId
          in: path
          required: true
          type: string
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	param := result.Document.Paths().Items()["/pets/{petId}"].Get().Parameters()[0].Value
	assert.Equal(t, "petId", param.Name())
	assert.Equal(t, "path", param.In())
	assert.True(t, param.Required())
}

// --- Body Parameters ---

func TestParseParameter_BodyParameter(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    post:
      parameters:
        - name: body
          in: body
          required: true
          schema:
            type: object
            properties:
              name:
                type: string
      responses:
        "201":
          description: "Created"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	param := result.Document.Paths().Items()["/pets"].Post().Parameters()[0].Value
	assert.Equal(t, "body", param.In())
	require.NotNil(t, param.Schema())
	assert.Equal(t, "object", param.Schema().Value.Type())
}

// --- Array Parameters ---

func TestParseParameter_ArrayParameter(t *testing.T) {
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
          collectionFormat: csv
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	param := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value
	assert.Equal(t, "array", param.Type())
	assert.Equal(t, "csv", param.CollectionFormat())
	require.NotNil(t, param.Items())
	assert.Equal(t, "string", param.Items().Type())
}

// --- Parameter Validation ---

func TestParseParameter_Validation(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: limit
          in: query
          type: integer
          minimum: 1
          maximum: 100
          default: 10
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	param := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value
	require.NotNil(t, param.Minimum())
	require.NotNil(t, param.Maximum())
	assert.Equal(t, float64(1), *param.Minimum())
	assert.Equal(t, float64(100), *param.Maximum())
	assert.Equal(t, 10, param.Default())
}

// --- Parameter with Enum ---

func TestParseParameter_Enum(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: status
          in: query
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
	param := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value
	require.Len(t, param.Enum(), 3)
	assert.Equal(t, "available", param.Enum()[0])
}

// --- Parameter $ref ---

func TestParseParameter_Ref(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - $ref: "#/parameters/limitParam"
      responses:
        "200":
          description: "OK"
parameters:
  limitParam:
    name: limit
    in: query
    type: integer
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	paramRef := result.Document.Paths().Items()["/pets"].Get().Parameters()[0]
	assert.Equal(t, "#/parameters/limitParam", paramRef.Ref)
}

// --- Header Parameters ---

func TestParseParameter_HeaderParameter(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: X-Request-ID
          in: header
          type: string
          required: true
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	param := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value
	assert.Equal(t, "X-Request-ID", param.Name())
	assert.Equal(t, "header", param.In())
}
