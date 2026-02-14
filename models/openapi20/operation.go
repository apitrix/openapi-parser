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
	parameters   []*RefParameter
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
func (o *Operation) Parameters() []*RefParameter     { return o.parameters }
func (o *Operation) Responses() *Responses           { return o.responses }
func (o *Operation) Schemes() []string               { return o.schemes }
func (o *Operation) Deprecated() bool                { return o.deprecated }
func (o *Operation) Security() []SecurityRequirement { return o.security }

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
func (o *Operation) SetExternalDocs(externalDocs *ExternalDocs) error {
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
func (o *Operation) SetConsumes(consumes []string) error {
	if err := o.Trix.RunHooks("consumes", o.consumes, consumes); err != nil {
		return err
	}
	o.consumes = consumes
	return nil
}
func (o *Operation) SetProduces(produces []string) error {
	if err := o.Trix.RunHooks("produces", o.produces, produces); err != nil {
		return err
	}
	o.produces = produces
	return nil
}
func (o *Operation) SetParameters(parameters []*RefParameter) error {
	if err := o.Trix.RunHooks("parameters", o.parameters, parameters); err != nil {
		return err
	}
	o.parameters = parameters
	return nil
}
func (o *Operation) SetResponses(responses *Responses) error {
	if err := o.Trix.RunHooks("responses", o.responses, responses); err != nil {
		return err
	}
	o.responses = responses
	return nil
}
func (o *Operation) SetSchemes(schemes []string) error {
	if err := o.Trix.RunHooks("schemes", o.schemes, schemes); err != nil {
		return err
	}
	o.schemes = schemes
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

// NewOperation creates a new Operation instance.
func NewOperation(
	tags []string, summary, description string, externalDocs *ExternalDocs,
	operationID string, consumes, produces []string,
	parameters []*RefParameter, responses *Responses,
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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (o *Operation) MarshalFields() []shared.Field { return o.marshalFields() }

func (o *Operation) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(o.marshalFields())
}

func (o *Operation) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(o.marshalFields())
}

var _ yaml.Marshaler = (*Operation)(nil)
