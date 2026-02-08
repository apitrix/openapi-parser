package openapi31x

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestParsePetstore31 tests parsing a complete OpenAPI 3.1 document.
func TestParsePetstore31(t *testing.T) {
	data, err := os.ReadFile("testdata/petstore31.yaml")
	require.NoError(t, err)

	doc, err := Parse(data)
	require.NoError(t, err)
	require.NotNil(t, doc)

	// Verify version
	assert.Equal(t, "3.1.0", doc.OpenAPI)

	// Verify info (with 3.1 additions)
	require.NotNil(t, doc.Info)
	assert.Equal(t, "Petstore - OpenAPI 3.1", doc.Info.Title)
	assert.Equal(t, "A pet store API demonstrating OpenAPI 3.1 features", doc.Info.Summary)
	assert.Equal(t, "1.0.0", doc.Info.Version)

	// Verify contact
	require.NotNil(t, doc.Info.Contact)
	assert.Equal(t, "API Support", doc.Info.Contact.Name)
	assert.Equal(t, "support@example.com", doc.Info.Contact.Email)

	// Verify license (with 3.1 additions)
	require.NotNil(t, doc.Info.License)
	assert.Equal(t, "Apache 2.0", doc.Info.License.Name)
	assert.Equal(t, "Apache-2.0", doc.Info.License.Identifier)

	// Verify jsonSchemaDialect
	assert.Equal(t, "https://json-schema.org/draft/2020-12/schema", doc.JsonSchemaDialect)

	// Verify servers
	require.Len(t, doc.Servers, 1)
	assert.Equal(t, "https://petstore.example.com/v1", doc.Servers[0].URL)

	// Verify webhooks
	require.NotNil(t, doc.Webhooks)
	require.Contains(t, doc.Webhooks, "newPet")
	webhook := doc.Webhooks["newPet"]
	require.NotNil(t, webhook.Value)
	require.NotNil(t, webhook.Value.Post)
	assert.Equal(t, "newPetWebhook", webhook.Value.Post.OperationID)

	// Verify paths
	require.NotNil(t, doc.Paths)
	require.Contains(t, doc.Paths.Items, "/pets")
	require.Contains(t, doc.Paths.Items, "/pets/{petId}")

	// Verify GET /pets
	getPets := doc.Paths.Items["/pets"].Get
	require.NotNil(t, getPets)
	assert.Equal(t, "listPets", getPets.OperationID)
	require.Len(t, getPets.Parameters, 1)
	require.NotNil(t, getPets.Parameters[0].Value)
	assert.Equal(t, "limit", getPets.Parameters[0].Value.Name)
	assert.Equal(t, "query", getPets.Parameters[0].Value.In)

	// Verify POST /pets
	postPets := doc.Paths.Items["/pets"].Post
	require.NotNil(t, postPets)
	assert.Equal(t, "createPets", postPets.OperationID)

	// Verify schemas in components
	require.NotNil(t, doc.Components)
	require.Contains(t, doc.Components.Schemas, "Pet")
	require.Contains(t, doc.Components.Schemas, "Error")
	require.Contains(t, doc.Components.Schemas, "PetList")

	// Verify Pet schema
	petSchema := doc.Components.Schemas["Pet"].Value
	require.NotNil(t, petSchema)
	assert.Equal(t, "object", petSchema.Type.Single)
	require.Contains(t, petSchema.Properties, "id")
	require.Contains(t, petSchema.Properties, "name")
	require.Contains(t, petSchema.Properties, "tag")

	// Verify type array (tag field: [string, null])
	tagSchema := petSchema.Properties["tag"].Value
	require.NotNil(t, tagSchema)
	require.Len(t, tagSchema.Type.Array, 2)
	assert.Contains(t, tagSchema.Type.Array, "string")
	assert.Contains(t, tagSchema.Type.Array, "null")

	// Verify const (status field)
	statusSchema := petSchema.Properties["status"].Value
	require.NotNil(t, statusSchema)
	assert.Equal(t, "available", statusSchema.Const)
	require.Len(t, statusSchema.Enum, 3)

	// Verify if/then/else (color field)
	colorSchema := petSchema.Properties["color"].Value
	require.NotNil(t, colorSchema)
	require.NotNil(t, colorSchema.If)
	require.NotNil(t, colorSchema.Then)
	require.NotNil(t, colorSchema.Else)

	// Verify unevaluatedProperties (metadata field)
	metadataSchema := petSchema.Properties["metadata"].Value
	require.NotNil(t, metadataSchema)
	require.NotNil(t, metadataSchema.UnevaluatedProperties)

	// Verify additionalProperties as boolean
	require.NotNil(t, metadataSchema.AdditionalPropertiesAllowed)
	assert.True(t, *metadataSchema.AdditionalPropertiesAllowed)

	// Verify prefixItems (PetList schema)
	petListSchema := doc.Components.Schemas["PetList"].Value
	require.NotNil(t, petListSchema)
	require.Len(t, petListSchema.PrefixItems, 1)

	// Verify pathItems in components
	require.NotNil(t, doc.Components.PathItems)
	require.Contains(t, doc.Components.PathItems, "PetOperations")

	// Verify tags
	require.Len(t, doc.Tags, 1)
	assert.Equal(t, "pets", doc.Tags[0].Name)

	// Verify externalDocs
	require.NotNil(t, doc.ExternalDocs)
	assert.Equal(t, "https://example.com", doc.ExternalDocs.URL)
}

