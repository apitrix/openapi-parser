package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for operation_parameters.go - parseOperationParameters function
// =============================================================================

// --- Single Parameter ---

func TestParseOperationParameters_Single(t *testing.T) {
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
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	params := doc.Paths.Items["/pets"].Get.Parameters
	assert.Len(t, params, 1)
}

// --- Multiple Parameters ---

func TestParseOperationParameters_Multiple(t *testing.T) {
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
        - name: offset
          in: query
          schema:
            type: integer
        - name: status
          in: query
          schema:
            type: string
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	params := doc.Paths.Items["/pets"].Get.Parameters
	assert.Len(t, params, 3)
}

// --- Different Locations ---

func TestParseOperationParameters_AllLocations(t *testing.T) {
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
        - name: filter
          in: query
          schema:
            type: string
        - name: X-Request-Id
          in: header
          schema:
            type: string
        - name: session
          in: cookie
          schema:
            type: string
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	params := doc.Paths.Items["/pets/{petId}"].Get.Parameters
	assert.Len(t, params, 4)
	assert.Equal(t, "path", params[0].Value.In)
	assert.Equal(t, "query", params[1].Value.In)
	assert.Equal(t, "header", params[2].Value.In)
	assert.Equal(t, "cookie", params[3].Value.In)
}

// --- Reference Parameter ---

func TestParseOperationParameters_Reference(t *testing.T) {
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "#/components/parameters/LimitParam", doc.Paths.Items["/pets"].Get.Parameters[0].Ref)
}

// --- No Parameters ---

func TestParseOperationParameters_None(t *testing.T) {
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
	assert.Empty(t, doc.Paths.Items["/pets"].Get.Parameters)
}
