package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLinkParameters(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  links:
    GetUser:
      operationId: getUser
      parameters:
        userId: '$response.body#/id'
        format: json
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	link := result.Document.Components().Links()["GetUser"].Value
	require.NotNil(t, link.Parameters())
	assert.Contains(t, link.Parameters(), "userId")
	assert.Contains(t, link.Parameters(), "format")
}
