package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for paths.go - parsePaths function
// =============================================================================

// --- Empty Paths ---

func TestParsePaths_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Paths.Items)
}

// --- Single Path ---

func TestParsePaths_Single(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Paths.Items, 1)
	assert.Contains(t, result.Document.Paths.Items, "/pets")
}

// --- Multiple Paths ---

func TestParsePaths_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
  /pets/{id}:
    get:
      responses:
        "200":
          description: "OK"
  /users:
    get:
      responses:
        "200":
          description: "OK"
  /orders:
    get:
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Paths.Items, 4)
}

// --- Path Templates ---

func TestParsePaths_Templates(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /users/{userId}:
    get:
      responses:
        "200":
          description: "OK"
  /users/{userId}/pets/{petId}:
    get:
      responses:
        "200":
          description: "OK"
  /stores/{storeId}/orders/{orderId}/items/{itemId}:
    get:
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Paths.Items, 3)
	assert.Contains(t, result.Document.Paths.Items, "/users/{userId}")
	assert.Contains(t, result.Document.Paths.Items, "/users/{userId}/pets/{petId}")
}

// --- Special Characters ---

func TestParsePaths_SpecialCharacters(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /api/v1/pets:
    get:
      responses:
        "200":
          description: "OK"
  /api/v2.0/pets:
    get:
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Contains(t, result.Document.Paths.Items, "/api/v1/pets")
	assert.Contains(t, result.Document.Paths.Items, "/api/v2.0/pets")
}

// --- Extensions ---

func TestParsePaths_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  x-custom: "value"
  /pets:
    get:
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Paths.VendorExtensions)
	assert.Equal(t, "value", result.Document.Paths.VendorExtensions["x-custom"])
}
