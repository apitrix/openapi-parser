package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for components_requestbodies.go - parseComponentsRequestBodies function
// =============================================================================

// --- Single RequestBody ---

func TestParseComponentsRequestBodies_Single(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  requestBodies:
    PetBody:
      content:
        application/json:
          schema:
            type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.RequestBodies, 1)
	assert.Contains(t, doc.Components.RequestBodies, "PetBody")
}

// --- Multiple RequestBodies ---

func TestParseComponentsRequestBodies_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  requestBodies:
    PetBody:
      content:
        application/json:
          schema:
            type: object
    UserBody:
      content:
        application/json:
          schema:
            type: object
    OrderBody:
      content:
        application/json:
          schema:
            type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.RequestBodies, 3)
}

// --- Empty ---

func TestParseComponentsRequestBodies_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  requestBodies: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, doc.Components.RequestBodies)
}

// --- With Required ---

func TestParseComponentsRequestBodies_Required(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  requestBodies:
    PetBody:
      required: true
      content:
        application/json:
          schema:
            type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	rb := doc.Components.RequestBodies["PetBody"].Value
	assert.True(t, rb.Required)
}