// TestParseVersion31x tests that versions 3.1.x are accepted.
func TestParseVersion31x(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    bool
	}{
		{"3.1.0", "3.1.0", true},
		{"3.1.1", "3.1.1", true},
		{"3.1.99", "3.1.99", true},
		{"3.2.0", "3.2.0", true},
		{"3.2.1", "3.2.1", true},
		{"3.0.0", "3.0.0", false},
		{"3.0.3", "3.0.3", false},
		{"2.0", "2.0", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := []byte(`openapi: "` + tt.version + `"
info:
  title: Test
  version: "1.0"
paths: {}
`)
			_, err := Parse(data)
			if tt.want {
				assert.NoError(t, err, "version %s should be accepted", tt.version)
			} else {
				assert.Error(t, err, "version %s should be rejected", tt.version)
				assert.Contains(t, err.Error(), "unsupported OpenAPI version")
			}
		})
	}
}

// TestParseInvalidVersion tests that invalid versions are rejected.
func TestParseInvalidVersion(t *testing.T) {
	data := []byte(`openapi: "3.0.3"
info:
  title: Test
  version: "1.0.0"
paths: {}
`)
	_, err := Parse(data)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported OpenAPI version")
}

// TestParseMissingInfo tests that missing info field is an error.
func TestParseMissingInfo(t *testing.T) {
	data := []byte(`openapi: "3.1.0"
paths: {}
`)
	_, err := Parse(data)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "info is required")
}

// TestParseLineColumnNumbers tests that line/column numbers are preserved.
func TestParseLineColumnNumbers(t *testing.T) {
	data := []byte(`openapi: "3.1.0"
info:
  title: Test API
  version: "1.0.0"
paths:
  /test:
    get:
      summary: Test operation
      operationId: testOp
      responses:
        "200":
          description: OK
`)
	doc, err := Parse(data)
	require.NoError(t, err)
	require.NotNil(t, doc)

	// Root starts at line 1
	assert.Equal(t, 1, doc.Trix.Source.Start.Line)

	// Info starts at line 3 (first content line of the mapping)
	require.NotNil(t, doc.Info)
	assert.Equal(t, 3, doc.Info.Trix.Source.Start.Line)

	// First path item
	require.NotNil(t, doc.Paths)
	pathItem := doc.Paths.Items["/test"]
	require.NotNil(t, pathItem)

	// GET operation
	require.NotNil(t, pathItem.Get)
	assert.Equal(t, "testOp", pathItem.Get.OperationID)
}

// TestParseUnknownFields tests unknown field detection.
func TestParseUnknownFields(t *testing.T) {
	data := []byte(`openapi: "3.1.0"
info:
  title: Test API
  version: "1.0.0"
  x-custom: value
  unknownInfoField: something
paths: {}
unknownRootField: something
`)
	result, err := ParseWithUnknownFields(data)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, result.Document)

	// Should detect unknown fields but NOT extensions
	assert.GreaterOrEqual(t, len(result.UnknownFields), 2)

	foundInfoUnknown := false
	foundRootUnknown := false
	for _, f := range result.UnknownFields {
		if f.Key == "unknownInfoField" {
			foundInfoUnknown = true
		}
		if f.Key == "unknownRootField" {
			foundRootUnknown = true
		}
	}
	assert.True(t, foundInfoUnknown, "should detect unknownInfoField")
	assert.True(t, foundRootUnknown, "should detect unknownRootField")

	// x-custom should NOT be in unknown fields
	for _, f := range result.UnknownFields {
		assert.NotEqual(t, "x-custom", f.Key, "extensions should not be reported as unknown")
	}
}

// TestParseExtensions tests that x- extensions are correctly parsed.
func TestParseExtensions(t *testing.T) {
	data := []byte(`openapi: "3.1.0"
info:
  title: Test API
  version: "1.0.0"
  x-info-ext: info-value
x-root-ext: root-value
paths: {}
`)
	doc, err := Parse(data)
	require.NoError(t, err)
	require.NotNil(t, doc)

	assert.Equal(t, "root-value", doc.VendorExtensions["x-root-ext"])
	assert.Equal(t, "info-value", doc.Info.VendorExtensions["x-info-ext"])
}

