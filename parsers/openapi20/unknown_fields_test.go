package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// =============================================================================
// Tests for unknown_fields.go - unknown field detection
// =============================================================================

// --- detectUnknownNodeFields ---

func TestDetectUnknownNodeFields_NoUnknown(t *testing.T) {
	// Arrange
	yamlContent := `name: "Test"
version: "1.0"`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	knownFields := map[string]struct{}{"name": {}, "version": {}}

	// Act
	result := detectUnknownNodeFields(docNode, knownFields, "info")

	// Assert
	assert.Empty(t, result)
}

func TestDetectUnknownNodeFields_WithUnknown(t *testing.T) {
	// Arrange
	yamlContent := `name: "Test"
unknownField: "value"
version: "1.0"`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	knownFields := map[string]struct{}{"name": {}, "version": {}}

	// Act
	result := detectUnknownNodeFields(docNode, knownFields, "info")

	// Assert
	require.Len(t, result, 1)
	assert.Equal(t, "unknownField", result[0].Name)
	assert.Equal(t, "info", result[0].Path)
}

func TestDetectUnknownNodeFields_IgnoresExtensions(t *testing.T) {
	// Arrange
	yamlContent := `name: "Test"
x-custom: "value"
x-another: true`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	knownFields := map[string]struct{}{"name": {}}

	// Act
	result := detectUnknownNodeFields(docNode, knownFields, "info")

	// Assert
	assert.Empty(t, result)
}

func TestDetectUnknownNodeFields_MultipleUnknown(t *testing.T) {
	// Arrange
	yamlContent := `name: "Test"
unknown1: "a"
version: "1.0"
unknown2: "b"`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	knownFields := map[string]struct{}{"name": {}, "version": {}}

	// Act
	result := detectUnknownNodeFields(docNode, knownFields, "info")

	// Assert
	require.Len(t, result, 2)
	assert.Equal(t, "unknown1", result[0].Name)
	assert.Equal(t, "unknown2", result[1].Name)
}

func TestDetectUnknownNodeFields_WithLocation(t *testing.T) {
	// Arrange
	yamlContent := `name: "Test"
unknownField: "value"`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	knownFields := map[string]struct{}{"name": {}}

	// Act
	result := detectUnknownNodeFields(docNode, knownFields, "root")

	// Assert
	require.Len(t, result, 1)
	assert.Greater(t, result[0].Line, 0)
}

// --- UnknownField struct ---

func TestUnknownField_Struct(t *testing.T) {
	// Arrange & Act
	uf := UnknownField{
		Name:   "testField",
		Path:   "info.contact",
		Line:   10,
		Column: 5,
	}

	// Assert
	assert.Equal(t, "testField", uf.Name)
	assert.Equal(t, "info.contact", uf.Path)
	assert.Equal(t, 10, uf.Line)
	assert.Equal(t, 5, uf.Column)
}

// --- Integration with Parse ---

func TestUnknownFields_Integration(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
  unknownInfo: "ignored"
paths:
  /pets:
    unknownPath: "ignored"
    get:
      unknownOp: "ignored"
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.NotEmpty(t, result.UnknownFields)

	// Check that we detected all unknown fields
	names := make([]string, len(result.UnknownFields))
	for i, f := range result.UnknownFields {
		names[i] = f.Name
	}
	assert.Contains(t, names, "unknownInfo")
	assert.Contains(t, names, "unknownPath")
	assert.Contains(t, names, "unknownOp")
}
