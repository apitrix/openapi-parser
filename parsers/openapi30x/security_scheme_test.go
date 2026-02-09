package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSecurityScheme_AllTypes(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
    http:
      type: http
      scheme: bearer
      bearerFormat: JWT
    oauth2:
      type: oauth2
      flows:
        clientCredentials:
          tokenUrl: https://example.com/token
          scopes: {}
    openIdConnect:
      type: openIdConnect
      openIdConnectUrl: https://example.com/.well-known/openid
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schemes := result.Document.Components.SecuritySchemes
	assert.Len(t, schemes, 4)
	assert.Equal(t, "apiKey", schemes["apiKey"].Value.Type)
	assert.Equal(t, "http", schemes["http"].Value.Type)
	assert.Equal(t, "oauth2", schemes["oauth2"].Value.Type)
	assert.Equal(t, "openIdConnect", schemes["openIdConnect"].Value.Type)
}
