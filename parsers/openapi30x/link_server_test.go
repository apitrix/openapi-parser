package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLinkServer(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	link := result.Document.Components().Links()["GetUser"].Value()
	require.NotNil(t, link.Server())
	assert.Equal(t, "https://api.example.com", link.Server().URL())
}
