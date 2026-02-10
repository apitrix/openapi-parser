package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for ref_path_item.go - path item reference parsing
// =============================================================================

// --- Basic Reference ---

func TestParseRefPathItem_Basic(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    $ref: '#/components/pathItems/Pets'
components:
  pathItems:
    Pets:
      get:
        responses:
          "200":
            description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := result.Document.Paths().Items()["/pets"]
	assert.Equal(t, "#/components/pathItems/Pets", ref.Ref())
}

// --- Multiple Path References ---

func TestParseRefPathItem_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    $ref: '#/components/pathItems/Pets'
  /users:
    $ref: '#/components/pathItems/Users'
components:
  pathItems:
    Pets:
      get:
        responses:
          "200":
            description: "OK"
    Users:
      get:
        responses:
          "200":
            description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "#/components/pathItems/Pets", result.Document.Paths().Items()["/pets"].Ref())
	assert.Equal(t, "#/components/pathItems/Users", result.Document.Paths().Items()["/users"].Ref())
}

// --- Mixed Inline and Reference ---

func TestParseRefPathItem_Mixed(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    $ref: '#/components/pathItems/Pets'
  /orders:
    get:
      responses:
        "200":
          description: "OK"
components:
  pathItems:
    Pets:
      get:
        responses:
          "200":
            description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "#/components/pathItems/Pets", result.Document.Paths().Items()["/pets"].Ref())
	assert.NotNil(t, result.Document.Paths().Items()["/orders"].Get())
}
