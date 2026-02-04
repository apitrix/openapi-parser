package openapi30

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_allof.go - ParseAllOf method
// =============================================================================

func TestParseSchemaAllOf_TwoSchemas(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    PetWithId:
      allOf:
        - type: object
          properties:
            id:
              type: integer
        - type: object
          properties:
            name:
              type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	allOf := doc.Components.Schemas["PetWithId"].Value.AllOf
	assert.Len(t, allOf, 2)
}

func TestParseSchemaAllOf_WithReferences(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Cat:
      allOf:
        - $ref: '#/components/schemas/Pet'
        - type: object
          properties:
            meow:
              type: boolean
    Pet:
      type: object
      properties:
        name:
          type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	allOf := doc.Components.Schemas["Cat"].Value.AllOf
	assert.Len(t, allOf, 2)
	assert.Equal(t, "#/components/schemas/Pet", allOf[0].Ref)
}

func TestParseSchemaAllOf_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    NoAllOf:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	allOf := doc.Components.Schemas["NoAllOf"].Value.AllOf
	assert.Nil(t, allOf)
}

func TestParseSchemaAllOf_MultipleReferences(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Combined:
      allOf:
        - $ref: '#/components/schemas/Base'
        - $ref: '#/components/schemas/Timestamped'
        - $ref: '#/components/schemas/Audited'
    Base:
      type: object
    Timestamped:
      type: object
    Audited:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	allOf := doc.Components.Schemas["Combined"].Value.AllOf
	assert.Len(t, allOf, 3)
}

func TestParseSchemaAllOf_MixedRefAndInline(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    ExtendedPet:
      allOf:
        - $ref: '#/components/schemas/Pet'
        - type: object
          properties:
            category:
              type: string
        - type: object
          properties:
            tags:
              type: array
              items:
                type: string
    Pet:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	allOf := doc.Components.Schemas["ExtendedPet"].Value.AllOf
	assert.Len(t, allOf, 3)
	// First is reference
	assert.Equal(t, "#/components/schemas/Pet", allOf[0].Ref)
	// Others are inline
	assert.NotEmpty(t, allOf[1].Value.Properties)
	assert.NotEmpty(t, allOf[2].Value.Properties)
}
