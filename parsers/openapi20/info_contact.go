package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

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
	contact := &openapi20models.Contact{}

	// All properties are simple - inline
	contact.Name = nodeGetString(node, "name")
	contact.URL = nodeGetString(node, "url")
	contact.Email = nodeGetString(node, "email")

	contact.Extensions = parseNodeExtensions(node)
	contact.NodeSource = cctx.nodeSource(node)

	// Detect unknown fields
	cctx.detectUnknown(node, contactKnownFieldsSet)

	return contact, nil
}
