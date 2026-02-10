package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for operation.go - parseOperation function
// =============================================================================

// --- Basic Operations ---

func TestParseOperation_Get(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      summary: "List pets"
      description: "Returns all pets"
      operationId: listPets
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	op := result.Document.Paths().Items()["/pets"].Get()
	require.NotNil(t, op)
	assert.Equal(t, "List pets", op.Summary())
	assert.Equal(t, "Returns all pets", op.Description())
	assert.Equal(t, "listPets", op.OperationID())
}

func TestParseOperation_AllMethods(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /resource:
    get:
      operationId: getResource
      responses:
        "200":
          description: "OK"
    post:
      operationId: createResource
      responses:
        "201":
          description: "Created"
    put:
      operationId: updateResource
      responses:
        "200":
          description: "OK"
    patch:
      operationId: patchResource
      responses:
        "200":
          description: "OK"
    delete:
      operationId: deleteResource
      responses:
        "204":
          description: "Deleted"
    options:
      operationId: optionsResource
      responses:
        "200":
          description: "OK"
    head:
      operationId: headResource
      responses:
        "200":
          description: "OK"
    trace:
      operationId: traceResource
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	path := result.Document.Paths().Items()["/resource"]
	assert.NotNil(t, path.Get())
	assert.NotNil(t, path.Post())
	assert.NotNil(t, path.Put())
	assert.NotNil(t, path.Patch())
	assert.NotNil(t, path.Delete())
	assert.NotNil(t, path.Options())
	assert.NotNil(t, path.Head())
	assert.NotNil(t, path.Trace())
}

// --- Tags ---

func TestParseOperation_Tags(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      tags:
        - pets
        - animals
        - store
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	tags := result.Document.Paths().Items()["/pets"].Get().Tags()
	assert.Len(t, tags, 3)
	assert.Contains(t, tags, "pets")
	assert.Contains(t, tags, "animals")
	assert.Contains(t, tags, "store")
}

func TestParseOperation_NoTags(t *testing.T) {
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
	tags := result.Document.Paths().Items()["/pets"].Get().Tags()
	assert.Empty(t, tags)
}

// --- Deprecated ---

func TestParseOperation_Deprecated(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /old-endpoint:
    get:
      deprecated: true
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.True(t, result.Document.Paths().Items()["/old-endpoint"].Get().Deprecated())
}

// --- Extensions ---

func TestParseOperation_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      x-internal: true
      x-rate-limit: 100
      x-roles:
        - admin
        - user
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ext := result.Document.Paths().Items()["/pets"].Get().VendorExtensions
	require.NotNil(t, ext)
	assert.Equal(t, true, ext["x-internal"])
	assert.Equal(t, 100, ext["x-rate-limit"])
}

// --- Complete Operation ---

func TestParseOperation_Complete(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets/{petId}:
    put:
      tags:
        - pets
      summary: "Update a pet"
      description: "Updates a pet in the store"
      operationId: updatePet
      deprecated: false
      parameters:
        - name: petId
          in: path
          required: true
          schema:
            type: string
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
      responses:
        "200":
          description: "Updated"
        "404":
          description: "Not found"
      security:
        - apiKey: []
      servers:
        - url: https://api.example.com
      externalDocs:
        url: https://docs.example.com
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	op := result.Document.Paths().Items()["/pets/{petId}"].Put()
	require.NotNil(t, op)
	assert.Len(t, op.Tags(), 1)
	assert.Equal(t, "Update a pet", op.Summary())
	assert.Equal(t, "updatePet", op.OperationID())
	assert.Len(t, op.Parameters(), 1)
	require.NotNil(t, op.RequestBody())
	assert.Len(t, op.Responses().Codes(), 2)
	assert.Len(t, op.Security(), 1)
	assert.Len(t, op.Servers(), 1)
	require.NotNil(t, op.ExternalDocs())
}

// --- Node Source ---

func TestParseOperation_NodeSource(t *testing.T) {
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
	op := result.Document.Paths().Items()["/pets"].Get()
	assert.Greater(t, op.Trix.Source.Start.Line, 0)
}

// --- Multiple Paths ---

func TestParseOperation_MultiplePaths(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      operationId: listPets
      responses:
        "200":
          description: "OK"
  /pets/{petId}:
    get:
      operationId: getPet
      responses:
        "200":
          description: "OK"
  /users:
    get:
      operationId: listUsers
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Paths().Items(), 3)
	assert.Equal(t, "listPets", result.Document.Paths().Items()["/pets"].Get().OperationID())
	assert.Equal(t, "getPet", result.Document.Paths().Items()["/pets/{petId}"].Get().OperationID())
	assert.Equal(t, "listUsers", result.Document.Paths().Items()["/users"].Get().OperationID())
}
