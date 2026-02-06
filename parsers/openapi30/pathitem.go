package openapi30

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

	pathItem := &openapi30models.PathItem{}
	var err error

	// Simple properties - inline
	pathItem.Ref = nodeGetString(node, "$ref")
	pathItem.Summary = nodeGetString(node, "summary")
	pathItem.Description = nodeGetString(node, "description")

	// HTTP method operations - complex but handled inline since same pattern
	pathItem.Get, err = parseOpenAPIPathsPathItemOperation(node, "get", ctx)
	if err != nil {
		return nil, err
	}

	pathItem.Put, err = parseOpenAPIPathsPathItemOperation(node, "put", ctx)
	if err != nil {
		return nil, err
	}

	pathItem.Post, err = parseOpenAPIPathsPathItemOperation(node, "post", ctx)
	if err != nil {
		return nil, err
	}

	pathItem.Delete, err = parseOpenAPIPathsPathItemOperation(node, "delete", ctx)
	if err != nil {
		return nil, err
	}

	pathItem.Options, err = parseOpenAPIPathsPathItemOperation(node, "options", ctx)
	if err != nil {
		return nil, err
	}

	pathItem.Head, err = parseOpenAPIPathsPathItemOperation(node, "head", ctx)
	if err != nil {
		return nil, err
	}

	pathItem.Patch, err = parseOpenAPIPathsPathItemOperation(node, "patch", ctx)
	if err != nil {
		return nil, err
	}

	pathItem.Trace, err = parseOpenAPIPathsPathItemOperation(node, "trace", ctx)
	if err != nil {
		return nil, err
	}

	// Complex properties - delegated to dedicated files
	pathItem.Servers, err = parsePathItemServers(node, ctx)
	if err != nil {
		return nil, err
	}

	pathItem.Parameters, err = parsePathItemParameters(node, ctx)
	if err != nil {
		return nil, err
	}

	pathItem.Extensions = parseNodeExtensions(node)
	pathItem.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, pathItemKnownFieldsSet)

	return pathItem, nil
}
