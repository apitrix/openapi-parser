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

	responses := &openapi20models.Responses{}
	responses.Codes = make(map[string]*openapi20models.ResponseRef)
	var err error

	for key, respNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		respRef, err := parseResponseRef(respNode, ctx.push(key))
		if err != nil {
			responses.Trix.Errors = append(responses.Trix.Errors, toParseError(err))
		}

		if key == "default" {
			responses.Default = respRef
		} else {
			responses.Codes[key] = respRef
		}
	}

	responses.VendorExtensions = parseNodeExtensions(node)
	responses.Trix.Source = ctx.nodeSource(node)

	return responses, err
}

// parseResponse parses a Response object from a yaml.Node.
func parseResponse(node *yaml.Node, ctx *ParseContext) (*openapi20models.Response, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "response must be an object")
	}

	resp := &openapi20models.Response{}
	var err error

	// Simple properties - inline
	resp.Description = nodeGetString(node, "description")

	// Complex property - Schema
	if schemaNode := nodeGetValue(node, "schema"); schemaNode != nil {
		resp.Schema, err = parseSchemaRef(schemaNode, ctx.push("schema"))
		if err != nil {
			resp.Trix.Errors = append(resp.Trix.Errors, toParseError(err))
		}
	}

	// Complex property - Headers
	if headersNode := nodeGetValue(node, "headers"); headersNode != nil {
		resp.Headers, err = parseHeaders(headersNode, ctx.push("headers"))
		if err != nil {
			resp.Trix.Errors = append(resp.Trix.Errors, toParseError(err))
		}
	}

	// Examples - map of mime type to example
	if examplesNode := nodeGetValue(node, "examples"); examplesNode != nil && nodeIsMapping(examplesNode) {
		resp.Examples = make(map[string]interface{})
		for key, exampleNode := range nodeMapPairs(examplesNode) {
			resp.Examples[key] = nodeToInterface(exampleNode)
		}
	}

	resp.VendorExtensions = parseNodeExtensions(node)
	resp.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	resp.Trix.Errors = append(resp.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, responseKnownFieldsSet))...)

	return resp, nil
}
