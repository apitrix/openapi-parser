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

	var errors []openapi30models.ParseError

	// All properties are complex (maps of refs) - delegated to dedicated files
	schemas, err := parseComponentsSchemas(node, ctx)
	if err != nil {
		return nil, err
	}

	responses, err := parseComponentsResponses(node, ctx)
	if err != nil {
		return nil, err
	}

	parameters, err := parseComponentsParameters(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	examples, err := parseComponentsExamples(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	requestBodies, err := parseComponentsRequestBodies(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	headers, err := parseComponentsHeaders(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	securitySchemes, err := parseComponentsSecuritySchemes(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	links, err := parseComponentsLinks(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	callbacks, err := parseComponentsCallbacks(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	// Create via constructor
	components := openapi30models.NewComponents(schemas, responses, parameters, examples, requestBodies, headers, securitySchemes, links, callbacks)

	components.VendorExtensions = parseNodeExtensions(node)
	components.Trix.Source = ctx.nodeSource(node)
	components.Trix.Errors = append(components.Trix.Errors, errors...)

	// Detect unknown fields
	components.Trix.Errors = append(components.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, componentsKnownFieldsSet))...)

	return components, nil
}
