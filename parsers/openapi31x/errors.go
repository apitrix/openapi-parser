package openapi31x

import (
	"errors"

	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/parsers/internal/shared"
)

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

// toParseError converts a parser-level ParseError (or generic error) into the
// model-level ParseError for attachment to a node's Trix.Errors slice.
func toParseError(err error) openapi31models.ParseError {
	var pe *ParseError
	if errors.As(err, &pe) {
		return openapi31models.ParseError{Message: pe.Message, Path: pe.Path}
	}
	return openapi31models.ParseError{Message: err.Error()}
}
