package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseHeaderSchema(t *testing.T) {
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
            X-Rate-Limit:
              schema:
                type: integer
                minimum: 0
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	header := doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Headers["X-Rate-Limit"].Value
	require.NotNil(t, header.Schema)
	assert.Equal(t, "integer", header.Schema.Value.Type.Single)
}
