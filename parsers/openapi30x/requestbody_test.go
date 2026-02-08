package openapi30x

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
	yaml := `openapi: "3.0.3"
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	rb := doc.Paths.Items["/pets"].Post.RequestBody.Value
	assert.Equal(t, "Pet to add", rb.Description)
}

// --- Required ---

func TestParseRequestBody_Required(t *testing.T) {
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	rb := doc.Paths.Items["/pets"].Post.RequestBody.Value
	assert.True(t, rb.Required)
}

func TestParseRequestBody_NotRequired(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	rb := doc.Paths.Items["/pets"].Post.RequestBody.Value
	assert.False(t, rb.Required)
}

// --- Multiple Content Types ---

func TestParseRequestBody_MultipleContentTypes(t *testing.T) {
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	content := doc.Paths.Items["/pets"].Post.RequestBody.Value.Content
	assert.Len(t, content, 4)
}

// --- Reference ---

func TestParseRequestBody_Reference(t *testing.T) {
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
      required: true
      content:
        application/json:
          schema:
            type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	rbRef := doc.Paths.Items["/pets"].Post.RequestBody
	assert.Equal(t, "#/components/requestBodies/PetBody", rbRef.Ref)
}

// --- Extensions ---

func TestParseRequestBody_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	rb := doc.Paths.Items["/pets"].Post.RequestBody.Value
	require.NotNil(t, rb.Extensions)
	assert.Equal(t, "value", rb.Extensions["x-custom"])
}

// --- With Schema ---

func TestParseRequestBody_ComplexSchema(t *testing.T) {
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := doc.Paths.Items["/pets"].Post.RequestBody.Value.Content["application/json"].Schema.Value
	assert.Equal(t, "object", schema.Type)
	assert.Len(t, schema.Required, 1)
	assert.Len(t, schema.Properties, 2)
}

// --- Empty / Missing ---

func TestParseRequestBody_Missing(t *testing.T) {
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Nil(t, doc.Paths.Items["/pets"].Get.RequestBody)
}
