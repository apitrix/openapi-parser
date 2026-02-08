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

	pathItem := &openapi20models.PathItem{}
	var err error

	// Handle $ref
	pathItem.Ref = nodeGetString(node, "$ref")

	// HTTP methods - delegated to operation parser
	if getNode := nodeGetValue(node, "get"); getNode != nil {
		pathItem.Get, err = parseOperation(getNode, ctx.push("get"))
		if err != nil {
			return nil, err
		}
	}

	if putNode := nodeGetValue(node, "put"); putNode != nil {
		pathItem.Put, err = parseOperation(putNode, ctx.push("put"))
		if err != nil {
			return nil, err
		}
	}

	if postNode := nodeGetValue(node, "post"); postNode != nil {
		pathItem.Post, err = parseOperation(postNode, ctx.push("post"))
		if err != nil {
			return nil, err
		}
	}

	if deleteNode := nodeGetValue(node, "delete"); deleteNode != nil {
		pathItem.Delete, err = parseOperation(deleteNode, ctx.push("delete"))
		if err != nil {
			return nil, err
		}
	}

	if optionsNode := nodeGetValue(node, "options"); optionsNode != nil {
		pathItem.Options, err = parseOperation(optionsNode, ctx.push("options"))
		if err != nil {
			return nil, err
		}
	}

	if headNode := nodeGetValue(node, "head"); headNode != nil {
		pathItem.Head, err = parseOperation(headNode, ctx.push("head"))
		if err != nil {
			return nil, err
		}
	}

	if patchNode := nodeGetValue(node, "patch"); patchNode != nil {
		pathItem.Patch, err = parseOperation(patchNode, ctx.push("patch"))
		if err != nil {
			return nil, err
		}
	}

	// Path-level parameters
	if paramsNode := nodeGetValue(node, "parameters"); paramsNode != nil {
		pathItem.Parameters, err = parseParameterRefs(paramsNode, ctx.push("parameters"))
		if err != nil {
			return nil, err
		}
	}

	pathItem.VendorExtensions = parseNodeExtensions(node)
	pathItem.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, pathItemKnownFieldsSet)

	return pathItem, nil
}
