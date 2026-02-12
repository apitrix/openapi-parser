package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for operation_requestbody.go - parseOperationRequestBody function
// =============================================================================

// --- Basic RequestBody ---

func TestParseOperationRequestBody_Basic(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    post:
      requestBody:
        content:
          application/json:
            schema:
              type: object
      responses:
        "201":
          description: "Created"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	rb := result.Document.Paths().Items()["/pets"].Post().RequestBody().Value()
	require.NotNil(t, rb)
	assert.NotNil(t, rb.Content()["application/json"])
}

// --- Required ---

func TestParseOperationRequestBody_Required(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    post:
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
      responses:
        "201":
          description: "Created"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	rb := result.Document.Paths().Items()["/pets"].Post().RequestBody().Value()
	assert.True(t, rb.Required())
}

// --- Reference ---

func TestParseOperationRequestBody_Reference(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    post:
      requestBody:
        $ref: '#/components/requestBodies/PetBody'
      responses:
        "201":
          description: "Created"
components:
  requestBodies:
    PetBody:
      content:
        application/json:
          schema:
            type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "#/components/requestBodies/PetBody", result.Document.Paths().Items()["/pets"].Post().RequestBody().Ref)
}

// --- No RequestBody ---

func TestParseOperationRequestBody_None(t *testing.T) {
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
	assert.Nil(t, result.Document.Paths().Items()["/pets"].Get().RequestBody())
}

// --- With Description ---

func TestParseOperationRequestBody_Description(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    post:
      requestBody:
        description: "Pet to add to the store"
        content:
          application/json:
            schema:
              type: object
      responses:
        "201":
          description: "Created"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	rb := result.Document.Paths().Items()["/pets"].Post().RequestBody().Value()
	assert.Equal(t, "Pet to add to the store", rb.Description())
}
