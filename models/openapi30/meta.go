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

// Node is embedded in all v30 types to provide source info and extensions.
type Node struct {
	NodeSource NodeSource             `json:"-" yaml:"-"`
	VendorExtensions map[string]interface{} `json:"-" yaml:"-"`
}
