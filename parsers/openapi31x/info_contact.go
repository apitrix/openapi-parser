package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseInfoContact parses the Info.Contact field.
func parseInfoContact(parent *yaml.Node, ctx *ParseContext) (*openapi31models.Contact, error) {
	node := nodeGetValue(parent, "contact")
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.push("contact").errorAt(node, "contact must be an object")
	}

	cctx := ctx.push("contact")

	// Create via constructor
	contact := openapi31models.NewContact(
		nodeGetString(node, "name"),
		nodeGetString(node, "url"),
		nodeGetString(node, "email"),
	)

	// Node-level fields
	contact.VendorExtensions = parseNodeExtensions(node)
	contact.Trix.Source = cctx.nodeSource(node)

	// Detect unknown fields
	contact.Trix.Errors = append(contact.Trix.Errors, unknownFieldParseErrors(cctx.detectUnknown(node, contactKnownFieldsSet))...)

	return contact, nil
}
