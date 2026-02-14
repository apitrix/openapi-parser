package openapi31

import (
	"openapi-parser/models/shared"
	"sort"

	"gopkg.in/yaml.v3"
)

// Callback is a map of possible out-of-band callbacks related to the parent operation.
// https://spec.openapis.org/oas/v3.1.0#callback-object
type Callback struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	paths map[string]*PathItem
}

func (c *Callback) Paths() map[string]*PathItem { return c.paths }

func (c *Callback) SetPaths(paths map[string]*PathItem) error {
	if err := c.Trix.RunHooks("paths", c.paths, paths); err != nil {
		return err
	}
	c.paths = paths
	return nil
}

// NewCallback creates a new Callback instance.
func NewCallback(paths map[string]*PathItem) *Callback {
	return &Callback{paths: paths}
}

func (c *Callback) marshalFields() []shared.Field {
	var fields []shared.Field
	if len(c.paths) > 0 {
		keys := make([]string, 0, len(c.paths))
		for k := range c.paths {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fields = append(fields, shared.Field{Key: k, Value: c.paths[k]})
		}
	}
	return shared.AppendExtensions(fields, c.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (c *Callback) MarshalFields() []shared.Field { return c.marshalFields() }

func (c *Callback) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(c.marshalFields())
}

func (c *Callback) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(c.marshalFields())
}

var _ yaml.Marshaler = (*Callback)(nil)
