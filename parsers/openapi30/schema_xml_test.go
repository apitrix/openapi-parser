package openapi30

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_xml.go - ParseXML method
// =============================================================================

// --- Basic XML ---

func TestParseSchemaXML_Basic(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      xml:
        name: pet
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	xml := doc.Components.Schemas["Pet"].Value.XML
	require.NotNil(t, xml)
	assert.Equal(t, "pet", xml.Name)
}

// --- Namespace ---

func TestParseSchemaXML_Namespace(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      xml:
        name: pet
        namespace: "http://example.com/pets"
        prefix: "pet"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	xml := doc.Components.Schemas["Pet"].Value.XML
	assert.Equal(t, "http://example.com/pets", xml.Namespace)
	assert.Equal(t, "pet", xml.Prefix)
}

// --- Attribute ---

func TestParseSchemaXML_Attribute(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      properties:
        id:
          type: integer
          xml:
            attribute: true
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	xml := doc.Components.Schemas["Pet"].Value.Properties["id"].Value.XML
	require.NotNil(t, xml)
	assert.True(t, xml.Attribute)
}

// --- Wrapped ---

func TestParseSchemaXML_Wrapped(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      properties:
        tags:
          type: array
          xml:
            wrapped: true
            name: tags
          items:
            type: string
            xml:
              name: tag
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	xml := doc.Components.Schemas["Pet"].Value.Properties["tags"].Value.XML
	require.NotNil(t, xml)
	assert.True(t, xml.Wrapped)
}

// --- Complete XML ---

func TestParseSchemaXML_Complete(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      xml:
        name: pet
        namespace: "http://example.com"
        prefix: "p"
        attribute: false
        wrapped: false
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	xml := doc.Components.Schemas["Pet"].Value.XML
	assert.Equal(t, "pet", xml.Name)
	assert.Equal(t, "http://example.com", xml.Namespace)
	assert.Equal(t, "p", xml.Prefix)
}

// --- Missing XML ---

func TestParseSchemaXML_Missing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	xml := doc.Components.Schemas["Pet"].Value.XML
	assert.Nil(t, xml)
}

// --- Extensions ---

func TestParseSchemaXML_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      xml:
        name: pet
        x-custom: "value"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	xml := doc.Components.Schemas["Pet"].Value.XML
	require.NotNil(t, xml.Extensions)
	assert.Equal(t, "value", xml.Extensions["x-custom"])
}
