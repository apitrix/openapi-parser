package openapi30

import (
	"strconv"

	"gopkg.in/yaml.v3"
)

// nodeGetValue gets a child node by key from a mapping node.
// Returns nil if not found or if node is not a mapping.
func nodeGetValue(node *yaml.Node, key string) *yaml.Node {
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

// nodeGetKeyNode gets the key node itself (for position info).
func nodeGetKeyNode(node *yaml.Node, key string) *yaml.Node {
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

// nodeToMap converts a mapping node to iterate over key-value pairs.
// Returns key string, value node pairs.
func nodeToMap(node *yaml.Node) map[string]*yaml.Node {
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

// nodeKeys returns all keys in a mapping node.
func nodeKeys(node *yaml.Node) []string {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}
	keys := make([]string, 0, len(node.Content)/2)
	for i := 0; i < len(node.Content)-1; i += 2 {
		keys = append(keys, node.Content[i].Value)
	}
	return keys
}

// nodeToSlice converts a sequence node to a slice of nodes.
func nodeToSlice(node *yaml.Node) []*yaml.Node {
	if node == nil || node.Kind != yaml.SequenceNode {
		return nil
	}
	return node.Content
}

// nodeGetString gets a string value from a mapping node by key.
func nodeGetString(node *yaml.Node, key string) string {
	valNode := nodeGetValue(node, key)
	if valNode == nil {
		return ""
	}
	return valNode.Value
}

// nodeGetBool gets a bool value from a mapping node by key.
func nodeGetBool(node *yaml.Node, key string) bool {
	valNode := nodeGetValue(node, key)
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

// nodeGetBoolPtr gets a *bool value from a mapping node by key.
func nodeGetBoolPtr(node *yaml.Node, key string) *bool {
	valNode := nodeGetValue(node, key)
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

// nodeGetInt gets an int value from a mapping node by key.
func nodeGetInt(node *yaml.Node, key string) int {
	valNode := nodeGetValue(node, key)
	if valNode == nil {
		return 0
	}
	i, _ := strconv.Atoi(valNode.Value)
	return i
}

// nodeGetFloat64 gets a float64 value from a mapping node by key.
func nodeGetFloat64(node *yaml.Node, key string) float64 {
	valNode := nodeGetValue(node, key)
	if valNode == nil {
		return 0
	}
	f, _ := strconv.ParseFloat(valNode.Value, 64)
	return f
}

// nodeGetFloat64Ptr gets a *float64 value from a mapping node by key.
func nodeGetFloat64Ptr(node *yaml.Node, key string) *float64 {
	valNode := nodeGetValue(node, key)
	if valNode == nil {
		return nil
	}
	f, err := strconv.ParseFloat(valNode.Value, 64)
	if err != nil {
		return nil
	}
	return &f
}

// nodeGetIntPtr gets a *int value from a mapping node by key.
func nodeGetIntPtr(node *yaml.Node, key string) *int {
	valNode := nodeGetValue(node, key)
	if valNode == nil {
		return nil
	}
	i, err := strconv.Atoi(valNode.Value)
	if err != nil {
		return nil
	}
	return &i
}

// nodeGetUint64Ptr gets a *uint64 value from a mapping node by key.
func nodeGetUint64Ptr(node *yaml.Node, key string) *uint64 {
	valNode := nodeGetValue(node, key)
	if valNode == nil {
		return nil
	}
	u, err := strconv.ParseUint(valNode.Value, 10, 64)
	if err != nil {
		return nil
	}
	return &u
}

// nodeGetStringSlice gets a string slice from a mapping node by key.
func nodeGetStringSlice(node *yaml.Node, key string) []string {
	valNode := nodeGetValue(node, key)
	if valNode == nil || valNode.Kind != yaml.SequenceNode {
		return nil
	}
	result := make([]string, 0, len(valNode.Content))
	for _, item := range valNode.Content {
		result = append(result, item.Value)
	}
	return result
}

// nodeGetStringMap gets a map[string]string from a mapping node by key.
func nodeGetStringMap(node *yaml.Node, key string) map[string]string {
	valNode := nodeGetValue(node, key)
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

// nodeGetAny gets the raw interface{} value by decoding the node.
func nodeGetAny(node *yaml.Node, key string) interface{} {
	valNode := nodeGetValue(node, key)
	if valNode == nil {
		return nil
	}
	return nodeToInterface(valNode)
}

// nodeToInterface converts a yaml.Node to interface{} (for Raw storage).
func nodeToInterface(node *yaml.Node) interface{} {
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
			result[i] = nodeToInterface(child)
		}
		return result
	case yaml.MappingNode:
		result := make(map[string]interface{})
		for i := 0; i < len(node.Content)-1; i += 2 {
			key := node.Content[i].Value
			result[key] = nodeToInterface(node.Content[i+1])
		}
		return result
	case yaml.DocumentNode:
		if len(node.Content) > 0 {
			return nodeToInterface(node.Content[0])
		}
		return nil
	case yaml.AliasNode:
		return nodeToInterface(node.Alias)
	default:
		return nil
	}
}

// nodeIsMapping checks if a node is a mapping node.
func nodeIsMapping(node *yaml.Node) bool {
	return node != nil && node.Kind == yaml.MappingNode
}

// nodeIsSequence checks if a node is a sequence node.
func nodeIsSequence(node *yaml.Node) bool {
	return node != nil && node.Kind == yaml.SequenceNode
}

// nodeIsScalar checks if a node is a scalar node.
func nodeIsScalar(node *yaml.Node) bool {
	return node != nil && node.Kind == yaml.ScalarNode
}

// parseNodeExtensions extracts extension fields (x-*) from a yaml.Node.
func parseNodeExtensions(node *yaml.Node) map[string]interface{} {
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
			extensions[key] = nodeToInterface(node.Content[i+1])
		}
	}
	return extensions
}

// nodeHasKey checks if a key exists in a mapping node.
func nodeHasKey(node *yaml.Node, key string) bool {
	return nodeGetValue(node, key) != nil
}

// nodeHasRef checks if the node contains a $ref key.
func nodeHasRef(node *yaml.Node) bool {
	return nodeHasKey(node, "$ref")
}

// nodeGetRef retrieves the $ref value from a node, returning empty string if not found.
func nodeGetRef(node *yaml.Node) string {
	return nodeGetString(node, "$ref")
}
