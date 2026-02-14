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

func (l *Link) SetOperationRef(operationRef string) error {
	if err := l.Trix.RunHooks("operationRef", l.operationRef, operationRef); err != nil {
		return err
	}
	l.operationRef = operationRef
	return nil
}
func (l *Link) SetOperationID(operationID string) error {
	if err := l.Trix.RunHooks("operationId", l.operationID, operationID); err != nil {
		return err
	}
	l.operationID = operationID
	return nil
}
func (l *Link) SetParameters(parameters map[string]interface{}) error {
	if err := l.Trix.RunHooks("parameters", l.parameters, parameters); err != nil {
		return err
	}
	l.parameters = parameters
	return nil
}
func (l *Link) SetRequestBody(requestBody interface{}) error {
	if err := l.Trix.RunHooks("requestBody", l.requestBody, requestBody); err != nil {
		return err
	}
	l.requestBody = requestBody
	return nil
}
func (l *Link) SetDescription(description string) error {
	if err := l.Trix.RunHooks("description", l.description, description); err != nil {
		return err
	}
	l.description = description
	return nil
}
func (l *Link) SetServer(server *Server) error {
	if err := l.Trix.RunHooks("server", l.server, server); err != nil {
		return err
	}
	l.server = server
	return nil
}

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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (l *Link) MarshalFields() []shared.Field { return l.marshalFields() }

func (l *Link) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(l.marshalFields())
}

func (l *Link) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(l.marshalFields())
}

var _ yaml.Marshaler = (*Link)(nil)
