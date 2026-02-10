package openapi31

// SchemaType represents a JSON Schema type field that can be either a single
// string or an array of strings (JSON Schema Draft 2020-12).
type SchemaType struct {
	// Single is set when the type is a single string value (e.g. "string").
	Single string
	// Array is set when the type is an array of strings (e.g. ["string", "null"]).
	Array []string
}

// IsEmpty returns true if no type was specified.
func (t SchemaType) IsEmpty() bool {
	return t.Single == "" && len(t.Array) == 0
}

// Values returns all type values as a slice, whether specified as single or array.
func (t SchemaType) Values() []string {
	if len(t.Array) > 0 {
		return t.Array
	}
	if t.Single != "" {
		return []string{t.Single}
	}
	return nil
}
