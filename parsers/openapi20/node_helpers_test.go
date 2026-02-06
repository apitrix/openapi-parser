package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// =============================================================================
// Tests for node_helpers.go - yaml.Node helper functions
// =============================================================================

// --- nodeGetValue ---

func TestNodeGetValue_Exists(t *testing.T) {
	// Arrange
	yamlContent := `key: value`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeGetValue(docNode, "key")

	// Assert
	require.NotNil(t, result)
	assert.Equal(t, "value", result.Value)
}

func TestNodeGetValue_NotExists(t *testing.T) {
	// Arrange
	yamlContent := `key: value`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeGetValue(docNode, "notexist")

	// Assert
	assert.Nil(t, result)
}

// --- nodeGetString ---

func TestNodeGetString_Exists(t *testing.T) {
	// Arrange
	yamlContent := `name: "John"`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeGetString(docNode, "name")

	// Assert
	assert.Equal(t, "John", result)
}

func TestNodeGetString_NotExists(t *testing.T) {
	// Arrange
	yamlContent := `other: value`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeGetString(docNode, "name")

	// Assert
	assert.Empty(t, result)
}

// --- nodeGetBool ---

func TestNodeGetBool_True(t *testing.T) {
	// Arrange
	yamlContent := `flag: true`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeGetBool(docNode, "flag")

	// Assert
	assert.True(t, result)
}

func TestNodeGetBool_False(t *testing.T) {
	// Arrange
	yamlContent := `flag: false`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeGetBool(docNode, "flag")

	// Assert
	assert.False(t, result)
}

// --- nodeGetStringSlice ---

func TestNodeGetStringSlice_Exists(t *testing.T) {
	// Arrange
	yamlContent := `tags:
  - a
  - b
  - c`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeGetStringSlice(docNode, "tags")

	// Assert
	assert.Equal(t, []string{"a", "b", "c"}, result)
}

func TestNodeGetStringSlice_Empty(t *testing.T) {
	// Arrange
	yamlContent := `tags: []`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeGetStringSlice(docNode, "tags")

	// Assert
	assert.Empty(t, result)
}

// --- nodeGetFloat64Ptr ---

func TestNodeGetFloat64Ptr_Exists(t *testing.T) {
	// Arrange
	yamlContent := `value: 3.14`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeGetFloat64Ptr(docNode, "value")

	// Assert
	require.NotNil(t, result)
	assert.Equal(t, 3.14, *result)
}

func TestNodeGetFloat64Ptr_NotExists(t *testing.T) {
	// Arrange
	yamlContent := `other: value`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeGetFloat64Ptr(docNode, "value")

	// Assert
	assert.Nil(t, result)
}

// --- nodeGetUint64Ptr ---

func TestNodeGetUint64Ptr_Exists(t *testing.T) {
	// Arrange
	yamlContent := `value: 100`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeGetUint64Ptr(docNode, "value")

	// Assert
	require.NotNil(t, result)
	assert.Equal(t, uint64(100), *result)
}

// --- nodeIsMapping ---

func TestNodeIsMapping_True(t *testing.T) {
	// Arrange
	yamlContent := `key: value`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeIsMapping(docNode)

	// Assert
	assert.True(t, result)
}

func TestNodeIsMapping_False(t *testing.T) {
	// Arrange
	yamlContent := `- a
- b`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeIsMapping(docNode)

	// Assert
	assert.False(t, result)
}

// --- nodeIsSequence ---

func TestNodeIsSequence_True(t *testing.T) {
	// Arrange
	yamlContent := `- a
- b`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeIsSequence(docNode)

	// Assert
	assert.True(t, result)
}

// --- nodeKeys ---

func TestNodeKeys(t *testing.T) {
	// Arrange
	yamlContent := `a: 1
b: 2
c: 3`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeKeys(docNode)

	// Assert
	assert.Equal(t, []string{"a", "b", "c"}, result)
}

// --- nodeHasRef ---

func TestNodeHasRef_True(t *testing.T) {
	// Arrange
	yamlContent := `$ref: "#/definitions/Pet"`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeHasRef(docNode)

	// Assert
	assert.True(t, result)
}

func TestNodeHasRef_False(t *testing.T) {
	// Arrange
	yamlContent := `type: string`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeHasRef(docNode)

	// Assert
	assert.False(t, result)
}

// --- nodeGetRef ---

func TestNodeGetRef(t *testing.T) {
	// Arrange
	yamlContent := `$ref: "#/definitions/Pet"`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := nodeGetRef(docNode)

	// Assert
	assert.Equal(t, "#/definitions/Pet", result)
}

// --- parseNodeExtensions ---

func TestParseNodeExtensions(t *testing.T) {
	// Arrange
	yamlContent := `type: string
x-custom: value
x-flag: true`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	// Act
	result := parseNodeExtensions(docNode)

	// Assert
	assert.Len(t, result, 2)
	assert.Equal(t, "value", result["x-custom"])
	assert.Equal(t, true, result["x-flag"])
}

// --- nodeToInterface ---

func TestNodeToInterface_String(t *testing.T) {
	// Arrange
	node := &yaml.Node{Kind: yaml.ScalarNode, Value: "test", Tag: "!!str"}

	// Act
	result := nodeToInterface(node)

	// Assert
	assert.Equal(t, "test", result)
}

func TestNodeToInterface_Int(t *testing.T) {
	// Arrange
	node := &yaml.Node{Kind: yaml.ScalarNode, Value: "42", Tag: "!!int"}

	// Act
	result := nodeToInterface(node)

	// Assert
	assert.Equal(t, 42, result)
}

func TestNodeToInterface_Bool(t *testing.T) {
	// Arrange
	node := &yaml.Node{Kind: yaml.ScalarNode, Value: "true", Tag: "!!bool"}

	// Act
	result := nodeToInterface(node)

	// Assert
	assert.Equal(t, true, result)
}
