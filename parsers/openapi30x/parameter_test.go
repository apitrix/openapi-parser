package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for parameter.go - parseParameter function
// =============================================================================

// --- Parameter Locations ---

func TestParseParameter_PathParam(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets/{petId}:
    get:
      parameters:
        - name: petId
          in: path
          required: true
          description: "The pet ID"
          schema:
            type: string
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/pets/{petId}"].Get.Parameters[0].Value
	assert.Equal(t, "petId", param.Name)
	assert.Equal(t, "path", param.In)
	assert.True(t, param.Required)
}

func TestParseParameter_QueryParam(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: limit
          in: query
          required: false
          schema:
            type: integer
            minimum: 1
            maximum: 100
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/pets"].Get.Parameters[0].Value
	assert.Equal(t, "limit", param.Name)
	assert.Equal(t, "query", param.In)
	assert.False(t, param.Required)
}

func TestParseParameter_HeaderParam(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: X-Request-ID
          in: header
          required: true
          schema:
            type: string
            format: uuid
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/pets"].Get.Parameters[0].Value
	assert.Equal(t, "X-Request-ID", param.Name)
	assert.Equal(t, "header", param.In)
}

func TestParseParameter_CookieParam(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: session_id
          in: cookie
          schema:
            type: string
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/pets"].Get.Parameters[0].Value
	assert.Equal(t, "session_id", param.Name)
	assert.Equal(t, "cookie", param.In)
}

// --- Multiple Parameters ---

func TestParseParameter_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets/{petId}:
    get:
      parameters:
        - name: petId
          in: path
          required: true
          schema:
            type: string
        - name: fields
          in: query
          schema:
            type: array
            items:
              type: string
        - name: X-Trace-Id
          in: header
          schema:
            type: string
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	params := result.Document.Paths.Items["/pets/{petId}"].Get.Parameters
	assert.Len(t, params, 3)
}

// --- Style and Explode ---

func TestParseParameter_StyleExplode(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: tags
          in: query
          style: form
          explode: true
          schema:
            type: array
            items:
              type: string
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/pets"].Get.Parameters[0].Value
	assert.Equal(t, "form", param.Style)
	require.NotNil(t, param.Explode)
	assert.True(t, *param.Explode)
}

func TestParseParameter_PipeDelimited(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: ids
          in: query
          style: pipeDelimited
          explode: false
          schema:
            type: array
            items:
              type: integer
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/pets"].Get.Parameters[0].Value
	assert.Equal(t, "pipeDelimited", param.Style)
}

// --- AllowEmptyValue, AllowReserved, Deprecated ---

func TestParseParameter_AllowEmptyValue(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /search:
    get:
      parameters:
        - name: q
          in: query
          allowEmptyValue: true
          schema:
            type: string
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/search"].Get.Parameters[0].Value
	assert.True(t, param.AllowEmptyValue)
}

func TestParseParameter_AllowReserved(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /files:
    get:
      parameters:
        - name: path
          in: query
          allowReserved: true
          schema:
            type: string
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/files"].Get.Parameters[0].Value
	assert.True(t, param.AllowReserved)
}

func TestParseParameter_Deprecated(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: oldParam
          in: query
          deprecated: true
          schema:
            type: string
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/pets"].Get.Parameters[0].Value
	assert.True(t, param.Deprecated)
}

// --- Examples ---

func TestParseParameter_Example(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: limit
          in: query
          schema:
            type: integer
          example: 10
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/pets"].Get.Parameters[0].Value
	assert.Equal(t, 10, param.Example)
}

func TestParseParameter_Examples(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: status
          in: query
          schema:
            type: string
          examples:
            available:
              value: "available"
            sold:
              value: "sold"
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/pets"].Get.Parameters[0].Value
	assert.Len(t, param.Examples, 2)
}

// --- Reference ---

func TestParseParameter_Reference(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - $ref: '#/components/parameters/LimitParam'
      responses:
        "200":
          description: "OK"
components:
  parameters:
    LimitParam:
      name: limit
      in: query
      schema:
        type: integer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	paramRef := result.Document.Paths.Items["/pets"].Get.Parameters[0]
	assert.Equal(t, "#/components/parameters/LimitParam", paramRef.Ref)
}

// --- Extensions ---

func TestParseParameter_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: query
          in: query
          x-custom: "value"
          x-internal: true
          schema:
            type: string
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/pets"].Get.Parameters[0].Value
	require.NotNil(t, param.VendorExtensions)
	assert.Equal(t, "value", param.VendorExtensions["x-custom"])
}
