package openapi20

import (
	"iter"

	"github.com/apitrix/openapi-parser/parsers/shared"

	"gopkg.in/yaml.v3"
)

// Thin wrappers delegating to shared.Node* functions.
// This preserves the unexported API so all callers within the package work unchanged.

func nodeGetValue(node *yaml.Node, key string) *yaml.Node {
	return shared.NodeGetValue(node, key)
}

func nodeGetKeyNode(node *yaml.Node, key string) *yaml.Node {
	return shared.NodeGetKeyNode(node, key)
}

func nodeToMap(node *yaml.Node) map[string]*yaml.Node {
	return shared.NodeToMap(node)
}

func nodeKeys(node *yaml.Node) []string {
	return shared.NodeKeys(node)
}

func nodeMapPairs(node *yaml.Node) iter.Seq2[string, *yaml.Node] {
	return shared.NodeMapPairs(node)
}

func nodeToSlice(node *yaml.Node) []*yaml.Node {
	return shared.NodeToSlice(node)
}

func nodeGetString(node *yaml.Node, key string) string {
	return shared.NodeGetString(node, key)
}

func nodeGetBool(node *yaml.Node, key string) bool {
	return shared.NodeGetBool(node, key)
}

func nodeGetBoolPtr(node *yaml.Node, key string) *bool {
	return shared.NodeGetBoolPtr(node, key)
}

func nodeGetInt(node *yaml.Node, key string) int {
	return shared.NodeGetInt(node, key)
}

func nodeGetFloat64(node *yaml.Node, key string) float64 {
	return shared.NodeGetFloat64(node, key)
}

func nodeGetFloat64Ptr(node *yaml.Node, key string) *float64 {
	return shared.NodeGetFloat64Ptr(node, key)
}

func nodeGetIntPtr(node *yaml.Node, key string) *int {
	return shared.NodeGetIntPtr(node, key)
}

func nodeGetUint64Ptr(node *yaml.Node, key string) *uint64 {
	return shared.NodeGetUint64Ptr(node, key)
}

func nodeGetStringSlice(node *yaml.Node, key string) []string {
	return shared.NodeGetStringSlice(node, key)
}

func nodeGetStringMap(node *yaml.Node, key string) map[string]string {
	return shared.NodeGetStringMap(node, key)
}

func nodeGetAny(node *yaml.Node, key string) interface{} {
	return shared.NodeGetAny(node, key)
}

func nodeToInterface(node *yaml.Node) interface{} {
	return shared.NodeToInterface(node)
}

func nodeIsMapping(node *yaml.Node) bool {
	return shared.NodeIsMapping(node)
}

func nodeIsSequence(node *yaml.Node) bool {
	return shared.NodeIsSequence(node)
}

func nodeIsScalar(node *yaml.Node) bool {
	return shared.NodeIsScalar(node)
}

func parseNodeExtensions(node *yaml.Node) map[string]interface{} {
	return shared.ParseNodeExtensions(node)
}

func nodeHasKey(node *yaml.Node, key string) bool {
	return shared.NodeHasKey(node, key)
}

func nodeHasRef(node *yaml.Node) bool {
	return shared.NodeHasRef(node)
}

func nodeGetRef(node *yaml.Node) string {
	return shared.NodeGetRef(node)
}
