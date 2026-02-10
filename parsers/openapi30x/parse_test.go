package openapi30x

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsePetstore(t *testing.T) {
	// Arrange
	data, err := os.ReadFile("testdata/petstore.yaml")
	require.NoError(t, err)

	// Act
	result, err := Parse(data)

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document)

	// Root
	assert.Equal(t, "3.0.3", result.Document.OpenAPIVersion())
	assert.Equal(t, "2024-01", result.Document.VendorExtensions["x-api-version"])

	// Info
	require.NotNil(t, result.Document.Info())
	assert.Equal(t, "Petstore API", result.Document.Info().Title())
	assert.Equal(t, "1.0.0", result.Document.Info().Version())
	assert.Equal(t, "A sample API for pets", result.Document.Info().Description())
	assert.Equal(t, "custom value", result.Document.Info().VendorExtensions["x-custom-info"])

	// Contact
	require.NotNil(t, result.Document.Info().Contact())
	assert.Equal(t, "API Support", result.Document.Info().Contact().Name())
	assert.Equal(t, "support@example.com", result.Document.Info().Contact().Email())

	// License
	require.NotNil(t, result.Document.Info().License())
	assert.Equal(t, "MIT", result.Document.Info().License().Name())
	assert.Equal(t, "https://opensource.org/licenses/MIT", result.Document.Info().License().URL())

	// Servers
	require.Len(t, result.Document.Servers(), 2)
	assert.Equal(t, "https://api.petstore.com/v1", result.Document.Servers()[0].URL())
	assert.Equal(t, "Production server", result.Document.Servers()[0].Description())
	assert.Equal(t, "https://staging-api.petstore.com/v1", result.Document.Servers()[1].URL())
	assert.Equal(t, "Staging server", result.Document.Servers()[1].Description())

	// Paths
	require.NotNil(t, result.Document.Paths())
	require.Len(t, result.Document.Paths().Items(), 2)
	require.Contains(t, result.Document.Paths().Items(), "/pets")
	require.Contains(t, result.Document.Paths().Items(), "/pets/{petId}")

	// /pets - GET
	petsPath := result.Document.Paths().Items()["/pets"]
	require.NotNil(t, petsPath.Get())
	assert.Equal(t, "List all pets", petsPath.Get().Summary())
	assert.Equal(t, "listPets", petsPath.Get().OperationID())
	assert.Equal(t, []string{"pets"}, petsPath.Get().Tags())

	// /pets GET - Parameters
	require.Len(t, petsPath.Get().Parameters(), 1)
	limitParam := petsPath.Get().Parameters()[0].Value
	require.NotNil(t, limitParam)
	assert.Equal(t, "limit", limitParam.Name())
	assert.Equal(t, "query", limitParam.In())
	assert.Equal(t, "How many items to return", limitParam.Description())
	assert.False(t, limitParam.Required())
	require.NotNil(t, limitParam.Schema())
	assert.Equal(t, "integer", limitParam.Schema().Value.Type())
	assert.Equal(t, "int32", limitParam.Schema().Value.Format())

	// /pets GET - Responses
	require.NotNil(t, petsPath.Get().Responses())
	require.Contains(t, petsPath.Get().Responses().Codes(), "200")
	require.NotNil(t, petsPath.Get().Responses().Default())

	resp200 := petsPath.Get().Responses().Codes()["200"].Value
	assert.Equal(t, "A list of pets", resp200.Description())
	require.Contains(t, resp200.Content(), "application/json")
	assert.Equal(t, "array", resp200.Content()["application/json"].Schema().Value.Type())

	// /pets - POST
	require.NotNil(t, petsPath.Post())
	assert.Equal(t, "Create a pet", petsPath.Post().Summary())
	assert.Equal(t, "createPet", petsPath.Post().OperationID())
	require.NotNil(t, petsPath.Post().RequestBody())
	assert.True(t, petsPath.Post().RequestBody().Value.Required())
	require.Contains(t, petsPath.Post().RequestBody().Value.Content(), "application/json")
	assert.Equal(t, "#/components/schemas/NewPet", petsPath.Post().RequestBody().Value.Content()["application/json"].Schema().Ref)

	// /pets/{petId} - GET
	petByIdPath := result.Document.Paths().Items()["/pets/{petId}"]
	require.NotNil(t, petByIdPath.Get())
	assert.Equal(t, "Get a pet by ID", petByIdPath.Get().Summary())
	assert.Equal(t, "getPetById", petByIdPath.Get().OperationID())
	require.Len(t, petByIdPath.Get().Parameters(), 1)
	petIdParam := petByIdPath.Get().Parameters()[0].Value
	assert.Equal(t, "petId", petIdParam.Name())
	assert.Equal(t, "path", petIdParam.In())
	assert.True(t, petIdParam.Required())
	assert.Equal(t, "The ID of the pet", petIdParam.Description())

	// /pets/{petId} GET - Responses
	require.Contains(t, petByIdPath.Get().Responses().Codes(), "200")
	require.Contains(t, petByIdPath.Get().Responses().Codes(), "404")
	assert.Equal(t, "Pet not found", petByIdPath.Get().Responses().Codes()["404"].Value.Description())

	// /pets/{petId} - DELETE
	require.NotNil(t, petByIdPath.Delete())
	assert.Equal(t, "Delete a pet", petByIdPath.Delete().Summary())
	assert.Equal(t, "deletePet", petByIdPath.Delete().OperationID())
	require.Contains(t, petByIdPath.Delete().Responses().Codes(), "204")
	assert.Equal(t, "Pet deleted", petByIdPath.Delete().Responses().Codes()["204"].Value.Description())

	// Components - Schemas
	require.NotNil(t, result.Document.Components())
	require.Len(t, result.Document.Components().Schemas(), 3)
	require.Contains(t, result.Document.Components().Schemas(), "Pet")
	require.Contains(t, result.Document.Components().Schemas(), "NewPet")
	require.Contains(t, result.Document.Components().Schemas(), "Error")

	petSchema := result.Document.Components().Schemas()["Pet"].Value
	assert.Equal(t, "object", petSchema.Type())
	assert.Equal(t, []string{"id", "name"}, petSchema.Required())
	require.Len(t, petSchema.Properties(), 3)
	assert.Equal(t, "integer", petSchema.Properties()["id"].Value.Type())
	assert.Equal(t, "string", petSchema.Properties()["name"].Value.Type())
	assert.Equal(t, "pet-extra", petSchema.VendorExtensions["x-schema-extension"])

	// Components - Security Schemes
	require.Contains(t, result.Document.Components().SecuritySchemes(), "bearerAuth")
	bearerScheme := result.Document.Components().SecuritySchemes()["bearerAuth"].Value
	assert.Equal(t, "http", bearerScheme.Type())
	assert.Equal(t, "bearer", bearerScheme.Scheme())
	assert.Equal(t, "JWT", bearerScheme.BearerFormat())

	// Security
	require.Len(t, result.Document.Security(), 1)
	require.Contains(t, result.Document.Security()[0], "bearerAuth")

	// Tags
	require.Len(t, result.Document.Tags(), 1)
	assert.Equal(t, "pets", result.Document.Tags()[0].Name())
	assert.Equal(t, "Pet operations", result.Document.Tags()[0].Description())
	require.NotNil(t, result.Document.Tags()[0].ExternalDocs())
	assert.Equal(t, "https://example.com/docs/pets", result.Document.Tags()[0].ExternalDocs().URL())
}

