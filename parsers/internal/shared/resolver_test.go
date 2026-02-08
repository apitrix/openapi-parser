package shared

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// =============================================================================
// SplitRef tests
// =============================================================================

func TestSplitRef_LocalPointer(t *testing.T) {
	file, pointer := SplitRef("#/components/schemas/Pet")
	assert.Equal(t, "", file)
	assert.Equal(t, "/components/schemas/Pet", pointer)
}

func TestSplitRef_FileOnly(t *testing.T) {
	file, pointer := SplitRef("./schemas/pet.yaml")
	assert.Equal(t, "./schemas/pet.yaml", file)
	assert.Equal(t, "", pointer)
}

func TestSplitRef_FileWithPointer(t *testing.T) {
	file, pointer := SplitRef("./common.yaml#/definitions/Error")
	assert.Equal(t, "./common.yaml", file)
	assert.Equal(t, "/definitions/Error", pointer)
}

func TestSplitRef_BareName(t *testing.T) {
	file, pointer := SplitRef("Pet")
	assert.Equal(t, "Pet", file)
	assert.Equal(t, "", pointer)
}

func TestSplitRef_EmptyRef(t *testing.T) {
	file, pointer := SplitRef("")
	assert.Equal(t, "", file)
	assert.Equal(t, "", pointer)
}

func TestIsExternalRef(t *testing.T) {
	assert.True(t, IsExternalRef("./schemas/pet.yaml"))
	assert.True(t, IsExternalRef("./common.yaml#/definitions/Error"))
	assert.False(t, IsExternalRef("#/components/schemas/Pet"))
	assert.False(t, IsExternalRef(""))
}

func TestIsLocalRef(t *testing.T) {
	assert.True(t, IsLocalRef("#/components/schemas/Pet"))
	assert.False(t, IsLocalRef("./schemas/pet.yaml"))
	assert.False(t, IsLocalRef(""))
}

// =============================================================================
// ResolveJSONPointer tests
// =============================================================================

func parseYAML(t *testing.T, data string) *yaml.Node {
	t.Helper()
	var node yaml.Node
	require.NoError(t, yaml.Unmarshal([]byte(data), &node))
	return &node
}

func TestResolveJSONPointer_Simple(t *testing.T) {
	root := parseYAML(t, `
components:
  schemas:
    Pet:
      type: object
`)
	node, err := ResolveJSONPointer(root, "/components/schemas/Pet")
	require.NoError(t, err)
	require.NotNil(t, node)
	assert.Equal(t, yaml.MappingNode, node.Kind)
}

func TestResolveJSONPointer_Nested(t *testing.T) {
	root := parseYAML(t, `
a:
  b:
    c: hello
`)
	node, err := ResolveJSONPointer(root, "/a/b/c")
	require.NoError(t, err)
	require.NotNil(t, node)
	assert.Equal(t, "hello", node.Value)
}

func TestResolveJSONPointer_Root(t *testing.T) {
	root := parseYAML(t, `type: object`)
	node, err := ResolveJSONPointer(root, "")
	require.NoError(t, err)
	require.NotNil(t, node)
}

func TestResolveJSONPointer_NotFound(t *testing.T) {
	root := parseYAML(t, `a: 1`)
	_, err := ResolveJSONPointer(root, "/b")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestResolveJSONPointer_EscapedTilde(t *testing.T) {
	root := parseYAML(t, `
"a/b":
  "c~d": value
`)
	// ~1 = /    ~0 = ~
	node, err := ResolveJSONPointer(root, "/a~1b/c~0d")
	require.NoError(t, err)
	require.NotNil(t, node)
	assert.Equal(t, "value", node.Value)
}

func TestResolveJSONPointer_SequenceIndex(t *testing.T) {
	root := parseYAML(t, `
items:
  - first
  - second
  - third
`)
	node, err := ResolveJSONPointer(root, "/items/1")
	require.NoError(t, err)
	assert.Equal(t, "second", node.Value)
}

func TestResolveJSONPointer_SequenceOutOfRange(t *testing.T) {
	root := parseYAML(t, `
items:
  - first
`)
	_, err := ResolveJSONPointer(root, "/items/5")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "out of range")
}

// =============================================================================
// RefResolver tests with afero in-memory filesystem
// =============================================================================

func TestRefResolver_LocalRef(t *testing.T) {
	root := parseYAML(t, `
components:
  schemas:
    Pet:
      type: object
      properties:
        name:
          type: string
`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())
	result, err := resolver.Resolve("#/components/schemas/Pet")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Circular)
	assert.Equal(t, yaml.MappingNode, result.Node.Kind)
}

func TestRefResolver_ExternalFileRef(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/base/pet.yaml", []byte(`type: object
properties:
  name:
    type: string
`), 0644)

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs("/base", root, fs)

	result, err := resolver.Resolve("pet.yaml")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Circular)
	assert.Equal(t, yaml.MappingNode, result.Node.Kind)
}

func TestRefResolver_ExternalFileWithPointer(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/base/schemas/common.yaml", []byte(`definitions:
  Error:
    type: object
    properties:
      message:
        type: string
`), 0644)

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs("/base", root, fs)

	result, err := resolver.Resolve("schemas/common.yaml#/definitions/Error")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, yaml.MappingNode, result.Node.Kind)
}

func TestRefResolver_FileCaching(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/base/schema.yaml", []byte("type: string"), 0644)

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs("/base", root, fs)

	// Resolve twice — second should use cache
	r1, err := resolver.Resolve("schema.yaml")
	require.NoError(t, err)
	r2, err := resolver.Resolve("schema.yaml")
	require.NoError(t, err)

	// Same node pointer = cache hit
	assert.Equal(t, r1.Node, r2.Node)
}

func TestRefResolver_CircularDetection(t *testing.T) {
	root := parseYAML(t, `
definitions:
  Node:
    type: object
    properties:
      children:
        type: array
`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())

	// Manually mark as visiting to simulate recursion
	canonical := resolver.canonicalize("#/definitions/Node")
	resolver.visiting[canonical] = true

	result, err := resolver.Resolve("#/definitions/Node")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.Circular, "should detect circular reference")
	assert.Nil(t, result.Node)
}

func TestRefResolver_MissingFile(t *testing.T) {
	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())

	_, err := resolver.Resolve("nonexistent.yaml")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to resolve external ref")
}

func TestRefResolver_MissingPointer(t *testing.T) {
	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())

	_, err := resolver.Resolve("#/nonexistent/path")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}
