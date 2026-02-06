package openapi30

import (
	"strings"

	"gopkg.in/yaml.v3"
)

// UnknownField represents a field in the source document that was not
// recognized as a valid OpenAPI field for its context.
type UnknownField struct {
	// Path is the JSON path to the parent object containing the unknown field.
	// Example: "paths./pets.get.responses.200"
	Path string

	// Key is the name of the unknown field.
	Key string

	// Line is the source line number (1-based) where the field was found.
	Line int

	// Column is the source column number (1-based) where the field was found.
	Column int
}

// detectUnknownNodeFields checks a yaml.Node for keys that are not in the
// known fields set and not extensions (x-*). Returns a slice of unknown fields.
func detectUnknownNodeFields(node *yaml.Node, knownFields map[string]struct{}, path string) []UnknownField {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}

	var unknown []UnknownField

	// Iterate through all keys in the mapping node
	for i := 0; i < len(node.Content)-1; i += 2 {
		keyNode := node.Content[i]
		key := keyNode.Value

		// Skip extensions (x-*)
		if isExtension(key) {
			continue
		}

		// Check if this is a known field
		if _, ok := knownFields[key]; !ok {
			unknown = append(unknown, UnknownField{
				Path:   path,
				Key:    key,
				Line:   keyNode.Line,
				Column: keyNode.Column,
			})
		}
	}

	return unknown
}

// isExtension checks if a field name is an OpenAPI extension (x-*).
func isExtension(key string) bool {
	return len(key) > 2 && key[0] == 'x' && key[1] == '-'
}

// formatPath converts a slice of path segments to a dot-separated path string.
func formatPath(segments []string) string {
	return strings.Join(segments, ".")
}

// UnknownFieldError is an error type that wraps unknown fields for reporting.
type UnknownFieldError struct {
	Fields []UnknownField
}

// Error implements the error interface.
func (e *UnknownFieldError) Error() string {
	if len(e.Fields) == 0 {
		return "no unknown fields"
	}
	if len(e.Fields) == 1 {
		f := e.Fields[0]
		return formatUnknownFieldMessage(f)
	}
	var sb strings.Builder
	sb.WriteString("multiple unknown fields found:\n")
	for _, f := range e.Fields {
		sb.WriteString("  - ")
		sb.WriteString(formatUnknownFieldMessage(f))
		sb.WriteString("\n")
	}
	return sb.String()
}

// formatUnknownFieldMessage formats a single unknown field for display.
func formatUnknownFieldMessage(f UnknownField) string {
	path := f.Path
	if path == "" {
		path = "(root)"
	}
	return "unknown field '" + f.Key + "' at " + path + " (line " + itoa(f.Line) + ", column " + itoa(f.Column) + ")"
}
