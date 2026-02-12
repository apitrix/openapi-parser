package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_prefixitems.go - prefixItems parsing
// =============================================================================

func TestParseSchema_PrefixItems_HappyPath(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Tuple:
      type: array
      prefixItems:
        - type: string
        - type: integer
        - type: boolean
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Tuple"].Value()
	require.Len(t, schema.PrefixItems(), 3)
	assert.Equal(t, "string", schema.PrefixItems()[0].Value().Type().Single)
	assert.Equal(t, "integer", schema.PrefixItems()[1].Value().Type().Single)
	assert.Equal(t, "boolean", schema.PrefixItems()[2].Value().Type().Single)
}

func TestParseSchema_PrefixItems_WithRef(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Tuple:
      type: array
      prefixItems:
        - $ref: '#/components/schemas/Name'
        - type: integer
    Name:
      type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Tuple"].Value()
	require.Len(t, schema.PrefixItems(), 2)
	assert.Equal(t, "#/components/schemas/Name", schema.PrefixItems()[0].Ref)
}

func TestParseSchema_PrefixItems_Empty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Simple:
      type: array
      items:
        type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Simple"].Value()
	assert.Nil(t, schema.PrefixItems())
}
