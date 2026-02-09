package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for link.go - parseLink function
// =============================================================================

// --- Basic Link ---

func TestParseLink_Basic(t *testing.T) {
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
            GetUser:
              operationId: getUser
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	link := result.Document.Paths.Items["/users/{id}"].Get.Responses.Codes["200"].Value.Links["GetUser"].Value
	require.NotNil(t, link)
	assert.Equal(t, "getUser", link.OperationID)
}

// --- OperationRef ---

func TestParseLink_OperationRef(t *testing.T) {
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
              operationRef: '#/paths/~1pets~1{petId}/get'
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	link := result.Document.Paths.Items["/users/{id}"].Get.Responses.Codes["200"].Value.Links["GetUserPets"].Value
	assert.Equal(t, "#/paths/~1pets~1{petId}/get", link.OperationRef)
}

// --- Parameters ---

func TestParseLink_Parameters(t *testing.T) {
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
              operationId: getUserPets
              parameters:
                userId: '$response.body#/id'
                limit: 10
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	link := result.Document.Paths.Items["/users/{id}"].Get.Responses.Codes["200"].Value.Links["GetUserPets"].Value
	assert.Len(t, link.Parameters, 2)
	assert.Contains(t, link.Parameters, "userId")
	assert.Contains(t, link.Parameters, "limit")
}

// --- RequestBody ---

func TestParseLink_RequestBody(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /users:
    post:
      responses:
        "201":
          description: "Created"
          links:
            UpdateUser:
              operationId: updateUser
              requestBody: '$response.body'
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	link := result.Document.Paths.Items["/users"].Post.Responses.Codes["201"].Value.Links["UpdateUser"].Value
	assert.Equal(t, "$response.body", link.RequestBody)
}

// --- Description ---

func TestParseLink_Description(t *testing.T) {
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
              operationId: getUserPets
              description: "Retrieves the pets owned by this user"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	link := result.Document.Paths.Items["/users/{id}"].Get.Responses.Codes["200"].Value.Links["GetUserPets"].Value
	assert.Equal(t, "Retrieves the pets owned by this user", link.Description)
}

// --- Server ---

func TestParseLink_Server(t *testing.T) {
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
              operationId: getUserPets
              server:
                url: https://pets.example.com
                description: "Pets API"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	link := result.Document.Paths.Items["/users/{id}"].Get.Responses.Codes["200"].Value.Links["GetUserPets"].Value
	require.NotNil(t, link.Server)
	assert.Equal(t, "https://pets.example.com", link.Server.URL)
}

// --- Multiple Links ---

func TestParseLink_Multiple(t *testing.T) {
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
              operationId: getUserPets
            GetUserOrders:
              operationId: getUserOrders
            GetUserAddresses:
              operationId: getUserAddresses
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	links := result.Document.Paths.Items["/users/{id}"].Get.Responses.Codes["200"].Value.Links
	assert.Len(t, links, 3)
}

// --- Reference ---

func TestParseLink_Reference(t *testing.T) {
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
	linkRef := result.Document.Paths.Items["/users/{id}"].Get.Responses.Codes["200"].Value.Links["GetUserPets"]
	assert.Equal(t, "#/components/links/GetUserPets", linkRef.Ref)
}

// --- Extensions ---

func TestParseLink_Extensions(t *testing.T) {
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
              operationId: getUserPets
              x-custom: "value"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	link := result.Document.Paths.Items["/users/{id}"].Get.Responses.Codes["200"].Value.Links["GetUserPets"].Value
	require.NotNil(t, link.VendorExtensions)
	assert.Equal(t, "value", link.VendorExtensions["x-custom"])
}
