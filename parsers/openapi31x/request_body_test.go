package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for requestbody.go - parseRequestBody function
// =============================================================================

// --- Basic RequestBody ---

func TestParseRequestBody_Basic(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    post:
      requestBody:
        description: "Pet to add"
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
	rb := result.Document.Paths().Items()["/pets"].Post().RequestBody().Value
	assert.Equal(t, "Pet to add", rb.Description())
}

// --- Required ---

func TestParseRequestBody_Required(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
	rb := result.Document.Paths().Items()["/pets"].Post().RequestBody().Value
	assert.True(t, rb.Required())
}

func TestParseRequestBody_NotRequired(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    post:
      requestBody:
        required: false
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
	rb := result.Document.Paths().Items()["/pets"].Post().RequestBody().Value
	assert.False(t, rb.Required())
}

// --- Multiple Content Types ---

func TestParseRequestBody_MultipleContentTypes(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
          application/xml:
            schema:
              type: object
          application/x-www-form-urlencoded:
            schema:
              type: object
          multipart/form-data:
            schema:
              type: object
      responses:
        "201":
          description: "Created"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	content := result.Document.Paths().Items()["/pets"].Post().RequestBody().Value.Content()
	assert.Len(t, content, 4)
}

// --- Reference ---

func TestParseRequestBody_Reference(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
      required: true
      content:
        application/json:
          schema:
            type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	rbRef := result.Document.Paths().Items()["/pets"].Post().RequestBody()
	assert.Equal(t, "#/components/requestBodies/PetBody", rbRef.Ref)
}

// --- Extensions ---

func TestParseRequestBody_Extensions(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    post:
      requestBody:
        x-custom: "value"
        x-internal: true
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
	rb := result.Document.Paths().Items()["/pets"].Post().RequestBody().Value
	require.NotNil(t, rb.VendorExtensions)
	assert.Equal(t, "value", rb.VendorExtensions["x-custom"])
}

// --- With Schema ---

func TestParseRequestBody_ComplexSchema(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
              required:
                - name
              properties:
                name:
                  type: string
                  minLength: 1
                tag:
                  type: string
      responses:
        "201":
          description: "Created"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Paths().Items()["/pets"].Post().RequestBody().Value.Content()["application/json"].Schema().Value
	assert.Equal(t, "object", schema.Type().Single)
	assert.Len(t, schema.Required(), 1)
	assert.Len(t, schema.Properties(), 2)
}

// --- Empty / Missing ---

func TestParseRequestBody_Missing(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
