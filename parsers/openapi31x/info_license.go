package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

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

	// Create via constructor
	license := openapi31models.NewLicense(
		nodeGetString(node, "name"),
		nodeGetString(node, "identifier"),
		nodeGetString(node, "url"),
	)

	// Node-level fields
	license.VendorExtensions = parseNodeExtensions(node)
	license.Trix.Source = lctx.nodeSource(node)

	// Detect unknown fields
	license.Trix.Errors = append(license.Trix.Errors, unknownFieldParseErrors(lctx.detectUnknown(node, licenseKnownFieldsSet))...)

	return license, nil
}
