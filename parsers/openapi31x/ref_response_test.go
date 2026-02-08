package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for ref_response.go - response reference parsing
// =============================================================================

// --- Basic Reference ---

func TestParseRefResponse_Basic(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
	ref := doc.Paths.Items["/pets"].Get.Responses.Codes["404"]
	assert.Equal(t, "#/components/responses/NotFound", ref.Ref)
}

// --- Multiple References ---

func TestParseRefResponse_Multiple(t *testing.T) {
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
        "400":
          $ref: '#/components/responses/BadRequest'
        "401":
          $ref: '#/components/responses/Unauthorized'
        "404":
          $ref: '#/components/responses/NotFound'
        "500":
          $ref: '#/components/responses/InternalError'
components:
  responses:
    BadRequest:
      description: "Bad request"
    Unauthorized:
      description: "Unauthorized"
    NotFound:
      description: "Not found"
    InternalError:
      description: "Internal error"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	codes := doc.Paths.Items["/pets"].Get.Responses.Codes
	assert.Equal(t, "#/components/responses/BadRequest", codes["400"].Ref)
	assert.Equal(t, "#/components/responses/Unauthorized", codes["401"].Ref)
	assert.Equal(t, "#/components/responses/NotFound", codes["404"].Ref)
	assert.Equal(t, "#/components/responses/InternalError", codes["500"].Ref)
}

// --- Default Response Reference ---

func TestParseRefResponse_Default(t *testing.T) {
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
        default:
          $ref: '#/components/responses/Error'
components:
  responses:
    Error:
      description: "Unexpected error"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := doc.Paths.Items["/pets"].Get.Responses.Default
	assert.Equal(t, "#/components/responses/Error", ref.Ref)
}
