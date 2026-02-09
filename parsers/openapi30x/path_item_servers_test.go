package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsePathItemServers(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    servers:
      - url: https://pets.example.com
      - url: https://pets-backup.example.com
    get:
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	pathItem := result.Document.Paths.Items["/pets"]
	assert.Len(t, pathItem.Servers, 2)
}
