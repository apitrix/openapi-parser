package openapi30x

import (
	"errors"
	"strings"

	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"
	"github.com/apitrix/openapi-parser/parsers/shared"
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
func toParseError(err error) openapi30models.ParseError {
	var pe *ParseError
	if errors.As(err, &pe) {
		return openapi30models.ParseError{Message: pe.Message, Path: pe.Path, Kind: "error"}
	}
	return openapi30models.ParseError{Message: err.Error(), Kind: "error"}
}

// unknownFieldParseErrors converts unknown fields into model-level ParseError entries
// with Kind "unknown_field" for attachment to a node's Trix.Errors slice.
func unknownFieldParseErrors(fields []UnknownField) []openapi30models.ParseError {
	if len(fields) == 0 {
		return nil
	}
	result := make([]openapi30models.ParseError, len(fields))
	for i, f := range fields {
		result[i] = openapi30models.ParseError{
			Message: "unknown field '" + f.Key + "'",
			Path:    pathFromString(f.Path),
			Kind:    "unknown_field",
		}
	}
	return result
}

// pathFromString converts a dot-separated path string to a []string.
func pathFromString(path string) []string {
	if path == "" {
		return nil
	}
	return strings.Split(path, ".")
}
