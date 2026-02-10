package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseOpenAPIPathsPathItem parses a PathItem object from a yaml.Node.
// OpenAPI 3.0.3 spec: https://spec.openapis.org/oas/v3.0.3#path-item-object
func parseOpenAPIPathsPathItem(node *yaml.Node, ctx *ParseContext) (*openapi30models.PathItem, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "pathItem must be an object")
	}

	var errors []openapi30models.ParseError

	// Simple properties
	ref := nodeGetString(node, "$ref")
	summary := nodeGetString(node, "summary")
	description := nodeGetString(node, "description")

	// HTTP method operations
	parseOp := func(method string) *openapi30models.Operation {
		op, err := parseOpenAPIPathsPathItemOperation(node, method, ctx)
		if err != nil {
			errors = append(errors, toParseError(err))
		}
		return op
	}

	get := parseOp("get")
	put := parseOp("put")
	post := parseOp("post")
	del := parseOp("delete")
	options := parseOp("options")
	head := parseOp("head")
	patch := parseOp("patch")
	trace := parseOp("trace")

	// Complex properties - delegated to dedicated files
	servers, err := parsePathItemServers(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	parameters, err := parsePathItemParameters(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	// Create via constructor
	pathItem := openapi30models.NewPathItem(ref, summary, description, get, put, post, del, options, head, patch, trace, servers, parameters)

	pathItem.VendorExtensions = parseNodeExtensions(node)
	pathItem.Trix.Source = ctx.nodeSource(node)
	pathItem.Trix.Errors = append(pathItem.Trix.Errors, errors...)

	// Detect unknown fields
	pathItem.Trix.Errors = append(pathItem.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, pathItemKnownFieldsSet))...)

	return pathItem, nil
}
