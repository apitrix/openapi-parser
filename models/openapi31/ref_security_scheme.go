package openapi31

import (
"encoding/json"

"gopkg.in/yaml.v3"
)

// SecuritySchemeRef represents a reference to a SecurityScheme or an inline SecurityScheme.
type SecuritySchemeRef struct {
Node                // embedded - provides VendorExtensions and Trix
Ref         string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
Summary     string  `json:"summary,omitempty" yaml:"summary,omitempty"`
Description string  `json:"description,omitempty" yaml:"description,omitempty"`
value    *SecurityScheme
circular bool
done     chan struct{} // closed when resolution completes; nil for inline
err      error         // resolution error, if any
}

// NewSecuritySchemeRef creates a new SecuritySchemeRef instance.
func NewSecuritySchemeRef(ref string) *SecuritySchemeRef {
return &SecuritySchemeRef{Ref: ref}
}

// Value returns the resolved SecurityScheme, blocking if background resolution is in progress.
func (r *SecuritySchemeRef) Value() *SecurityScheme {
if r.done != nil {
<-r.done
}
return r.value
}

// Circular returns true if a circular reference was detected, blocking if resolution is in progress.
func (r *SecuritySchemeRef) Circular() bool {
if r.done != nil {
<-r.done
}
return r.circular
}

// ResolveErr returns the resolution error, blocking if resolution is in progress.
func (r *SecuritySchemeRef) ResolveErr() error {
if r.done != nil {
<-r.done
}
return r.err
}

// RawValue returns the value without blocking. For use by the resolver.
func (r *SecuritySchemeRef) RawValue() *SecurityScheme { return r.value }

// RawCircular returns the circular flag without blocking. For use by the resolver.
func (r *SecuritySchemeRef) RawCircular() bool { return r.circular }

// SetValue sets the resolved value.
func (r *SecuritySchemeRef) SetValue(v *SecurityScheme) { r.value = v }

// SetCircular sets the circular flag.
func (r *SecuritySchemeRef) SetCircular(c bool) { r.circular = c }

// SetResolveErr sets the resolution error.
func (r *SecuritySchemeRef) SetResolveErr(err error) { r.err = err }

// InitDone initializes the done channel, signaling this ref needs async resolution.
func (r *SecuritySchemeRef) InitDone() { r.done = make(chan struct{}) }

// MarkDone closes the done channel, unblocking any waiters.
func (r *SecuritySchemeRef) MarkDone() {
if r.done != nil {
close(r.done)
}
}

// Done returns the done channel for waiting on resolution.
func (r *SecuritySchemeRef) Done() <-chan struct{} { return r.done }

func (r *SecuritySchemeRef) MarshalJSON() ([]byte, error) {
if r.Ref != "" {
m := map[string]string{"$ref": r.Ref}
if r.Summary != "" {
m["summary"] = r.Summary
}
if r.Description != "" {
m["description"] = r.Description
}
return json.Marshal(m)
}
if r.value != nil {
return r.value.MarshalJSON()
}
return []byte("null"), nil
}

func (r *SecuritySchemeRef) MarshalYAML() (interface{}, error) {
if r.Ref != "" {
content := []*yaml.Node{
{Kind: yaml.ScalarNode, Tag: "!!str", Value: "$ref"},
{Kind: yaml.ScalarNode, Tag: "!!str", Value: r.Ref},
}
if r.Summary != "" {
content = append(content,
&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: "summary"},
&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: r.Summary},
)
}
if r.Description != "" {
content = append(content,
&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: "description"},
&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: r.Description},
)
}
return &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map", Content: content}, nil
}
if r.value != nil {
return r.value.MarshalYAML()
}
return nil, nil
}

var _ yaml.Marshaler = (*SecuritySchemeRef)(nil)