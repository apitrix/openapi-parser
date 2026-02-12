package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseParameterExamples(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: status
          in: query
          schema:
            type: string
          examples:
            available:
              value: "available"
            sold:
              value: "sold"
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value()
	require.NotNil(t, param.Examples())
	assert.Len(t, param.Examples(), 2)
}
