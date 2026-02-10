package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseServerVariables(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	vars := result.Document.Servers()[0].Variables()
	require.NotNil(t, vars)
	assert.Contains(t, vars, "env")
	assert.Equal(t, "prod", vars["env"].Default())
}
