package openapi20

import "openapi-parser/models/shared"

// Type aliases for shared meta types.
type Location = shared.Location
type NodeSource = shared.NodeSource
type ParseError = shared.ParseError
type Trix = shared.Trix

// Node is embedded in all v20 types to provide vendor extensions and library metadata.
type Node = shared.Node
