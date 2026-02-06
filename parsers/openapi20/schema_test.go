package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema.go - parseSchema function
// =============================================================================

// --- Basic Schema Types ---

func TestParseSchema_StringType(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Name:
    type: string
    minLength: 1
    maxLength: 100
    pattern: "^[a-zA-Z]+$"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := doc.Definitions["Name"].Value
	assert.Equal(t, "string", schema.Type)
	require.NotNil(t, schema.MinLength)
	require.NotNil(t, schema.MaxLength)
	assert.Equal(t, uint64(1), *schema.MinLength)
	assert.Equal(t, uint64(100), *schema.MaxLength)
	assert.Equal(t, "^[a-zA-Z]+$", schema.Pattern)
}

func TestParseSchema_IntegerType(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Age:
    type: integer
    format: int32
    minimum: 0
    maximum: 150
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := doc.Definitions["Age"].Value
	assert.Equal(t, "integer", schema.Type)
	assert.Equal(t, "int32", schema.Format)
	require.NotNil(t, schema.Minimum)
	require.NotNil(t, schema.Maximum)
}

// --- Object Schema ---

func TestParseSchema_Object(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Pet:
    type: object
    required:
      - name
    properties:
      id:
        type: integer
      name:
        type: string
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := doc.Definitions["Pet"].Value
	assert.Equal(t, "object", schema.Type)
	assert.Equal(t, []string{"name"}, schema.Required)
	require.Contains(t, schema.Properties, "id")
	require.Contains(t, schema.Properties, "name")
}

// --- Array Schema ---

func TestParseSchema_Array(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  PetList:
    type: array
    items:
      $ref: "#/definitions/Pet"
    minItems: 0
    maxItems: 100
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := doc.Definitions["PetList"].Value
	assert.Equal(t, "array", schema.Type)
	require.NotNil(t, schema.Items)
	assert.Equal(t, "#/definitions/Pet", schema.Items.Ref)
}

// --- AllOf Composition ---

func TestParseSchema_AllOf(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Dog:
    allOf:
      - $ref: "#/definitions/Pet"
      - type: object
        properties:
          breed:
            type: string
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := doc.Definitions["Dog"].Value
	require.Len(t, schema.AllOf, 2)
	assert.Equal(t, "#/definitions/Pet", schema.AllOf[0].Ref)
	assert.Equal(t, "object", schema.AllOf[1].Value.Type)
}

// --- AdditionalProperties ---

func TestParseSchema_AdditionalPropertiesSchema(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  StringMap:
    type: object
    additionalProperties:
      type: string
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := doc.Definitions["StringMap"].Value
	require.NotNil(t, schema.AdditionalProperties)
	assert.Equal(t, "string", schema.AdditionalProperties.Value.Type)
}

func TestParseSchema_AdditionalPropertiesBoolean(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Closed:
    type: object
    additionalProperties: false
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := doc.Definitions["Closed"].Value
	require.NotNil(t, schema.AdditionalPropertiesAllowed)
	assert.False(t, *schema.AdditionalPropertiesAllowed)
}

// --- Discriminator ---

func TestParseSchema_Discriminator(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Pet:
    type: object
    discriminator: petType
    properties:
      petType:
        type: string
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := doc.Definitions["Pet"].Value
	assert.Equal(t, "petType", schema.Discriminator)
}

// --- Enum ---

func TestParseSchema_Enum(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Status:
    type: string
    enum:
      - active
      - inactive
      - pending
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := doc.Definitions["Status"].Value
	require.Len(t, schema.Enum, 3)
	assert.Equal(t, "active", schema.Enum[0])
}

// --- XML ---

func TestParseSchema_XML(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Pet:
    type: object
    xml:
      name: pet
      namespace: "http://example.com"
      prefix: p
      wrapped: true
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := doc.Definitions["Pet"].Value
	require.NotNil(t, schema.XML)
	assert.Equal(t, "pet", schema.XML.Name)
	assert.Equal(t, "http://example.com", schema.XML.Namespace)
	assert.True(t, schema.XML.Wrapped)
}

// --- ReadOnly ---

func TestParseSchema_ReadOnly(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Pet:
    type: object
    properties:
      id:
        type: integer
        readOnly: true
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := doc.Definitions["Pet"].Value.Properties["id"].Value
	assert.True(t, schema.ReadOnly)
}

// --- Example ---

func TestParseSchema_Example(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Name:
    type: string
    example: "John Doe"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schema := doc.Definitions["Name"].Value
	assert.Equal(t, "John Doe", schema.Example)
}

// --- $ref ---

func TestParseSchema_Ref(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
          schema:
            $ref: "#/definitions/Pet"
definitions:
  Pet:
    type: object
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	schemaRef := doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Schema
	assert.Equal(t, "#/definitions/Pet", schemaRef.Ref)
}
