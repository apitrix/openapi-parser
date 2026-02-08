package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_unevaluated.go - unevaluatedItems/unevaluatedProperties
// =============================================================================

func TestParseSchema_UnevaluatedProperties_HappyPath(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Strict:
      type: object
      properties:
        name:
          type: string
      unevaluatedProperties:
        type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["Strict"].Value
	require.NotNil(t, schema.UnevaluatedProperties)
	assert.Equal(t, "string", schema.UnevaluatedProperties.Value.Type.Single)
}

func TestParseSchema_UnevaluatedItems_HappyPath(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    StrictArray:
      type: array
      items:
        type: string
      unevaluatedItems:
        type: integer
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["StrictArray"].Value
	require.NotNil(t, schema.UnevaluatedItems)
	assert.Equal(t, "integer", schema.UnevaluatedItems.Value.Type.Single)
}

func TestParseSchema_NoUnevaluated(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Normal:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["Normal"].Value
	assert.Nil(t, schema.UnevaluatedProperties)
	assert.Nil(t, schema.UnevaluatedItems)
}
