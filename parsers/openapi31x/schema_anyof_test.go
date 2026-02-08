package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_anyof.go - ParseAnyOf method
// =============================================================================

func TestParseSchemaAnyOf_TwoSchemas(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    PetOrError:
      anyOf:
        - type: object
          properties:
            name:
              type: string
        - type: object
          properties:
            error:
              type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	anyOf := doc.Components.Schemas["PetOrError"].Value.AnyOf
	assert.Len(t, anyOf, 2)
}

func TestParseSchemaAnyOf_WithReferences(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Result:
      anyOf:
        - $ref: '#/components/schemas/Success'
        - $ref: '#/components/schemas/Error'
    Success:
      type: object
    Error:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	anyOf := doc.Components.Schemas["Result"].Value.AnyOf
	assert.Len(t, anyOf, 2)
	assert.Equal(t, "#/components/schemas/Success", anyOf[0].Ref)
	assert.Equal(t, "#/components/schemas/Error", anyOf[1].Ref)
}

func TestParseSchemaAnyOf_Empty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Simple:
      type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	anyOf := doc.Components.Schemas["Simple"].Value.AnyOf
	assert.Nil(t, anyOf)
}

func TestParseSchemaAnyOf_MixedTypes(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    FlexibleId:
      anyOf:
        - type: string
        - type: integer
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	anyOf := doc.Components.Schemas["FlexibleId"].Value.AnyOf
	assert.Len(t, anyOf, 2)
	assert.Equal(t, "string", anyOf[0].Value.Type.Single)
	assert.Equal(t, "integer", anyOf[1].Value.Type.Single)
}

func TestParseSchemaAnyOf_ComplexSchemas(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Flexible:
      anyOf:
        - type: object
          properties:
            type:
              type: string
              enum: ["cat"]
            meow:
              type: boolean
        - type: object
          properties:
            type:
              type: string
              enum: ["dog"]
            bark:
              type: boolean
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	anyOf := doc.Components.Schemas["Flexible"].Value.AnyOf
	assert.Len(t, anyOf, 2)
	// Both have properties
	assert.NotEmpty(t, anyOf[0].Value.Properties)
	assert.NotEmpty(t, anyOf[1].Value.Properties)
}
