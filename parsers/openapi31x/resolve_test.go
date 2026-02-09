package openapi31x

import (
	"testing"

	"openapi-parser/parsers/internal/shared"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

func parseForResolve(t *testing.T, data string) (*ParseResult, *yaml.Node) {
	t.Helper()
	result, err := Parse([]byte(data))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	var rootNode yaml.Node
	if err := yaml.Unmarshal([]byte(data), &rootNode); err != nil {
		t.Fatalf("yaml unmarshal error: %v", err)
	}

	var docNode *yaml.Node
	if rootNode.Kind == yaml.DocumentNode && len(rootNode.Content) > 0 {
		docNode = rootNode.Content[0]
	} else {
		docNode = &rootNode
	}
	return result, docNode
}

func TestResolve_LocalSchemaRef(t *testing.T) {
	spec := `openapi: "3.1.0"
info:
  title: Test
  version: "1.0"
  summary: Test spec
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
	r := shared.NewRefResolverWithFs("/base", docNode, afero.NewMemMapFs())

	if err := resolveDocument(result.Document, r, make(map[string]bool)); err != nil {
		t.Fatalf("resolveDocument() error: %v", err)
	}

	pet := result.Document.Components.Schemas["Pet"]
	if pet == nil || pet.Value == nil {
		t.Fatal("Pet schema should be populated")
	}

	resp := result.Document.Paths.Items["/pets"].Get.Responses.Codes["200"]
	if resp == nil || resp.Value == nil {
		t.Fatal("200 response should be populated")
	}
	schema := resp.Value.Content["application/json"].Schema
	if schema == nil || schema.Value == nil {
		t.Fatal("response schema ref should be resolved")
	}
}

func TestResolve_ExternalFileRef(t *testing.T) {
	fs := afero.NewMemMapFs()

	afero.WriteFile(fs, "/base/models.yaml", []byte(`Tag:
  type: object
  properties:
    id:
      type: integer
    name:
      type: string
`), 0644)

	spec := `openapi: "3.1.0"
info:
  title: Test
  version: "1.0"
  summary: Test spec
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

	if err := resolveDocument(result.Document, r, make(map[string]bool)); err != nil {
		t.Fatalf("resolveDocument() error: %v", err)
	}

	pet := result.Document.Components.Schemas["Pet"]
	if pet == nil || pet.Value == nil {
		t.Fatal("Pet schema should be populated")
	}
	tagRef := pet.Value.Properties["tag"]
	if tagRef == nil {
		t.Fatal("Pet.tag property should exist")
	}
	if tagRef.Value == nil {
		t.Fatal("Pet.tag ref Value should be resolved from external file")
	}
	if tagRef.Value.Properties["name"] == nil {
		t.Error("Tag schema should have 'name' property")
	}
}

func TestResolve_ExternalResponseRef(t *testing.T) {
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

	spec := `openapi: "3.1.0"
info:
  title: Test
  version: "1.0"
  summary: Test spec
paths:
  /pets:
    get:
      operationId: listPets
      responses:
        "404":
          $ref: './responses.yaml#/NotFound'`

	result, docNode := parseForResolve(t, spec)
	r := shared.NewRefResolverWithFs("/base", docNode, fs)

	if err := resolveDocument(result.Document, r, make(map[string]bool)); err != nil {
		t.Fatalf("resolveDocument() error: %v", err)
	}

	resp := result.Document.Paths.Items["/pets"].Get.Responses.Codes["404"]
	if resp == nil {
		t.Fatal("404 response should exist")
	}
	if resp.Value == nil {
		t.Fatal("404 response Value should be resolved from external file")
	}
	if resp.Value.Description != "The requested resource was not found" {
		t.Errorf("unexpected description: %q", resp.Value.Description)
	}
}

func TestResolve_CircularSchemaRef(t *testing.T) {
	spec := `openapi: "3.1.0"
info:
  title: Test
  version: "1.0"
  summary: Test spec
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
	r := shared.NewRefResolverWithFs("/base", docNode, afero.NewMemMapFs())

	if err := resolveDocument(result.Document, r, make(map[string]bool)); err != nil {
		t.Fatalf("resolveDocument() error: %v", err)
	}

	t.Run("TreeNode self-reference", func(t *testing.T) {
		treeNode := result.Document.Components.Schemas["TreeNode"]
		if treeNode == nil || treeNode.Value == nil {
			t.Fatal("TreeNode schema should be populated")
		}
		children := treeNode.Value.Properties["children"]
		if children == nil || children.Value == nil {
			t.Fatal("children property should exist")
		}
		items := children.Value.Items
		if items == nil {
			t.Fatal("children.items should exist")
		}
		if !items.Circular {
			t.Error("TreeNode self-reference should be marked circular")
		}
		if items.Value != nil {
			t.Error("circular ref should not have Value populated")
		}
	})

	t.Run("Person self-reference", func(t *testing.T) {
		person := result.Document.Components.Schemas["Person"]
		if person == nil || person.Value == nil {
			t.Fatal("Person schema should be populated")
		}
		bestFriend := person.Value.Properties["bestFriend"]
		if bestFriend == nil {
			t.Fatal("bestFriend property should exist")
		}
		if !bestFriend.Circular {
			t.Error("Person self-reference should be marked circular")
		}
	})
}
