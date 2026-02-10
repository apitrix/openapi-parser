package openapi31x

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
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  links:
    GetUser:
      operationId: getUser
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().Links(), 1)
	assert.Contains(t, result.Document.Components().Links(), "GetUser")
}

// --- Multiple Links ---

func TestParseComponentsLinks_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().Links(), 3)
}

// --- Empty ---

func TestParseComponentsLinks_Empty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  links: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Components().Links())
}

// --- With Parameters ---

func TestParseComponentsLinks_WithParameters(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	link := result.Document.Components().Links()["GetUserPets"].Value
	assert.Contains(t, link.Parameters(), "userId")
}
