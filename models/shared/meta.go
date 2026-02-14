package shared

// Location represents a position in the source file.
type Location struct {
	Line   int `json:"-" yaml:"-"` // 1-based line number
	Column int `json:"-" yaml:"-"` // 1-based column number
}

// NodeSource contains source location and raw parsed data for a node.
type NodeSource struct {
	Start Location    `json:"-" yaml:"-"` // Start position
	End   Location    `json:"-" yaml:"-"` // End position
	Raw   interface{} `json:"-" yaml:"-"` // Raw parsed data (map/slice/scalar)
}

// ParseError represents a parsing error associated with a specific node.
type ParseError struct {
	Message string   `json:"-" yaml:"-"` // Human-readable error message
	Path    []string `json:"-" yaml:"-"` // JSON path where the error occurred
	Kind    string   `json:"-" yaml:"-"` // Error kind: "error" or "unknown_field"
}

// HookFunc is called before a field is set.
// Return a non-nil error to reject the change.
type HookFunc func(field string, oldVal, newVal interface{}) error

// Trix contains all library-level metadata and functionality.
// Everything under Trix is provided by the apitrix library,
// not part of the OpenAPI specification itself.
type Trix struct {
	Source NodeSource            `json:"-" yaml:"-"` // Source location info
	Errors []ParseError          `json:"-" yaml:"-"` // Parsing errors attached to this node
	hooks  map[string][]HookFunc // lazy-initialized; nil until first OnSet
}

// OnSet registers a hook that runs before a field is set.
func (t *Trix) OnSet(field string, fn HookFunc) {
	if t.hooks == nil {
		t.hooks = make(map[string][]HookFunc)
	}
	t.hooks[field] = append(t.hooks[field], fn)
}

// RunHooks runs all registered hooks for the given field.
// Returns the first error encountered, or nil.
func (t *Trix) RunHooks(field string, oldVal, newVal interface{}) error {
	if t.hooks == nil {
		return nil
	}
	for _, fn := range t.hooks[field] {
		if err := fn(field, oldVal, newVal); err != nil {
			return err
		}
	}
	return nil
}

// ElementBase is embedded in all OpenAPI types to provide vendor extensions and library metadata.
type ElementBase struct {
	VendorExtensions map[string]interface{} `json:"-" yaml:"-"`
	Trix             Trix                   `json:"-" yaml:"-"`
}
