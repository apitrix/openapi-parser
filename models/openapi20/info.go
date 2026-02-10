package openapi20

// Info provides metadata about the API.
// https://swagger.io/specification/v2/#info-object
type Info struct {
	Node // embedded - provides VendorExtensions and Trix

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

// NewInfo creates a new Info instance.
func NewInfo(title, description, termsOfService, version string, contact *Contact, license *License) *Info {
	return &Info{
		title: title, description: description, termsOfService: termsOfService,
		version: version, contact: contact, license: license,
	}
}
