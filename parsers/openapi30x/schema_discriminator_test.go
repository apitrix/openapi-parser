package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_discriminator.go - ParseDiscriminator method
// =============================================================================

func TestParseSchemaDiscriminator_Basic(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      oneOf:
        - $ref: '#/components/schemas/Cat'
        - $ref: '#/components/schemas/Dog'
      discriminator:
        propertyName: petType
    Cat:
      type: object
    Dog:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	disc := result.Document.Components().Schemas()["Pet"].Value().Discriminator()
	require.NotNil(t, disc)
	assert.Equal(t, "petType", disc.PropertyName())
}

func TestParseSchemaDiscriminator_WithMapping(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      oneOf:
        - $ref: '#/components/schemas/Cat'
        - $ref: '#/components/schemas/Dog'
      discriminator:
        propertyName: petType
        mapping:
          cat: '#/components/schemas/Cat'
          dog: '#/components/schemas/Dog'
          kitty: '#/components/schemas/Cat'
    Cat:
      type: object
    Dog:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	disc := result.Document.Components().Schemas()["Pet"].Value().Discriminator()
	require.NotNil(t, disc)
	assert.Equal(t, "petType", disc.PropertyName())
	assert.Len(t, disc.Mapping(), 3)
	assert.Equal(t, "#/components/schemas/Cat", disc.Mapping()["cat"])
	assert.Equal(t, "#/components/schemas/Dog", disc.Mapping()["dog"])
	assert.Equal(t, "#/components/schemas/Cat", disc.Mapping()["kitty"])
}

func TestParseSchemaDiscriminator_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Simple:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	disc := result.Document.Components().Schemas()["Simple"].Value().Discriminator()
	assert.Nil(t, disc)
}

func TestParseSchemaDiscriminator_EmptyMapping(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      oneOf:
        - $ref: '#/components/schemas/Cat'
        - $ref: '#/components/schemas/Dog'
      discriminator:
        propertyName: type
    Cat:
      type: object
    Dog:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	disc := result.Document.Components().Schemas()["Pet"].Value().Discriminator()
	require.NotNil(t, disc)
	assert.Equal(t, "type", disc.PropertyName())
	assert.Empty(t, disc.Mapping())
}

func TestParseSchemaDiscriminator_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      oneOf:
        - $ref: '#/components/schemas/Cat'
        - $ref: '#/components/schemas/Dog'
      discriminator:
        propertyName: type
        x-custom: "value"
    Cat:
      type: object
    Dog:
      type: object
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	disc := result.Document.Components().Schemas()["Pet"].Value().Discriminator()
	require.NotNil(t, disc)
	require.NotNil(t, disc.VendorExtensions)
	assert.Equal(t, "value", disc.VendorExtensions["x-custom"])
}
