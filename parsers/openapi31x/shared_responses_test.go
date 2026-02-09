package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSharedResponses(t *testing.T) {
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
        "500":
          $ref: '#/components/responses/InternalError'
components:
  responses:
    BadRequest:
      description: "Bad request"
    InternalError:
      description: "Internal error"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	// Shared responses in components
	assert.Len(t, result.Document.Components.Responses, 2)
	// References in operation
	assert.Equal(t, "#/components/responses/BadRequest", result.Document.Paths.Items["/pets"].Get.Responses.Codes["400"].Ref)
}
