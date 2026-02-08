package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseOpenAPIComponents parses the OpenAPI.Components field.
func parseOpenAPIComponents(node *yaml.Node, ctx *ParseContext) (*openapi31models.Components, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "components must be an object")
	}

	components := &openapi31models.Components{}
	var err error

	// All properties are complex (maps of refs) - delegated to dedicated files
	components.Schemas, err = parseComponentsSchemas(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}

	components.Responses, err = parseComponentsResponses(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
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

	// PathItems - new in 3.1
	components.PathItems, err = parseComponentsPathItems(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}

	components.VendorExtensions = parseNodeExtensions(node)
	components.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	components.Trix.Errors = append(components.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, componentsKnownFieldsSet))...)

	return components, nil
}
