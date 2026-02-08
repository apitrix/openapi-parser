package openapi30

// Link represents a possible design-time link for a response.
// https://spec.openapis.org/oas/v3.0.3#link-object
type Link struct {
	Node // embedded - provides VendorExtensions and Trix

	OperationRef string                 `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
	OperationID  string                 `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  interface{}            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Server       *Server                `json:"server,omitempty" yaml:"server,omitempty"`
}
