package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseOpenAPIInfo parses the OpenAPI.Info field.
func parseOpenAPIInfo(node *yaml.Node, ctx *ParseContext) (*openapi30models.Info, error) {
	if node == nil {
		return nil, ctx.errorf("info is required")
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "info must be an object")
	}

	// Parse sub-objects first
	var errs []openapi30models.ParseError

	contact, err := parseInfoContact(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	license, err := parseInfoLicense(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor
	info := openapi30models.NewInfo(
		nodeGetString(node, "title"),
		nodeGetString(node, "description"),
		nodeGetString(node, "termsOfService"),
		nodeGetString(node, "version"),
		contact,
		license,
	)

	// Node-level fields (VendorExtensions + Trix) are still public via embedding
	info.VendorExtensions = parseNodeExtensions(node)
	info.Trix.Source = ctx.nodeSource(node)
	info.Trix.Errors = append(info.Trix.Errors, errs...)

	// Detect unknown fields
	info.Trix.Errors = append(info.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, infoKnownFieldsSet))...)

	return info, nil
}
