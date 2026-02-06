package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseSwaggerInfo parses the Swagger.Info field.
func parseSwaggerInfo(node *yaml.Node, ctx *ParseContext) (*openapi20models.Info, error) {
	if node == nil {
		return nil, ctx.errorf("info is required")
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "info must be an object")
	}

	info := &openapi20models.Info{}
	var err error

	// Simple properties - inline
	info.Title = nodeGetString(node, "title")
	info.Description = nodeGetString(node, "description")
	info.TermsOfService = nodeGetString(node, "termsOfService")
	info.Version = nodeGetString(node, "version")

	// Complex properties - delegated to dedicated files
	info.Contact, err = parseInfoContact(node, ctx)
	if err != nil {
		return nil, err
	}

	info.License, err = parseInfoLicense(node, ctx)
	if err != nil {
		return nil, err
	}

	info.Extensions = parseNodeExtensions(node)
	info.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, infoKnownFieldsSet)

	return info, nil
}
