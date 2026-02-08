package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseOpenAPIInfo parses the OpenAPI.Info field.
func parseOpenAPIInfo(node *yaml.Node, ctx *ParseContext) (*openapi31models.Info, error) {
	if node == nil {
		return nil, ctx.errorf("info is required")
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "info must be an object")
	}

	info := &openapi31models.Info{}
	var err error

	// Simple properties - inline
	info.Title = nodeGetString(node, "title")
	info.Summary = nodeGetString(node, "summary") // New in 3.1
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

	info.VendorExtensions = parseNodeExtensions(node)
	info.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, infoKnownFieldsSet)

	return info, nil
}
