package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parseRequestBodyRef parses a RequestBodyRef from a yaml.Node.
func parseRequestBodyRef(node *yaml.Node, ctx *ParseContext) (*shared.RefWithMeta[openapi31models.RequestBody], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "requestBody must be an object")
	}

	ref := &shared.RefWithMeta[openapi31models.RequestBody]{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Summary = nodeGetString(node, "summary")
		ref.Description = nodeGetString(node, "description")
		return ref, nil
	}

	body, err := parseSharedRequestBody(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.SetValue(body)

	return ref, nil
}
