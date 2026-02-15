package openapi31x

import (
	"strings"

	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseContext holds parsing state and provides utilities for error reporting
// and reference resolution. Uses a stack-based path to avoid allocations.
type ParseContext struct {
	Root          *yaml.Node          // Document root for $ref resolution
	pathStack     *[]string           // Shared backing slice for path segments
	depth         int                 // Current depth in the path stack
	Components    *yaml.Node          // Cached components section
	unknownFields *[]UnknownField     // Pointer to shared slice for accumulating unknown fields
	config        *ParseConfig // Feature flags controlling parsing behavior
}

// newParseContext creates a new ParseContext from the document root.
func newParseContext(root *yaml.Node, cfg *ParseConfig) *ParseContext {
	unknown := make([]UnknownField, 0)
	pathStack := make([]string, 0, 16) // Pre-allocate for typical depth
	ctx := &ParseContext{
		Root:          root,
		pathStack:     &pathStack,
		depth:         0,
		unknownFields: &unknown,
		config:        cfg,
	}

	// Cache components section if present
	ctx.Components = nodeGetValue(root, "components")

	return ctx
}

// Push creates a new context with the given path segment appended.
// Uses a shared backing slice to avoid allocations on each call.
func (ctx *ParseContext) Push(segment string) *ParseContext {
	stack := *ctx.pathStack
	newDepth := ctx.depth + 1

	if newDepth > len(stack) {
		// Grow the backing slice
		*ctx.pathStack = append(stack, segment)
	} else {
		// Reuse existing slot
		stack[ctx.depth] = segment
	}

	return &ParseContext{
		Root:          ctx.Root,
		pathStack:     ctx.pathStack,
		depth:         newDepth,
		Components:    ctx.Components,
		unknownFields: ctx.unknownFields, // Share the same slice
		config:        ctx.config,
	}
}

// Path returns the current path as a slice.
func (ctx *ParseContext) path() []string {
	return (*ctx.pathStack)[:ctx.depth]
}

// push is a convenience method that returns *ParseContext for internal use.
func (ctx *ParseContext) push(segment string) *ParseContext {
	return ctx.Push(segment)
}

// Errorf creates a ParseError with the current path.
// Implements *ParseContext interface.
func (ctx *ParseContext) Errorf(format string, args ...interface{}) error {
	return newParseError(ctx.path(), format, args...)
}

// errorf is an alias for Errorf for internal use.
func (ctx *ParseContext) errorf(format string, args ...interface{}) error {
	return ctx.Errorf(format, args...)
}

// errorAt creates a ParseError with line/column info from a node.
func (ctx *ParseContext) errorAt(node *yaml.Node, format string, args ...interface{}) error {
	err := newParseError(ctx.path(), format, args...)
	if node != nil {
		err.Line = node.Line
		err.Column = node.Column
	}
	return err
}

// CurrentPath returns the current path as a dot-separated string.
func (ctx *ParseContext) CurrentPath() string {
	return strings.Join(ctx.path(), ".")
}

// nodeSource creates a NodeSource for the given yaml.Node with full position info.
func (ctx *ParseContext) nodeSource(node *yaml.Node) openapi31models.NodeSource {
	if node == nil {
		return openapi31models.NodeSource{}
	}
	return openapi31models.NodeSource{
		Start: openapi31models.Location{
			Line:   node.Line,
			Column: node.Column,
		},
		Raw: nodeToInterface(node),
	}
}

// detectUnknown checks a node for unknown fields and records them.
// Returns nil immediately if unknown field detection is disabled in config.
// knownFields is the precomputed set of valid field names for this object type.
// Returns the unknown fields found so callers can also attach them to Trix.Errors.
func (ctx *ParseContext) detectUnknown(node *yaml.Node, knownFields map[string]struct{}) []UnknownField {
	if !ctx.config.DetectUnknownFields {
		return nil
	}
	unknown := detectUnknownNodeFields(node, knownFields, ctx.CurrentPath())
	if len(unknown) > 0 && ctx.unknownFields != nil {
		*ctx.unknownFields = append(*ctx.unknownFields, unknown...)
	}
	return unknown
}

// UnknownFieldsResult returns all unknown fields accumulated during parsing.
func (ctx *ParseContext) UnknownFieldsResult() []UnknownField {
	if ctx.unknownFields == nil {
		return nil
	}
	return *ctx.unknownFields
}
