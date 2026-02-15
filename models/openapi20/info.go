package openapi20

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Info provides metadata about the API.
// https://swagger.io/specification/v2/#info-object
type Info struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	title          string
	description    string
	termsOfService string
	contact        *Contact
	license        *License
	version        string
}

func (i *Info) Title() string          { return i.title }
func (i *Info) Description() string    { return i.description }
func (i *Info) TermsOfService() string { return i.termsOfService }
func (i *Info) Contact() *Contact      { return i.contact }
func (i *Info) License() *License      { return i.license }
func (i *Info) Version() string        { return i.version }

func (i *Info) SetTitle(title string) error {
	if err := i.Trix.RunHooks("title", i.title, title); err != nil {
		return err
	}
	i.title = title
	return nil
}
func (i *Info) SetDescription(description string) error {
	if err := i.Trix.RunHooks("description", i.description, description); err != nil {
		return err
	}
	i.description = description
	return nil
}
func (i *Info) SetTermsOfService(termsOfService string) error {
	if err := i.Trix.RunHooks("termsOfService", i.termsOfService, termsOfService); err != nil {
		return err
	}
	i.termsOfService = termsOfService
	return nil
}
func (i *Info) SetContact(contact *Contact) error {
	if err := i.Trix.RunHooks("contact", i.contact, contact); err != nil {
		return err
	}
	i.contact = contact
	return nil
}
func (i *Info) SetLicense(license *License) error {
	if err := i.Trix.RunHooks("license", i.license, license); err != nil {
		return err
	}
	i.license = license
	return nil
}
func (i *Info) SetVersion(version string) error {
	if err := i.Trix.RunHooks("version", i.version, version); err != nil {
		return err
	}
	i.version = version
	return nil
}

// NewInfo creates a new Info instance.
func NewInfo(title, description, termsOfService, version string, contact *Contact, license *License) *Info {
	return &Info{
		title: title, description: description, termsOfService: termsOfService,
		version: version, contact: contact, license: license,
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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (i *Info) MarshalFields() []shared.Field { return i.marshalFields() }

func (i *Info) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(i.marshalFields())
}

func (i *Info) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(i.marshalFields())
}

var _ yaml.Marshaler = (*Info)(nil)
