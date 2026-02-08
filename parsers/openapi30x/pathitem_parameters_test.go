package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsePathItemParameters(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets/{id}:
    parameters:
      - name: id
        in: path
        required: true
        schema:
          type: integer
    get:
      responses:
        "200":
          description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	pathItem := doc.Paths.Items["/pets/{id}"]
	assert.Len(t, pathItem.Parameters, 1)
}
