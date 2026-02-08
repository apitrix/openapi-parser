package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseOpenAPIComponents parses the OpenAPI.Components field.
func parseOpenAPIComponents(node *yaml.Node, ctx *ParseContext) (*openapi30models.Components, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "components must be an object")
	}

	components := &openapi30models.Components{}
	var err error

	// All properties are complex (maps of refs) - delegated to dedicated files
	components.Schemas, err = parseComponentsSchemas(node, ctx)
	if err != nil {
		return nil, err
	}

	components.Responses, err = parseComponentsResponses(node, ctx)
	if err != nil {
		return nil, err
	}

	components.Parameters, err = parseComponentsParameters(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}

	components.Examples, err = parseComponentsExamples(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}

	components.RequestBodies, err = parseComponentsRequestBodies(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}

	components.Headers, err = parseComponentsHeaders(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}

	components.SecuritySchemes, err = parseComponentsSecuritySchemes(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}

	components.Links, err = parseComponentsLinks(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}

	components.Callbacks, err = parseComponentsCallbacks(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}

	components.VendorExtensions = parseNodeExtensions(node)
	components.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, componentsKnownFieldsSet)

	return components, nil
}
