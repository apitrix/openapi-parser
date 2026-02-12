package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseMediaTypeSchema(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
          content:
            application/json:
              schema:
                type: array
                items:
                  type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	mt := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Content()["application/json"]
	require.NotNil(t, mt.Schema())
	assert.Equal(t, "array", mt.Schema().Value().Type().Single)
}
