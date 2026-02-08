package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for mediatype.go - parseMediaType function
// =============================================================================

// --- Basic MediaType ---

func TestParseMediaType_Basic(t *testing.T) {
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
          content:
            application/json:
              schema:
                type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	mt := doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Content["application/json"]
	require.NotNil(t, mt)
	assert.NotNil(t, mt.Schema)
}

// --- Multiple Content Types ---

func TestParseMediaType_Multiple(t *testing.T) {
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
          content:
            application/json:
              schema:
                type: object
            application/xml:
              schema:
                type: object
            text/plain:
              schema:
                type: string
            text/html:
              schema:
                type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	content := doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Content
	assert.Len(t, content, 4)
}

// --- Example ---

func TestParseMediaType_Example(t *testing.T) {
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
          content:
            application/json:
              schema:
                type: object
              example:
                name: "Fluffy"
                status: "available"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	mt := doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Content["application/json"]
	assert.NotNil(t, mt.Example)
}

// --- Examples (multiple) ---

func TestParseMediaType_Examples(t *testing.T) {
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
          content:
            application/json:
              schema:
                type: object
              examples:
                cat:
                  summary: "A cat"
                  value:
                    name: "Fluffy"
                    type: "cat"
                dog:
                  summary: "A dog"
                  value:
                    name: "Buddy"
                    type: "dog"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	mt := doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Content["application/json"]
	assert.Len(t, mt.Examples, 2)
	assert.Contains(t, mt.Examples, "cat")
	assert.Contains(t, mt.Examples, "dog")
}

// --- Encoding ---

func TestParseMediaType_Encoding(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /upload:
    post:
      requestBody:
        content:
          multipart/form-data:
            schema:
              type: object
              properties:
                file:
                  type: string
                  format: binary
                metadata:
                  type: object
            encoding:
              file:
                contentType: application/octet-stream
              metadata:
                contentType: application/json
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	enc := doc.Paths.Items["/upload"].Post.RequestBody.Value.Content["multipart/form-data"].Encoding
	assert.Len(t, enc, 2)
	assert.Contains(t, enc, "file")
	assert.Contains(t, enc, "metadata")
}

// --- Schema Reference ---

func TestParseMediaType_SchemaReference(t *testing.T) {
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
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Pet'
components:
  schemas:
    Pet:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	mt := doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Content["application/json"]
	assert.Equal(t, "#/components/schemas/Pet", mt.Schema.Ref)
}

// --- Extensions ---

func TestParseMediaType_Extensions(t *testing.T) {
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
          content:
            application/json:
              schema:
                type: object
              x-custom: "value"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	mt := doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Content["application/json"]
	require.NotNil(t, mt.VendorExtensions)
	assert.Equal(t, "value", mt.VendorExtensions["x-custom"])
}

// --- Complex Media Types ---

func TestParseMediaType_Wildcards(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /files:
    get:
      responses:
        "200":
          description: "OK"
          content:
            "*/*":
              schema:
                type: string
                format: binary
            "image/*":
              schema:
                type: string
                format: binary
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	content := doc.Paths.Items["/files"].Get.Responses.Codes["200"].Value.Content
	assert.Contains(t, content, "*/*")
	assert.Contains(t, content, "image/*")
}
