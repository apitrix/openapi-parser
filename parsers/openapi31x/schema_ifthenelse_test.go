package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_ifthenelse.go - if/then/else parsing
// =============================================================================

func TestParseSchema_IfThenElse_HappyPath(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Conditional:
      type: object
      if:
        properties:
          kind:
            const: dog
      then:
        properties:
          bark:
            type: boolean
      else:
        properties:
          purr:
            type: boolean
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Conditional"].Value()
	require.NotNil(t, schema.If())
	require.NotNil(t, schema.Then())
	require.NotNil(t, schema.Else())
	assert.Contains(t, schema.If().Value().Properties(), "kind")
	assert.Contains(t, schema.Then().Value().Properties(), "bark")
	assert.Contains(t, schema.Else().Value().Properties(), "purr")
}

func TestParseSchema_IfOnly(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Test:
      type: object
      if:
        properties:
          status:
            const: active
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Test"].Value()
	require.NotNil(t, schema.If())
	assert.Nil(t, schema.Then())
	assert.Nil(t, schema.Else())
}

func TestParseSchema_NoIfThenElse(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Simple:
      type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Simple"].Value()
	assert.Nil(t, schema.If())
	assert.Nil(t, schema.Then())
	assert.Nil(t, schema.Else())
}
