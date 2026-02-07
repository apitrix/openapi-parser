package shared

import (
	"iter"
	"strconv"

	"gopkg.in/yaml.v3"
)

// NodeGetValue gets a child node by key from a mapping node.
// Returns nil if not found or if node is not a mapping.
func NodeGetValue(node *yaml.Node, key string) *yaml.Node {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}
	// MappingNode Content is [key1, val1, key2, val2, ...]
	for i := 0; i < len(node.Content)-1; i += 2 {
		if node.Content[i].Value == key {
			return node.Content[i+1]
		}
	}
	return nil
}

// NodeGetKeyNode gets the key node itself (for position info).
func NodeGetKeyNode(node *yaml.Node, key string) *yaml.Node {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}
	for i := 0; i < len(node.Content)-1; i += 2 {
		if node.Content[i].Value == key {
			return node.Content[i]
		}
	}
	return nil
}

// NodeToMap converts a mapping node to iterate over key-value pairs.
// Returns key string, value node pairs.
func NodeToMap(node *yaml.Node) map[string]*yaml.Node {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}
	result := make(map[string]*yaml.Node)
	for i := 0; i < len(node.Content)-1; i += 2 {
		key := node.Content[i].Value
		result[key] = node.Content[i+1]
	}
	return result
}

// NodeKeys returns all keys in a mapping node.
func NodeKeys(node *yaml.Node) []string {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}
	keys := make([]string, 0, len(node.Content)/2)
	for i := 0; i < len(node.Content)-1; i += 2 {
		keys = append(keys, node.Content[i].Value)
	}
	return keys
}

// NodeMapPairs iterates over key-value pairs in a mapping node (single-pass).
// Returns a range-over-func iterator for use with Go 1.23+ range syntax.
func NodeMapPairs(node *yaml.Node) iter.Seq2[string, *yaml.Node] {
	return func(yield func(string, *yaml.Node) bool) {
		if node == nil || node.Kind != yaml.MappingNode {
			return
		}
		for i := 0; i < len(node.Content)-1; i += 2 {
			if !yield(node.Content[i].Value, node.Content[i+1]) {
				return
			}
		}
	}
}

// NodeToSlice converts a sequence node to a slice of nodes.
func NodeToSlice(node *yaml.Node) []*yaml.Node {
	if node == nil || node.Kind != yaml.SequenceNode {
		return nil
	}
	return node.Content
}

// NodeGetString gets a string value from a mapping node by key.
func NodeGetString(node *yaml.Node, key string) string {
	valNode := NodeGetValue(node, key)
	if valNode == nil {
		return ""
	}
	return valNode.Value
}

// NodeGetBool gets a bool value from a mapping node by key.
func NodeGetBool(node *yaml.Node, key string) bool {
	valNode := NodeGetValue(node, key)
	if valNode == nil {
		return false
	}
	// Handle various bool representations
	switch valNode.Value {
	case "true", "True", "TRUE", "yes", "Yes", "YES", "on", "On", "ON":
		return true
	default:
		return false
	}
}

// NodeGetBoolPtr gets a *bool value from a mapping node by key.
func NodeGetBoolPtr(node *yaml.Node, key string) *bool {
	valNode := NodeGetValue(node, key)
	if valNode == nil {
		return nil
	}
	// Handle various bool representations
	var b bool
	switch valNode.Value {
	case "true", "True", "TRUE", "yes", "Yes", "YES", "on", "On", "ON":
		b = true
	default:
		b = false
	}
	return &b
}

// NodeGetInt gets an int value from a mapping node by key.
func NodeGetInt(node *yaml.Node, key string) int {
	valNode := NodeGetValue(node, key)
	if valNode == nil {
		return 0
	}
	i, _ := strconv.Atoi(valNode.Value)
	return i
}

// NodeGetFloat64 gets a float64 value from a mapping node by key.
func NodeGetFloat64(node *yaml.Node, key string) float64 {
	valNode := NodeGetValue(node, key)
	if valNode == nil {
		return 0
	}
	f, _ := strconv.ParseFloat(valNode.Value, 64)
	return f
}

// NodeGetFloat64Ptr gets a *float64 value from a mapping node by key.
func NodeGetFloat64Ptr(node *yaml.Node, key string) *float64 {
	valNode := NodeGetValue(node, key)
	if valNode == nil {
		return nil
	}
	f, err := strconv.ParseFloat(valNode.Value, 64)
	if err != nil {
		return nil
	}
	return &f
}

