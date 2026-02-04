package openapi30

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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Examples, 1)
	assert.Contains(t, doc.Components.Examples, "PetExample")
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Examples, 3)
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, doc.Components.Examples)
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ex := doc.Components.Examples["Complete"].Value
	assert.Equal(t, "A complete example", ex.Summary)
	assert.Equal(t, "This is a detailed description", ex.Description)
}
