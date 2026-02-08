package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_oneof.go - ParseOneOf method
// =============================================================================

func TestParseSchemaOneOf_TwoSchemas(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      oneOf:
        - type: object
          properties:
            bark:
              type: boolean
        - type: object
          properties:
            meow:
              type: boolean
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	oneOf := doc.Components.Schemas["Pet"].Value.OneOf
	assert.Len(t, oneOf, 2)
}

func TestParseSchemaOneOf_WithReferences(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      oneOf:
        - $ref: '#/components/schemas/Cat'
        - $ref: '#/components/schemas/Dog'
    Cat:
      type: object
    Dog:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	oneOf := doc.Components.Schemas["Pet"].Value.OneOf
	assert.Len(t, oneOf, 2)
	assert.Equal(t, "#/components/schemas/Cat", oneOf[0].Ref)
	assert.Equal(t, "#/components/schemas/Dog", oneOf[1].Ref)
}

func TestParseSchemaOneOf_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Simple:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	oneOf := doc.Components.Schemas["Simple"].Value.OneOf
	assert.Nil(t, oneOf)
}

func TestParseSchemaOneOf_WithDiscriminator(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      oneOf:
        - $ref: '#/components/schemas/Cat'
        - $ref: '#/components/schemas/Dog'
      discriminator:
        propertyName: petType
    Cat:
      type: object
    Dog:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["Pet"].Value
	assert.Len(t, schema.OneOf, 2)
	require.NotNil(t, schema.Discriminator)
	assert.Equal(t, "petType", schema.Discriminator.PropertyName)
}

func TestParseSchemaOneOf_ManyOptions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Value:
      oneOf:
        - type: string
        - type: integer
        - type: number
        - type: boolean
        - type: array
          items:
            type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	oneOf := doc.Components.Schemas["Value"].Value.OneOf
	assert.Len(t, oneOf, 5)
}
