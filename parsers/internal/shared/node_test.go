package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gopkg.in/yaml.v3"
)

// mkMapping builds a yaml MappingNode from alternating key, value strings.
func mkMapping(pairs ...string) *yaml.Node {
	node := &yaml.Node{Kind: yaml.MappingNode}
	for i := 0; i+1 < len(pairs); i += 2 {
		node.Content = append(node.Content,
			&yaml.Node{Kind: yaml.ScalarNode, Value: pairs[i]},
			&yaml.Node{Kind: yaml.ScalarNode, Value: pairs[i+1]},
		)
	}
	return node
}

func mkScalar(v string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Value: v}
}

func mkSequence(items ...string) *yaml.Node {
	node := &yaml.Node{Kind: yaml.SequenceNode}
	for _, s := range items {
		node.Content = append(node.Content, mkScalar(s))
	}
	return node
}

// --- NodeGetValue ---

func TestNodeGetValue(t *testing.T) {
	t.Run("key found", func(t *testing.T) {
		node := mkMapping("name", "Fido", "type", "object")
		got := NodeGetValue(node, "type")
		require.NotNil(t, got)
		assert.Equal(t, "object", got.Value)
	})
	t.Run("key not found", func(t *testing.T) {
		node := mkMapping("name", "Fido")
		assert.Nil(t, NodeGetValue(node, "missing"))
	})
	t.Run("nil node", func(t *testing.T) {
		assert.Nil(t, NodeGetValue(nil, "key"))
	})
	t.Run("non-mapping node", func(t *testing.T) {
		assert.Nil(t, NodeGetValue(mkScalar("hello"), "key"))
	})
}

// --- NodeGetKeyNode ---

func TestNodeGetKeyNode(t *testing.T) {
	t.Run("key found", func(t *testing.T) {
		node := mkMapping("name", "Fido")
		got := NodeGetKeyNode(node, "name")
		require.NotNil(t, got)
		assert.Equal(t, "name", got.Value)
	})
	t.Run("key not found", func(t *testing.T) {
		node := mkMapping("name", "Fido")
		assert.Nil(t, NodeGetKeyNode(node, "missing"))
	})
}

// --- NodeToMap ---

func TestNodeToMap(t *testing.T) {
	t.Run("normal mapping", func(t *testing.T) {
		node := mkMapping("a", "1", "b", "2")
		m := NodeToMap(node)
		assert.Len(t, m, 2)
		assert.Equal(t, "1", m["a"].Value)
		assert.Equal(t, "2", m["b"].Value)
	})
	t.Run("nil node", func(t *testing.T) {
		assert.Nil(t, NodeToMap(nil))
	})
	t.Run("non-mapping", func(t *testing.T) {
		assert.Nil(t, NodeToMap(mkScalar("hello")))
	})
}

// --- NodeKeys ---

func TestNodeKeys(t *testing.T) {
	node := mkMapping("x", "1", "y", "2", "z", "3")
	keys := NodeKeys(node)
	assert.ElementsMatch(t, []string{"x", "y", "z"}, keys)
}

func TestNodeKeys_NilNode(t *testing.T) {
	assert.Nil(t, NodeKeys(nil))
}

// --- NodeMapPairs ---

func TestNodeMapPairs(t *testing.T) {
	node := mkMapping("a", "1", "b", "2")
	collected := map[string]string{}
	for k, v := range NodeMapPairs(node) {
		collected[k] = v.Value
	}
	assert.Equal(t, map[string]string{"a": "1", "b": "2"}, collected)
}

// --- NodeToSlice ---

func TestNodeToSlice(t *testing.T) {
	node := mkSequence("alpha", "beta")
	got := NodeToSlice(node)
	require.Len(t, got, 2)
	assert.Equal(t, "alpha", got[0].Value)
	assert.Equal(t, "beta", got[1].Value)
}

func TestNodeToSlice_NilNode(t *testing.T) {
	assert.Nil(t, NodeToSlice(nil))
}

// --- NodeGetString ---

func TestNodeGetString(t *testing.T) {
	node := mkMapping("title", "My API")
	assert.Equal(t, "My API", NodeGetString(node, "title"))
	assert.Equal(t, "", NodeGetString(node, "missing"))
}

// --- NodeGetBool ---

func TestNodeGetBool(t *testing.T) {
	node := mkMapping("required", "true", "deprecated", "false")
	assert.True(t, NodeGetBool(node, "required"))
	assert.False(t, NodeGetBool(node, "deprecated"))
	assert.False(t, NodeGetBool(node, "missing"))
}

// --- NodeGetBoolPtr ---

func TestNodeGetBoolPtr(t *testing.T) {
	node := mkMapping("nullable", "true")
	got := NodeGetBoolPtr(node, "nullable")
	require.NotNil(t, got)
	assert.True(t, *got)

	assert.Nil(t, NodeGetBoolPtr(node, "missing"))
}

// --- NodeGetInt ---

func TestNodeGetInt(t *testing.T) {
	node := mkMapping("maxItems", "10")
	assert.Equal(t, 10, NodeGetInt(node, "maxItems"))
	assert.Equal(t, 0, NodeGetInt(node, "missing"))
}

// --- NodeGetFloat64 ---

func TestNodeGetFloat64(t *testing.T) {
	node := mkMapping("minimum", "3.14")
	assert.InDelta(t, 3.14, NodeGetFloat64(node, "minimum"), 0.001)
	assert.InDelta(t, 0.0, NodeGetFloat64(node, "missing"), 0.001)
}

