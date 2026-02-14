package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parsePathItem parses a PathItem object from a yaml.Node.
func parsePathItem(node *yaml.Node, ctx *ParseContext) (*openapi20models.PathItem, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "path item must be an object")
	}

	var err error
	var errors []error

	// HTTP methods - delegated to operation parser
	var get, put, post, del, options, head, patch *openapi20models.Operation

	if getNode := nodeGetValue(node, "get"); getNode != nil {
		get, err = parseOperation(getNode, ctx.push("get"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	if putNode := nodeGetValue(node, "put"); putNode != nil {
		put, err = parseOperation(putNode, ctx.push("put"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	if postNode := nodeGetValue(node, "post"); postNode != nil {
		post, err = parseOperation(postNode, ctx.push("post"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	if deleteNode := nodeGetValue(node, "delete"); deleteNode != nil {
		del, err = parseOperation(deleteNode, ctx.push("delete"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	if optionsNode := nodeGetValue(node, "options"); optionsNode != nil {
		options, err = parseOperation(optionsNode, ctx.push("options"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	if headNode := nodeGetValue(node, "head"); headNode != nil {
		head, err = parseOperation(headNode, ctx.push("head"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	if patchNode := nodeGetValue(node, "patch"); patchNode != nil {
		patch, err = parseOperation(patchNode, ctx.push("patch"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	// Path-level parameters
	var parameters []*openapi20models.RefParameter
	if paramsNode := nodeGetValue(node, "parameters"); paramsNode != nil {
		parameters, err = parseParameterRefs(paramsNode, ctx.push("parameters"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	pathItem := openapi20models.NewPathItem(
		nodeGetString(node, "$ref"),
		get, put, post, del, options, head, patch,
		parameters,
	)

	for _, e := range errors {
		pathItem.Trix.Errors = append(pathItem.Trix.Errors, toParseError(e))
	}

	pathItem.VendorExtensions = parseNodeExtensions(node)
	pathItem.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	pathItem.Trix.Errors = append(pathItem.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, pathItemKnownFieldsSet))...)

	return pathItem, nil
}
