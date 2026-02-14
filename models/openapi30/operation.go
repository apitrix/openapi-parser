package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Operation describes a single API operation on a path.
// https://spec.openapis.org/oas/v3.0.3#operation-object
type Operation struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	tags         []string
	summary      string
	description  string
	externalDocs *ExternalDocumentation
	operationID  string
	parameters   []*shared.Ref[Parameter]
	requestBody  *shared.Ref[RequestBody]
	responses    *Responses
	callbacks    map[string]*shared.Ref[Callback]
	deprecated   bool
	security     []SecurityRequirement
	servers      []*Server
}

func (o *Operation) Tags() []string                              { return o.tags }
func (o *Operation) Summary() string                             { return o.summary }
func (o *Operation) Description() string                         { return o.description }
func (o *Operation) ExternalDocs() *ExternalDocumentation        { return o.externalDocs }
func (o *Operation) OperationID() string                         { return o.operationID }
func (o *Operation) Parameters() []*shared.Ref[Parameter]        { return o.parameters }
func (o *Operation) RequestBody() *shared.Ref[RequestBody]       { return o.requestBody }
func (o *Operation) Responses() *Responses                       { return o.responses }
func (o *Operation) Callbacks() map[string]*shared.Ref[Callback] { return o.callbacks }
func (o *Operation) Deprecated() bool                            { return o.deprecated }
func (o *Operation) Security() []SecurityRequirement             { return o.security }
func (o *Operation) Servers() []*Server                          { return o.servers }

// NewOperation creates a new Operation instance.
func NewOperation(
	tags []string, summary, description string, externalDocs *ExternalDocumentation,
	operationID string, parameters []*shared.Ref[Parameter], requestBody *shared.Ref[RequestBody],
	responses *Responses, callbacks map[string]*shared.Ref[Callback], deprecated bool,
	security []SecurityRequirement, servers []*Server,
) *Operation {
	return &Operation{
		tags: tags, summary: summary, description: description, externalDocs: externalDocs,
		operationID: operationID, parameters: parameters, requestBody: requestBody,
		responses: responses, callbacks: callbacks, deprecated: deprecated,
		security: security, servers: servers,
	}
}

func (o *Operation) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "tags", Value: o.tags},
		{Key: "summary", Value: o.summary},
		{Key: "description", Value: o.description},
		{Key: "externalDocs", Value: o.externalDocs},
		{Key: "operationId", Value: o.operationID},
		{Key: "parameters", Value: o.parameters},
		{Key: "requestBody", Value: o.requestBody},
		{Key: "responses", Value: o.responses},
		{Key: "callbacks", Value: o.callbacks},
		{Key: "deprecated", Value: o.deprecated},
		{Key: "security", Value: o.security},
		{Key: "servers", Value: o.servers},
	}
	return shared.AppendExtensions(fields, o.VendorExtensions)
}

func (o *Operation) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(o.marshalFields())
}

func (o *Operation) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(o.marshalFields())
}

var _ yaml.Marshaler = (*Operation)(nil)
