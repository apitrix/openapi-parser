package openapi30

import (
	"strings"

	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseContext holds parsing state and provides utilities for error reporting
// and reference resolution.
type ParseContext struct {
	Root          *yaml.Node      // Document root for $ref resolution
	Path          []string        // Current JSON path (e.g., ["paths", "/pets", "get"])
	Components    *yaml.Node      // Cached components section
	unknownFields *[]UnknownField // Pointer to shared slice for accumulating unknown fields
}

// newParseContext creates a new ParseContext from the document root.
func newParseContext(root *yaml.Node) *ParseContext {
	unknown := make([]UnknownField, 0)
	ctx := &ParseContext{
		Root:          root,
		Path:          []string{},
		unknownFields: &unknown,
	}

	// Cache components section if present
	ctx.Components = nodeGetValue(root, "components")

	return ctx
}

// Push creates a new context with the given path segment appended.
// Implements *ParseContext interface.
func (ctx *ParseContext) Push(segment string) *ParseContext {
	newPath := make([]string, len(ctx.Path), len(ctx.Path)+1)
	copy(newPath, ctx.Path)
	newPath = append(newPath, segment)

	return &ParseContext{
		Root:          ctx.Root,
		Path:          newPath,
		Components:    ctx.Components,
		unknownFields: ctx.unknownFields, // Share the same slice
	}
}

// push is a convenience method that returns *ParseContext for internal use.
func (ctx *ParseContext) push(segment string) *ParseContext {
	return ctx.Push(segment)
}

// Errorf creates a ParseError with the current path.
// Implements *ParseContext interface.
func (ctx *ParseContext) Errorf(format string, args ...interface{}) error {
	return newParseError(ctx.Path, format, args...)
}

// errorf is an alias for Errorf for internal use.
func (ctx *ParseContext) errorf(format string, args ...interface{}) error {
	return ctx.Errorf(format, args...)
}

// errorAt creates a ParseError with line/column info from a node.
func (ctx *ParseContext) errorAt(node *yaml.Node, format string, args ...interface{}) error {
	err := newParseError(ctx.Path, format, args...)
	if node != nil {
		err.Line = node.Line
		err.Column = node.Column
	}
	return err
}

// CurrentPath returns the current path as a dot-separated string.
func (ctx *ParseContext) CurrentPath() string {
	return strings.Join(ctx.Path, ".")
}

// nodeSource creates a NodeSource for the given yaml.Node with full position info.
func (ctx *ParseContext) nodeSource(node *yaml.Node) openapi30models.NodeSource {
	if node == nil {
		return openapi30models.NodeSource{}
	}
	return openapi30models.NodeSource{
		Start: openapi30models.Location{
			Line:   node.Line,
			Column: node.Column,
		},
		Raw: nodeToInterface(node),
	}
}

// detectUnknown checks a node for unknown fields and records them.
// knownFields is the list of valid field names for this object type.
func (ctx *ParseContext) detectUnknown(node *yaml.Node, knownFields []string) {
	if ctx.unknownFields == nil {
		return
	}
	unknown := detectUnknownNodeFields(node, knownFields, ctx.CurrentPath())
	if len(unknown) > 0 {
		*ctx.unknownFields = append(*ctx.unknownFields, unknown...)
	}
}

// UnknownFieldsResult returns all unknown fields accumulated during parsing.
func (ctx *ParseContext) UnknownFieldsResult() []UnknownField {
	if ctx.unknownFields == nil {
		return nil
	}
	return *ctx.unknownFields
}
