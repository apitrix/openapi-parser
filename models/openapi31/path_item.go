package openapi31

// PathItem describes operations available on a single path.
// https://spec.openapis.org/oas/v3.1.0#path-item-object
type PathItem struct {
	Node // embedded - provides VendorExtensions and Trix

	ref         string
	summary     string
	description string
	get         *Operation
	put         *Operation
	post        *Operation
	delete      *Operation
	options     *Operation
	head        *Operation
	patch       *Operation
	trace       *Operation
	servers     []*Server
	parameters  []*ParameterRef
}

func (p *PathItem) Ref() string                 { return p.ref }
func (p *PathItem) Summary() string             { return p.summary }
func (p *PathItem) Description() string         { return p.description }
func (p *PathItem) Get() *Operation             { return p.get }
func (p *PathItem) Put() *Operation             { return p.put }
func (p *PathItem) Post() *Operation            { return p.post }
func (p *PathItem) Delete() *Operation          { return p.delete }
func (p *PathItem) Options() *Operation         { return p.options }
func (p *PathItem) Head() *Operation            { return p.head }
func (p *PathItem) Patch() *Operation           { return p.patch }
func (p *PathItem) Trace() *Operation           { return p.trace }
func (p *PathItem) Servers() []*Server          { return p.servers }
func (p *PathItem) Parameters() []*ParameterRef { return p.parameters }

// SetProperty sets a named property on the PathItem.
// Used by parsers for post-construction field assignment.
func (p *PathItem) SetProperty(name string, value interface{}) {
	switch name {
	case "$ref":
		p.ref = value.(string)
	case "summary":
		p.summary = value.(string)
	case "description":
		p.description = value.(string)
	case "get":
		p.get = value.(*Operation)
	case "put":
		p.put = value.(*Operation)
	case "post":
		p.post = value.(*Operation)
	case "delete":
		p.delete = value.(*Operation)
	case "options":
		p.options = value.(*Operation)
	case "head":
		p.head = value.(*Operation)
	case "patch":
		p.patch = value.(*Operation)
	case "trace":
		p.trace = value.(*Operation)
	case "servers":
		p.servers = value.([]*Server)
	case "parameters":
		p.parameters = value.([]*ParameterRef)
	}
}

// NewPathItem creates a new PathItem instance.
func NewPathItem() *PathItem {
	return &PathItem{}
}
