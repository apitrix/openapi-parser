package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseResponseHeaders(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
          headers:
            X-Rate-Limit:
              schema:
                type: integer
            X-Request-Id:
              schema:
                type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	resp := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value
	assert.Len(t, resp.Headers(), 2)
}
