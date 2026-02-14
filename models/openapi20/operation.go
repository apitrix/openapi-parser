package openapi20

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Operation describes a single API operation on a path.
// https://swagger.io/specification/v2/#operation-object
type Operation struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	tags         []string
	summary      string
	description  string
	externalDocs *ExternalDocs
	operationID  string
	consumes     []string
	produces     []string
	parameters   []*shared.Ref[Parameter]
	responses    *Responses
	schemes      []string
	deprecated   bool
	security     []SecurityRequirement
}

func (o *Operation) Tags() []string                       { return o.tags }
func (o *Operation) Summary() string                      { return o.summary }
func (o *Operation) Description() string                  { return o.description }
func (o *Operation) ExternalDocs() *ExternalDocs          { return o.externalDocs }
func (o *Operation) OperationID() string                  { return o.operationID }
func (o *Operation) Consumes() []string                   { return o.consumes }
func (o *Operation) Produces() []string                   { return o.produces }
func (o *Operation) Parameters() []*shared.Ref[Parameter] { return o.parameters }
func (o *Operation) Responses() *Responses                { return o.responses }
func (o *Operation) Schemes() []string                    { return o.schemes }
func (o *Operation) Deprecated() bool                     { return o.deprecated }
func (o *Operation) Security() []SecurityRequirement      { return o.security }

// NewOperation creates a new Operation instance.
func NewOperation(
	tags []string, summary, description string, externalDocs *ExternalDocs,
	operationID string, consumes, produces []string,
	parameters []*shared.Ref[Parameter], responses *Responses,
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

func (o *Operation) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "tags", Value: o.tags},
		{Key: "summary", Value: o.summary},
		{Key: "description", Value: o.description},
		{Key: "externalDocs", Value: o.externalDocs},
		{Key: "operationId", Value: o.operationID},
		{Key: "consumes", Value: o.consumes},
		{Key: "produces", Value: o.produces},
		{Key: "parameters", Value: o.parameters},
		{Key: "responses", Value: o.responses},
		{Key: "schemes", Value: o.schemes},
		{Key: "deprecated", Value: o.deprecated},
		{Key: "security", Value: o.security},
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
