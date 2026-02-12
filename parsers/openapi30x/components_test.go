package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for components.go - parseComponents function
// =============================================================================

// --- All Component Types ---

func TestParseComponents_AllTypes(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
  responses:
    NotFound:
      description: "Not found"
  parameters:
    LimitParam:
      name: limit
      in: query
      schema:
        type: integer
  examples:
    PetExample:
      value:
        name: "Fluffy"
  requestBodies:
    PetBody:
      content:
        application/json:
          schema:
            type: object
  headers:
    X-Rate-Limit:
      schema:
        type: integer
  securitySchemes:
    apiKey:
      type: apiKey
      in: header
      name: X-API-Key
  links:
    GetPet:
      operationId: getPet
  callbacks:
    onEvent:
      '{$url}':
        post:
          responses:
            "200":
              description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	c := result.Document.Components()
	require.NotNil(t, c)
	assert.Len(t, c.Schemas(), 1)
	assert.Len(t, c.Responses(), 1)
	assert.Len(t, c.Parameters(), 1)
	assert.Len(t, c.Examples(), 1)
	assert.Len(t, c.RequestBodies(), 1)
	assert.Len(t, c.Headers(), 1)
	assert.Len(t, c.SecuritySchemes(), 1)
	assert.Len(t, c.Links(), 1)
	assert.Len(t, c.Callbacks(), 1)
}

// --- Empty Components ---

func TestParseComponents_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Components())
}

// --- Missing Components ---

func TestParseComponents_Missing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	// Components can be nil or empty
	if result.Document.Components() != nil {
		assert.Empty(t, result.Document.Components().Schemas())
	}
}

// --- Multiple Schemas ---

func TestParseComponents_MultipleSchemas(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
    Owner:
      type: object
    Category:
      type: object
    Tag:
      type: object
    Order:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components().Schemas(), 5)
}

// --- Security Schemes Types ---

func TestParseComponents_SecuritySchemeTypes(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    apiKey:
      type: apiKey
      in: header
      name: X-API-Key
    basicAuth:
      type: http
      scheme: basic
    bearerAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT
    oauth2:
      type: oauth2
      flows:
        clientCredentials:
          tokenUrl: https://example.com/token
          scopes: {}
    openId:
      type: openIdConnect
      openIdConnectUrl: https://example.com/.well-known/openid
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schemes := result.Document.Components().SecuritySchemes()
	assert.Len(t, schemes, 5)
	assert.Equal(t, "apiKey", schemes["apiKey"].Value().Type())
	assert.Equal(t, "http", schemes["basicAuth"].Value().Type())
	assert.Equal(t, "http", schemes["bearerAuth"].Value().Type())
	assert.Equal(t, "oauth2", schemes["oauth2"].Value().Type())
	assert.Equal(t, "openIdConnect", schemes["openId"].Value().Type())
}

// --- Extensions ---

func TestParseComponents_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  x-custom: "value"
  x-internal: true
  schemas:
    Pet:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Components().VendorExtensions)
	assert.Equal(t, "value", result.Document.Components().VendorExtensions["x-custom"])
}

// --- Cross-References ---

func TestParseComponents_CrossReferences(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
          $ref: '#/components/schemas/Owner'
        tags:
          type: array
          items:
            $ref: '#/components/schemas/Tag'
    Owner:
      type: object
      properties:
        name:
          type: string
    Tag:
      type: object
      properties:
        name:
          type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	pet := result.Document.Components().Schemas()["Pet"].Value()
	assert.Equal(t, "#/components/schemas/Owner", pet.Properties()["owner"].Ref)
	assert.Equal(t, "#/components/schemas/Tag", pet.Properties()["tags"].Value().Items().Ref)
}
