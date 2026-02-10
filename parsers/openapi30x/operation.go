package openapi30x

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
	var errors []openapi30models.ParseError

	// Simple properties
	tags := nodeGetStringSlice(node, "tags")
	summary := nodeGetString(node, "summary")
	description := nodeGetString(node, "description")
	operationID := nodeGetString(node, "operationId")
	deprecated := nodeGetBool(node, "deprecated")

	// Complex properties - delegated to dedicated files
	externalDocs, err := parseOperationExternalDocs(node, opCtx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	parameters, err := parseOperationParameters(node, opCtx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	requestBody, err := parseOperationRequestBody(node, opCtx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	responses, err := parseOperationResponses(node, opCtx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	callbacks, err := parseOperationCallbacks(node, opCtx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	security, err := parseOperationSecurity(node, opCtx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	servers, err := parseOperationServers(node, opCtx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	// Create via constructor
	op := openapi30models.NewOperation(tags, summary, description, externalDocs, operationID, parameters, requestBody, responses, callbacks, deprecated, security, servers)

	op.VendorExtensions = parseNodeExtensions(node)
	op.Trix.Source = opCtx.nodeSource(node)
	op.Trix.Errors = append(op.Trix.Errors, errors...)

	// Detect unknown fields
	op.Trix.Errors = append(op.Trix.Errors, unknownFieldParseErrors(opCtx.detectUnknown(node, operationKnownFieldsSet))...)

	return op, nil
}
