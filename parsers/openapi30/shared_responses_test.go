package openapi30

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSharedResponses(t *testing.T) {
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	// Shared responses in components
	assert.Len(t, doc.Components.Responses, 2)
	// References in operation
	assert.Equal(t, "#/components/responses/BadRequest", doc.Paths.Items["/pets"].Get.Responses.Codes["400"].Ref)
}
