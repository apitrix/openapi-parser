package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseOAuthFlow(t *testing.T) {
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
          refreshUrl: https://example.com/oauth/refresh
          scopes:
            read: "Read access"
            write: "Write access"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	flow := doc.Components.SecuritySchemes["oauth2"].Value.Flows.ClientCredentials
	require.NotNil(t, flow)
	assert.Equal(t, "https://example.com/oauth/token", flow.TokenURL)
	assert.Equal(t, "https://example.com/oauth/refresh", flow.RefreshURL)
	assert.Len(t, flow.Scopes, 2)
}
