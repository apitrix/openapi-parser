package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseParameterSchema(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: limit
          in: query
          schema:
            type: integer
            minimum: 1
            maximum: 100
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value()
	require.NotNil(t, param.Schema())
	assert.Equal(t, "integer", param.Schema().Value().Type().Single)
}
