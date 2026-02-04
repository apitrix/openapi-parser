package openapi30

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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Schemas, 1)
	assert.Contains(t, doc.Components.Schemas, "Pet")
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Schemas, 5)
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, doc.Components.Schemas)
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "string", doc.Components.Schemas["StringSchema"].Value.Type)
	assert.Equal(t, "integer", doc.Components.Schemas["IntSchema"].Value.Type)
	assert.Equal(t, "number", doc.Components.Schemas["NumberSchema"].Value.Type)
	assert.Equal(t, "boolean", doc.Components.Schemas["BoolSchema"].Value.Type)
	assert.Equal(t, "array", doc.Components.Schemas["ArraySchema"].Value.Type)
	assert.Equal(t, "object", doc.Components.Schemas["ObjectSchema"].Value.Type)
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	pet := doc.Components.Schemas["Pet"].Value
	assert.Equal(t, "#/components/schemas/User", pet.Properties["owner"].Ref)
	assert.Equal(t, "#/components/schemas/Tag", pet.Properties["tags"].Value.Items.Ref)
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Schemas, 4)
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["Pet"].Value
	assert.Len(t, schema.Required, 2)
	assert.Len(t, schema.Properties, 3)
	assert.Len(t, schema.Properties["status"].Value.Enum, 3)
}
