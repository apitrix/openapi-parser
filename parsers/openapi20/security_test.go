package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for security.go - parseSecurityScheme, parseSecurityRequirement
// =============================================================================

// --- API Key Security ---

func TestParseSecurityScheme_ApiKey(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
securityDefinitions:
  api_key:
    type: apiKey
    name: X-API-Key
    in: header
    description: "API key authentication"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	scheme := result.Document.SecurityDefinitions["api_key"]
	assert.Equal(t, "apiKey", scheme.Type)
	assert.Equal(t, "X-API-Key", scheme.Name)
	assert.Equal(t, "header", scheme.In)
	assert.Equal(t, "API key authentication", scheme.Description)
}

// --- Basic Authentication ---

func TestParseSecurityScheme_Basic(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
securityDefinitions:
  basic_auth:
    type: basic
    description: "HTTP Basic authentication"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	scheme := result.Document.SecurityDefinitions["basic_auth"]
	assert.Equal(t, "basic", scheme.Type)
}

// --- OAuth2 Implicit ---

func TestParseSecurityScheme_OAuth2Implicit(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
securityDefinitions:
  oauth2_implicit:
    type: oauth2
    flow: implicit
    authorizationUrl: "https://example.com/oauth/authorize"
    scopes:
      read: "Read access"
      write: "Write access"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	scheme := result.Document.SecurityDefinitions["oauth2_implicit"]
	assert.Equal(t, "oauth2", scheme.Type)
	assert.Equal(t, "implicit", scheme.Flow)
	assert.Equal(t, "https://example.com/oauth/authorize", scheme.AuthorizationURL)
	require.Contains(t, scheme.Scopes, "read")
	assert.Equal(t, "Read access", scheme.Scopes["read"])
}

// --- OAuth2 Password ---

func TestParseSecurityScheme_OAuth2Password(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
securityDefinitions:
  oauth2_password:
    type: oauth2
    flow: password
    tokenUrl: "https://example.com/oauth/token"
    scopes:
      read: "Read access"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	scheme := result.Document.SecurityDefinitions["oauth2_password"]
	assert.Equal(t, "password", scheme.Flow)
	assert.Equal(t, "https://example.com/oauth/token", scheme.TokenURL)
}

// --- OAuth2 AccessCode ---

func TestParseSecurityScheme_OAuth2AccessCode(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
securityDefinitions:
  oauth2_accesscode:
    type: oauth2
    flow: accessCode
    authorizationUrl: "https://example.com/oauth/authorize"
    tokenUrl: "https://example.com/oauth/token"
    scopes:
      admin: "Admin access"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	scheme := result.Document.SecurityDefinitions["oauth2_accesscode"]
	assert.Equal(t, "accessCode", scheme.Flow)
	assert.NotEmpty(t, scheme.AuthorizationURL)
	assert.NotEmpty(t, scheme.TokenURL)
}

// --- Security Requirement Simple ---

func TestParseSecurityRequirement_Simple(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      security:
        - api_key: []
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	security := result.Document.Paths.Items["/pets"].Get.Security
	require.Len(t, security, 1)
	require.Contains(t, security[0], "api_key")
	assert.Empty(t, security[0]["api_key"])
}

// --- Security Requirement With Scopes ---

func TestParseSecurityRequirement_WithScopes(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
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
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	security := result.Document.Paths.Items["/pets"].Get.Security
	require.Len(t, security, 1)
	scopes := security[0]["oauth2"]
	assert.Equal(t, []string{"read", "write"}, scopes)
}

// --- Multiple Security Requirements ---

func TestParseSecurityRequirement_Multiple(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      security:
        - api_key: []
        - oauth2:
            - read
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	security := result.Document.Paths.Items["/pets"].Get.Security
	require.Len(t, security, 2)
}

// --- Global Security ---

func TestParseSecurityRequirement_Global(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
security:
  - api_key: []
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.Len(t, result.Document.Security, 1)
	assert.Contains(t, result.Document.Security[0], "api_key")
}

// --- Security Scheme Extensions ---

func TestParseSecurityScheme_Extensions(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
securityDefinitions:
  api_key:
    type: apiKey
    name: X-API-Key
    in: header
    x-custom: "value"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	ext := result.Document.SecurityDefinitions["api_key"].VendorExtensions
	assert.Equal(t, "value", ext["x-custom"])
}
