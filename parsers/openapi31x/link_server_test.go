package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLinkServer(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  links:
    GetUser:
      operationId: getUser
      server:
        url: https://api.example.com
        description: "API server"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	link := doc.Components.Links["GetUser"].Value
	require.NotNil(t, link.Server)
	assert.Equal(t, "https://api.example.com", link.Server.URL)
}
