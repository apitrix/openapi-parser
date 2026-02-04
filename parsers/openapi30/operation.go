package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseOpenAPIPathsPathItemOperation parses an Operation for a given HTTP method.
// This file handles all HTTP methods (get, put, post, delete, options, head, patch, trace)
// since they all parse to the same Operation type with identical validation/migration rules.
// OpenAPI 3.0.3 spec: https://spec.openapis.org/oas/v3.0.3#operation-object
func parseOpenAPIPathsPathItemOperation(parent *yaml.Node, method string, ctx *ParseContext) (*openapi30models.Operation, error) {
	node := nodeGetValue(parent, method)
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.push(method).errorAt(node, "operation must be an object")
	}

	opCtx := ctx.push(method)
	op := &openapi30models.Operation{}
	var err error

	// Simple properties - inline
	op.Tags = nodeGetStringSlice(node, "tags")
	op.Summary = nodeGetString(node, "summary")
	op.Description = nodeGetString(node, "description")
	op.OperationID = nodeGetString(node, "operationId")
	op.Deprecated = nodeGetBool(node, "deprecated")

	// Complex properties - delegated to dedicated files
	op.ExternalDocs, err = parseOperationExternalDocs(node, opCtx)
	if err != nil {
		return nil, err
	}

	op.Parameters, err = parseOperationParameters(node, opCtx)
	if err != nil {
		return nil, err
	}

	op.RequestBody, err = parseOperationRequestBody(node, opCtx)
	if err != nil {
		return nil, err
	}

	op.Responses, err = parseOperationResponses(node, opCtx)
	if err != nil {
		return nil, err
	}

	op.Callbacks, err = parseOperationCallbacks(node, opCtx)
	if err != nil {
		return nil, err
	}

	op.Security, err = parseOperationSecurity(node, opCtx)
	if err != nil {
		return nil, err
	}

	op.Servers, err = parseOperationServers(node, opCtx)
	if err != nil {
		return nil, err
	}

	op.Extensions = parseNodeExtensions(node)
	op.NodeSource = opCtx.nodeSource(node)

	// Detect unknown fields
	opCtx.detectUnknown(node, operationKnownFields)

	return op, nil
}
