package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_additionalproperties.go - ParseAdditionalProperties method
// =============================================================================

func TestParseSchemaAdditionalProperties_True(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Open"].Value
	require.NotNil(t, schema.AdditionalPropertiesAllowed())
	assert.True(t, *schema.AdditionalPropertiesAllowed())
}

func TestParseSchemaAdditionalProperties_False(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Closed"].Value
	require.NotNil(t, schema.AdditionalPropertiesAllowed())
	assert.False(t, *schema.AdditionalPropertiesAllowed())
}

func TestParseSchemaAdditionalProperties_Schema(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["StringMap"].Value
	require.NotNil(t, schema.AdditionalProperties())
	assert.Equal(t, "string", schema.AdditionalProperties().Value.Type())
}

func TestParseSchemaAdditionalProperties_ComplexSchema(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["ObjectMap"].Value
	require.NotNil(t, schema.AdditionalProperties())
	addProps := schema.AdditionalProperties().Value
	assert.Equal(t, "object", addProps.Type())
	assert.Len(t, addProps.Properties(), 2)
}

func TestParseSchemaAdditionalProperties_Reference(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["PetMap"].Value
	require.NotNil(t, schema.AdditionalProperties())
	assert.Equal(t, "#/components/schemas/Pet", schema.AdditionalProperties().Ref)
}

func TestParseSchemaAdditionalProperties_Missing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    NoAddProps:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["NoAddProps"].Value
	assert.Nil(t, schema.AdditionalProperties())
	assert.Nil(t, schema.AdditionalPropertiesAllowed())
}

func TestParseSchemaAdditionalProperties_WithProperties(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Mixed"].Value
	// Has named properties
	assert.Len(t, schema.Properties(), 1)
	// And additional properties schema
	require.NotNil(t, schema.AdditionalProperties())
	assert.Equal(t, "integer", schema.AdditionalProperties().Value.Type())
}
