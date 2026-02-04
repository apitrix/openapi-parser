package openapi30

import (
	"fmt"
	"strings"
)

// ParseError represents an error that occurred during parsing,
// including the JSON path and source location where the error occurred.
type ParseError struct {
	Path    []string // JSON path where error occurred
	Message string   // Error message
	Cause   error    // Wrapped error if any
	Line    int      // Line number (1-based, 0 if unknown)
	Column  int      // Column number (1-based, 0 if unknown)
}

// newParseError creates a new ParseError with the given path and message.
func newParseError(path []string, format string, args ...interface{}) *ParseError {
	// Make a copy of the path to avoid mutation
	pathCopy := make([]string, len(path))
	copy(pathCopy, path)

	return &ParseError{
		Path:    pathCopy,
		Message: fmt.Sprintf(format, args...),
	}
}

// newParseErrorWithCause creates a new ParseError that wraps another error.
func newParseErrorWithCause(path []string, cause error, format string, args ...interface{}) *ParseError {
	err := newParseError(path, format, args...)
	err.Cause = cause
	return err
}

// Error implements the error interface.
func (e *ParseError) Error() string {
	pathStr := "(root)"
	if len(e.Path) > 0 {
		pathStr = strings.Join(e.Path, ".")
	}

	// Include line:column if available
	location := ""
	if e.Line > 0 {
		if e.Column > 0 {
			location = fmt.Sprintf(" [line %d, col %d]", e.Line, e.Column)
		} else {
			location = fmt.Sprintf(" [line %d]", e.Line)
		}
	}

	if e.Cause != nil {
		return fmt.Sprintf("parse error at %s%s: %s: %v", pathStr, location, e.Message, e.Cause)
	}
	return fmt.Sprintf("parse error at %s%s: %s", pathStr, location, e.Message)
}

// Unwrap returns the underlying cause error for use with errors.Is/As.
func (e *ParseError) Unwrap() error {
	return e.Cause
}
