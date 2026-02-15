package openapi20

import (
	openapi20models "github.com/apitrix/openapi-parser/models/openapi20"

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

	var err error

	// Complex properties - delegated to dedicated files
	var contact *openapi20models.Contact
	contact, err = parseInfoContact(node, ctx)
	var contactErr error
	if err != nil {
		contactErr = err
	}

	var license *openapi20models.License
	license, err = parseInfoLicense(node, ctx)
	var licenseErr error
	if err != nil {
		licenseErr = err
	}

	info := openapi20models.NewInfo(
		nodeGetString(node, "title"),
		nodeGetString(node, "description"),
		nodeGetString(node, "termsOfService"),
		nodeGetString(node, "version"),
		contact,
		license,
	)

	if contactErr != nil {
		info.Trix.Errors = append(info.Trix.Errors, toParseError(contactErr))
	}
	if licenseErr != nil {
		info.Trix.Errors = append(info.Trix.Errors, toParseError(licenseErr))
	}

	info.VendorExtensions = parseNodeExtensions(node)
	info.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	info.Trix.Errors = append(info.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, infoKnownFieldsSet))...)

	return info, nil
}
