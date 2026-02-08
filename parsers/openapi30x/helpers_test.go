package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for helpers.go - helper functions
// =============================================================================

// --- Node Helpers via Integration ---

// These tests verify helper functions indirectly through parsing

func TestHelpers_StringParsing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test Title"
  description: "Test Description"
  version: "1.0.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "Test Title", doc.Info.Title)
	assert.Equal(t, "Test Description", doc.Info.Description)
}

func TestHelpers_BoolParsing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Test:
      type: object
      deprecated: true
      nullable: true
      readOnly: false
      writeOnly: false
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["Test"].Value
	assert.True(t, schema.Deprecated)
	assert.True(t, schema.Nullable)
	assert.False(t, schema.ReadOnly)
	assert.False(t, schema.WriteOnly)
}

func TestHelpers_IntParsing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Test:
      type: string
      minLength: 5
      maxLength: 100
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["Test"].Value
	require.NotNil(t, schema.MinLength)
	require.NotNil(t, schema.MaxLength)
	assert.Equal(t, uint64(5), *schema.MinLength)
	assert.Equal(t, uint64(100), *schema.MaxLength)
}

func TestHelpers_FloatParsing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Test:
      type: number
      minimum: 0.5
      maximum: 99.9
      multipleOf: 0.1
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["Test"].Value
	require.NotNil(t, schema.Minimum)
	require.NotNil(t, schema.Maximum)
	require.NotNil(t, schema.MultipleOf)
	assert.Equal(t, 0.5, *schema.Minimum)
	assert.Equal(t, 99.9, *schema.Maximum)
	assert.Equal(t, 0.1, *schema.MultipleOf)
}

func TestHelpers_ArrayParsing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Test:
      type: object
      required:
        - id
        - name
        - status
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Components.Schemas["Test"].Value
	assert.Len(t, schema.Required, 3)
	assert.Contains(t, schema.Required, "id")
	assert.Contains(t, schema.Required, "name")
	assert.Contains(t, schema.Required, "status")
}

func TestHelpers_MapParsing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Test:
      type: object
      properties:
        a:
          type: string
        b:
          type: integer
        c:
          type: boolean
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	props := doc.Components.Schemas["Test"].Value.Properties
	assert.Len(t, props, 3)
	assert.Contains(t, props, "a")
	assert.Contains(t, props, "b")
	assert.Contains(t, props, "c")
}

func TestHelpers_ExtensionParsing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
  x-string: "value"
  x-number: 42
  x-bool: true
  x-array:
    - a
    - b
  x-object:
    key: value
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ext := doc.Info.VendorExtensions
	assert.Equal(t, "value", ext["x-string"])
	assert.Equal(t, 42, ext["x-number"])
	assert.Equal(t, true, ext["x-bool"])
}

func TestHelpers_EnumParsing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Status:
      type: string
      enum:
        - pending
        - active
        - inactive
        - deleted
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	enum := doc.Components.Schemas["Status"].Value.Enum
	assert.Len(t, enum, 4)
}

func TestHelpers_DefaultParsing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    StringDefault:
      type: string
      default: "hello"
    IntDefault:
      type: integer
      default: 42
    BoolDefault:
      type: boolean
      default: true
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "hello", doc.Components.Schemas["StringDefault"].Value.Default)
	assert.Equal(t, 42, doc.Components.Schemas["IntDefault"].Value.Default)
	assert.Equal(t, true, doc.Components.Schemas["BoolDefault"].Value.Default)
}
