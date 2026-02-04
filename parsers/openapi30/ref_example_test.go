package openapi30

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for ref_example.go - example reference parsing
// =============================================================================

// --- Basic Reference ---

func TestParseRefExample_Basic(t *testing.T) {
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
                pet:
                  $ref: '#/components/examples/PetExample'
components:
  examples:
    PetExample:
      value:
        name: "Fluffy"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Content["application/json"].Examples["pet"]
	assert.Equal(t, "#/components/examples/PetExample", ref.Ref)
}

// --- In Parameter ---

func TestParseRefExample_InParameter(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: status
          in: query
          schema:
            type: string
          examples:
            available:
              $ref: '#/components/examples/AvailableStatus'
      responses:
        "200":
          description: "OK"
components:
  examples:
    AvailableStatus:
      value: "available"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := doc.Paths.Items["/pets"].Get.Parameters[0].Value.Examples["available"]
	assert.Equal(t, "#/components/examples/AvailableStatus", ref.Ref)
}

// --- Multiple References ---

func TestParseRefExample_Multiple(t *testing.T) {
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
                cat:
                  $ref: '#/components/examples/CatExample'
                dog:
                  $ref: '#/components/examples/DogExample'
components:
  examples:
    CatExample:
      value:
        type: "cat"
    DogExample:
      value:
        type: "dog"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	examples := doc.Paths.Items["/pets"].Get.Responses.Codes["200"].Value.Content["application/json"].Examples
	assert.Equal(t, "#/components/examples/CatExample", examples["cat"].Ref)
	assert.Equal(t, "#/components/examples/DogExample", examples["dog"].Ref)
}
