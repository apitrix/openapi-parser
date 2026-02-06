package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for xml.go - parseXML
// =============================================================================

// --- Basic XML ---

func TestParseXML_Basic(t *testing.T) {
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
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	xml := doc.Definitions["Pet"].Value.XML
	require.NotNil(t, xml)
	assert.Equal(t, "pet", xml.Name)
}

// --- XML with Namespace ---

func TestParseXML_Namespace(t *testing.T) {
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
      namespace: "http://example.com/schema"
      prefix: ex
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	xml := doc.Definitions["Pet"].Value.XML
	assert.Equal(t, "http://example.com/schema", xml.Namespace)
	assert.Equal(t, "ex", xml.Prefix)
}

// --- XML Attribute ---

func TestParseXML_Attribute(t *testing.T) {
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
        xml:
          attribute: true
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	xml := doc.Definitions["Pet"].Value.Properties["id"].Value.XML
	require.NotNil(t, xml)
	assert.True(t, xml.Attribute)
}

// --- XML Wrapped ---

func TestParseXML_Wrapped(t *testing.T) {
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
      tags:
        type: array
        items:
          type: string
        xml:
          name: tag
          wrapped: true
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	xml := doc.Definitions["Pet"].Value.Properties["tags"].Value.XML
	require.NotNil(t, xml)
	assert.True(t, xml.Wrapped)
	assert.Equal(t, "tag", xml.Name)
}

// --- XML All Properties ---

func TestParseXML_AllProperties(t *testing.T) {
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
      prefix: ex
      attribute: false
      wrapped: true
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	xml := doc.Definitions["Pet"].Value.XML
	assert.Equal(t, "pet", xml.Name)
	assert.Equal(t, "http://example.com", xml.Namespace)
	assert.Equal(t, "ex", xml.Prefix)
	assert.False(t, xml.Attribute)
	assert.True(t, xml.Wrapped)
}

// --- XML Extensions ---

func TestParseXML_Extensions(t *testing.T) {
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
      x-custom: "value"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	xml := doc.Definitions["Pet"].Value.XML
	require.NotNil(t, xml.Extensions)
	assert.Equal(t, "value", xml.Extensions["x-custom"])
}
