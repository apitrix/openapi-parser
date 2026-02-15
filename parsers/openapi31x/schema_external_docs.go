package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseExternalDocs parses the Schema.ExternalDocs field.
// Complex property: nested ExternalDocs object
func (p *schemaParser) ParseExternalDocs(parent *yaml.Node, c *ParseContext) (*openapi31models.ExternalDocumentation, error) {
	node := nodeGetValue(parent, "externalDocs")
	if node == nil {
		return nil, nil
	}
	return parseSharedExternalDocs(node, c.Push("externalDocs"))
}
