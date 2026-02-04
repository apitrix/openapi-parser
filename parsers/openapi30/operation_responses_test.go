package openapi30

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for operation_responses.go - parseOperationResponses function
// =============================================================================

// --- Single Response ---

func TestParseOperationResponses_Single(t *testing.T) {
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
	codes := doc.Paths.Items["/pets"].Get.Responses.Codes
	assert.Len(t, codes, 1)
}

// --- Multiple Status Codes ---

func TestParseOperationResponses_Multiple(t *testing.T) {
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
        "400":
          description: "Bad Request"
        "404":
          description: "Not Found"
        "500":
          description: "Internal Error"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	codes := doc.Paths.Items["/pets"].Get.Responses.Codes
	assert.Len(t, codes, 4)
}

// --- Default Response ---

func TestParseOperationResponses_Default(t *testing.T) {
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
        default:
          description: "Unexpected error"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	defaultResp := doc.Paths.Items["/pets"].Get.Responses.Default
	require.NotNil(t, defaultResp)
	assert.Equal(t, "Unexpected error", defaultResp.Value.Description)
}

// --- With Content ---

func TestParseOperationResponses_WithContent(t *testing.T) {
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
                type: array
                items:
                  type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	content := doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Content
	assert.Contains(t, content, "application/json")
}

// --- Reference Response ---

func TestParseOperationResponses_Reference(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "404":
          $ref: '#/components/responses/NotFound'
components:
  responses:
    NotFound:
      description: "Not found"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "#/components/responses/NotFound", doc.Paths.Items["/pets"].Get.Responses.Codes["404"].Ref)
}
