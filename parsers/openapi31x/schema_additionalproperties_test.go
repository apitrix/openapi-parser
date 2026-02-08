package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_additionalproperties.go - ParseAdditionalProperties method
// =============================================================================

func TestParseSchemaAdditionalProperties_True(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Open:
      type: object
      additionalProperties: true
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["Open"].Value
	require.NotNil(t, schema.AdditionalPropertiesAllowed)
	assert.True(t, *schema.AdditionalPropertiesAllowed)
}

func TestParseSchemaAdditionalProperties_False(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Closed:
      type: object
      additionalProperties: false
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["Closed"].Value
	require.NotNil(t, schema.AdditionalPropertiesAllowed)
	assert.False(t, *schema.AdditionalPropertiesAllowed)
}

func TestParseSchemaAdditionalProperties_Schema(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    StringMap:
      type: object
      additionalProperties:
        type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["StringMap"].Value
	require.NotNil(t, schema.AdditionalProperties)
	assert.Equal(t, "string", schema.AdditionalProperties.Value.Type.Single)
}

func TestParseSchemaAdditionalProperties_ComplexSchema(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    ObjectMap:
      type: object
      additionalProperties:
        type: object
        properties:
          id:
            type: integer
          name:
            type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["ObjectMap"].Value
	require.NotNil(t, schema.AdditionalProperties)
	addProps := schema.AdditionalProperties.Value
	assert.Equal(t, "object", addProps.Type.Single)
	assert.Len(t, addProps.Properties, 2)
}

func TestParseSchemaAdditionalProperties_Reference(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    PetMap:
      type: object
      additionalProperties:
        $ref: '#/components/schemas/Pet'
    Pet:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["PetMap"].Value
	require.NotNil(t, schema.AdditionalProperties)
	assert.Equal(t, "#/components/schemas/Pet", schema.AdditionalProperties.Ref)
}

func TestParseSchemaAdditionalProperties_Missing(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    NoAddProps:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["NoAddProps"].Value
	assert.Nil(t, schema.AdditionalProperties)
	assert.Nil(t, schema.AdditionalPropertiesAllowed)
}

func TestParseSchemaAdditionalProperties_WithProperties(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Mixed:
      type: object
      properties:
        name:
          type: string
      additionalProperties:
        type: integer
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["Mixed"].Value
	// Has named properties
	assert.Len(t, schema.Properties, 1)
	// And additional properties schema
	require.NotNil(t, schema.AdditionalProperties)
	assert.Equal(t, "integer", schema.AdditionalProperties.Value.Type.Single)
}
