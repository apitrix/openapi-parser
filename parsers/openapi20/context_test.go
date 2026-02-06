package openapi20

import (
	"testing"

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
	ctx := newParseContext(node)

	// Assert
	assert.NotNil(t, ctx)
	assert.Equal(t, node, ctx.Root)
	assert.Empty(t, ctx.Path)
}

// --- Push ---

func TestParseContext_Push(t *testing.T) {
	// Arrange
	node := &yaml.Node{}
	ctx := newParseContext(node)

	// Act
	child := ctx.push("paths")
	grandchild := child.push("/pets")

	// Assert
	assert.Equal(t, []string{"paths"}, child.Path)
	assert.Equal(t, []string{"paths", "/pets"}, grandchild.Path)
}

// --- CurrentPath ---

func TestParseContext_CurrentPath(t *testing.T) {
	// Arrange
	node := &yaml.Node{}
	ctx := newParseContext(node)
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
	ctx := newParseContext(node)
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
	ctx := newParseContext(node)
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

	ctx := newParseContext(docNode)
	ctx.unknownFields = &[]UnknownField{}

	// Act
	ctx.detectUnknown(docNode, map[string]struct{}{"known": {}})

	// Assert
	assert.Len(t, *ctx.unknownFields, 1)
	assert.Equal(t, "unknown", (*ctx.unknownFields)[0].Name)
}

// --- UnknownFieldsResult ---

func TestParseContext_UnknownFieldsResult(t *testing.T) {
	// Arrange
	fields := []UnknownField{{Name: "test", Path: "root"}}
	ctx := &ParseContext{unknownFields: &fields}

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
	ctx := newParseContext(node)

	// Act
	source := ctx.nodeSource(node)

	// Assert
	assert.NotNil(t, source)
	assert.Equal(t, 10, source.Start.Line)
	assert.Equal(t, 5, source.Start.Column)
}
