package openapi31

// Operation describes a single API operation on a path.
// https://spec.openapis.org/oas/v3.1.0#operation-object
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

// SetProperty sets a named property on the Operation.
// Used by parsers for post-construction field assignment.
func (o *Operation) SetProperty(name string, value interface{}) {
	switch name {
	case "tags":
		o.tags = value.([]string)
	case "summary":
		o.summary = value.(string)
	case "description":
		o.description = value.(string)
	case "externalDocs":
		o.externalDocs = value.(*ExternalDocumentation)
	case "operationId":
		o.operationID = value.(string)
	case "parameters":
		o.parameters = value.([]*ParameterRef)
	case "requestBody":
		o.requestBody = value.(*RequestBodyRef)
	case "responses":
		o.responses = value.(*Responses)
	case "callbacks":
		o.callbacks = value.(map[string]*CallbackRef)
	case "deprecated":
		o.deprecated = value.(bool)
	case "security":
		o.security = value.([]SecurityRequirement)
	case "servers":
		o.servers = value.([]*Server)
	}
}

// NewOperation creates a new Operation instance.
func NewOperation() *Operation {
	return &Operation{}
}
