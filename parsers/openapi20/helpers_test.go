package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// =============================================================================
// Tests for helpers.go - map helper functions
// =============================================================================

// --- getString ---

func TestGetString_Exists(t *testing.T) {
	// Arrange
	m := map[string]interface{}{"key": "value"}

	// Act
	result := getString(m, "key")

	// Assert
	assert.Equal(t, "value", result)
}

func TestGetString_NotExists(t *testing.T) {
	// Arrange
	m := map[string]interface{}{}

	// Act
	result := getString(m, "key")

	// Assert
	assert.Empty(t, result)
}

func TestGetString_WrongType(t *testing.T) {
	// Arrange
	m := map[string]interface{}{"key": 123}

	// Act
	result := getString(m, "key")

	// Assert
	assert.Empty(t, result)
}

// --- getBool ---

func TestGetBool_True(t *testing.T) {
	// Arrange
	m := map[string]interface{}{"key": true}

	// Act
	result := getBool(m, "key")

	// Assert
	assert.True(t, result)
}

func TestGetBool_False(t *testing.T) {
	// Arrange
	m := map[string]interface{}{"key": false}

	// Act
	result := getBool(m, "key")

	// Assert
	assert.False(t, result)
}

func TestGetBool_NotExists(t *testing.T) {
	// Arrange
	m := map[string]interface{}{}

	// Act
	result := getBool(m, "key")

	// Assert
	assert.False(t, result)
}

// --- getInt ---

func TestGetInt_Exists(t *testing.T) {
	// Arrange
	m := map[string]interface{}{"key": 42}

	// Act
	result := getInt(m, "key")

	// Assert
	assert.Equal(t, 42, result)
}

func TestGetInt_FromFloat(t *testing.T) {
	// Arrange
	m := map[string]interface{}{"key": float64(42)}

	// Act
	result := getInt(m, "key")

	// Assert
	assert.Equal(t, 42, result)
}

func TestGetInt_NotExists(t *testing.T) {
	// Arrange
	m := map[string]interface{}{}

	// Act
	result := getInt(m, "key")

	// Assert
	assert.Equal(t, 0, result)
}

// --- getFloat64 ---

func TestGetFloat64_Exists(t *testing.T) {
	// Arrange
	m := map[string]interface{}{"key": 3.14}

	// Act
	result := getFloat64(m, "key")

	// Assert
	assert.Equal(t, 3.14, result)
}

func TestGetFloat64_FromInt(t *testing.T) {
	// Arrange
	m := map[string]interface{}{"key": 42}

	// Act
	result := getFloat64(m, "key")

	// Assert
	assert.Equal(t, float64(42), result)
}

// --- getStringSlice ---

func TestGetStringSlice_Exists(t *testing.T) {
	// Arrange
	m := map[string]interface{}{"key": []interface{}{"a", "b", "c"}}

	// Act
	result := getStringSlice(m, "key")

	// Assert
	assert.Equal(t, []string{"a", "b", "c"}, result)
}

func TestGetStringSlice_NotExists(t *testing.T) {
	// Arrange
	m := map[string]interface{}{}

	// Act
	result := getStringSlice(m, "key")

	// Assert
	assert.Nil(t, result)
}

// --- getMap ---

func TestGetMap_Exists(t *testing.T) {
	// Arrange
	inner := map[string]interface{}{"inner": "value"}
	m := map[string]interface{}{"key": inner}

	// Act
	result := getMap(m, "key")

	// Assert
	assert.Equal(t, inner, result)
}

func TestGetMap_NotExists(t *testing.T) {
	// Arrange
	m := map[string]interface{}{}

	// Act
	result := getMap(m, "key")

	// Assert
	assert.Nil(t, result)
}

// --- hasRef ---

func TestHasRef_True(t *testing.T) {
	// Arrange
	m := map[string]interface{}{"$ref": "#/definitions/Pet"}

	// Act
	result := hasRef(m)

	// Assert
	assert.True(t, result)
}

func TestHasRef_False(t *testing.T) {
	// Arrange
	m := map[string]interface{}{"type": "string"}

	// Act
	result := hasRef(m)

	// Assert
	assert.False(t, result)
}

// --- getRef ---

func TestGetRef_Exists(t *testing.T) {
	// Arrange
	m := map[string]interface{}{"$ref": "#/definitions/Pet"}

	// Act
	result := getRef(m)

	// Assert
	assert.Equal(t, "#/definitions/Pet", result)
}

// --- parseExtensions ---

func TestParseExtensions_WithExtensions(t *testing.T) {
	// Arrange
	m := map[string]interface{}{
		"type":     "string",
		"x-custom": "value",
		"x-flag":   true,
	}

	// Act
	result := parseExtensions(m)

	// Assert
	assert.Len(t, result, 2)
	assert.Equal(t, "value", result["x-custom"])
	assert.Equal(t, true, result["x-flag"])
}

func TestParseExtensions_NoExtensions(t *testing.T) {
	// Arrange
	m := map[string]interface{}{
		"type": "string",
	}

	// Act
	result := parseExtensions(m)

	// Assert
	assert.Empty(t, result)
}
