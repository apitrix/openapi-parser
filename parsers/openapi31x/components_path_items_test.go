package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for components_pathitems.go - parseComponentsPathItems
// =============================================================================

func TestParseComponentsPathItems_HappyPath(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  pathItems:
    SharedOps:
      get:
        summary: Shared GET
        operationId: sharedGet
        responses:
          "200":
            description: OK
    AnotherOp:
      post:
        summary: Another POST
        operationId: anotherPost
        responses:
          "201":
            description: Created
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Components)
	require.Len(t, result.Document.Components.PathItems, 2)
	assert.NotNil(t, result.Document.Components.PathItems["SharedOps"].Value.Get)
	assert.NotNil(t, result.Document.Components.PathItems["AnotherOp"].Value.Post)
}

func TestParseComponentsPathItems_WithRef(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  pathItems:
    Shared:
      $ref: '#/components/pathItems/Inline'
      summary: ref summary
    Inline:
      get:
        operationId: inlineGet
        responses:
          "200":
            description: OK
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := result.Document.Components.PathItems["Shared"]
	assert.Equal(t, "#/components/pathItems/Inline", ref.Ref)
	assert.Equal(t, "ref summary", ref.Summary)
}

func TestParseComponentsPathItems_Empty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Test:
      type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Nil(t, result.Document.Components.PathItems)
}

func TestParseComponentsPathItems_Extensions(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  pathItems:
    Ops:
      x-custom: value
      get:
        operationId: testGet
        responses:
          "200":
            description: OK
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ops := result.Document.Components.PathItems["Ops"]
	assert.Equal(t, "value", ops.Value.VendorExtensions["x-custom"])
}
