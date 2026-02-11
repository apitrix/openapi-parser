package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Info provides metadata about the API.
// https://spec.openapis.org/oas/v3.1.0#info-object
type Info struct {
	Node // embedded - provides VendorExtensions and Trix

	title          string
	summary        string
	description    string
	termsOfService string
	contact        *Contact
	license        *License
	version        string
}

func (i *Info) Title() string          { return i.title }
func (i *Info) Summary() string        { return i.summary }
func (i *Info) Description() string    { return i.description }
func (i *Info) TermsOfService() string { return i.termsOfService }
func (i *Info) Contact() *Contact      { return i.contact }
func (i *Info) License() *License      { return i.license }
func (i *Info) Version() string        { return i.version }

// NewInfo creates a new Info instance with all spec-defined fields.
func NewInfo(title, summary, description, termsOfService, version string, contact *Contact, license *License) *Info {
	return &Info{
		title: title, summary: summary, description: description,
		termsOfService: termsOfService, version: version,
		contact: contact, license: license,
	}
}

func (i *Info) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "title", Value: i.title},
		{Key: "summary", Value: i.summary},
		{Key: "description", Value: i.description},
		{Key: "termsOfService", Value: i.termsOfService},
		{Key: "contact", Value: i.contact},
		{Key: "license", Value: i.license},
		{Key: "version", Value: i.version},
	}
	return shared.AppendExtensions(fields, i.VendorExtensions)
}

func (i *Info) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(i.marshalFields())
}

func (i *Info) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(i.marshalFields())
}

var _ yaml.Marshaler = (*Info)(nil)
