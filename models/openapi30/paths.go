package openapi30

import (
	"openapi-parser/models/shared"
	"sort"

	"gopkg.in/yaml.v3"
)

// Paths holds the relative paths to individual endpoints.
// https://spec.openapis.org/oas/v3.0.3#paths-object
type Paths struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	// items maps paths to their definitions.
	items map[string]*PathItem
}

func (p *Paths) Items() map[string]*PathItem { return p.items }

func (p *Paths) SetItems(items map[string]*PathItem) error {
	if err := p.Trix.RunHooks("items", p.items, items); err != nil {
		return err
	}
	p.items = items
	return nil
}

// NewPaths creates a new Paths instance with the given items map.
func NewPaths(items map[string]*PathItem) *Paths {
	return &Paths{items: items}
}

// MarshalJSON serializes Paths as a flat JSON object with path keys.
func (p *Paths) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(p.marshalFields())
}

// MarshalYAML serializes Paths as a flat YAML mapping with path keys.
func (p *Paths) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(p.marshalFields())
}

func (p *Paths) marshalFields() []shared.Field {
	keys := make([]string, 0, len(p.items))
	for k := range p.items {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fields := make([]shared.Field, 0, len(keys)+len(p.VendorExtensions))
	for _, k := range keys {
		fields = append(fields, shared.Field{Key: k, Value: p.items[k]})
	}
	return shared.AppendExtensions(fields, p.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (p *Paths) MarshalFields() []shared.Field { return p.marshalFields() }

var _ yaml.Marshaler = (*Paths)(nil)
