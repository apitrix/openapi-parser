package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for components_schemas.go - parseComponentsSchemas function
// =============================================================================

// --- Single Schema ---

func TestParseComponentsSchemas_Single(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().Schemas(), 1)
	assert.Contains(t, result.Document.Components().Schemas(), "Pet")
}

// --- Multiple Schemas ---

func TestParseComponentsSchemas_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
    User:
      type: object
    Order:
      type: object
    Category:
      type: object
    Tag:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().Schemas(), 5)
}

// --- Empty ---

func TestParseComponentsSchemas_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Components().Schemas())
}

// --- Schema Types ---

func TestParseComponentsSchemas_AllTypes(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    StringSchema:
      type: string
    IntSchema:
      type: integer
    NumberSchema:
      type: number
    BoolSchema:
      type: boolean
    ArraySchema:
      type: array
      items:
        type: string
    ObjectSchema:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "string", result.Document.Components().Schemas()["StringSchema"].Value.Type())
	assert.Equal(t, "integer", result.Document.Components().Schemas()["IntSchema"].Value.Type())
	assert.Equal(t, "number", result.Document.Components().Schemas()["NumberSchema"].Value.Type())
	assert.Equal(t, "boolean", result.Document.Components().Schemas()["BoolSchema"].Value.Type())
	assert.Equal(t, "array", result.Document.Components().Schemas()["ArraySchema"].Value.Type())
	assert.Equal(t, "object", result.Document.Components().Schemas()["ObjectSchema"].Value.Type())
}

// --- Reference Between Schemas ---

func TestParseComponentsSchemas_References(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      properties:
        owner:
          $ref: '#/components/schemas/User'
        tags:
          type: array
          items:
            $ref: '#/components/schemas/Tag'
    User:
      type: object
    Tag:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	pet := result.Document.Components().Schemas()["Pet"].Value
	assert.Equal(t, "#/components/schemas/User", pet.Properties()["owner"].Ref)
	assert.Equal(t, "#/components/schemas/Tag", pet.Properties()["tags"].Value.Items().Ref)
}

// --- Special Names ---

func TestParseComponentsSchemas_SpecialNames(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    "Pet.Model":
      type: object
    "user-response":
      type: object
    "API_Error":
      type: object
    "123Schema":
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().Schemas(), 4)
}

// --- Complex Schema ---

func TestParseComponentsSchemas_Complex(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      required:
        - id
        - name
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string
          minLength: 1
        status:
          type: string
          enum:
            - available
            - pending
            - sold
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Pet"].Value
	assert.Len(t, schema.Required(), 2)
	assert.Len(t, schema.Properties(), 3)
	assert.Len(t, schema.Properties()["status"].Value.Enum(), 3)
}
