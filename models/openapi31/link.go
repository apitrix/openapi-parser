package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Link represents a possible design-time link for a response.
// https://spec.openapis.org/oas/v3.1.0#link-object
type Link struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	operationRef string
	operationID  string
	parameters   map[string]interface{}
	requestBody  interface{}
	description  string
	server       *Server
}

func (l *Link) OperationRef() string               { return l.operationRef }
func (l *Link) OperationID() string                { return l.operationID }
func (l *Link) Parameters() map[string]interface{} { return l.parameters }
func (l *Link) RequestBody() interface{}           { return l.requestBody }
func (l *Link) Description() string                { return l.description }
func (l *Link) Server() *Server                    { return l.server }

// NewLink creates a new Link instance.
func NewLink(operationRef, operationID, description string, parameters map[string]interface{}, requestBody interface{}, server *Server) *Link {
	return &Link{
		operationRef: operationRef, operationID: operationID,
		parameters: parameters, requestBody: requestBody,
		description: description, server: server,
	}
}

func (l *Link) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "operationRef", Value: l.operationRef},
		{Key: "operationId", Value: l.operationID},
		{Key: "parameters", Value: l.parameters},
		{Key: "requestBody", Value: l.requestBody},
		{Key: "description", Value: l.description},
		{Key: "server", Value: l.server},
	}
	return shared.AppendExtensions(fields, l.VendorExtensions)
}

func (l *Link) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(l.marshalFields())
}

func (l *Link) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(l.marshalFields())
}

var _ yaml.Marshaler = (*Link)(nil)
