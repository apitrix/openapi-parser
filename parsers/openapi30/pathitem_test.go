package openapi30

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for pathitem.go - parsePathItem function
// =============================================================================

// --- Operations ---

func TestParsePathItem_AllOperations(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /resource:
    get:
      responses:
        "200":
          description: "OK"
    put:
      responses:
        "200":
          description: "OK"
    post:
      responses:
        "201":
          description: "Created"
    delete:
      responses:
        "204":
          description: "Deleted"
    options:
      responses:
        "200":
          description: "OK"
    head:
      responses:
        "200":
          description: "OK"
    patch:
      responses:
        "200":
          description: "OK"
    trace:
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	path := doc.Paths.Items["/resource"]
	assert.NotNil(t, path.Get)
	assert.NotNil(t, path.Put)
	assert.NotNil(t, path.Post)
	assert.NotNil(t, path.Delete)
	assert.NotNil(t, path.Options)
	assert.NotNil(t, path.Head)
	assert.NotNil(t, path.Patch)
	assert.NotNil(t, path.Trace)
}

// --- Summary and Description ---

func TestParsePathItem_SummaryDescription(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    summary: "Pet operations"
    description: "Operations for managing pets"
    get:
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	path := doc.Paths.Items["/pets"]
	assert.Equal(t, "Pet operations", path.Summary)
	assert.Equal(t, "Operations for managing pets", path.Description)
}

// --- Parameters ---

func TestParsePathItem_Parameters(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets/{petId}:
    parameters:
      - name: petId
        in: path
        required: true
        schema:
          type: string
      - name: version
        in: header
        schema:
          type: string
    get:
      responses:
        "200":
          description: "OK"
    put:
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	params := doc.Paths.Items["/pets/{petId}"].Parameters
	assert.Len(t, params, 2)
}

// --- Servers ---

func TestParsePathItem_Servers(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    servers:
      - url: https://pets.example.com
      - url: https://backup.pets.example.com
    get:
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	servers := doc.Paths.Items["/pets"].Servers
	assert.Len(t, servers, 2)
}

// --- Reference ---

func TestParsePathItem_Reference(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	path := doc.Paths.Items["/pets"]
	assert.Equal(t, "#/components/pathItems/Pets", path.Ref)
}

// --- Extensions ---

func TestParsePathItem_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    x-internal: true
    x-rate-limit: 100
    get:
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ext := doc.Paths.Items["/pets"].Extensions
	require.NotNil(t, ext)
	assert.Equal(t, true, ext["x-internal"])
}

// --- Multiple Paths ---

func TestParsePathItem_MultiplePaths(t *testing.T) {
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
  /pets/{petId}:
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Paths.Items, 4)
}

// --- Path with Template ---

func TestParsePathItem_PathTemplate(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /users/{userId}/pets/{petId}/medical-records/{recordId}:
    get:
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Contains(t, doc.Paths.Items, "/users/{userId}/pets/{petId}/medical-records/{recordId}")
}

// --- Empty Path ---

func TestParsePathItem_EmptyPaths(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, doc.Paths.Items)
}
