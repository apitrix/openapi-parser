package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseInfoLicense parses the Info.License field.
func parseInfoLicense(parent *yaml.Node, ctx *ParseContext) (*openapi30models.License, error) {
	node := nodeGetValue(parent, "license")
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.push("license").errorAt(node, "license must be an object")
	}

	lctx := ctx.push("license")
	license := &openapi30models.License{}

	// All properties are simple - inline
	license.Name = nodeGetString(node, "name")
	license.URL = nodeGetString(node, "url")

	license.Extensions = parseNodeExtensions(node)
	license.NodeSource = lctx.nodeSource(node)

	// Detect unknown fields
	lctx.detectUnknown(node, licenseKnownFieldsSet)

	return license, nil
}
