package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema.go - parseSharedSchema and schemaParser methods
// Comprehensive schema parsing tests
// =============================================================================

// --- Basic Types ---

func TestParseSchema_StringType(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Name:
      type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Name"].Value()
	assert.Equal(t, "string", schema.Type().Single)
}

func TestParseSchema_IntegerType(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Count:
      type: integer
      format: int32
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Count"].Value()
	assert.Equal(t, "integer", schema.Type().Single)
	assert.Equal(t, "int32", schema.Format())
}

func TestParseSchema_NumberType(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Price:
      type: number
      format: double
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Price"].Value()
	assert.Equal(t, "number", schema.Type().Single)
	assert.Equal(t, "double", schema.Format())
}

func TestParseSchema_BooleanType(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Active:
      type: boolean
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Active"].Value()
	assert.Equal(t, "boolean", schema.Type().Single)
}

func TestParseSchema_ArrayType(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Tags"].Value()
	assert.Equal(t, "array", schema.Type().Single)
	require.NotNil(t, schema.Items())
}

func TestParseSchema_ObjectType(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
	schema := result.Document.Components().Schemas()["Pet"].Value()
	assert.Equal(t, "object", schema.Type().Single)
	assert.Len(t, schema.Properties(), 1)
}

// --- String Constraints ---

func TestParseSchema_StringConstraints(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Username:
      type: string
      minLength: 3
      maxLength: 50
      pattern: "^[a-zA-Z0-9_]+$"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Username"].Value()
	require.NotNil(t, schema.MinLength())
	assert.Equal(t, uint64(3), *schema.MinLength())
	require.NotNil(t, schema.MaxLength())
	assert.Equal(t, uint64(50), *schema.MaxLength())
	assert.Equal(t, "^[a-zA-Z0-9_]+$", schema.Pattern())
}

// --- Number Constraints ---

func TestParseSchema_NumberConstraints(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Age:
      type: integer
      minimum: 0
      maximum: 150
      exclusiveMinimum: 1
      exclusiveMaximum: 149
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Age"].Value()
	require.NotNil(t, schema.Minimum())
	assert.Equal(t, float64(0), *schema.Minimum())
	require.NotNil(t, schema.Maximum())
	assert.Equal(t, float64(150), *schema.Maximum())
	require.NotNil(t, schema.ExclusiveMinimum())
	assert.Equal(t, float64(1), *schema.ExclusiveMinimum())
	require.NotNil(t, schema.ExclusiveMaximum())
	assert.Equal(t, float64(149), *schema.ExclusiveMaximum())
}

func TestParseSchema_MultipleOf(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Even:
      type: integer
      multipleOf: 2
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Even"].Value()
	require.NotNil(t, schema.MultipleOf())
	assert.Equal(t, float64(2), *schema.MultipleOf())
}

// --- Array Constraints ---

func TestParseSchema_ArrayConstraints(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
      minItems: 1
      maxItems: 10
      uniqueItems: true
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Tags"].Value()
	require.NotNil(t, schema.MinItems())
	assert.Equal(t, uint64(1), *schema.MinItems())
	require.NotNil(t, schema.MaxItems())
	assert.Equal(t, uint64(10), *schema.MaxItems())
	assert.True(t, schema.UniqueItems())
}

// --- Object Constraints ---

func TestParseSchema_ObjectConstraints(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Config:
      type: object
      minProperties: 1
      maxProperties: 20
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Config"].Value()
	require.NotNil(t, schema.MinProperties())
	assert.Equal(t, uint64(1), *schema.MinProperties())
	require.NotNil(t, schema.MaxProperties())
	assert.Equal(t, uint64(20), *schema.MaxProperties())
}

// --- Required Fields ---

func TestParseSchema_RequiredFields(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      required:
        - id
        - name
        - status
      properties:
        id:
          type: integer
        name:
          type: string
        status:
          type: string
        tag:
          type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Pet"].Value()
	assert.Len(t, schema.Required(), 3)
	assert.Contains(t, schema.Required(), "id")
	assert.Contains(t, schema.Required(), "name")
	assert.Contains(t, schema.Required(), "status")
}

// --- Type Array (replaces Nullable in 3.1), ReadOnly, WriteOnly, Deprecated ---

func TestParseSchema_TypeArray(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    OptionalName:
      type:
        - string
        - "null"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["OptionalName"].Value()
	require.Len(t, schema.Type().Array, 2)
	assert.Contains(t, schema.Type().Array, "string")
	assert.Contains(t, schema.Type().Array, "null")
}

func TestParseSchema_ReadWriteOnly(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    User:
      type: object
      properties:
        id:
          type: integer
          readOnly: true
        password:
          type: string
          writeOnly: true
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["User"].Value()
	assert.True(t, schema.Properties()["id"].Value().ReadOnly())
	assert.True(t, schema.Properties()["password"].Value().WriteOnly())
}

func TestParseSchema_Deprecated(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    OldApi:
      type: object
      deprecated: true
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["OldApi"].Value()
	assert.True(t, schema.Deprecated())
}

// --- Enum and Default ---

func TestParseSchema_Enum(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Status:
      type: string
      enum:
        - available
        - pending
        - sold
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Status"].Value()
	assert.Len(t, schema.Enum(), 3)
}

func TestParseSchema_Default(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Count:
      type: integer
      default: 10
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Count"].Value()
	assert.Equal(t, 10, schema.Default())
}

func TestParseSchema_Example(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Name:
      type: string
      example: "John Doe"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Name"].Value()
	assert.Equal(t, "John Doe", schema.Example())
}

// --- Title and Description ---

func TestParseSchema_TitleDescription(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      title: "Pet Model"
      description: "A representation of a pet in the system"
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Pet"].Value()
	assert.Equal(t, "Pet Model", schema.Title())
	assert.Equal(t, "A representation of a pet in the system", schema.Description())
}

// --- Extensions ---

func TestParseSchema_Extensions(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      x-internal: true
      x-model-name: "PetModel"
      x-tags:
        - internal
        - deprecated
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Pet"].Value()
	require.NotNil(t, schema.VendorExtensions)
	assert.Equal(t, true, schema.VendorExtensions["x-internal"])
	assert.Equal(t, "PetModel", schema.VendorExtensions["x-model-name"])
}

// --- Node Source ---

func TestParseSchema_NodeSource(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Pet"].Value()
	assert.Greater(t, schema.Trix.Source.Start.Line, 0)
	assert.Greater(t, schema.Trix.Source.Start.Column, 0)
}

// --- Empty/Minimal Schema ---

func TestParseSchema_Empty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Any:
      {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Any"].Value()
	require.NotNil(t, schema)
	assert.True(t, schema.Type().IsEmpty())
}

// --- Complex Nested Schema ---

func TestParseSchema_DeeplyNested(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Order:
      type: object
      properties:
        items:
          type: array
          items:
            type: object
            properties:
              product:
                type: object
                properties:
                  name:
                    type: string
                  price:
                    type: number
              quantity:
                type: integer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Order"].Value()
	items := schema.Properties()["items"].Value()
	require.NotNil(t, items.Items())
	itemSchema := items.Items().Value()
	require.NotNil(t, itemSchema.Properties()["product"])
}
