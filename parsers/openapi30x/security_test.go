package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for security.go - parseSecurity function
// =============================================================================

// --- Basic Security Requirement ---

func TestParseSecurity_SingleScheme(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
security:
  - apiKey: []
paths: {}
components:
  securitySchemes:
    apiKey:
      type: apiKey
      in: header
      name: X-API-Key
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.Len(t, result.Document.Security(), 1)
	assert.Contains(t, result.Document.Security()[0], "apiKey")
}

// --- Multiple Schemes (AND) ---

func TestParseSecurity_MultipleSchemes(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
security:
  - apiKey: []
    oauth2:
      - read
paths: {}
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
            read: Read access
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.Len(t, result.Document.Security(), 1)
	// Both schemes required (AND)
	assert.Contains(t, result.Document.Security()[0], "apiKey")
	assert.Contains(t, result.Document.Security()[0], "oauth2")
}

// --- Alternative Schemes (OR) ---

func TestParseSecurity_AlternativeSchemes(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
security:
  - apiKey: []
  - oauth2:
      - read
  - bearerAuth: []
paths: {}
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
            read: Read access
    bearerAuth:
      type: http
      scheme: bearer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	// Three alternatives (OR)
	assert.Len(t, result.Document.Security(), 3)
}

// --- Scopes ---

func TestParseSecurity_Scopes(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
security:
  - oauth2:
      - read
      - write
      - admin
paths: {}
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
            admin: Admin
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	scopes := result.Document.Security()[0]["oauth2"]
	assert.Len(t, scopes, 3)
	assert.Contains(t, scopes, "read")
	assert.Contains(t, scopes, "write")
	assert.Contains(t, scopes, "admin")
}

// --- Empty Security (public) ---

func TestParseSecurity_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
security: []
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Security())
}

// --- Optional Security ---

func TestParseSecurity_Optional(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
security:
  - {}
  - apiKey: []
paths: {}
components:
  securitySchemes:
    apiKey:
      type: apiKey
      in: header
      name: X-API-Key
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	// First is empty (no auth required), second requires apiKey
	assert.Len(t, result.Document.Security(), 2)
	assert.Empty(t, result.Document.Security()[0])
	assert.Contains(t, result.Document.Security()[1], "apiKey")
}

// --- Operation-Level Security ---

func TestParseSecurity_OperationLevel(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
  /private:
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
    bearerAuth:
      type: http
      scheme: bearer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	// Global security
	assert.Len(t, result.Document.Security(), 1)
	// Public endpoint overrides to empty
	assert.Empty(t, result.Document.Paths().Items()["/public"].Get().Security())
	// Private endpoint has its own security
	assert.Len(t, result.Document.Paths().Items()["/private"].Get().Security(), 1)
}

// --- Missing Security ---

func TestParseSecurity_Missing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Nil(t, result.Document.Security())
}
