package openapi30

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSecuritySchemeFlows(t *testing.T) {
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
          scopes:
            read: Read
        authorizationCode:
          authorizationUrl: https://example.com/auth
          tokenUrl: https://example.com/token
          scopes:
            write: Write
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	flows := doc.Components.SecuritySchemes["oauth2"].Value.Flows
	require.NotNil(t, flows)
	assert.NotNil(t, flows.Implicit)
	assert.NotNil(t, flows.AuthorizationCode)
}
