package openapi30

// Operation describes a single API operation on a path.
// https://spec.openapis.org/oas/v3.0.3#operation-object
type Operation struct {
	Node // embedded - provides VendorExtensions and Trix

	tags         []string
	summary      string
	description  string
	externalDocs *ExternalDocumentation
	operationID  string
	parameters   []*ParameterRef
	requestBody  *RequestBodyRef
	responses    *Responses
	callbacks    map[string]*CallbackRef
	deprecated   bool
	security     []SecurityRequirement
	servers      []*Server
}

func (o *Operation) Tags() []string                       { return o.tags }
func (o *Operation) Summary() string                      { return o.summary }
func (o *Operation) Description() string                  { return o.description }
func (o *Operation) ExternalDocs() *ExternalDocumentation { return o.externalDocs }
func (o *Operation) OperationID() string                  { return o.operationID }
func (o *Operation) Parameters() []*ParameterRef          { return o.parameters }
func (o *Operation) RequestBody() *RequestBodyRef         { return o.requestBody }
func (o *Operation) Responses() *Responses                { return o.responses }
func (o *Operation) Callbacks() map[string]*CallbackRef   { return o.callbacks }
func (o *Operation) Deprecated() bool                     { return o.deprecated }
func (o *Operation) Security() []SecurityRequirement      { return o.security }
func (o *Operation) Servers() []*Server                   { return o.servers }

// NewOperation creates a new Operation instance.
func NewOperation(
	tags []string, summary, description string, externalDocs *ExternalDocumentation,
	operationID string, parameters []*ParameterRef, requestBody *RequestBodyRef,
	responses *Responses, callbacks map[string]*CallbackRef, deprecated bool,
	security []SecurityRequirement, servers []*Server,
) *Operation {
	return &Operation{
		tags: tags, summary: summary, description: description, externalDocs: externalDocs,
		operationID: operationID, parameters: parameters, requestBody: requestBody,
		responses: responses, callbacks: callbacks, deprecated: deprecated,
		security: security, servers: servers,
	}
}
