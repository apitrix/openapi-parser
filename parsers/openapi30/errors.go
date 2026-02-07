package openapi30

import "openapi-parser/parsers/internal/shared"

// ParseError represents an error that occurred during parsing,
// including the JSON path and source location where the error occurred.
// This is a type alias for backward compatibility.
type ParseError = shared.ParseError

// newParseError creates a new ParseError with the given path and message.
func newParseError(path []string, format string, args ...interface{}) *ParseError {
	return shared.NewParseError(path, format, args...)
}

// newParseErrorWithCause creates a new ParseError that wraps another error.
func newParseErrorWithCause(path []string, cause error, format string, args ...interface{}) *ParseError {
	return shared.NewParseErrorWithCause(path, cause, format, args...)
}
