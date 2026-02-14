package openapi20

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// PathItem describes operations available on a single path.
// https://swagger.io/specification/v2/#path-item-object
type PathItem struct {
	Node // embedded - provides VendorExtensions and Trix

	ref        string
	get        *Operation
	put        *Operation
	post       *Operation
	delete     *Operation
	options    *Operation
	head       *Operation
	patch      *Operation
	parameters []*shared.Ref[Parameter]
}

func (pi *PathItem) Ref() string                 { return pi.ref }
func (pi *PathItem) Get() *Operation             { return pi.get }
func (pi *PathItem) Put() *Operation             { return pi.put }
func (pi *PathItem) Post() *Operation            { return pi.post }
func (pi *PathItem) Delete() *Operation          { return pi.delete }
func (pi *PathItem) Options() *Operation         { return pi.options }
func (pi *PathItem) Head() *Operation            { return pi.head }
func (pi *PathItem) Patch() *Operation           { return pi.patch }
func (pi *PathItem) Parameters() []*shared.Ref[Parameter] { return pi.parameters }

// NewPathItem creates a new PathItem instance.
func NewPathItem(
	ref string,
	get, put, post, del, options, head, patch *Operation,
	parameters []*shared.Ref[Parameter],
) *PathItem {
	return &PathItem{
		ref: ref,
		get: get, put: put, post: post, delete: del,
		options: options, head: head, patch: patch,
		parameters: parameters,
	}
}

func (pi *PathItem) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "$ref", Value: pi.ref},
		{Key: "get", Value: pi.get},
		{Key: "put", Value: pi.put},
		{Key: "post", Value: pi.post},
		{Key: "delete", Value: pi.delete},
		{Key: "options", Value: pi.options},
		{Key: "head", Value: pi.head},
		{Key: "patch", Value: pi.patch},
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
