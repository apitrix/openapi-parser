package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseResponseLinks(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /users/{id}:
    get:
      responses:
        "200":
          description: "OK"
          links:
            GetPets:
              operationId: getUserPets
            GetOrders:
              operationId: getUserOrders
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	resp := result.Document.Paths().Items()["/users/{id}"].Get().Responses().Codes()["200"].Value()
	assert.Len(t, resp.Links(), 2)
}
