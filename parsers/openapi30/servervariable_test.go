package openapi30

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseServerVariable(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
servers:
  - url: https://{host}:{port}
    variables:
      host:
        default: api.example.com
      port:
        default: "443"
        enum:
          - "443"
          - "8443"
        description: "Port number"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	portVar := doc.Servers[0].Variables["port"]
	require.NotNil(t, portVar)
	assert.Equal(t, "443", portVar.Default)
	assert.Len(t, portVar.Enum, 2)
	assert.Equal(t, "Port number", portVar.Description)
}
