package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for operation.go - parseOperation function
// =============================================================================

// --- Basic Operation ---

func TestParseOperation_Basic(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      summary: "List pets"
      description: "Returns all pets"
      operationId: "listPets"
      responses:
        "200":
          description: "OK"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	op := doc.Paths.Items["/pets"].Get
	assert.Equal(t, "List pets", op.Summary)
	assert.Equal(t, "Returns all pets", op.Description)
	assert.Equal(t, "listPets", op.OperationID)
}

// --- Tags ---

func TestParseOperation_Tags(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      tags:
        - pets
        - animals
      responses:
        "200":
          description: "OK"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, []string{"pets", "animals"}, doc.Paths.Items["/pets"].Get.Tags)
}

// --- Consumes/Produces ---

func TestParseOperation_ConsumesProduces(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    post:
      consumes:
        - application/json
      produces:
        - application/json
      responses:
        "201":
          description: "Created"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	op := doc.Paths.Items["/pets"].Post
	assert.Equal(t, []string{"application/json"}, op.Consumes)
	assert.Equal(t, []string{"application/json"}, op.Produces)
}

// --- Parameters ---

func TestParseOperation_Parameters(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - name: limit
          in: query
          type: integer
        - name: offset
          in: query
          type: integer
      responses:
        "200":
          description: "OK"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	params := doc.Paths.Items["/pets"].Get.Parameters
	require.Len(t, params, 2)
	assert.Equal(t, "limit", params[0].Value.Name)
	assert.Equal(t, "offset", params[1].Value.Name)
}

// --- Responses ---

func TestParseOperation_Responses(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "Success"
        "404":
          description: "Not found"
        default:
          description: "Error"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	responses := doc.Paths.Items["/pets"].Get.Responses
	require.NotNil(t, responses.Default)
	assert.Equal(t, "Error", responses.Default.Value.Description)
	assert.Equal(t, "Success", responses.Codes["200"].Value.Description)
	assert.Equal(t, "Not found", responses.Codes["404"].Value.Description)
}

// --- Deprecated ---

func TestParseOperation_Deprecated(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      deprecated: true
      responses:
        "200":
          description: "OK"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.True(t, doc.Paths.Items["/pets"].Get.Deprecated)
}

// --- Security ---

func TestParseOperation_Security(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      security:
        - api_key: []
      responses:
        "200":
          description: "OK"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.Len(t, doc.Paths.Items["/pets"].Get.Security, 1)
	assert.Contains(t, doc.Paths.Items["/pets"].Get.Security[0], "api_key")
}

// --- External Docs ---

func TestParseOperation_ExternalDocs(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      externalDocs:
        description: "More info"
        url: "https://example.com"
      responses:
        "200":
          description: "OK"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, doc.Paths.Items["/pets"].Get.ExternalDocs)
	assert.Equal(t, "https://example.com", doc.Paths.Items["/pets"].Get.ExternalDocs.URL)
}

// --- Extensions ---

func TestParseOperation_Extensions(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      x-custom: "value"
      responses:
        "200":
          description: "OK"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, doc.Paths.Items["/pets"].Get.Extensions)
	assert.Equal(t, "value", doc.Paths.Items["/pets"].Get.Extensions["x-custom"])
}
