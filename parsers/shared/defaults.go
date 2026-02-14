package shared

import "gopkg.in/yaml.v3"

// DefaultServersURL is the OpenAPI 3.x default server URL when servers is absent or empty.
const DefaultServersURL = "/"

// DefaultBasePath is the Swagger 2.0 default basePath when absent.
const DefaultBasePath = "/"

// ApplySpecDefaults returns true if the config has ApplySpecDefaults enabled.
// Callers use this to decide whether to fill in spec defaults for absent fields.
func ApplySpecDefaults(cfg *ParseConfig) bool {
	return cfg != nil && cfg.ApplySpecDefaults
}

// ServersAbsentOrEmpty returns true if the servers node is nil, or is a sequence
// with no elements. Used with ApplySpecDefaults to apply the OpenAPI 3.x default.
func ServersAbsentOrEmpty(node *yaml.Node) bool {
	if node == nil {
		return true
	}
	if node.Kind != yaml.SequenceNode {
		return true
	}
	return len(node.Content) == 0
}
