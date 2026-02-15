package openapi30

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// PathItem describes operations available on a single path.
// https://spec.openapis.org/oas/v3.0.3#path-item-object
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

func (pi *PathItem) Ref() string                 { return pi.ref }
func (pi *PathItem) Summary() string             { return pi.summary }
func (pi *PathItem) Description() string         { return pi.description }
func (pi *PathItem) Get() *Operation             { return pi.get }
func (pi *PathItem) Put() *Operation             { return pi.put }
func (pi *PathItem) Post() *Operation            { return pi.post }
func (pi *PathItem) Delete() *Operation          { return pi.delete }
func (pi *PathItem) Options() *Operation         { return pi.options }
func (pi *PathItem) Head() *Operation            { return pi.head }
func (pi *PathItem) Patch() *Operation           { return pi.patch }
func (pi *PathItem) Trace() *Operation           { return pi.trace }
func (pi *PathItem) Servers() []*Server          { return pi.servers }
func (pi *PathItem) Parameters() []*RefParameter { return pi.parameters }

func (pi *PathItem) SetRef(ref string) error {
	if err := pi.Trix.RunHooks("$ref", pi.ref, ref); err != nil {
		return err
	}
	pi.ref = ref
	return nil
}
func (pi *PathItem) SetSummary(summary string) error {
	if err := pi.Trix.RunHooks("summary", pi.summary, summary); err != nil {
		return err
	}
	pi.summary = summary
	return nil
}
func (pi *PathItem) SetDescription(description string) error {
	if err := pi.Trix.RunHooks("description", pi.description, description); err != nil {
		return err
	}
	pi.description = description
	return nil
}
func (pi *PathItem) SetGet(get *Operation) error {
	if err := pi.Trix.RunHooks("get", pi.get, get); err != nil {
		return err
	}
	pi.get = get
	return nil
}
func (pi *PathItem) SetPut(put *Operation) error {
	if err := pi.Trix.RunHooks("put", pi.put, put); err != nil {
		return err
	}
	pi.put = put
	return nil
}
func (pi *PathItem) SetPost(post *Operation) error {
	if err := pi.Trix.RunHooks("post", pi.post, post); err != nil {
		return err
	}
	pi.post = post
	return nil
}
func (pi *PathItem) SetDelete(delete *Operation) error {
	if err := pi.Trix.RunHooks("delete", pi.delete, delete); err != nil {
		return err
	}
	pi.delete = delete
	return nil
}
func (pi *PathItem) SetOptions(options *Operation) error {
	if err := pi.Trix.RunHooks("options", pi.options, options); err != nil {
		return err
	}
	pi.options = options
	return nil
}
func (pi *PathItem) SetHead(head *Operation) error {
	if err := pi.Trix.RunHooks("head", pi.head, head); err != nil {
		return err
	}
	pi.head = head
	return nil
}
func (pi *PathItem) SetPatch(patch *Operation) error {
	if err := pi.Trix.RunHooks("patch", pi.patch, patch); err != nil {
		return err
	}
	pi.patch = patch
	return nil
}
func (pi *PathItem) SetTrace(trace *Operation) error {
	if err := pi.Trix.RunHooks("trace", pi.trace, trace); err != nil {
		return err
	}
	pi.trace = trace
	return nil
}
func (pi *PathItem) SetServers(servers []*Server) error {
	if err := pi.Trix.RunHooks("servers", pi.servers, servers); err != nil {
		return err
	}
	pi.servers = servers
	return nil
}
func (pi *PathItem) SetParameters(parameters []*RefParameter) error {
	if err := pi.Trix.RunHooks("parameters", pi.parameters, parameters); err != nil {
		return err
	}
	pi.parameters = parameters
	return nil
}

// NewPathItem creates a new PathItem instance.
func NewPathItem(
	ref, summary, description string,
	get, put, post, del, options, head, patch, trace *Operation,
	servers []*Server, parameters []*RefParameter,
) *PathItem {
	return &PathItem{
		ref: ref, summary: summary, description: description,
		get: get, put: put, post: post, delete: del,
		options: options, head: head, patch: patch, trace: trace,
		servers: servers, parameters: parameters,
	}
}

func (pi *PathItem) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "$ref", Value: pi.ref},
		{Key: "summary", Value: pi.summary},
		{Key: "description", Value: pi.description},
		{Key: "get", Value: pi.get},
		{Key: "put", Value: pi.put},
		{Key: "post", Value: pi.post},
		{Key: "delete", Value: pi.delete},
		{Key: "options", Value: pi.options},
		{Key: "head", Value: pi.head},
		{Key: "patch", Value: pi.patch},
		{Key: "trace", Value: pi.trace},
		{Key: "servers", Value: pi.servers},
		{Key: "parameters", Value: pi.parameters},
	}
	return shared.AppendExtensions(fields, pi.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (pi *PathItem) MarshalFields() []shared.Field { return pi.marshalFields() }

func (pi *PathItem) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(pi.marshalFields())
}

func (pi *PathItem) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(pi.marshalFields())
}

var _ yaml.Marshaler = (*PathItem)(nil)