func TestParseInvalidVersion(t *testing.T) {
	// Arrange
	data := []byte(`{"openapi": "2.0", "info": {"title": "Test", "version": "1.0"}}`)

	// Act
	_, err := Parse(data)

	// Assert
	assert.Error(t, err)
}

func TestParseMissingInfo(t *testing.T) {
	// Arrange
	data := []byte(`{"openapi": "3.0.0"}`)

	// Act
	result, err := Parse(data)

	// Assert — info errors are now collected on result.Document.Trix.Errors, not returned
	require.NoError(t, err)
	require.NotNil(t, result.Document)
	assert.NotEmpty(t, result.Document.Trix.Errors, "missing info should produce a Trix error")
}

func TestParseLineColumnNumbers(t *testing.T) {
	// Arrange
	data := []byte(`openapi: "3.0.3"
info:
  title: "Test API"
  version: "1.0.0"
paths:
  /pets:
    get:
      summary: "Get pets"
      responses:
        "200":
          description: "Success"
`)

	// Act
	result, err := Parse(data)

	// Assert
	require.NoError(t, err)

	assert.NotZero(t, result.Document.Trix.Source.Start.Line, "root line")
	assert.NotZero(t, result.Document.Trix.Source.Start.Column, "root column")
	assert.NotZero(t, result.Document.Info().Trix.Source.Start.Line, "info line")
	assert.NotZero(t, result.Document.Paths().Trix.Source.Start.Line, "paths line")

	petsPath := result.Document.Paths().Items()["/pets"]
	require.NotNil(t, petsPath)
	assert.NotZero(t, petsPath.Trix.Source.Start.Line, "/pets line")
	assert.NotZero(t, petsPath.Get().Trix.Source.Start.Line, "GET operation line")
}
