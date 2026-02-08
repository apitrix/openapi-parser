package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_not.go - ParseNot method
// =============================================================================

func TestParseSchemaNot_SimpleType(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    NotString:
      not:
        type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	not := doc.Components.Schemas["NotString"].Value.Not
	require.NotNil(t, not)
	assert.Equal(t, "string", not.Value.Type)
}

func TestParseSchemaNot_Reference(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    NotPet:
      not:
        $ref: '#/components/schemas/Pet'
    Pet:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	not := doc.Components.Schemas["NotPet"].Value.Not
	require.NotNil(t, not)
	assert.Equal(t, "#/components/schemas/Pet", not.Ref)
}

func TestParseSchemaNot_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	not := doc.Components.Schemas["Simple"].Value.Not
	assert.Nil(t, not)
}

func TestParseSchemaNot_ComplexSchema(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    NotComplex:
      not:
        type: object
        properties:
          forbidden:
            type: string
        required:
          - forbidden
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	not := doc.Components.Schemas["NotComplex"].Value.Not
	require.NotNil(t, not)
	assert.Equal(t, "object", not.Value.Type)
	assert.NotEmpty(t, not.Value.Properties)
}
