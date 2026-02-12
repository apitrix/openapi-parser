package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_dependentschemas.go - dependentSchemas parsing
// =============================================================================

func TestParseSchema_DependentSchemas_HappyPath(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Address:
      type: object
      dependentSchemas:
        street:
          required:
            - city
            - state
        country:
          properties:
            zipCode:
              type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Address"].Value()
	require.NotNil(t, schema.DependentSchemas())
	require.Contains(t, schema.DependentSchemas(), "street")
	require.Contains(t, schema.DependentSchemas(), "country")
	assert.Len(t, schema.DependentSchemas()["street"].Value().Required(), 2)
}

func TestParseSchema_DependentSchemas_Empty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Simple:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Simple"].Value()
	assert.Nil(t, schema.DependentSchemas())
}
