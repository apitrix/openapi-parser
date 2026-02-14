package openapi30

import "openapi-parser/models/shared"

// Location represents a position in the source file.
type Location = shared.Location

// NodeSource contains source location and raw parsed data for a node.
type NodeSource = shared.NodeSource

// ParseError represents a parsing error associated with a specific node.
type ParseError = shared.ParseError

// Trix contains all library-level metadata and functionality.
// Everything under Trix is provided by the apitrix library,
// not part of the OpenAPI specification itself.
type Trix struct {
	Source NodeSource   `json:"-" yaml:"-"` // Source location info
	Errors []ParseError `json:"-" yaml:"-"` // Parsing errors attached to this node

	// Resolved reference fields — populated by the resolver, not part of the spec.
	// These are nil/empty unless the resolver has run.

	// ResolvedMapping holds discriminator mapping values resolved to schema refs (Discriminator only).
	ResolvedMapping map[string]*shared.Ref[Schema] `json:"-" yaml:"-"`

	// ResolvedOperation holds the operation resolved from operationRef (Link only).
	ResolvedOperation *Operation `json:"-" yaml:"-"`
}

// Node is embedded in all v30 types to provide vendor extensions and library metadata.
type Node struct {
	VendorExtensions map[string]interface{} `json:"-" yaml:"-"`
	Trix             Trix                   `json:"-" yaml:"-"`
}
