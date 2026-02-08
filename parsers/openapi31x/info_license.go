package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseInfoLicense parses the Info.License field.
func parseInfoLicense(parent *yaml.Node, ctx *ParseContext) (*openapi31models.License, error) {
	node := nodeGetValue(parent, "license")
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.push("license").errorAt(node, "license must be an object")
	}

	lctx := ctx.push("license")
	license := &openapi31models.License{}

	// All properties are simple - inline
	license.Name = nodeGetString(node, "name")
	license.Identifier = nodeGetString(node, "identifier") // New in 3.1
	license.URL = nodeGetString(node, "url")

	license.VendorExtensions = parseNodeExtensions(node)
	license.Trix.Source = lctx.nodeSource(node)

	// Detect unknown fields
	lctx.detectUnknown(node, licenseKnownFieldsSet)

	return license, nil
}
