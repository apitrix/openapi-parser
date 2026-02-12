package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"openapi-parser/parsers/shared"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

// helper to parse YAML into a document and model in one step
func parseForResolve(t *testing.T, data string) (*ParseResult, *yaml.Node) {
	t.Helper()
	result, err := Parse([]byte(data))
	require.NoError(t, err, "parse error")

	var rootNode yaml.Node
	err = yaml.Unmarshal([]byte(data), &rootNode)
	require.NoError(t, err, "yaml unmarshal error")

	var docNode *yaml.Node
	if rootNode.Kind == yaml.DocumentNode && len(rootNode.Content) > 0 {
		docNode = rootNode.Content[0]
	} else {
		docNode = &rootNode
	}
	return result, docNode
}

func TestResolve_LocalSchemaRef(t *testing.T) {
	// Arrange
	spec := `openapi: "3.0.3"
info:
  title: Test
  version: "1.0"
paths:
  /pets:
    get:
      operationId: listPets
      responses:
        "200":
          description: ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Pet'
components:
  schemas:
    Pet:
      type: object
      properties:
        name:
          type: string`

	result, docNode := parseForResolve(t, spec)
	_ = shared.NewRefResolverWithFs("/base", docNode, afero.NewMemMapFs())

	// Act
	err := Resolve(result.Document, docNode, "/base")

	// Assert
	require.NoError(t, err)

	pet := result.Document.Components().Schemas()["Pet"]
	require.NotNil(t, pet)
	require.NotNil(t, pet.Value(), "Pet schema should be populated")

	resp := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"]
	require.NotNil(t, resp)
	require.NotNil(t, resp.Value(), "200 response should be populated")

	schema := resp.Value().Content()["application/json"].Schema()
	require.NotNil(t, schema, "response schema should exist")
	assert.NotNil(t, schema.Value(), "schema ref Value should be resolved")
}

func TestResolve_ExternalFileRef(t *testing.T) {
	// Arrange
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/base/models.yaml", []byte(`Tag:
  type: object
  properties:
    id:
      type: integer
    name:
      type: string
`), 0644)

	spec := `openapi: "3.0.3"
info:
  title: Test
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      properties:
        tag:
          $ref: './models.yaml#/Tag'`

	result, docNode := parseForResolve(t, spec)
	r := shared.NewRefResolverWithFs("/base", docNode, fs)

	// Act
	err := resolveDocument(result.Document, r, make(map[string]bool))

	// Assert
	require.NoError(t, err)

	pet := result.Document.Components().Schemas()["Pet"]
	require.NotNil(t, pet)
	require.NotNil(t, pet.Value(), "Pet schema should be populated")

	tagRef := pet.Value().Properties()["tag"]
	require.NotNil(t, tagRef, "Pet.tag property should exist")
	require.NotNil(t, tagRef.Value(), "Pet.tag ref Value should be resolved from external file")
	assert.NotNil(t, tagRef.Value().Properties()["name"], "Tag schema should have 'name' property")
	assert.NotNil(t, tagRef.Value().Properties()["id"], "Tag schema should have 'id' property")
}

func TestResolve_ExternalResponseRef(t *testing.T) {
	// Arrange
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/base/responses.yaml", []byte(`NotFound:
  description: The requested resource was not found
  content:
    application/json:
      schema:
        type: object
        properties:
          error:
            type: string
`), 0644)

	spec := `openapi: "3.0.3"
info:
  title: Test
  version: "1.0"
paths:
  /pets:
    get:
      operationId: listPets
      responses:
        "404":
          $ref: './responses.yaml#/NotFound'`

	result, docNode := parseForResolve(t, spec)
	r := shared.NewRefResolverWithFs("/base", docNode, fs)

	// Act
	err := resolveDocument(result.Document, r, make(map[string]bool))

	// Assert
	require.NoError(t, err)

	resp := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["404"]
	require.NotNil(t, resp, "404 response should exist")
	require.NotNil(t, resp.Value(), "404 response Value should be resolved from external file")
	assert.Equal(t, "The requested resource was not found", resp.Value().Description())
}

func TestResolve_CircularSchemaRef(t *testing.T) {
	// Arrange
	spec := `openapi: "3.0.3"
info:
  title: Test
  version: "1.0"
paths: {}
components:
  schemas:
    TreeNode:
      type: object
      properties:
        value:
          type: string
        children:
          type: array
          items:
            $ref: '#/components/schemas/TreeNode'
    Person:
      type: object
      properties:
        name:
          type: string
        bestFriend:
          $ref: '#/components/schemas/Person'`

	result, docNode := parseForResolve(t, spec)
	_ = shared.NewRefResolverWithFs("/base", docNode, afero.NewMemMapFs())

	// Act
	err := resolveDocument(result.Document, shared.NewRefResolverWithFs("/base", docNode, afero.NewMemMapFs()), make(map[string]bool))

	// Assert
	require.NoError(t, err)

	t.Run("TreeNode self-reference", func(t *testing.T) {
		// Assert
		treeNode := result.Document.Components().Schemas()["TreeNode"]
		require.NotNil(t, treeNode)
		require.NotNil(t, treeNode.Value(), "TreeNode schema should be populated")

		children := treeNode.Value().Properties()["children"]
		require.NotNil(t, children)
		require.NotNil(t, children.Value(), "children property should be resolved")

		items := children.Value().Items()
		require.NotNil(t, items, "children.items should exist")
		assert.True(t, items.Circular(), "TreeNode self-reference should be marked circular")
		assert.Nil(t, items.Value(), "circular ref should not have Value populated")
	})

	t.Run("Person self-reference", func(t *testing.T) {
		// Assert
		person := result.Document.Components().Schemas()["Person"]
		require.NotNil(t, person)
		require.NotNil(t, person.Value(), "Person schema should be populated")

		bestFriend := person.Value().Properties()["bestFriend"]
		require.NotNil(t, bestFriend, "bestFriend property should exist")
		assert.True(t, bestFriend.Circular(), "Person self-reference should be marked circular")
	})
}
