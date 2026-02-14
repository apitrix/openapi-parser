package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// PathItem describes operations available on a single path.
// https://spec.openapis.org/oas/v3.1.0#path-item-object
type PathItem struct {
	ElementBase // embedded - provides VendorExtensions and Trix

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
	parameters  []*shared.RefWithMeta[Parameter]
}

func (p *PathItem) Ref() string                                  { return p.ref }
func (p *PathItem) Summary() string                              { return p.summary }
func (p *PathItem) Description() string                          { return p.description }
func (p *PathItem) Get() *Operation                              { return p.get }
func (p *PathItem) Put() *Operation                              { return p.put }
func (p *PathItem) Post() *Operation                             { return p.post }
func (p *PathItem) Delete() *Operation                           { return p.delete }
func (p *PathItem) Options() *Operation                          { return p.options }
func (p *PathItem) Head() *Operation                             { return p.head }
func (p *PathItem) Patch() *Operation                            { return p.patch }
func (p *PathItem) Trace() *Operation                            { return p.trace }
func (p *PathItem) Servers() []*Server                           { return p.servers }
func (p *PathItem) Parameters() []*shared.RefWithMeta[Parameter] { return p.parameters }

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
		p.parameters = value.([]*shared.RefWithMeta[Parameter])
	}
}

// NewPathItem creates a new PathItem instance.
func NewPathItem() *PathItem {
	return &PathItem{}
}

func (p *PathItem) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "$ref", Value: p.ref},
		{Key: "summary", Value: p.summary},
		{Key: "description", Value: p.description},
		{Key: "get", Value: p.get},
		{Key: "put", Value: p.put},
		{Key: "post", Value: p.post},
		{Key: "delete", Value: p.delete},
		{Key: "options", Value: p.options},
		{Key: "head", Value: p.head},
		{Key: "patch", Value: p.patch},
		{Key: "trace", Value: p.trace},
		{Key: "servers", Value: p.servers},
		{Key: "parameters", Value: p.parameters},
	}
	return shared.AppendExtensions(fields, p.VendorExtensions)
}

func (p *PathItem) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(p.marshalFields())
}

func (p *PathItem) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(p.marshalFields())
}

var _ yaml.Marshaler = (*PathItem)(nil)
