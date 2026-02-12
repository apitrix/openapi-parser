package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for ref_link.go - link reference parsing
// =============================================================================

// --- Basic Reference ---

func TestParseRefLink_Basic(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
            GetUserPets:
              $ref: '#/components/links/GetUserPets'
components:
  links:
    GetUserPets:
      operationId: getUserPets
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := result.Document.Paths().Items()["/users/{id}"].Get().Responses().Codes()["200"].Value().Links()["GetUserPets"]
	assert.Equal(t, "#/components/links/GetUserPets", ref.Ref)
}

// --- Multiple References ---

func TestParseRefLink_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
            GetUserPets:
              $ref: '#/components/links/GetUserPets'
            GetUserOrders:
              $ref: '#/components/links/GetUserOrders'
components:
  links:
    GetUserPets:
      operationId: getUserPets
    GetUserOrders:
      operationId: getUserOrders
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	links := result.Document.Paths().Items()["/users/{id}"].Get().Responses().Codes()["200"].Value().Links()
	assert.Equal(t, "#/components/links/GetUserPets", links["GetUserPets"].Ref)
	assert.Equal(t, "#/components/links/GetUserOrders", links["GetUserOrders"].Ref)
}

// --- Mixed Inline and Reference ---

func TestParseRefLink_Mixed(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
            GetUserPets:
              $ref: '#/components/links/GetUserPets'
            InlineLink:
              operationId: inlineOp
components:
  links:
    GetUserPets:
      operationId: getUserPets
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	links := result.Document.Paths().Items()["/users/{id}"].Get().Responses().Codes()["200"].Value().Links()
	assert.Equal(t, "#/components/links/GetUserPets", links["GetUserPets"].Ref)
	assert.Equal(t, "inlineOp", links["InlineLink"].Value().OperationID())
}
