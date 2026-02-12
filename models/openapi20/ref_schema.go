package openapi20

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// SchemaRef represents a reference to a Schema or an inline Schema.
type SchemaRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	value    *Schema
	circular bool
	done     chan struct{} // closed when resolution completes; nil for inline
	err      error         // resolution error, if any
}

// NewSchemaRef creates a new SchemaRef instance.
func NewSchemaRef(ref string) *SchemaRef {
	return &SchemaRef{Ref: ref}
}

// Value returns the resolved Schema, blocking if background resolution is in progress.
func (r *SchemaRef) Value() *Schema {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

// Circular returns true if a circular reference was detected, blocking if resolution is in progress.
func (r *SchemaRef) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

// ResolveErr returns the resolution error, blocking if resolution is in progress.
func (r *SchemaRef) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

// RawValue returns the value without blocking. For use by the resolver.
func (r *SchemaRef) RawValue() *Schema { return r.value }

// RawCircular returns the circular flag without blocking. For use by the resolver.
func (r *SchemaRef) RawCircular() bool { return r.circular }

// SetValue sets the resolved value.
func (r *SchemaRef) SetValue(v *Schema) { r.value = v }

// SetCircular sets the circular flag.
func (r *SchemaRef) SetCircular(c bool) { r.circular = c }

// SetResolveErr sets the resolution error.
func (r *SchemaRef) SetResolveErr(err error) { r.err = err }

// InitDone initializes the done channel, signaling this ref needs async resolution.
func (r *SchemaRef) InitDone() { r.done = make(chan struct{}) }

// MarkDone closes the done channel, unblocking any waiters.
func (r *SchemaRef) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}

// Done returns the done channel for waiting on resolution.
func (r *SchemaRef) Done() <-chan struct{} { return r.done }

func (r *SchemaRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return json.Marshal(struct {
			Ref string `json:"$ref"`
		}{Ref: r.Ref})
	}
	if r.value != nil {
		return r.value.MarshalJSON()
	}
	return []byte("null"), nil
}

func (r *SchemaRef) MarshalYAML() (interface{}, error) {
	if r.Ref != "" {
		return &yaml.Node{
			Kind: yaml.MappingNode,
			Tag:  "!!map",
			Content: []*yaml.Node{
				{Kind: yaml.ScalarNode, Tag: "!!str", Value: "$ref"},
				{Kind: yaml.ScalarNode, Tag: "!!str", Value: r.Ref},
			},
		}, nil
	}
	if r.value != nil {
		return r.value.MarshalYAML()
	}
	return nil, nil
}

var _ yaml.Marshaler = (*SchemaRef)(nil)
