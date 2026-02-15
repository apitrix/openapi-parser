package openapi31

import (
	"github.com/apitrix/openapi-parser/models/shared"
	"sort"

	"gopkg.in/yaml.v3"
)

// Paths holds the relative paths to individual endpoints.
// https://spec.openapis.org/oas/v3.1.0#paths-object
type Paths struct {
	ElementBase // embedded - provides VendorExtensions and Trix

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

func (p *Paths) marshalFields() []shared.Field {
	var fields []shared.Field
	if len(p.items) > 0 {
		keys := make([]string, 0, len(p.items))
		for k := range p.items {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fields = append(fields, shared.Field{Key: k, Value: p.items[k]})
		}
	}
	return shared.AppendExtensions(fields, p.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (p *Paths) MarshalFields() []shared.Field { return p.marshalFields() }

func (p *Paths) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(p.marshalFields())
}

func (p *Paths) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(p.marshalFields())
}

var _ yaml.Marshaler = (*Paths)(nil)
