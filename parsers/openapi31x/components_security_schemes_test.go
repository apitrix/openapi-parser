package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for components_securityschemes.go - parseComponentsSecuritySchemes function
// =============================================================================

// --- API Key ---

func TestParseComponentsSecuritySchemes_ApiKey(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
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
	scheme := result.Document.Components().SecuritySchemes()["apiKey"].Value()
	assert.Equal(t, "apiKey", scheme.Type())
	assert.Equal(t, "header", scheme.In())
	assert.Equal(t, "X-API-Key", scheme.Name())
}

// --- HTTP Basic ---

func TestParseComponentsSecuritySchemes_HTTPBasic(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    basic:
      type: http
      scheme: basic
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	scheme := result.Document.Components().SecuritySchemes()["basic"].Value()
	assert.Equal(t, "http", scheme.Type())
	assert.Equal(t, "basic", scheme.Scheme())
}

// --- HTTP Bearer ---

func TestParseComponentsSecuritySchemes_HTTPBearer(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    bearer:
      type: http
      scheme: bearer
      bearerFormat: JWT
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	scheme := result.Document.Components().SecuritySchemes()["bearer"].Value()
	assert.Equal(t, "http", scheme.Type())
	assert.Equal(t, "bearer", scheme.Scheme())
	assert.Equal(t, "JWT", scheme.BearerFormat())
}

// --- OAuth2 ---

func TestParseComponentsSecuritySchemes_OAuth2(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    oauth2:
      type: oauth2
      flows:
        clientCredentials:
          tokenUrl: https://example.com/token
          scopes: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	scheme := result.Document.Components().SecuritySchemes()["oauth2"].Value()
	assert.Equal(t, "oauth2", scheme.Type())
	assert.NotNil(t, scheme.Flows())
}

// --- OpenID Connect ---

func TestParseComponentsSecuritySchemes_OpenIdConnect(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    oidc:
      type: openIdConnect
      openIdConnectUrl: https://example.com/.well-known/openid
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	scheme := result.Document.Components().SecuritySchemes()["oidc"].Value()
	assert.Equal(t, "openIdConnect", scheme.Type())
	assert.Equal(t, "https://example.com/.well-known/openid", scheme.OpenIDConnectURL())
}

// --- Multiple Schemes ---

func TestParseComponentsSecuritySchemes_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
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
          scopes: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().SecuritySchemes(), 3)
}

// --- Empty ---

func TestParseComponentsSecuritySchemes_Empty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Components().SecuritySchemes())
}

// --- With Description ---

func TestParseComponentsSecuritySchemes_WithDescription(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    apiKey:
      type: apiKey
      in: header
      name: X-API-Key
      description: "API key for authentication"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	scheme := result.Document.Components().SecuritySchemes()["apiKey"].Value()
	assert.Equal(t, "API key for authentication", scheme.Description())
}