// --- NodeGetFloat64Ptr ---

func TestNodeGetFloat64Ptr(t *testing.T) {
	node := mkMapping("maximum", "99.5")
	got := NodeGetFloat64Ptr(node, "maximum")
	require.NotNil(t, got)
	assert.InDelta(t, 99.5, *got, 0.001)

	assert.Nil(t, NodeGetFloat64Ptr(node, "missing"))
}

// --- NodeGetIntPtr ---

func TestNodeGetIntPtr(t *testing.T) {
	node := mkMapping("minLength", "5")
	got := NodeGetIntPtr(node, "minLength")
	require.NotNil(t, got)
	assert.Equal(t, 5, *got)

	assert.Nil(t, NodeGetIntPtr(node, "missing"))
}

// --- NodeGetUint64Ptr ---

func TestNodeGetUint64Ptr(t *testing.T) {
	node := mkMapping("maxLength", "255")
	got := NodeGetUint64Ptr(node, "maxLength")
	require.NotNil(t, got)
	assert.Equal(t, uint64(255), *got)

	assert.Nil(t, NodeGetUint64Ptr(node, "missing"))
}

// --- NodeGetStringSlice ---

func TestNodeGetStringSlice(t *testing.T) {
	// Build node with a sequence value
	seqNode := mkSequence("read", "write")
	node := &yaml.Node{
		Kind: yaml.MappingNode,
		Content: []*yaml.Node{
			mkScalar("scopes"), seqNode,
		},
	}
	got := NodeGetStringSlice(node, "scopes")
	assert.Equal(t, []string{"read", "write"}, got)
	assert.Nil(t, NodeGetStringSlice(node, "missing"))
}

// --- NodeGetStringMap ---

func TestNodeGetStringMap(t *testing.T) {
	inner := mkMapping("en", "English", "fr", "French")
	node := &yaml.Node{
		Kind: yaml.MappingNode,
		Content: []*yaml.Node{
			mkScalar("langs"), inner,
		},
	}
	got := NodeGetStringMap(node, "langs")
	assert.Equal(t, map[string]string{"en": "English", "fr": "French"}, got)
	assert.Nil(t, NodeGetStringMap(node, "missing"))
}

// --- NodeGetAny ---

func TestNodeGetAny(t *testing.T) {
	node := mkMapping("example", "hello")
	got := NodeGetAny(node, "example")
	assert.Equal(t, "hello", got)
	assert.Nil(t, NodeGetAny(node, "missing"))
}

// --- NodeToInterface ---

func TestNodeToInterface_Scalar(t *testing.T) {
	got := NodeToInterface(mkScalar("hello"))
	assert.Equal(t, "hello", got)
}

func TestNodeToInterface_Mapping(t *testing.T) {
	node := mkMapping("a", "hello")
	got := NodeToInterface(node)
	m, ok := got.(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, "hello", m["a"])
}

func TestNodeToInterface_Sequence(t *testing.T) {
	node := mkSequence("x", "y")
	got := NodeToInterface(node)
	s, ok := got.([]interface{})
	require.True(t, ok)
	assert.Len(t, s, 2)
}

func TestNodeToInterface_Nil(t *testing.T) {
	assert.Nil(t, NodeToInterface(nil))
}

// --- NodeIsMapping / NodeIsSequence / NodeIsScalar ---

func TestNodeIsMapping(t *testing.T) {
	assert.True(t, NodeIsMapping(mkMapping("k", "v")))
	assert.False(t, NodeIsMapping(mkScalar("v")))
	assert.False(t, NodeIsMapping(nil))
}

func TestNodeIsSequence(t *testing.T) {
	assert.True(t, NodeIsSequence(mkSequence("a")))
	assert.False(t, NodeIsSequence(mkScalar("v")))
}

func TestNodeIsScalar(t *testing.T) {
	assert.True(t, NodeIsScalar(mkScalar("v")))
	assert.False(t, NodeIsScalar(mkMapping("k", "v")))
}

// --- ParseNodeExtensions ---

func TestParseNodeExtensions(t *testing.T) {
	t.Run("has extensions", func(t *testing.T) {
		node := mkMapping("type", "object", "x-custom", "hello", "x-internal", "true")
		got := ParseNodeExtensions(node)
		assert.Len(t, got, 2)
		assert.Equal(t, "hello", got["x-custom"])
	})
	t.Run("no extensions", func(t *testing.T) {
		node := mkMapping("type", "string")
		assert.Nil(t, ParseNodeExtensions(node))
	})
	t.Run("nil node", func(t *testing.T) {
		assert.Nil(t, ParseNodeExtensions(nil))
	})
}

// --- NodeHasKey ---

func TestNodeHasKey(t *testing.T) {
	node := mkMapping("name", "test")
	assert.True(t, NodeHasKey(node, "name"))
	assert.False(t, NodeHasKey(node, "absent"))
}

// --- NodeHasRef / NodeGetRef ---

func TestNodeHasRef(t *testing.T) {
	ref := mkMapping("$ref", "#/components/schemas/Pet")
	assert.True(t, NodeHasRef(ref))

	noRef := mkMapping("type", "object")
	assert.False(t, NodeHasRef(noRef))
}

func TestNodeGetRef(t *testing.T) {
	node := mkMapping("$ref", "#/defs/Pet")
	assert.Equal(t, "#/defs/Pet", NodeGetRef(node))
	assert.Equal(t, "", NodeGetRef(mkMapping("type", "object")))
}
