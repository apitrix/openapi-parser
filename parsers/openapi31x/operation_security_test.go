package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for operation_security.go - parseOperationSecurity function
// =============================================================================

// --- Single Security Scheme ---

func TestParseOperationSecurity_Single(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      security:
        - apiKey: []
      responses:
        "200":
          description: "OK"
components:
  securitySchemes:
    apiKey:
      type: apiKey
      in: header
      name: X-API-Key
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	security := result.Document.Paths().Items()["/pets"].Get().Security()
	assert.Len(t, security, 1)
}

// --- Multiple Alternatives (OR) ---

func TestParseOperationSecurity_Alternatives(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      security:
        - apiKey: []
        - bearer: []
        - oauth2:
            - read
      responses:
        "200":
          description: "OK"
components:
  securitySchemes:
    apiKey:
      type: apiKey
      in: header
      name: X-API-Key
    bearer:
      type: http
      scheme: bearer
    oauth2:
      type: oauth2
      flows:
        clientCredentials:
          tokenUrl: https://example.com/token
          scopes:
            read: Read
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	security := result.Document.Paths().Items()["/pets"].Get().Security()
	assert.Len(t, security, 3)
}

// --- Empty Security (Public) ---

func TestParseOperationSecurity_Empty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
security:
  - apiKey: []
paths:
  /public:
    get:
      security: []
      responses:
        "200":
          description: "OK"
components:
  securitySchemes:
    apiKey:
      type: apiKey
      in: header
      name: X-API-Key
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	security := result.Document.Paths().Items()["/public"].Get().Security()
	assert.Empty(t, security)
}

// --- No Operation Security ---

func TestParseOperationSecurity_None(t *testing.T) {
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
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	// Uses global security (nil means inherit)
	assert.Nil(t, result.Document.Paths().Items()["/pets"].Get().Security())
}

// --- With Scopes ---

func TestParseOperationSecurity_WithScopes(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      security:
        - oauth2:
            - read
            - write
      responses:
        "200":
          description: "OK"
components:
  securitySchemes:
    oauth2:
      type: oauth2
      flows:
        clientCredentials:
          tokenUrl: https://example.com/token
          scopes:
            read: Read
            write: Write
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	scopes := result.Document.Paths().Items()["/pets"].Get().Security()[0]["oauth2"]
	assert.Len(t, scopes, 2)
}
