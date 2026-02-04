package openapi30

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseServerVariables(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
servers:
  - url: https://{env}.example.com
    variables:
      env:
        default: prod
        enum:
          - prod
          - staging
        description: "Environment"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	vars := doc.Servers[0].Variables
	require.NotNil(t, vars)
	assert.Contains(t, vars, "env")
	assert.Equal(t, "prod", vars["env"].Default)
}
