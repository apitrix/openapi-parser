package openapi30

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

// Trix contains all library-level metadata and functionality.
// Everything under Trix is provided by the apitrix library,
// not part of the OpenAPI specification itself.
type Trix struct {
	Source NodeSource `json:"-" yaml:"-"` // Source location info
}

// Node is embedded in all v30 types to provide vendor extensions and library metadata.
type Node struct {
	VendorExtensions map[string]interface{} `json:"-" yaml:"-"`
	Trix             Trix                   `json:"-" yaml:"-"`
}
