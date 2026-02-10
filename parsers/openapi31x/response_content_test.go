package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseResponseContent(t *testing.T) {
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
            text/plain:
              schema:
                type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	resp := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value
	assert.Len(t, resp.Content(), 2)
}