// NodeGetIntPtr gets a *int value from a mapping node by key.
func NodeGetIntPtr(node *yaml.Node, key string) *int {
	valNode := NodeGetValue(node, key)
	if valNode == nil {
		return nil
	}
	i, err := strconv.Atoi(valNode.Value)
	if err != nil {
		return nil
	}
	return &i
}

// NodeGetUint64Ptr gets a *uint64 value from a mapping node by key.
func NodeGetUint64Ptr(node *yaml.Node, key string) *uint64 {
	valNode := NodeGetValue(node, key)
	if valNode == nil {
		return nil
	}
	u, err := strconv.ParseUint(valNode.Value, 10, 64)
	if err != nil {
		return nil
	}
	return &u
}

// NodeGetStringSlice gets a string slice from a mapping node by key.
func NodeGetStringSlice(node *yaml.Node, key string) []string {
	valNode := NodeGetValue(node, key)
	if valNode == nil || valNode.Kind != yaml.SequenceNode {
		return nil
	}
	result := make([]string, 0, len(valNode.Content))
	for _, item := range valNode.Content {
		result = append(result, item.Value)
	}
	return result
}

// NodeGetStringMap gets a map[string]string from a mapping node by key.
func NodeGetStringMap(node *yaml.Node, key string) map[string]string {
	valNode := NodeGetValue(node, key)
	if valNode == nil || valNode.Kind != yaml.MappingNode {
		return nil
	}
	result := make(map[string]string)
	for i := 0; i < len(valNode.Content)-1; i += 2 {
		k := valNode.Content[i].Value
		v := valNode.Content[i+1].Value
		result[k] = v
	}
	return result
}

// NodeGetAny gets the raw interface{} value by decoding the node.
func NodeGetAny(node *yaml.Node, key string) interface{} {
	valNode := NodeGetValue(node, key)
	if valNode == nil {
		return nil
	}
	return NodeToInterface(valNode)
}

// NodeToInterface converts a yaml.Node to interface{} (for Raw storage).
func NodeToInterface(node *yaml.Node) interface{} {
	if node == nil {
		return nil
	}
	switch node.Kind {
	case yaml.ScalarNode:
		// Try to decode as proper type
		var v interface{}
		_ = node.Decode(&v)
		return v
	case yaml.SequenceNode:
		result := make([]interface{}, len(node.Content))
		for i, child := range node.Content {
			result[i] = NodeToInterface(child)
		}
		return result
	case yaml.MappingNode:
		result := make(map[string]interface{})
		for i := 0; i < len(node.Content)-1; i += 2 {
			key := node.Content[i].Value
			result[key] = NodeToInterface(node.Content[i+1])
		}
		return result
	case yaml.DocumentNode:
		if len(node.Content) > 0 {
			return NodeToInterface(node.Content[0])
		}
		return nil
	case yaml.AliasNode:
		return NodeToInterface(node.Alias)
	default:
		return nil
	}
}

// NodeIsMapping checks if a node is a mapping node.
func NodeIsMapping(node *yaml.Node) bool {
	return node != nil && node.Kind == yaml.MappingNode
}

// NodeIsSequence checks if a node is a sequence node.
func NodeIsSequence(node *yaml.Node) bool {
	return node != nil && node.Kind == yaml.SequenceNode
}

// NodeIsScalar checks if a node is a scalar node.
func NodeIsScalar(node *yaml.Node) bool {
	return node != nil && node.Kind == yaml.ScalarNode
}

// ParseNodeExtensions extracts extension fields (x-*) from a yaml.Node.
func ParseNodeExtensions(node *yaml.Node) map[string]interface{} {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}

	var extensions map[string]interface{}
	for i := 0; i < len(node.Content)-1; i += 2 {
		key := node.Content[i].Value
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			if extensions == nil {
				extensions = make(map[string]interface{})
			}
			extensions[key] = NodeToInterface(node.Content[i+1])
		}
	}
	return extensions
}

// NodeHasKey checks if a key exists in a mapping node.
func NodeHasKey(node *yaml.Node, key string) bool {
	return NodeGetValue(node, key) != nil
}

// NodeHasRef checks if the node contains a $ref key.
func NodeHasRef(node *yaml.Node) bool {
	return NodeHasKey(node, "$ref")
}

// NodeGetRef retrieves the $ref value from a node, returning empty string if not found.
func NodeGetRef(node *yaml.Node) string {
	return NodeGetString(node, "$ref")
}
