package openapi20

import (
	openapi20models "github.com/apitrix/openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseInfoContact parses the Info.Contact field.
func parseInfoContact(parent *yaml.Node, ctx *ParseContext) (*openapi20models.Contact, error) {
	node := nodeGetValue(parent, "contact")
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.push("contact").errorAt(node, "contact must be an object")
	}

	cctx := ctx.push("contact")

	contact := openapi20models.NewContact(
		nodeGetString(node, "name"),
		nodeGetString(node, "url"),
		nodeGetString(node, "email"),
	)

	contact.VendorExtensions = parseNodeExtensions(node)
	contact.Trix.Source = cctx.nodeSource(node)

	// Detect unknown fields
	contact.Trix.Errors = append(contact.Trix.Errors, unknownFieldParseErrors(cctx.detectUnknown(node, contactKnownFieldsSet))...)

	return contact, nil
}
