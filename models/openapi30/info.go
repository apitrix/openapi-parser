package openapi30

// Info provides metadata about the API.
// https://spec.openapis.org/oas/v3.0.3#info-object
type Info struct {
	Node // embedded - provides VendorExtensions and Trix

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
