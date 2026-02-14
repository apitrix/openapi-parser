package openapi20

import (
	"openapi-parser/models/shared"
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

	var err error
	var errors []error

	// Complex property - ExternalDocs (parsed first for constructor)
	var externalDocs *openapi20models.ExternalDocs
	if edNode := nodeGetValue(node, "externalDocs"); edNode != nil {
		externalDocs, err = parseExternalDocs(edNode, ctx.push("externalDocs"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	// Complex property - Parameters
	var parameters []*shared.Ref[openapi20models.Parameter]
	if paramsNode := nodeGetValue(node, "parameters"); paramsNode != nil {
		parameters, err = parseParameterRefs(paramsNode, ctx.push("parameters"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	// Complex property - Responses
	var responses *openapi20models.Responses
	if respNode := nodeGetValue(node, "responses"); respNode != nil {
		responses, err = parseResponses(respNode, ctx.push("responses"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	// Complex property - Security
	var security []openapi20models.SecurityRequirement
	if secNode := nodeGetValue(node, "security"); secNode != nil {
		security, err = parseSecurityRequirements(secNode, ctx.push("security"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	op := openapi20models.NewOperation(
		nodeGetStringSlice(node, "tags"),
		nodeGetString(node, "summary"),
		nodeGetString(node, "description"),
		externalDocs,
		nodeGetString(node, "operationId"),
		nodeGetStringSlice(node, "consumes"),
		nodeGetStringSlice(node, "produces"),
		parameters,
		responses,
		nodeGetStringSlice(node, "schemes"),
		nodeGetBool(node, "deprecated"),
		security,
	)

	for _, e := range errors {
		op.Trix.Errors = append(op.Trix.Errors, toParseError(e))
	}

	op.VendorExtensions = parseNodeExtensions(node)
	op.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	op.Trix.Errors = append(op.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, operationKnownFieldsSet))...)

	return op, nil
}
