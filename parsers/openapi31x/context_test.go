package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for context.go - ParseContext functionality
// =============================================================================

// --- Path Tracking ---

func TestParseContext_PathInErrors(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses: "invalid"
`
	_, err := Parse([]byte(yaml))
	require.Error(t, err)
	// Error should contain path information
	assert.Contains(t, err.Error(), "responses")
}

// --- Complex Path ---

func TestParseContext_DeepPath(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets/{id}:
    get:
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
      responses:
        "200":
          description: "OK"
          content:
            application/json:
              schema:
                type: object
                properties:
                  nested:
                    type: object
                    properties:
                      deep:
                        type: string
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	// Verify deep schema is accessible
	resp := doc.Paths.Items["/pets/{id}"].Get.Responses.Codes["200"].Value
	schema := resp.Content["application/json"].Schema.Value
	assert.NotNil(t, schema.Properties["nested"])
}

// --- Multiple Paths ---

func TestParseContext_MultiplePaths(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /a:
    get:
      responses:
        "200":
          description: "OK"
  /b:
    get:
      responses:
        "200":
          description: "OK"
  /c:
    get:
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Paths.Items, 3)
}

// --- Node Source Line/Column ---

func TestParseContext_NodeSourceAccuracy(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      summary: "Get pets"
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)

	// Check that node sources are recorded
	op := doc.Paths.Items["/pets"].Get
	assert.Greater(t, op.NodeSource.Start.Line, 0)
	assert.Greater(t, op.NodeSource.Start.Column, 0)
}

// --- Extensions Preserved ---

func TestParseContext_ExtensionsPreserved(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  x-info-ext: "info"
paths:
  /pets:
    x-path-ext: "path"
    get:
      x-op-ext: "operation"
      responses:
        "200":
          description: "OK"
          x-resp-ext: "response"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)

	assert.Equal(t, "info", doc.Info.Extensions["x-info-ext"])
	assert.Equal(t, "path", doc.Paths.Items["/pets"].Extensions["x-path-ext"])
	assert.Equal(t, "operation", doc.Paths.Items["/pets"].Get.Extensions["x-op-ext"])
	assert.Equal(t, "response", doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Extensions["x-resp-ext"])
}
