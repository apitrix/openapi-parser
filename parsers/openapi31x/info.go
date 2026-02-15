package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

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

	// Parse sub-objects first
	var errs []openapi31models.ParseError

	contact, err := parseInfoContact(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	license, err := parseInfoLicense(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor
	info := openapi31models.NewInfo(
		nodeGetString(node, "title"),
		nodeGetString(node, "summary"),
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
