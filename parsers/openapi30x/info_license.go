package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

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

	// Create via constructor
	license := openapi30models.NewLicense(
		nodeGetString(node, "name"),
		nodeGetString(node, "url"),
	)

	// Node-level fields
	license.VendorExtensions = parseNodeExtensions(node)
	license.Trix.Source = lctx.nodeSource(node)

	// Detect unknown fields
	license.Trix.Errors = append(license.Trix.Errors, unknownFieldParseErrors(lctx.detectUnknown(node, licenseKnownFieldsSet))...)

	return license, nil
}
