package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parsePathItemAdditionalOperations parses the PathItem.additionalOperations field (OpenAPI 3.2).
func parsePathItemAdditionalOperations(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi31models.Operation, error) {
	node := nodeGetValue(parent, "additionalOperations")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	ops := make(map[string]*openapi31models.Operation)
	actx := ctx.push("additionalOperations")
	for name, opNode := range nodeMapPairs(node) {
		op, err := parseOperationFromNode(opNode, actx.push(name))
		if err != nil {
			return nil, err
		}
		if op != nil {
			ops[name] = op
		}
	}
	return ops, nil
}

// parseOpenAPIPathsPathItem parses a PathItem object from a yaml.Node.
// OpenAPI 3.1.0 spec: https://spec.openapis.org/oas/v3.1.0#path-item-object
func parseOpenAPIPathsPathItem(node *yaml.Node, ctx *ParseContext) (*openapi31models.PathItem, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "pathItem must be an object")
	}

	pathItem := openapi31models.NewPathItem()
	var err error

	// Simple properties
	if ref := nodeGetString(node, "$ref"); ref != "" {
		pathItem.SetProperty("$ref", ref)
	}
	if summary := nodeGetString(node, "summary"); summary != "" {
		pathItem.SetProperty("summary", summary)
	}
	if desc := nodeGetString(node, "description"); desc != "" {
		pathItem.SetProperty("description", desc)
	}

	// HTTP method operations
	get, err := parseOpenAPIPathsPathItemOperation(node, "get", ctx)
	if err != nil {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(err))
	}
	if get != nil {
		pathItem.SetProperty("get", get)
	}

	put, err := parseOpenAPIPathsPathItemOperation(node, "put", ctx)
	if err != nil {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(err))
	}
	if put != nil {
		pathItem.SetProperty("put", put)
	}

	post, err := parseOpenAPIPathsPathItemOperation(node, "post", ctx)
	if err != nil {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(err))
	}
	if post != nil {
		pathItem.SetProperty("post", post)
	}

	del, err := parseOpenAPIPathsPathItemOperation(node, "delete", ctx)
	if err != nil {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(err))
	}
	if del != nil {
		pathItem.SetProperty("delete", del)
	}

	options, err := parseOpenAPIPathsPathItemOperation(node, "options", ctx)
	if err != nil {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(err))
	}
	if options != nil {
		pathItem.SetProperty("options", options)
	}

	head, err := parseOpenAPIPathsPathItemOperation(node, "head", ctx)
	if err != nil {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(err))
	}
	if head != nil {
		pathItem.SetProperty("head", head)
	}

	patch, err := parseOpenAPIPathsPathItemOperation(node, "patch", ctx)
	if err != nil {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(err))
	}
	if patch != nil {
		pathItem.SetProperty("patch", patch)
	}

	trace, err := parseOpenAPIPathsPathItemOperation(node, "trace", ctx)
	if err != nil {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(err))
	}
	if trace != nil {
		pathItem.SetProperty("trace", trace)
	}

	// OpenAPI 3.2: query and additionalOperations
	query, err := parseOpenAPIPathsPathItemOperation(node, "query", ctx)
	if err != nil {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(err))
	}
	if query != nil {
		pathItem.SetProperty("query", query)
	}

	additionalOps, err := parsePathItemAdditionalOperations(node, ctx)
	if err != nil {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(err))
	}
	if additionalOps != nil {
		pathItem.SetProperty("additionalOperations", additionalOps)
	}

	// Complex properties - delegated to dedicated files
	servers, err := parsePathItemServers(node, ctx)
	if err != nil {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(err))
	}
	if servers != nil {
		pathItem.SetProperty("servers", servers)
	}

	parameters, err := parsePathItemParameters(node, ctx)
	if err != nil {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(err))
	}
	if parameters != nil {
		pathItem.SetProperty("parameters", parameters)
	}

	pathItem.VendorExtensions = parseNodeExtensions(node)
	pathItem.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	pathItem.Trix.Errors = append(pathItem.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, pathItemKnownFieldsSet))...)

	return pathItem, nil
}
