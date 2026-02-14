package openapi30

import (
	"openapi-parser/models/shared"

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
	parameters  []*shared.Ref[Parameter]
}

func (pi *PathItem) Ref() string                          { return pi.ref }
func (pi *PathItem) Summary() string                      { return pi.summary }
func (pi *PathItem) Description() string                  { return pi.description }
func (pi *PathItem) Get() *Operation                      { return pi.get }
func (pi *PathItem) Put() *Operation                      { return pi.put }
func (pi *PathItem) Post() *Operation                     { return pi.post }
func (pi *PathItem) Delete() *Operation                   { return pi.delete }
func (pi *PathItem) Options() *Operation                  { return pi.options }
func (pi *PathItem) Head() *Operation                     { return pi.head }
func (pi *PathItem) Patch() *Operation                    { return pi.patch }
func (pi *PathItem) Trace() *Operation                    { return pi.trace }
func (pi *PathItem) Servers() []*Server                   { return pi.servers }
func (pi *PathItem) Parameters() []*shared.Ref[Parameter] { return pi.parameters }

// NewPathItem creates a new PathItem instance.
func NewPathItem(
	ref, summary, description string,
	get, put, post, del, options, head, patch, trace *Operation,
	servers []*Server, parameters []*shared.Ref[Parameter],
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

func (pi *PathItem) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(pi.marshalFields())
}

func (pi *PathItem) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(pi.marshalFields())
}

var _ yaml.Marshaler = (*PathItem)(nil)
