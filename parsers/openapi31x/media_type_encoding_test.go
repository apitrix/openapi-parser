package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseMediaTypeEncoding(t *testing.T) {
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
              file:
                contentType: application/octet-stream
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	mt := result.Document.Paths().Items()["/upload"].Post().RequestBody().Value().Content()["multipart/form-data"]
	require.NotNil(t, mt.Encoding())
	assert.Contains(t, mt.Encoding(), "file")
}
