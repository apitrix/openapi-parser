package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_properties.go - ParseProperties method
// =============================================================================

func TestParseSchemaProperties_SingleProperty(t *testing.T) {
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
        name:
          type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	props := result.Document.Components().Schemas()["Pet"].Value.Properties()
	assert.Len(t, props, 1)
	assert.Contains(t, props, "name")
}

func TestParseSchemaProperties_MultipleProperties(t *testing.T) {
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
        id:
          type: integer
        name:
          type: string
        tag:
          type: string
        status:
          type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	props := result.Document.Components().Schemas()["Pet"].Value.Properties()
	assert.Len(t, props, 4)
}

func TestParseSchemaProperties_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Empty:
      type: object
      properties: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	props := result.Document.Components().Schemas()["Empty"].Value.Properties()
	assert.Empty(t, props)
}

func TestParseSchemaProperties_Missing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    NoProps:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	props := result.Document.Components().Schemas()["NoProps"].Value.Properties()
	assert.Nil(t, props)
}

func TestParseSchemaProperties_NestedObjects(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    User:
      type: object
      properties:
        address:
          type: object
          properties:
            street:
              type: string
            city:
              type: string
            zipCode:
              type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	userProps := result.Document.Components().Schemas()["User"].Value.Properties()
	address := userProps["address"].Value
	require.NotNil(t, address)
	assert.Len(t, address.Properties(), 3)
}

func TestParseSchemaProperties_WithReferences(t *testing.T) {
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
          $ref: '#/components/schemas/Owner'
        category:
          $ref: '#/components/schemas/Category'
    Owner:
      type: object
    Category:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	props := result.Document.Components().Schemas()["Pet"].Value.Properties()
	assert.Equal(t, "#/components/schemas/Owner", props["owner"].Ref)
	assert.Equal(t, "#/components/schemas/Category", props["category"].Ref)
}

func TestParseSchemaProperties_SpecialNames(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Config:
      type: object
      properties:
        "property-with-dashes":
          type: string
        "property.with.dots":
          type: string
        "123numeric":
          type: integer
        "_underscore":
          type: boolean
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	props := result.Document.Components().Schemas()["Config"].Value.Properties()
	assert.Len(t, props, 4)
	assert.Contains(t, props, "property-with-dashes")
	assert.Contains(t, props, "property.with.dots")
	assert.Contains(t, props, "123numeric")
	assert.Contains(t, props, "_underscore")
}

func TestParseSchemaProperties_MixedTypes(t *testing.T) {
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
        stringProp:
          type: string
        intProp:
          type: integer
        numProp:
          type: number
        boolProp:
          type: boolean
        arrayProp:
          type: array
          items:
            type: string
        objectProp:
          type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	props := result.Document.Components().Schemas()["Mixed"].Value.Properties()
	assert.Equal(t, "string", props["stringProp"].Value.Type())
	assert.Equal(t, "integer", props["intProp"].Value.Type())
	assert.Equal(t, "number", props["numProp"].Value.Type())
	assert.Equal(t, "boolean", props["boolProp"].Value.Type())
	assert.Equal(t, "array", props["arrayProp"].Value.Type())
	assert.Equal(t, "object", props["objectProp"].Value.Type())
}
