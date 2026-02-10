package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSet(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		checks map[string]bool
	}{
		{
			"normal fields",
			[]string{"type", "description", "required"},
			map[string]bool{"type": true, "description": true, "required": true, "unknown": false},
		},
		{"empty slice", []string{}, map[string]bool{"anything": false}},
		{"nil slice", nil, map[string]bool{"anything": false}},
		{"single element", []string{"only"}, map[string]bool{"only": true, "other": false}},
		{"duplicates", []string{"a", "a", "b"}, map[string]bool{"a": true, "b": true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange & Act
			result := ToSet(tt.input)

			// Assert
			for key, wantPresent := range tt.checks {
				_, got := result[key]
				assert.Equal(t, wantPresent, got, "ToSet result[%q]", key)
			}
		})
	}
}

func TestToSet_LengthMatchesUniqueInputs(t *testing.T) {
	// Arrange & Act
	result := ToSet([]string{"x", "y", "z"})

	// Assert
	assert.Len(t, result, 3)
}
