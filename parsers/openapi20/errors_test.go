package openapi20

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// =============================================================================
// Tests for errors.go - ParseError
// =============================================================================

// --- NewParseError ---

func TestNewParseError_Basic(t *testing.T) {
	// Arrange & Act
	err := newParseError([]string{"info", "title"}, "title is required")

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, []string{"info", "title"}, err.Path)
	assert.Contains(t, err.Error(), "title is required")
}

// --- ParseError Error() ---

func TestParseError_Error(t *testing.T) {
	// Arrange
	err := &ParseError{
		Message: "invalid value",
		Path:    []string{"paths", "/pets", "get"},
	}

	// Act
	msg := err.Error()

	// Assert
	assert.Contains(t, msg, "paths./pets.get")
	assert.Contains(t, msg, "invalid value")
}

// --- ParseError with Location ---

func TestParseError_WithLocation(t *testing.T) {
	// Arrange
	err := &ParseError{
		Message: "invalid value",
		Path:    []string{"info"},
		Line:    10,
		Column:  5,
	}

	// Act
	msg := err.Error()

	// Assert
	assert.Contains(t, msg, "10")
	assert.Contains(t, msg, "5")
}

// --- ParseError Unwrap ---

func TestParseError_Unwrap(t *testing.T) {
	// Arrange
	cause := errors.New("underlying error")
	err := &ParseError{
		Message: "parsing failed",
		Path:    []string{"info"},
		Cause:   cause,
	}

	// Act
	unwrapped := err.Unwrap()

	// Assert
	assert.Equal(t, cause, unwrapped)
}

// --- ParseError Is ---

func TestParseError_Is(t *testing.T) {
	// Arrange
	cause := errors.New("underlying error")
	err := &ParseError{
		Message: "parsing failed",
		Cause:   cause,
	}

	// Act & Assert
	assert.True(t, errors.Is(err, cause))
}

// --- ParseError without Cause ---

func TestParseError_NilCause(t *testing.T) {
	// Arrange
	err := &ParseError{
		Message: "simple error",
		Path:    []string{"info"},
	}

	// Act
	unwrapped := err.Unwrap()

	// Assert
	assert.Nil(t, unwrapped)
}

// --- ParseError at Root ---

func TestParseError_AtRoot(t *testing.T) {
	// Arrange
	err := &ParseError{
		Message: "invalid document",
		Path:    []string{},
	}

	// Act
	msg := err.Error()

	// Assert
	assert.Contains(t, msg, "(root)")
}
