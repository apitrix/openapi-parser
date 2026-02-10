package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for example.go - parseExample function
// =============================================================================

// --- Basic Example ---

func TestParseExample_Basic(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  examples:
    PetExample:
      value:
        name: "Fluffy"
        status: "available"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ex := result.Document.Components().Examples()["PetExample"].Value
	require.NotNil(t, ex)
	assert.NotNil(t, ex.Value)
}

// --- With Summary and Description ---

func TestParseExample_SummaryDescription(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  examples:
    PetExample:
      summary: "A sample pet"
      description: "This is a detailed description of the example pet"
      value:
        name: "Fluffy"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ex := result.Document.Components().Examples()["PetExample"].Value
	assert.Equal(t, "A sample pet", ex.Summary())
	assert.Equal(t, "This is a detailed description of the example pet", ex.Description())
}

// --- External Value ---

func TestParseExample_ExternalValue(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  examples:
    LargeExample:
      summary: "A large example"
      externalValue: "https://example.com/examples/large.json"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ex := result.Document.Components().Examples()["LargeExample"].Value
	assert.Equal(t, "https://example.com/examples/large.json", ex.ExternalValue())
}

// --- Multiple Examples ---

func TestParseExample_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  examples:
    CatExample:
      value:
        type: "cat"
    DogExample:
      value:
        type: "dog"
    BirdExample:
      value:
        type: "bird"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().Examples(), 3)
}

// --- Complex Value ---

func TestParseExample_ComplexValue(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  examples:
    ComplexExample:
      value:
        id: 123
        name: "Fluffy"
        tags:
          - name: "cute"
          - name: "friendly"
        owner:
          name: "John"
          email: "john@example.com"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ex := result.Document.Components().Examples()["ComplexExample"].Value
	require.NotNil(t, ex.Value)
}

// --- In Parameter ---

func TestParseExample_InParameter(t *testing.T) {
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
              summary: "Available pets"
              value: "available"
            sold:
              summary: "Sold pets"
              value: "sold"
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	examples := result.Document.Paths().Items()["/pets"].Get().Parameters()[0].Value.Examples()
	assert.Len(t, examples, 2)
}

// --- In MediaType ---

func TestParseExample_InMediaType(t *testing.T) {
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
              schema:
                type: object
              examples:
                cat:
                  value:
                    name: "Fluffy"
                dog:
                  value:
                    name: "Buddy"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	examples := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value.Content()["application/json"].Examples()
	assert.Len(t, examples, 2)
}

// --- Reference ---

func TestParseExample_Reference(t *testing.T) {
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	exRef := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value.Content()["application/json"].Examples()["pet"]
	assert.Equal(t, "#/components/examples/PetExample", exRef.Ref)
}

// --- Extensions ---

func TestParseExample_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  examples:
    PetExample:
      value:
        name: "Fluffy"
      x-custom: "value"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ex := result.Document.Components().Examples()["PetExample"].Value
	require.NotNil(t, ex.VendorExtensions)
	assert.Equal(t, "value", ex.VendorExtensions["x-custom"])
}
