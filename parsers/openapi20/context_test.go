package openapi20

import (
	"testing"

	"openapi-parser/parsers/shared"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

// =============================================================================
// Tests for context.go - ParseContext
// =============================================================================

// --- newParseContext ---

func TestNewParseContext(t *testing.T) {
	// Arrange
	node := &yaml.Node{}

	// Act
	ctx := newParseContext(node, shared.All())

	// Assert
	assert.NotNil(t, ctx)
	assert.Equal(t, node, ctx.Root)
	assert.Empty(t, ctx.path())
}

// --- Push ---

func TestParseContext_Push(t *testing.T) {
	// Arrange
	node := &yaml.Node{}
	ctx := newParseContext(node, shared.All())

	// Act
	child := ctx.push("paths")
	grandchild := child.push("/pets")

	// Assert
	assert.Equal(t, []string{"paths"}, child.path())
	assert.Equal(t, []string{"paths", "/pets"}, grandchild.path())
}

// --- CurrentPath ---

func TestParseContext_CurrentPath(t *testing.T) {
	// Arrange
	node := &yaml.Node{}
	ctx := newParseContext(node, shared.All())
	ctx = ctx.push("paths")
	ctx = ctx.push("/pets")
	ctx = ctx.push("get")

	// Act
	path := ctx.CurrentPath()

	// Assert
	assert.Equal(t, "paths./pets.get", path)
}

// --- errorf ---

func TestParseContext_Errorf(t *testing.T) {
	// Arrange
	node := &yaml.Node{}
	ctx := newParseContext(node, shared.All())
	ctx = ctx.push("info")

	// Act
	err := ctx.errorf("missing required field: %s", "title")

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "missing required field: title")
	assert.Contains(t, err.Error(), "info")
}

// --- errorAt ---

func TestParseContext_ErrorAt(t *testing.T) {
	// Arrange
	node := &yaml.Node{Line: 5, Column: 10}
	ctx := newParseContext(node, shared.All())
	ctx = ctx.push("info")

	// Act
	err := ctx.errorAt(node, "invalid value")

	// Assert
	assert.Error(t, err)
	parseErr, ok := err.(*ParseError)
	assert.True(t, ok)
	assert.Equal(t, 5, parseErr.Line)
	assert.Equal(t, 10, parseErr.Column)
}

// --- detectUnknown ---

func TestParseContext_DetectUnknown(t *testing.T) {
	// Arrange
	yamlContent := `known: value
unknown: value
x-extension: value
`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	ctx := newParseContext(docNode, shared.All())

	// Act
	ctx.detectUnknown(docNode, map[string]struct{}{"known": {}})

	// Assert
	assert.Len(t, *ctx.unknownFields, 1)
	assert.Equal(t, "unknown", (*ctx.unknownFields)[0].Name)
}

func TestParseContext_DetectUnknown_Disabled(t *testing.T) {
	// Arrange
	yamlContent := `known: value
unknown: value
`
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yamlContent), &node)
	docNode := node.Content[0]

	ctx := newParseContext(docNode, shared.None())

	// Act
	result := ctx.detectUnknown(docNode, map[string]struct{}{"known": {}})

	// Assert - no unknown fields detected when disabled
	assert.Nil(t, result)
	assert.Empty(t, *ctx.unknownFields)
}

// --- UnknownFieldsResult ---

func TestParseContext_UnknownFieldsResult(t *testing.T) {
	// Arrange
	fields := []UnknownField{{Name: "test", Path: "root"}}
	ctx := &ParseContext{
		unknownFields: &fields,
		config:        shared.All(),
	}

	// Act
	result := ctx.UnknownFieldsResult()

	// Assert
	assert.Equal(t, fields, result)
}

// --- nodeSource ---

func TestParseContext_NodeSource(t *testing.T) {
	// Arrange
	node := &yaml.Node{
		Kind:   yaml.MappingNode,
		Line:   10,
		Column: 5,
	}
	ctx := newParseContext(node, shared.All())

	// Act
	source := ctx.nodeSource(node)

	// Assert
	assert.NotNil(t, source)
	assert.Equal(t, 10, source.Start.Line)
	assert.Equal(t, 5, source.Start.Column)
}
