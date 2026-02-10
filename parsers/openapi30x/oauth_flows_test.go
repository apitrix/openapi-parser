package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for oauthflows.go - parseOAuthFlows function
// =============================================================================

// --- Implicit Flow ---

func TestParseOAuthFlows_Implicit(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    oauth2:
      type: oauth2
      flows:
        implicit:
          authorizationUrl: https://example.com/oauth/authorize
          scopes:
            read: "Read access"
            write: "Write access"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	flows := result.Document.Components().SecuritySchemes()["oauth2"].Value.Flows()
	require.NotNil(t, flows.Implicit())
	assert.Equal(t, "https://example.com/oauth/authorize", flows.Implicit().AuthorizationURL())
	assert.Len(t, flows.Implicit().Scopes(), 2)
}

// --- Password Flow ---

func TestParseOAuthFlows_Password(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    oauth2:
      type: oauth2
      flows:
        password:
          tokenUrl: https://example.com/oauth/token
          scopes:
            read: "Read access"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	flows := result.Document.Components().SecuritySchemes()["oauth2"].Value.Flows()
	require.NotNil(t, flows.Password())
	assert.Equal(t, "https://example.com/oauth/token", flows.Password().TokenURL())
}

// --- Client Credentials Flow ---

func TestParseOAuthFlows_ClientCredentials(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
          tokenUrl: https://example.com/oauth/token
          scopes:
            admin: "Admin access"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	flows := result.Document.Components().SecuritySchemes()["oauth2"].Value.Flows()
	require.NotNil(t, flows.ClientCredentials())
	assert.Equal(t, "https://example.com/oauth/token", flows.ClientCredentials().TokenURL())
}

// --- Authorization Code Flow ---

func TestParseOAuthFlows_AuthorizationCode(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    oauth2:
      type: oauth2
      flows:
        authorizationCode:
          authorizationUrl: https://example.com/oauth/authorize
          tokenUrl: https://example.com/oauth/token
          refreshUrl: https://example.com/oauth/refresh
          scopes:
            read: "Read"
            write: "Write"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	flows := result.Document.Components().SecuritySchemes()["oauth2"].Value.Flows()
	require.NotNil(t, flows.AuthorizationCode())
	assert.Equal(t, "https://example.com/oauth/authorize", flows.AuthorizationCode().AuthorizationURL())
	assert.Equal(t, "https://example.com/oauth/token", flows.AuthorizationCode().TokenURL())
	assert.Equal(t, "https://example.com/oauth/refresh", flows.AuthorizationCode().RefreshURL())
}

// --- Multiple Flows ---

func TestParseOAuthFlows_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    oauth2:
      type: oauth2
      flows:
        implicit:
          authorizationUrl: https://example.com/auth
          scopes: {}
        password:
          tokenUrl: https://example.com/token
          scopes: {}
        clientCredentials:
          tokenUrl: https://example.com/token
          scopes: {}
        authorizationCode:
          authorizationUrl: https://example.com/auth
          tokenUrl: https://example.com/token
          scopes: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	flows := result.Document.Components().SecuritySchemes()["oauth2"].Value.Flows()
	assert.NotNil(t, flows.Implicit())
	assert.NotNil(t, flows.Password())
	assert.NotNil(t, flows.ClientCredentials())
	assert.NotNil(t, flows.AuthorizationCode())
}

// --- Scopes ---

func TestParseOAuthFlows_Scopes(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
          scopes:
            read: "Read access to resources"
            write: "Write access to resources"
            delete: "Delete access"
            admin: "Admin access"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	scopes := result.Document.Components().SecuritySchemes()["oauth2"].Value.Flows().ClientCredentials().Scopes()
	assert.Len(t, scopes, 4)
	assert.Equal(t, "Read access to resources", scopes["read"])
}

// --- Empty Scopes ---

func TestParseOAuthFlows_EmptyScopes(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	scopes := result.Document.Components().SecuritySchemes()["oauth2"].Value.Flows().ClientCredentials().Scopes()
	assert.Empty(t, scopes)
}

// --- Extensions ---

func TestParseOAuthFlows_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    oauth2:
      type: oauth2
      flows:
        x-custom: "value"
        clientCredentials:
          tokenUrl: https://example.com/token
          scopes: {}
          x-flow-custom: "flow-value"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	flows := result.Document.Components().SecuritySchemes()["oauth2"].Value.Flows()
	require.NotNil(t, flows.VendorExtensions)
	assert.Equal(t, "value", flows.VendorExtensions["x-custom"])
}
