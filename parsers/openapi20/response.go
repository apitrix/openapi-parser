package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseResponses parses a Responses container from a yaml.Node.
func parseResponses(node *yaml.Node, ctx *ParseContext) (*openapi20models.Responses, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "responses must be an object")
	}

	codes := make(map[string]*openapi20models.ResponseRef)
	var defaultResp *openapi20models.ResponseRef
	var err error

	for key, respNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		respRef, parseErr := parseResponseRef(respNode, ctx.push(key))
		if parseErr != nil {
			err = parseErr
		}

		if key == "default" {
			defaultResp = respRef
		} else {
			codes[key] = respRef
		}
	}

	responses := openapi20models.NewResponses(defaultResp, codes)

	responses.VendorExtensions = parseNodeExtensions(node)
	responses.Trix.Source = ctx.nodeSource(node)

	if err != nil {
		responses.Trix.Errors = append(responses.Trix.Errors, toParseError(err))
	}

	return responses, nil
}

// parseResponse parses a Response object from a yaml.Node.
func parseResponse(node *yaml.Node, ctx *ParseContext) (*openapi20models.Response, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "response must be an object")
	}

	var err error

	// Complex property - Schema (parsed first for constructor)
	var schema *openapi20models.SchemaRef
	var schemaErr error
	if schemaNode := nodeGetValue(node, "schema"); schemaNode != nil {
		schema, err = parseSchemaRef(schemaNode, ctx.push("schema"))
		if err != nil {
			schemaErr = err
		}
	}

	// Complex property - Headers (parsed first for constructor)
	var headers map[string]*openapi20models.Header
	var headersErr error
	if headersNode := nodeGetValue(node, "headers"); headersNode != nil {
		headers, err = parseHeaders(headersNode, ctx.push("headers"))
		if err != nil {
			headersErr = err
		}
	}

	// Examples - map of mime type to example
	var examples map[string]interface{}
	if examplesNode := nodeGetValue(node, "examples"); examplesNode != nil && nodeIsMapping(examplesNode) {
		examples = make(map[string]interface{})
		for key, exampleNode := range nodeMapPairs(examplesNode) {
			examples[key] = nodeToInterface(exampleNode)
		}
	}

	resp := openapi20models.NewResponse(
		nodeGetString(node, "description"),
		schema,
		headers,
		examples,
	)

	if schemaErr != nil {
		resp.Trix.Errors = append(resp.Trix.Errors, toParseError(schemaErr))
	}
	if headersErr != nil {
		resp.Trix.Errors = append(resp.Trix.Errors, toParseError(headersErr))
	}

	resp.VendorExtensions = parseNodeExtensions(node)
	resp.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	resp.Trix.Errors = append(resp.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, responseKnownFieldsSet))...)

	return resp, nil
}
