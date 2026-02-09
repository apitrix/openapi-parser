package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseParameterContent(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: filter
          in: query
          content:
            application/json:
              schema:
                type: object
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	param := result.Document.Paths.Items["/pets"].Get.Parameters[0].Value
	require.NotNil(t, param.Content)
	assert.Contains(t, param.Content, "application/json")
}
