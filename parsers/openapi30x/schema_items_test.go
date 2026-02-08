package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_items.go - ParseItems method
// =============================================================================

func TestParseSchemaItems_StringArray(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Tags:
      type: array
      items:
        type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	items := doc.Components.Schemas["Tags"].Value.Items
	require.NotNil(t, items)
	assert.Equal(t, "string", items.Value.Type)
}

func TestParseSchemaItems_ObjectArray(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pets:
      type: array
      items:
        type: object
        properties:
          name:
            type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	items := doc.Components.Schemas["Pets"].Value.Items
	require.NotNil(t, items)
	assert.Equal(t, "object", items.Value.Type)
	assert.NotEmpty(t, items.Value.Properties)
}

func TestParseSchemaItems_Reference(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    PetList:
      type: array
      items:
        $ref: '#/components/schemas/Pet'
    Pet:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	items := doc.Components.Schemas["PetList"].Value.Items
	require.NotNil(t, items)
	assert.Equal(t, "#/components/schemas/Pet", items.Ref)
}

func TestParseSchemaItems_NestedArray(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Matrix:
      type: array
      items:
        type: array
        items:
          type: integer
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	items := doc.Components.Schemas["Matrix"].Value.Items
	require.NotNil(t, items)
	assert.Equal(t, "array", items.Value.Type)
	require.NotNil(t, items.Value.Items)
	assert.Equal(t, "integer", items.Value.Items.Value.Type)
}

func TestParseSchemaItems_Missing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    NotArray:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	items := doc.Components.Schemas["NotArray"].Value.Items
	assert.Nil(t, items)
}

func TestParseSchemaItems_EnumArray(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Statuses:
      type: array
      items:
        type: string
        enum:
          - pending
          - active
          - inactive
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	items := doc.Components.Schemas["Statuses"].Value.Items
	require.NotNil(t, items)
	assert.Len(t, items.Value.Enum, 3)
}
