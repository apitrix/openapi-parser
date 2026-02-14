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
	parameters  []*RefParameter
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
func (p *PathItem) Parameters() []*RefParameter { return p.parameters }

func (p *PathItem) SetRef(ref string) error {
	if err := p.Trix.RunHooks("$ref", p.ref, ref); err != nil {
		return err
	}
	p.ref = ref
	return nil
}
func (p *PathItem) SetSummary(summary string) error {
	if err := p.Trix.RunHooks("summary", p.summary, summary); err != nil {
		return err
	}
	p.summary = summary
	return nil
}
func (p *PathItem) SetDescription(description string) error {
	if err := p.Trix.RunHooks("description", p.description, description); err != nil {
		return err
	}
	p.description = description
	return nil
}
func (p *PathItem) SetGet(get *Operation) error {
	if err := p.Trix.RunHooks("get", p.get, get); err != nil {
		return err
	}
	p.get = get
	return nil
}
func (p *PathItem) SetPut(put *Operation) error {
	if err := p.Trix.RunHooks("put", p.put, put); err != nil {
		return err
	}
	p.put = put
	return nil
}
func (p *PathItem) SetPost(post *Operation) error {
	if err := p.Trix.RunHooks("post", p.post, post); err != nil {
		return err
	}
	p.post = post
	return nil
}
func (p *PathItem) SetDelete(delete *Operation) error {
	if err := p.Trix.RunHooks("delete", p.delete, delete); err != nil {
		return err
	}
	p.delete = delete
	return nil
}
func (p *PathItem) SetOptions(options *Operation) error {
	if err := p.Trix.RunHooks("options", p.options, options); err != nil {
		return err
	}
	p.options = options
	return nil
}
func (p *PathItem) SetHead(head *Operation) error {
	if err := p.Trix.RunHooks("head", p.head, head); err != nil {
		return err
	}
	p.head = head
	return nil
}
func (p *PathItem) SetPatch(patch *Operation) error {
	if err := p.Trix.RunHooks("patch", p.patch, patch); err != nil {
		return err
	}
	p.patch = patch
	return nil
}
func (p *PathItem) SetTrace(trace *Operation) error {
	if err := p.Trix.RunHooks("trace", p.trace, trace); err != nil {
		return err
	}
	p.trace = trace
	return nil
}
func (p *PathItem) SetServers(servers []*Server) error {
	if err := p.Trix.RunHooks("servers", p.servers, servers); err != nil {
		return err
	}
	p.servers = servers
	return nil
}
func (p *PathItem) SetParameters(parameters []*RefParameter) error {
	if err := p.Trix.RunHooks("parameters", p.parameters, parameters); err != nil {
		return err
	}
	p.parameters = parameters
	return nil
}

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
		p.parameters = value.([]*RefParameter)
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
