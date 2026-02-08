package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseMediaTypeExamples(t *testing.T) {
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
          content:
            application/json:
              examples:
                pet1:
                  value:
                    name: "Fluffy"
                pet2:
                  value:
                    name: "Buddy"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	mt := doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Content["application/json"]
	require.NotNil(t, mt.Examples)
	assert.Len(t, mt.Examples, 2)
}
