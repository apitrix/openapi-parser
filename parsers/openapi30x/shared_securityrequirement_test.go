package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSharedSecurityRequirement(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
security:
  - apiKey: []
  - oauth2:
      - read
      - write
paths:
  /pets:
    get:
      security:
        - bearerAuth: []
      responses:
        "200":
          description: "OK"
components:
  securitySchemes:
    apiKey:
      type: apiKey
      in: header
      name: X-API-Key
    oauth2:
      type: oauth2
      flows:
        clientCredentials:
          tokenUrl: https://example.com/token
          scopes:
            read: Read
            write: Write
    bearerAuth:
      type: http
      scheme: bearer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	// Global security
	assert.Len(t, result.Document.Security, 2)
	// Operation-level security
	assert.Len(t, result.Document.Paths.Items["/pets"].Get.Security, 1)
}
