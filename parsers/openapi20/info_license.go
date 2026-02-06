package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseInfoLicense parses the Info.License field.
func parseInfoLicense(parent *yaml.Node, ctx *ParseContext) (*openapi20models.License, error) {
	node := nodeGetValue(parent, "license")
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.push("license").errorAt(node, "license must be an object")
	}

	lctx := ctx.push("license")
	license := &openapi20models.License{}

	// All properties are simple - inline
	license.Name = nodeGetString(node, "name")
	license.URL = nodeGetString(node, "url")

	license.Extensions = parseNodeExtensions(node)
	license.NodeSource = lctx.nodeSource(node)

	// Detect unknown fields
	lctx.detectUnknown(node, licenseKnownFields)

	return license, nil
}
