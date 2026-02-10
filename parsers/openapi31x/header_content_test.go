package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseHeaderContent(t *testing.T) {
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
          headers:
            X-Custom:
              content:
                application/json:
                  schema:
                    type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	header := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value.Headers()["X-Custom"].Value
	require.NotNil(t, header.Content())
	assert.Contains(t, header.Content(), "application/json")
}
