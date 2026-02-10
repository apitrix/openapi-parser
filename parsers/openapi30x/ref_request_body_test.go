package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for ref_request_body.go - request body reference parsing
// =============================================================================

// --- Basic Reference ---

func TestParseRefRequestBody_Basic(t *testing.T) {
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
	ref := result.Document.Paths().Items()["/pets"].Post().RequestBody()
	assert.Equal(t, "#/components/requestBodies/PetBody", ref.Ref)
}

// --- Multiple Operations ---

func TestParseRefRequestBody_MultipleOperations(t *testing.T) {
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
  /pets/{id}:
    put:
      requestBody:
        $ref: '#/components/requestBodies/PetBody'
      responses:
        "200":
          description: "OK"
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
	assert.Equal(t, "#/components/requestBodies/PetBody", result.Document.Paths().Items()["/pets/{id}"].Put().RequestBody().Ref)
}

// --- Different Request Bodies ---

func TestParseRefRequestBody_Different(t *testing.T) {
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
  /users:
    post:
      requestBody:
        $ref: '#/components/requestBodies/UserBody'
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
    UserBody:
      content:
        application/json:
          schema:
            type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "#/components/requestBodies/PetBody", result.Document.Paths().Items()["/pets"].Post().RequestBody().Ref)
	assert.Equal(t, "#/components/requestBodies/UserBody", result.Document.Paths().Items()["/users"].Post().RequestBody().Ref)
}
