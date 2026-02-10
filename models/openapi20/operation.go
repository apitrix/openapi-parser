package openapi20

// Operation describes a single API operation on a path.
// https://swagger.io/specification/v2/#operation-object
type Operation struct {
	Node // embedded - provides VendorExtensions and Trix

	Tags         []string              `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocs         `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationID  string                `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Consumes     []string              `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces     []string              `json:"produces,omitempty" yaml:"produces,omitempty"`
	Parameters   []*ParameterRef       `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Responses    *Responses            `json:"responses" yaml:"responses"`
	Schemes      []string              `json:"schemes,omitempty" yaml:"schemes,omitempty"`
	Deprecated   bool                  `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     []SecurityRequirement `json:"security,omitempty" yaml:"security,omitempty"`
}

// NewOperation creates a new Operation instance.
func NewOperation() *Operation {
	return &Operation{}
}
