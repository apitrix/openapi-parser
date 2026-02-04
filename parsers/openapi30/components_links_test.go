package openapi30

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for components_links.go - parseComponentsLinks function
// =============================================================================

// --- Single Link ---

func TestParseComponentsLinks_Single(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  links:
    GetUser:
      operationId: getUser
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Links, 1)
	assert.Contains(t, doc.Components.Links, "GetUser")
}

// --- Multiple Links ---

func TestParseComponentsLinks_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  links:
    GetUser:
      operationId: getUser
    GetPets:
      operationId: getPets
    GetOrders:
      operationId: getOrders
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Links, 3)
}

// --- Empty ---

func TestParseComponentsLinks_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  links: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, doc.Components.Links)
}

// --- With Parameters ---

func TestParseComponentsLinks_WithParameters(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  links:
    GetUserPets:
      operationId: getUserPets
      parameters:
        userId: '$response.body#/id'
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	link := doc.Components.Links["GetUserPets"].Value
	assert.Contains(t, link.Parameters, "userId")
}
