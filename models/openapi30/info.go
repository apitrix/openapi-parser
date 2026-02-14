package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Info provides metadata about the API.
// https://spec.openapis.org/oas/v3.0.3#info-object
type Info struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	title          string
	description    string
	termsOfService string
	contact        *Contact
	license        *License
	version        string
}

// Title returns the title of the API.
func (i *Info) Title() string { return i.title }

// Description returns the description of the API.
func (i *Info) Description() string { return i.description }

// TermsOfService returns the URL to the Terms of Service.
func (i *Info) TermsOfService() string { return i.termsOfService }

// Contact returns the contact information for the API.
func (i *Info) Contact() *Contact { return i.contact }

// License returns the license information for the API.
func (i *Info) License() *License { return i.license }

// Version returns the version of the API.
func (i *Info) Version() string { return i.version }

// NewInfo creates a new Info instance with all spec-defined fields.
func NewInfo(title, description, termsOfService, version string, contact *Contact, license *License) *Info {
	return &Info{
		title:          title,
		description:    description,
		termsOfService: termsOfService,
		version:        version,
		contact:        contact,
		license:        license,
	}
}

func (i *Info) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "title", Value: i.title},
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
