package openapi31x

import (
	"openapi-parser/parsers/internal/shared"

	"gopkg.in/yaml.v3"
)

// UnknownField represents a field in the source document that was not
// recognized as a valid OpenAPI field for its context.
// This is a type alias for backward compatibility.
type UnknownField = shared.UnknownField

// UnknownFieldError is an error type that wraps unknown fields for reporting.
// This is a type alias for backward compatibility.
type UnknownFieldError = shared.UnknownFieldError

// detectUnknownNodeFields checks a yaml.Node for keys that are not in the
// known fields set and not extensions (x-*). Returns a slice of unknown fields.
func detectUnknownNodeFields(node *yaml.Node, knownFields map[string]struct{}, path string) []UnknownField {
	return shared.DetectUnknownNodeFields(node, knownFields, path)
}

// isExtension checks if a field name is an OpenAPI extension (x-*).
func isExtension(key string) bool {
	return shared.IsExtension(key)
}

// formatPath converts a slice of path segments to a dot-separated path string.
func formatPath(segments []string) string {
	return shared.FormatPath(segments)
}

// formatUnknownFieldMessage formats a single unknown field for display.
func formatUnknownFieldMessage(f UnknownField) string {
	return shared.FormatUnknownFieldMessage(f)
}