// TestParseSchemaTypeArray tests parsing of type arrays.
func TestParseSchemaTypeArray(t *testing.T) {
	data := []byte(`openapi: "3.1.0"
info:
  title: Test API
  version: "1.0.0"
paths: {}
components:
  schemas:
    NullableString:
      type:
        - string
        - "null"
    SingleType:
      type: string
    NoType: {}
`)
	doc, err := Parse(data)
	require.NoError(t, err)
	require.NotNil(t, doc)

	// Type array
	nullableStr := doc.Components.Schemas["NullableString"].Value
	require.NotNil(t, nullableStr)
	assert.Empty(t, nullableStr.Type.Single)
	assert.Equal(t, []string{"string", "null"}, nullableStr.Type.Array)
	assert.Equal(t, []string{"string", "null"}, nullableStr.Type.Values())

	// Single type
	singleType := doc.Components.Schemas["SingleType"].Value
	require.NotNil(t, singleType)
	assert.Equal(t, "string", singleType.Type.Single)
	assert.Empty(t, singleType.Type.Array)
	assert.Equal(t, []string{"string"}, singleType.Type.Values())

	// No type
	noType := doc.Components.Schemas["NoType"].Value
	require.NotNil(t, noType)
	assert.True(t, noType.Type.IsEmpty())
	assert.Nil(t, noType.Type.Values())
}

// TestParseRefWithSummaryDescription tests that $ref with summary/description works.
func TestParseRefWithSummaryDescription(t *testing.T) {
	data := []byte(`openapi: "3.1.0"
info:
  title: Test API
  version: "1.0.0"
paths:
  /test:
    get:
      summary: Test
      operationId: test
      parameters:
        - $ref: '#/components/parameters/Limit'
          summary: The limit parameter
          description: Override description
      responses:
        "200":
          $ref: '#/components/responses/Success'
          summary: Overridden success
          description: Overridden success description
components:
  parameters:
    Limit:
      name: limit
      in: query
      schema:
        type: integer
  responses:
    Success:
      description: Success
`)
	doc, err := Parse(data)
	require.NoError(t, err)
	require.NotNil(t, doc)

	// Parameter $ref with summary/description
	params := doc.Paths.Items["/test"].Get.Parameters
	require.Len(t, params, 1)
	assert.Equal(t, "#/components/parameters/Limit", params[0].Ref)
	assert.Equal(t, "The limit parameter", params[0].Summary)
	assert.Equal(t, "Override description", params[0].Description)

	// Response $ref with summary/description
	resp := doc.Paths.Items["/test"].Get.Responses.Codes["200"]
	require.NotNil(t, resp)
	assert.Equal(t, "#/components/responses/Success", resp.Ref)
	assert.Equal(t, "Overridden success", resp.Summary)
	assert.Equal(t, "Overridden success description", resp.Description)
}

// TestParseWebhooks tests webhook parsing.
func TestParseWebhooks(t *testing.T) {
	data := []byte(`openapi: "3.1.0"
info:
  title: Test API
  version: "1.0.0"
webhooks:
  orderCreated:
    post:
      summary: Order created
      operationId: orderCreated
      responses:
        "200":
          description: OK
  orderDeleted:
    $ref: '#/components/pathItems/DeleteWebhook'
    summary: Webhook for deletion
components:
  pathItems:
    DeleteWebhook:
      delete:
        summary: Delete webhook
        operationId: deleteWebhook
        responses:
          "200":
            description: OK
`)
	doc, err := Parse(data)
	require.NoError(t, err)
	require.NotNil(t, doc)

	require.Len(t, doc.Webhooks, 2)

	// Inline webhook
	orderCreated := doc.Webhooks["orderCreated"]
	require.NotNil(t, orderCreated.Value)
	require.NotNil(t, orderCreated.Value.Post)
	assert.Equal(t, "orderCreated", orderCreated.Value.Post.OperationID)

	// $ref webhook
	orderDeleted := doc.Webhooks["orderDeleted"]
	assert.Equal(t, "#/components/pathItems/DeleteWebhook", orderDeleted.Ref)
	assert.Equal(t, "Webhook for deletion", orderDeleted.Summary)
}

// TestParseComponentsPathItems tests components/pathItems parsing.
func TestParseComponentsPathItems(t *testing.T) {
	data := []byte(`openapi: "3.1.0"
info:
  title: Test API
  version: "1.0.0"
paths: {}
components:
  pathItems:
    SharedOps:
      get:
        summary: Shared GET
        operationId: sharedGet
        responses:
          "200":
            description: OK
`)
	doc, err := Parse(data)
	require.NoError(t, err)
	require.NotNil(t, doc)

	require.Contains(t, doc.Components.PathItems, "SharedOps")
	sharedOps := doc.Components.PathItems["SharedOps"]
	require.NotNil(t, sharedOps.Value)
	require.NotNil(t, sharedOps.Value.Get)
	assert.Equal(t, "sharedGet", sharedOps.Value.Get.OperationID)
}
