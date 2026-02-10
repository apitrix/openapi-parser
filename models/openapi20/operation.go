package openapi20

// Operation describes a single API operation on a path.
// https://swagger.io/specification/v2/#operation-object
type Operation struct {
	Node // embedded - provides VendorExtensions and Trix

	tags         []string
	summary      string
	description  string
	externalDocs *ExternalDocs
	operationID  string
	consumes     []string
	produces     []string
	parameters   []*ParameterRef
	responses    *Responses
	schemes      []string
	deprecated   bool
	security     []SecurityRequirement
}

func (o *Operation) Tags() []string                  { return o.tags }
func (o *Operation) Summary() string                 { return o.summary }
func (o *Operation) Description() string             { return o.description }
func (o *Operation) ExternalDocs() *ExternalDocs     { return o.externalDocs }
func (o *Operation) OperationID() string             { return o.operationID }
func (o *Operation) Consumes() []string              { return o.consumes }
func (o *Operation) Produces() []string              { return o.produces }
func (o *Operation) Parameters() []*ParameterRef     { return o.parameters }
func (o *Operation) Responses() *Responses           { return o.responses }
func (o *Operation) Schemes() []string               { return o.schemes }
func (o *Operation) Deprecated() bool                { return o.deprecated }
func (o *Operation) Security() []SecurityRequirement { return o.security }

// NewOperation creates a new Operation instance.
func NewOperation(
	tags []string, summary, description string, externalDocs *ExternalDocs,
	operationID string, consumes, produces []string,
	parameters []*ParameterRef, responses *Responses,
	schemes []string, deprecated bool, security []SecurityRequirement,
) *Operation {
	return &Operation{
		tags: tags, summary: summary, description: description,
		externalDocs: externalDocs, operationID: operationID,
		consumes: consumes, produces: produces,
		parameters: parameters, responses: responses,
		schemes: schemes, deprecated: deprecated, security: security,
	}
}
