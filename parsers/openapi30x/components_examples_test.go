package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for components_examples.go - parseComponentsExamples function
// =============================================================================

// --- Single Example ---

func TestParseComponentsExamples_Single(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  examples:
    PetExample:
      value:
        name: "Fluffy"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components.Examples, 1)
	assert.Contains(t, result.Document.Components.Examples, "PetExample")
}

// --- Multiple Examples ---

func TestParseComponentsExamples_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  examples:
    Cat:
      value:
        type: cat
    Dog:
      value:
        type: dog
    Bird:
      value:
        type: bird
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components.Examples, 3)
}

// --- Empty ---

func TestParseComponentsExamples_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  examples: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Components.Examples)
}

// --- With All Fields ---

func TestParseComponentsExamples_AllFields(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  examples:
    Complete:
      summary: "A complete example"
      description: "This is a detailed description"
      value:
        id: 1
        name: "Test"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ex := result.Document.Components.Examples["Complete"].Value
	assert.Equal(t, "A complete example", ex.Summary)
	assert.Equal(t, "This is a detailed description", ex.Description)
}
