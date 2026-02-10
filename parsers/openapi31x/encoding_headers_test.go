package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseEncodingHeaders(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /upload:
    post:
      requestBody:
        content:
          multipart/form-data:
            schema:
              type: object
            encoding:
              data:
                headers:
                  X-Custom:
                    schema:
                      type: string
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	enc := result.Document.Paths().Items()["/upload"].Post().RequestBody().Value.Content()["multipart/form-data"].Encoding()["data"]
	require.NotNil(t, enc.Headers())
	assert.Contains(t, enc.Headers(), "X-Custom")
}
