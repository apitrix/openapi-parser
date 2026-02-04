package openapi30

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseRequestBodyContent(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    post:
      requestBody:
        content:
          application/json:
            schema:
              type: object
          application/xml:
            schema:
              type: object
      responses:
        "201":
          description: "Created"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	rb := doc.Paths.Items["/pets"].Post.RequestBody.Value
	assert.Len(t, rb.Content, 2)
}
