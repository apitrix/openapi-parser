package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseOperation parses an Operation object from a yaml.Node.
func parseOperation(node *yaml.Node, ctx *ParseContext) (*openapi20models.Operation, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "operation must be an object")
	}

	op := &openapi20models.Operation{}
	var err error

	// Simple properties - inline
	op.Tags = nodeGetStringSlice(node, "tags")
	op.Summary = nodeGetString(node, "summary")
	op.Description = nodeGetString(node, "description")
	op.OperationID = nodeGetString(node, "operationId")
	op.Consumes = nodeGetStringSlice(node, "consumes")
	op.Produces = nodeGetStringSlice(node, "produces")
	op.Schemes = nodeGetStringSlice(node, "schemes")
	op.Deprecated = nodeGetBool(node, "deprecated")

	// Complex property - ExternalDocs
	if edNode := nodeGetValue(node, "externalDocs"); edNode != nil {
		op.ExternalDocs, err = parseExternalDocs(edNode, ctx.push("externalDocs"))
		if err != nil {
			return nil, err
		}
	}

	// Complex property - Parameters
	if paramsNode := nodeGetValue(node, "parameters"); paramsNode != nil {
		op.Parameters, err = parseParameterRefs(paramsNode, ctx.push("parameters"))
		if err != nil {
			return nil, err
		}
	}

	// Complex property - Responses
	if respNode := nodeGetValue(node, "responses"); respNode != nil {
		op.Responses, err = parseResponses(respNode, ctx.push("responses"))
		if err != nil {
			return nil, err
		}
	}

	// Complex property - Security
	if secNode := nodeGetValue(node, "security"); secNode != nil {
		op.Security, err = parseSecurityRequirements(secNode, ctx.push("security"))
		if err != nil {
			return nil, err
		}
	}

	op.VendorExtensions = parseNodeExtensions(node)
	op.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, operationKnownFieldsSet)

	return op, nil
}
