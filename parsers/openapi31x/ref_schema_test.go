package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for ref_schema.go - schema reference parsing
// =============================================================================

// --- Basic Reference ---

func TestParseRefSchema_Basic(t *testing.T) {
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
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Pet'
components:
  schemas:
    Pet:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schemaRef := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"].Value().Content()["application/json"].Schema()
	assert.Equal(t, "#/components/schemas/Pet", schemaRef.Ref)
}

// --- In Property ---

func TestParseRefSchema_InProperty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      properties:
        owner:
          $ref: '#/components/schemas/User'
    User:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := result.Document.Components().Schemas()["Pet"].Value().Properties()["owner"]
	assert.Equal(t, "#/components/schemas/User", ref.Ref)
}

// --- In Array Items ---

func TestParseRefSchema_InItems(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    PetList:
      type: array
      items:
        $ref: '#/components/schemas/Pet'
    Pet:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := result.Document.Components().Schemas()["PetList"].Value().Items()
	assert.Equal(t, "#/components/schemas/Pet", ref.Ref)
}

// --- In AllOf ---

func TestParseRefSchema_InAllOf(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Cat:
      allOf:
        - $ref: '#/components/schemas/Pet'
        - type: object
          properties:
            meow:
              type: boolean
    Pet:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := result.Document.Components().Schemas()["Cat"].Value().AllOf()[0]
	assert.Equal(t, "#/components/schemas/Pet", ref.Ref)
}

// --- Multiple References ---

func TestParseRefSchema_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      properties:
        owner:
          $ref: '#/components/schemas/User'
        category:
          $ref: '#/components/schemas/Category'
        tags:
          type: array
          items:
            $ref: '#/components/schemas/Tag'
    User:
      type: object
    Category:
      type: object
    Tag:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	pet := result.Document.Components().Schemas()["Pet"].Value()
	assert.Equal(t, "#/components/schemas/User", pet.Properties()["owner"].Ref)
	assert.Equal(t, "#/components/schemas/Category", pet.Properties()["category"].Ref)
	assert.Equal(t, "#/components/schemas/Tag", pet.Properties()["tags"].Value().Items().Ref)
}
