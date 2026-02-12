package openapi20

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// ResponseRef represents a reference to a Response or an inline Response.
type ResponseRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	value    *Response
	circular bool
	done     chan struct{} // closed when resolution completes; nil for inline
	err      error         // resolution error, if any
}

// NewResponseRef creates a new ResponseRef instance.
func NewResponseRef(ref string) *ResponseRef {
	return &ResponseRef{Ref: ref}
}

// Value returns the resolved Response, blocking if background resolution is in progress.
func (r *ResponseRef) Value() *Response {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

// Circular returns true if a circular reference was detected, blocking if resolution is in progress.
func (r *ResponseRef) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

// ResolveErr returns the resolution error, blocking if resolution is in progress.
func (r *ResponseRef) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

// RawValue returns the value without blocking. For use by the resolver.
func (r *ResponseRef) RawValue() *Response { return r.value }

// RawCircular returns the circular flag without blocking. For use by the resolver.
func (r *ResponseRef) RawCircular() bool { return r.circular }

// SetValue sets the resolved value.
func (r *ResponseRef) SetValue(v *Response) { r.value = v }

// SetCircular sets the circular flag.
func (r *ResponseRef) SetCircular(c bool) { r.circular = c }

// SetResolveErr sets the resolution error.
func (r *ResponseRef) SetResolveErr(err error) { r.err = err }

// InitDone initializes the done channel, signaling this ref needs async resolution.
func (r *ResponseRef) InitDone() { r.done = make(chan struct{}) }

// MarkDone closes the done channel, unblocking any waiters.
func (r *ResponseRef) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}

// Done returns the done channel for waiting on resolution.
func (r *ResponseRef) Done() <-chan struct{} { return r.done }

func (r *ResponseRef) MarshalJSON() ([]byte, error) {
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

func (r *ResponseRef) MarshalYAML() (interface{}, error) {
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

var _ yaml.Marshaler = (*ResponseRef)(nil)
