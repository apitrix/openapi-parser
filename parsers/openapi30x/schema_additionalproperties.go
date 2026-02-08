package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// AdditionalPropertiesResult represents the polymorphic additionalProperties field
// which can be either a boolean or a schema reference.
type AdditionalPropertiesResult struct {
	Allowed   *bool          // If set, additionalProperties is a boolean
	SchemaRef *openapi30models.SchemaRef // If set, additionalProperties is a schema
}

// ParseAdditionalProperties parses the Schema.AdditionalProperties field.
// Complex property: polymorphic - can be bool OR SchemaRef
func (p *schemaParser) ParseAdditionalProperties(parent *yaml.Node, c *ParseContext) (*AdditionalPropertiesResult, error) {
	node := nodeGetValue(parent, "additionalProperties")
	if node == nil {
		return nil, nil
	}

	pctx := c.Push("additionalProperties")

	// Check if it's a boolean
	if nodeIsScalar(node) {
		// Try to interpret as boolean
		allowed := nodeGetBool(parent, "additionalProperties")
		return &AdditionalPropertiesResult{
			Allowed:   &allowed,
			SchemaRef: nil,
		}, nil
	}

	// Otherwise, it should be a schema reference
	if nodeIsMapping(node) {
		schemaRef, err := parseSchemaRef(node, pctx)
		if err != nil {
			return nil, err
		}
		return &AdditionalPropertiesResult{
			Allowed:   nil,
			SchemaRef: schemaRef,
		}, nil
	}

	return nil, pctx.errorAt(node, "additionalProperties must be a boolean or schema object")
}
