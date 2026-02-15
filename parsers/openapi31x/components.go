package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

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

	components := openapi31models.NewComponents()
	var err error

	// All properties are complex (maps of refs) - delegated to dedicated files
	schemas, err := parseComponentsSchemas(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}
	if schemas != nil {
		components.SetProperty("schemas", schemas)
	}

	responses, err := parseComponentsResponses(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}
	if responses != nil {
		components.SetProperty("responses", responses)
	}

	parameters, err := parseComponentsParameters(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}
	if parameters != nil {
		components.SetProperty("parameters", parameters)
	}

	examples, err := parseComponentsExamples(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}
	if examples != nil {
		components.SetProperty("examples", examples)
	}

	requestBodies, err := parseComponentsRequestBodies(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}
	if requestBodies != nil {
		components.SetProperty("requestBodies", requestBodies)
	}

	headers, err := parseComponentsHeaders(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}
	if headers != nil {
		components.SetProperty("headers", headers)
	}

	securitySchemes, err := parseComponentsSecuritySchemes(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}
	if securitySchemes != nil {
		components.SetProperty("securitySchemes", securitySchemes)
	}

	links, err := parseComponentsLinks(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}
	if links != nil {
		components.SetProperty("links", links)
	}

	callbacks, err := parseComponentsCallbacks(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}
	if callbacks != nil {
		components.SetProperty("callbacks", callbacks)
	}

	// PathItems - new in 3.1
	pathItems, err := parseComponentsPathItems(node, ctx)
	if err != nil {
		components.Trix.Errors = append(components.Trix.Errors, toParseError(err))
	}
	if pathItems != nil {
		components.SetProperty("pathItems", pathItems)
	}

	components.VendorExtensions = parseNodeExtensions(node)
	components.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	components.Trix.Errors = append(components.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, componentsKnownFieldsSet))...)

	return components, nil
}
