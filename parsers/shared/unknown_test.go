package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gopkg.in/yaml.v3"
)

// helper to parse YAML into a mapping node
func yamlMapping(t *testing.T, data string) *yaml.Node {
	t.Helper()
	var root yaml.Node
	require.NoError(t, yaml.Unmarshal([]byte(data), &root))
	if root.Kind == yaml.DocumentNode && len(root.Content) > 0 {
		return root.Content[0]
	}
	return &root
}

func TestDetectUnknownNodeFields_NoUnknown(t *testing.T) {
	// Arrange
	node := yamlMapping(t, "type: object\ndescription: A pet")
	known := ToSet([]string{"type", "description"})

	// Act
	result := DetectUnknownNodeFields(node, known, "schemas.Pet")

	// Assert
	assert.Empty(t, result)
}

func TestDetectUnknownNodeFields_WithUnknown(t *testing.T) {
	// Arrange
	node := yamlMapping(t, "type: object\nfoo: bar\ndescription: test")
	known := ToSet([]string{"type", "description"})

	// Act
	result := DetectUnknownNodeFields(node, known, "schemas.Pet")

	// Assert
	require.Len(t, result, 1)
	assert.Equal(t, "foo", result[0].Key)
	assert.Equal(t, "foo", result[0].Name)
	assert.Equal(t, "schemas.Pet", result[0].Path)
}

func TestDetectUnknownNodeFields_SkipsExtensions(t *testing.T) {
	// Arrange
	node := yamlMapping(t, "type: object\nx-custom: value\nx-internal: true")
	known := ToSet([]string{"type"})

	// Act & Assert
	assert.Empty(t, DetectUnknownNodeFields(node, known, "root"))
}

func TestDetectUnknownNodeFields_NilNode(t *testing.T) {
	assert.Nil(t, DetectUnknownNodeFields(nil, ToSet([]string{}), ""))
}

func TestDetectUnknownNodeFields_NonMappingNode(t *testing.T) {
	node := &yaml.Node{Kind: yaml.SequenceNode}
	assert.Nil(t, DetectUnknownNodeFields(node, ToSet([]string{}), ""))
}

func TestDetectUnknownNodeFields_LineAndColumn(t *testing.T) {
	// Arrange
	node := yamlMapping(t, "type: object\nunknown_field: value")
	known := ToSet([]string{"type"})

	// Act
	result := DetectUnknownNodeFields(node, known, "test")

	// Assert
	require.Len(t, result, 1)
	assert.Greater(t, result[0].Line, 0)
	assert.Greater(t, result[0].Column, 0)
}

func TestIsExtension(t *testing.T) {
	tests := []struct {
		key  string
		want bool
	}{
		{"x-custom", true},
		{"x-a", true},
		{"x-", false},
		{"x", false},
		{"type", false},
		{"", false},
		{"X-Custom", false},
	}
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			assert.Equal(t, tt.want, IsExtension(tt.key))
		})
	}
}

func TestFormatPath(t *testing.T) {
	tests := []struct {
		name     string
		segments []string
		want     string
	}{
		{"multiple segments", []string{"paths", "/pets", "get"}, "paths./pets.get"},
		{"single segment", []string{"info"}, "info"},
		{"empty", []string{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, FormatPath(tt.segments))
		})
	}
}

func TestUnknownFieldError_Error_NoFields(t *testing.T) {
	err := &UnknownFieldError{Fields: nil}
	assert.Equal(t, "no unknown fields", err.Error())
}

func TestUnknownFieldError_Error_SingleField(t *testing.T) {
	err := &UnknownFieldError{
		Fields: []UnknownField{{Path: "info", Key: "foo", Line: 3, Column: 1}},
	}
	got := err.Error()
	assert.Contains(t, got, "foo")
	assert.Contains(t, got, "info")
}

func TestUnknownFieldError_Error_MultipleFields(t *testing.T) {
	err := &UnknownFieldError{
		Fields: []UnknownField{
			{Path: "info", Key: "foo", Line: 3, Column: 1},
			{Path: "paths", Key: "bar", Line: 10, Column: 5},
		},
	}
	got := err.Error()
	assert.Contains(t, got, "multiple unknown fields")
	assert.Contains(t, got, "foo")
	assert.Contains(t, got, "bar")
}

func TestFormatUnknownFieldMessage(t *testing.T) {
	t.Run("with path", func(t *testing.T) {
		f := UnknownField{Path: "info", Key: "badField", Line: 5, Column: 3}
		got := FormatUnknownFieldMessage(f)
		assert.Contains(t, got, "badField")
		assert.Contains(t, got, "info")
		assert.Contains(t, got, "5")
	})
	t.Run("empty path uses root", func(t *testing.T) {
		f := UnknownField{Path: "", Key: "x", Line: 1, Column: 1}
		assert.Contains(t, FormatUnknownFieldMessage(f), "(root)")
	})
}
