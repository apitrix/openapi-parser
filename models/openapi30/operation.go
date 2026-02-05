package openapi30

// Operation describes a single API operation on a path.
// https://spec.openapis.org/oas/v3.0.3#operation-object
type Operation struct {
	Node // embedded - provides NodeSource and Extensions

	Tags         []string                `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string                  `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                  `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation  `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationID  string                  `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   []*ParameterRef         `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  *RequestBodyRef         `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    *Responses              `json:"responses" yaml:"responses"`
	Callbacks    map[string]*CallbackRef `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   bool                    `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     []SecurityRequirement   `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      []*Server               `json:"servers,omitempty" yaml:"servers,omitempty"`
}
