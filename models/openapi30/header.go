package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Header represents a Header Object.
// https://spec.openapis.org/oas/v3.0.3#header-object
type Header struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	description     string
	required        bool
	deprecated      bool
	allowEmptyValue bool
	style           string
	explode         *bool
	allowReserved   bool
	schema          *RefSchema
	example         interface{}
	examples        map[string]*RefExample
	content         map[string]*MediaType
}

func (h *Header) Description() string              { return h.description }
func (h *Header) Required() bool                   { return h.required }
func (h *Header) Deprecated() bool                 { return h.deprecated }
func (h *Header) AllowEmptyValue() bool            { return h.allowEmptyValue }
func (h *Header) Style() string                    { return h.style }
func (h *Header) Explode() *bool                   { return h.explode }
func (h *Header) AllowReserved() bool              { return h.allowReserved }
func (h *Header) Schema() *RefSchema               { return h.schema }
func (h *Header) Example() interface{}             { return h.example }
func (h *Header) Examples() map[string]*RefExample { return h.examples }
func (h *Header) Content() map[string]*MediaType   { return h.content }

func (h *Header) SetDescription(description string) error {
	if err := h.Trix.RunHooks("description", h.description, description); err != nil {
		return err
	}
	h.description = description
	return nil
}
func (h *Header) SetRequired(required bool) error {
	if err := h.Trix.RunHooks("required", h.required, required); err != nil {
		return err
	}
	h.required = required
	return nil
}
func (h *Header) SetDeprecated(deprecated bool) error {
	if err := h.Trix.RunHooks("deprecated", h.deprecated, deprecated); err != nil {
		return err
	}
	h.deprecated = deprecated
	return nil
}
func (h *Header) SetAllowEmptyValue(allowEmptyValue bool) error {
	if err := h.Trix.RunHooks("allowEmptyValue", h.allowEmptyValue, allowEmptyValue); err != nil {
		return err
	}
	h.allowEmptyValue = allowEmptyValue
	return nil
}
func (h *Header) SetStyle(style string) error {
	if err := h.Trix.RunHooks("style", h.style, style); err != nil {
		return err
	}
	h.style = style
	return nil
}
func (h *Header) SetExplode(explode *bool) error {
	if err := h.Trix.RunHooks("explode", h.explode, explode); err != nil {
		return err
	}
	h.explode = explode
	return nil
}
func (h *Header) SetAllowReserved(allowReserved bool) error {
	if err := h.Trix.RunHooks("allowReserved", h.allowReserved, allowReserved); err != nil {
		return err
	}
	h.allowReserved = allowReserved
	return nil
}
func (h *Header) SetSchema(schema *RefSchema) error {
	if err := h.Trix.RunHooks("schema", h.schema, schema); err != nil {
		return err
	}
	h.schema = schema
	return nil
}
func (h *Header) SetExample(example interface{}) error {
	if err := h.Trix.RunHooks("example", h.example, example); err != nil {
		return err
	}
	h.example = example
	return nil
}
func (h *Header) SetExamples(examples map[string]*RefExample) error {
	if err := h.Trix.RunHooks("examples", h.examples, examples); err != nil {
		return err
	}
	h.examples = examples
	return nil
}
func (h *Header) SetContent(content map[string]*MediaType) error {
	if err := h.Trix.RunHooks("content", h.content, content); err != nil {
		return err
	}
	h.content = content
	return nil
}

// NewHeader creates a new Header instance.
func NewHeader(
	description string, required, deprecated, allowEmptyValue bool,
	style string, explode *bool, allowReserved bool,
	schema *RefSchema, example interface{}, examples map[string]*RefExample,
	content map[string]*MediaType,
) *Header {
	return &Header{
		description: description, required: required, deprecated: deprecated,
		allowEmptyValue: allowEmptyValue, style: style, explode: explode,
		allowReserved: allowReserved, schema: schema, example: example,
		examples: examples, content: content,
	}
}

func (h *Header) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "description", Value: h.description},
		{Key: "required", Value: h.required},
		{Key: "deprecated", Value: h.deprecated},
		{Key: "allowEmptyValue", Value: h.allowEmptyValue},
		{Key: "style", Value: h.style},
		{Key: "explode", Value: h.explode},
		{Key: "allowReserved", Value: h.allowReserved},
		{Key: "schema", Value: h.schema},
		{Key: "example", Value: h.example},
		{Key: "examples", Value: h.examples},
		{Key: "content", Value: h.content},
	}
	return shared.AppendExtensions(fields, h.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (h *Header) MarshalFields() []shared.Field { return h.marshalFields() }

func (h *Header) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(h.marshalFields())
}

func (h *Header) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(h.marshalFields())
}

var _ yaml.Marshaler = (*Header)(nil)
