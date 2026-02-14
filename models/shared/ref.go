package shared

import (
	"encoding/json"
	"strings"

	"gopkg.in/yaml.v3"
)

// Ref represents a JSON Reference ($ref) or an inline value of type T.
// Used by OpenAPI 2.0 and 3.0 models where $ref has no additional fields.
type Ref[T any] struct {
	ElementBase
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	value    *T
	circular bool
	done     chan struct{} // closed when resolution completes; nil for inline
	err      error         // resolution error, if any
}

// NewRef creates a new Ref with the given $ref string.
func NewRef[T any](ref string) *Ref[T] {
	return &Ref[T]{Ref: ref}
}

// Value returns the resolved value, blocking if background resolution is in progress.
func (r *Ref[T]) Value() *T {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

// Circular returns true if a circular reference was detected, blocking if resolution is in progress.
func (r *Ref[T]) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

// ResolveErr returns the resolution error, blocking if resolution is in progress.
func (r *Ref[T]) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

// RawValue returns the value without blocking. For use by the resolver.
func (r *Ref[T]) RawValue() *T { return r.value }

// RawCircular returns the circular flag without blocking. For use by the resolver.
func (r *Ref[T]) RawCircular() bool { return r.circular }

// SetValue sets the resolved value.
func (r *Ref[T]) SetValue(v *T) { r.value = v }

// SetCircular sets the circular flag.
func (r *Ref[T]) SetCircular(c bool) { r.circular = c }

// SetResolveErr sets the resolution error and adds it to Trix.Errors
// so it appears in ParseResult.Errors.
func (r *Ref[T]) SetResolveErr(err error) {
	r.err = err
	if err != nil {
		r.Trix.Errors = append(r.Trix.Errors, ParseError{
			Path:    refPathFromRef(r.Ref),
			Message: err.Error(),
			Kind:    "resolve_error",
		})
	}
}

// InitDone initializes the done channel, signaling this ref needs async resolution.
func (r *Ref[T]) InitDone() { r.done = make(chan struct{}) }

// MarkDone closes the done channel, unblocking any waiters.
func (r *Ref[T]) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}

// Done returns the done channel for waiting on resolution.
func (r *Ref[T]) Done() <-chan struct{} { return r.done }

// MarshalJSON serializes the ref: if $ref is set, emits {"$ref":"..."};
// if an inline value is set, delegates to the value; otherwise emits null.
func (r *Ref[T]) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return json.Marshal(struct {
			Ref string `json:"$ref"`
		}{Ref: r.Ref})
	}
	if r.value != nil {
		return json.Marshal(r.value)
	}
	return []byte("null"), nil
}

// MarshalYAML serializes the ref to YAML.
func (r *Ref[T]) MarshalYAML() (interface{}, error) {
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
		if m, ok := any(r.value).(yaml.Marshaler); ok {
			return m.MarshalYAML()
		}
	}
	return nil, nil
}

var _ yaml.Marshaler = (*Ref[struct{}])(nil)

// =============================================================================
// RefWithMeta — OpenAPI 3.1 variant with summary and description on $ref
// =============================================================================

// RefWithMeta is like Ref but adds Summary and Description fields
// as permitted by OpenAPI 3.1's extended $ref syntax.
type RefWithMeta[T any] struct {
	ElementBase
	Ref         string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	value       *T
	circular    bool
	done        chan struct{} // closed when resolution completes; nil for inline
	err         error         // resolution error, if any
}

// NewRefWithMeta creates a new RefWithMeta with the given $ref string.
func NewRefWithMeta[T any](ref string) *RefWithMeta[T] {
	return &RefWithMeta[T]{Ref: ref}
}

// Value returns the resolved value, blocking if background resolution is in progress.
func (r *RefWithMeta[T]) Value() *T {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

// Circular returns true if a circular reference was detected, blocking if resolution is in progress.
func (r *RefWithMeta[T]) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

// ResolveErr returns the resolution error, blocking if resolution is in progress.
func (r *RefWithMeta[T]) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

// RawValue returns the value without blocking. For use by the resolver.
func (r *RefWithMeta[T]) RawValue() *T { return r.value }

// RawCircular returns the circular flag without blocking. For use by the resolver.
func (r *RefWithMeta[T]) RawCircular() bool { return r.circular }

// SetValue sets the resolved value.
func (r *RefWithMeta[T]) SetValue(v *T) { r.value = v }

// SetCircular sets the circular flag.
func (r *RefWithMeta[T]) SetCircular(c bool) { r.circular = c }

// SetResolveErr sets the resolution error and adds it to Trix.Errors
// so it appears in ParseResult.Errors.
func (r *RefWithMeta[T]) SetResolveErr(err error) {
	r.err = err
	if err != nil {
		r.Trix.Errors = append(r.Trix.Errors, ParseError{
			Path:    refPathFromRef(r.Ref),
			Message: err.Error(),
			Kind:    "resolve_error",
		})
	}
}

// InitDone initializes the done channel, signaling this ref needs async resolution.
func (r *RefWithMeta[T]) InitDone() { r.done = make(chan struct{}) }

// MarkDone closes the done channel, unblocking any waiters.
func (r *RefWithMeta[T]) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}

// Done returns the done channel for waiting on resolution.
func (r *RefWithMeta[T]) Done() <-chan struct{} { return r.done }

// MarshalJSON serializes the ref: if $ref is set, emits {"$ref":"...", "summary":"...", "description":"..."};
// if an inline value is set, delegates to the value; otherwise emits null.
func (r *RefWithMeta[T]) MarshalJSON() ([]byte, error) {
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
		return json.Marshal(r.value)
	}
	return []byte("null"), nil
}

// MarshalYAML serializes the ref to YAML, including summary and description when present.
func (r *RefWithMeta[T]) MarshalYAML() (interface{}, error) {
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
		if m, ok := any(r.value).(yaml.Marshaler); ok {
			return m.MarshalYAML()
		}
	}
	return nil, nil
}

var _ yaml.Marshaler = (*RefWithMeta[struct{}])(nil)

// refPathFromRef parses a $ref string into a JSON path.
// "#/components/schemas/Pet" -> ["components", "schemas", "Pet"]
// "file.yaml#/definitions/X" -> ["definitions", "X"] (pointer part only)
func refPathFromRef(ref string) []string {
	if ref == "" {
		return nil
	}
	// Split on # to get pointer part
	_, pointer := splitRef(ref)
	if pointer == "" || pointer == "/" {
		return nil
	}
	// Remove leading / and split
	pointer = strings.TrimPrefix(pointer, "/")
	if pointer == "" {
		return nil
	}
	return strings.Split(pointer, "/")
}

func splitRef(ref string) (filePath, pointer string) {
	if idx := strings.Index(ref, "#"); idx >= 0 {
		return ref[:idx], ref[idx+1:]
	}
	return ref, ""
}
