package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseOpenAPIPathsPathItemOperation parses an Operation for a given HTTP method.
// This file handles all HTTP methods (get, put, post, delete, options, head, patch, trace)
// since they all parse to the same Operation type with identical validation/migration rules.
// OpenAPI 3.1.0 spec: https://spec.openapis.org/oas/v3.1.0#operation-object
func parseOpenAPIPathsPathItemOperation(parent *yaml.Node, method string, ctx *ParseContext) (*openapi31models.Operation, error) {
	node := nodeGetValue(parent, method)
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.push(method).errorAt(node, "operation must be an object")
	}

	opCtx := ctx.push(method)
	op := openapi31models.NewOperation()

	// Simple properties
	if tags := nodeGetStringSlice(node, "tags"); tags != nil {
		op.SetProperty("tags", tags)
	}
	if summary := nodeGetString(node, "summary"); summary != "" {
		op.SetProperty("summary", summary)
	}
	if desc := nodeGetString(node, "description"); desc != "" {
		op.SetProperty("description", desc)
	}
	if opID := nodeGetString(node, "operationId"); opID != "" {
		op.SetProperty("operationId", opID)
	}
	if nodeGetBool(node, "deprecated") {
		op.SetProperty("deprecated", true)
	}

	// Complex properties - delegated to dedicated files
	var err error

	externalDocs, err := parseOperationExternalDocs(node, opCtx)
	if err != nil {
		op.Trix.Errors = append(op.Trix.Errors, toParseError(err))
	}
	if externalDocs != nil {
		op.SetProperty("externalDocs", externalDocs)
	}

	parameters, err := parseOperationParameters(node, opCtx)
	if err != nil {
		op.Trix.Errors = append(op.Trix.Errors, toParseError(err))
	}
	if parameters != nil {
		op.SetProperty("parameters", parameters)
	}

	requestBody, err := parseOperationRequestBody(node, opCtx)
	if err != nil {
		op.Trix.Errors = append(op.Trix.Errors, toParseError(err))
	}
	if requestBody != nil {
		op.SetProperty("requestBody", requestBody)
	}

	responses, err := parseOperationResponses(node, opCtx)
	if err != nil {
		op.Trix.Errors = append(op.Trix.Errors, toParseError(err))
	}
	if responses != nil {
		op.SetProperty("responses", responses)
	}

	callbacks, err := parseOperationCallbacks(node, opCtx)
	if err != nil {
		op.Trix.Errors = append(op.Trix.Errors, toParseError(err))
	}
	if callbacks != nil {
		op.SetProperty("callbacks", callbacks)
	}

	security, err := parseOperationSecurity(node, opCtx)
	if err != nil {
		op.Trix.Errors = append(op.Trix.Errors, toParseError(err))
	}
	if security != nil {
		op.SetProperty("security", security)
	}

	servers, err := parseOperationServers(node, opCtx)
	if err != nil {
		op.Trix.Errors = append(op.Trix.Errors, toParseError(err))
	}
	if servers != nil {
		op.SetProperty("servers", servers)
	}

	op.VendorExtensions = parseNodeExtensions(node)
	op.Trix.Source = opCtx.nodeSource(node)

	// Detect unknown fields
	op.Trix.Errors = append(op.Trix.Errors, unknownFieldParseErrors(opCtx.detectUnknown(node, operationKnownFieldsSet))...)

	return op, nil
}
