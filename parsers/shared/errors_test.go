package shared

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewParseError(t *testing.T) {
	tests := []struct {
		name    string
		path    []string
		format  string
		args    []interface{}
		wantMsg string
	}{
		{"simple message", []string{"components", "schemas", "Pet"}, "invalid type", nil, "invalid type"},
		{"formatted message", []string{"paths", "/pets"}, "expected %s, got %s", []interface{}{"object", "array"}, "expected object, got array"},
		{"empty path", nil, "root error", nil, "root error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange & Act
			err := NewParseError(tt.path, tt.format, tt.args...)

			// Assert
			assert.Equal(t, tt.wantMsg, err.Message)
			assert.Len(t, err.Path, len(tt.path))
		})
	}
}

func TestNewParseError_PathIsCopied(t *testing.T) {
	// Arrange
	path := []string{"a", "b", "c"}

	// Act
	err := NewParseError(path, "test")
	path[0] = "mutated"

	// Assert — the error's path should not be affected
	assert.Equal(t, "a", err.Path[0], "ParseError path should be a copy, not a reference")
}

func TestNewParseErrorWithCause(t *testing.T) {
	// Arrange
	cause := fmt.Errorf("underlying problem")

	// Act
	err := NewParseErrorWithCause([]string{"info"}, cause, "wrapping %s", "cause")

	// Assert
	assert.Equal(t, "wrapping cause", err.Message)
	assert.Equal(t, cause, err.Cause)
}

func TestParseError_Error_EmptyPath(t *testing.T) {
	// Arrange
	err := &ParseError{Path: nil, Message: "something went wrong"}

	// Act
	got := err.Error()

	// Assert
	assert.Contains(t, got, "(root)")
	assert.Contains(t, got, "something went wrong")
}

func TestParseError_Error_WithPath(t *testing.T) {
	// Arrange
	err := &ParseError{Path: []string{"components", "schemas", "Pet"}, Message: "invalid type"}

	// Act & Assert
	assert.Contains(t, err.Error(), "components.schemas.Pet")
}

func TestParseError_Error_WithLineAndColumn(t *testing.T) {
	// Arrange
	err := &ParseError{Path: []string{"info"}, Message: "missing title", Line: 5, Column: 3}

	// Act
	got := err.Error()

	// Assert
	assert.Contains(t, got, "line 5")
	assert.Contains(t, got, "col 3")
}

func TestParseError_Error_WithLineOnly(t *testing.T) {
	// Arrange
	err := &ParseError{Path: []string{"info"}, Message: "missing title", Line: 10}

	// Act
	got := err.Error()

	// Assert
	assert.Contains(t, got, "line 10")
	assert.NotContains(t, got, "col")
}

func TestParseError_Error_WithCause(t *testing.T) {
	// Arrange
	cause := fmt.Errorf("file not found")
	err := &ParseError{Path: []string{"paths"}, Message: "loading failed", Cause: cause}

	// Act
	got := err.Error()

	// Assert
	assert.Contains(t, got, "loading failed")
	assert.Contains(t, got, "file not found")
}

func TestParseError_Unwrap(t *testing.T) {
	// Arrange
	cause := fmt.Errorf("root cause")
	err := NewParseErrorWithCause([]string{}, cause, "wrap")

	// Act & Assert
	assert.True(t, errors.Is(err, cause))
}

func TestParseError_Unwrap_NilCause(t *testing.T) {
	// Arrange
	err := NewParseError([]string{}, "no cause")

	// Act & Assert
	require.Nil(t, err.Unwrap())
}
