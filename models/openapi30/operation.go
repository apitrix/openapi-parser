package openapi30

import (
	"github.com/apitrix/openapi-parser/models/shared"

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
	parameters   []*RefParameter
	requestBody  *RefRequestBody
	responses    *Responses
	callbacks    map[string]*RefCallback
	deprecated   bool
	security     []SecurityRequirement
	servers      []*Server
}

func (o *Operation) Tags() []string                       { return o.tags }
func (o *Operation) Summary() string                      { return o.summary }
func (o *Operation) Description() string                  { return o.description }
func (o *Operation) ExternalDocs() *ExternalDocumentation { return o.externalDocs }
func (o *Operation) OperationID() string                  { return o.operationID }
func (o *Operation) Parameters() []*RefParameter          { return o.parameters }
func (o *Operation) RequestBody() *RefRequestBody         { return o.requestBody }
func (o *Operation) Responses() *Responses                { return o.responses }
func (o *Operation) Callbacks() map[string]*RefCallback   { return o.callbacks }
func (o *Operation) Deprecated() bool                     { return o.deprecated }
func (o *Operation) Security() []SecurityRequirement      { return o.security }
func (o *Operation) Servers() []*Server                   { return o.servers }

func (o *Operation) SetTags(tags []string) error {
	if err := o.Trix.RunHooks("tags", o.tags, tags); err != nil {
		return err
	}
	o.tags = tags
	return nil
}
func (o *Operation) SetSummary(summary string) error {
	if err := o.Trix.RunHooks("summary", o.summary, summary); err != nil {
		return err
	}
	o.summary = summary
	return nil
}
func (o *Operation) SetDescription(description string) error {
	if err := o.Trix.RunHooks("description", o.description, description); err != nil {
		return err
	}
	o.description = description
	return nil
}
func (o *Operation) SetExternalDocs(externalDocs *ExternalDocumentation) error {
	if err := o.Trix.RunHooks("externalDocs", o.externalDocs, externalDocs); err != nil {
		return err
	}
	o.externalDocs = externalDocs
	return nil
}
func (o *Operation) SetOperationID(operationID string) error {
	if err := o.Trix.RunHooks("operationId", o.operationID, operationID); err != nil {
		return err
	}
	o.operationID = operationID
	return nil
}
func (o *Operation) SetParameters(parameters []*RefParameter) error {
	if err := o.Trix.RunHooks("parameters", o.parameters, parameters); err != nil {
		return err
	}
	o.parameters = parameters
	return nil
}
func (o *Operation) SetRequestBody(requestBody *RefRequestBody) error {
	if err := o.Trix.RunHooks("requestBody", o.requestBody, requestBody); err != nil {
		return err
	}
	o.requestBody = requestBody
	return nil
}
func (o *Operation) SetResponses(responses *Responses) error {
	if err := o.Trix.RunHooks("responses", o.responses, responses); err != nil {
		return err
	}
	o.responses = responses
	return nil
}
func (o *Operation) SetCallbacks(callbacks map[string]*RefCallback) error {
	if err := o.Trix.RunHooks("callbacks", o.callbacks, callbacks); err != nil {
		return err
	}
	o.callbacks = callbacks
	return nil
}
func (o *Operation) SetDeprecated(deprecated bool) error {
	if err := o.Trix.RunHooks("deprecated", o.deprecated, deprecated); err != nil {
		return err
	}
	o.deprecated = deprecated
	return nil
}
func (o *Operation) SetSecurity(security []SecurityRequirement) error {
	if err := o.Trix.RunHooks("security", o.security, security); err != nil {
		return err
	}
	o.security = security
	return nil
}
func (o *Operation) SetServers(servers []*Server) error {
	if err := o.Trix.RunHooks("servers", o.servers, servers); err != nil {
		return err
	}
	o.servers = servers
	return nil
}

// NewOperation creates a new Operation instance.
func NewOperation(
	tags []string, summary, description string, externalDocs *ExternalDocumentation,
	operationID string, parameters []*RefParameter, requestBody *RefRequestBody,
	responses *Responses, callbacks map[string]*RefCallback, deprecated bool,
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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (o *Operation) MarshalFields() []shared.Field { return o.marshalFields() }

func (o *Operation) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(o.marshalFields())
}

func (o *Operation) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(o.marshalFields())
}

var _ yaml.Marshaler = (*Operation)(nil)
