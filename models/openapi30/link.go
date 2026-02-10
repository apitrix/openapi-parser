package openapi30

// Link represents a possible design-time link for a response.
// https://spec.openapis.org/oas/v3.0.3#link-object
type Link struct {
	Node // embedded - provides VendorExtensions and Trix

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
func NewLink(operationRef, operationID string, parameters map[string]interface{}, requestBody interface{}, description string, server *Server) *Link {
	return &Link{operationRef: operationRef, operationID: operationID, parameters: parameters, requestBody: requestBody, description: description, server: server}
}
