package openapi20

import (
	"testing"

	"openapi-parser/parsers/shared"

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
	spec := `swagger: "2.0"
info:
  title: Test
  version: "1.0"
basePath: /
paths:
  /pets:
    get:
      operationId: listPets
      responses:
        "200":
          description: ok
          schema:
            $ref: '#/definitions/Pet'
definitions:
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

	pet := result.Document.Definitions()["Pet"]
	if pet == nil || pet.Value() == nil {
		t.Fatal("Pet definition should be populated")
	}

	resp := result.Document.Paths().Items()["/pets"].Get().Responses().Codes()["200"]
	if resp == nil || resp.Value() == nil {
		t.Fatal("200 response should exist")
	}
	schema := resp.Value().Schema()
	if schema == nil || schema.Value() == nil {
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

	spec := `swagger: "2.0"
info:
  title: Test
  version: "1.0"
basePath: /
paths: {}
definitions:
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

	pet := result.Document.Definitions()["Pet"]
	if pet == nil || pet.Value() == nil {
		t.Fatal("Pet definition should be populated")
	}
	tagRef := pet.Value().Properties()["tag"]
	if tagRef == nil {
		t.Fatal("Pet.tag property should exist")
	}
	if tagRef.Value() == nil {
		t.Fatal("Pet.tag ref Value should be resolved from external file")
	}
	if tagRef.Value().Properties()["name"] == nil {
		t.Error("Tag schema should have 'name' property")
	}
}

func TestResolve_CircularSchemaRef(t *testing.T) {
	spec := `swagger: "2.0"
info:
  title: Test
  version: "1.0"
basePath: /
paths: {}
definitions:
  TreeNode:
    type: object
    properties:
      value:
        type: string
      children:
        type: array
        items:
          $ref: '#/definitions/TreeNode'
  Person:
    type: object
    properties:
      name:
        type: string
      bestFriend:
        $ref: '#/definitions/Person'`

	result, docNode := parseForResolve(t, spec)
	r := shared.NewRefResolverWithFs("/base", docNode, afero.NewMemMapFs())

	if err := resolveDocument(result.Document, r, make(map[string]bool)); err != nil {
		t.Fatalf("resolveDocument() error: %v", err)
	}

	t.Run("TreeNode self-reference", func(t *testing.T) {
		treeNode := result.Document.Definitions()["TreeNode"]
		if treeNode == nil || treeNode.Value() == nil {
			t.Fatal("TreeNode schema should be populated")
		}
		children := treeNode.Value().Properties()["children"]
		if children == nil || children.Value() == nil {
			t.Fatal("children property should exist")
		}
		items := children.Value().Items()
		if items == nil {
			t.Fatal("children.items should exist")
		}
		if !items.Circular() {
			t.Error("TreeNode self-reference should be marked circular")
		}
	})

	t.Run("Person self-reference", func(t *testing.T) {
		person := result.Document.Definitions()["Person"]
		if person == nil || person.Value() == nil {
			t.Fatal("Person schema should be populated")
		}
		bestFriend := person.Value().Properties()["bestFriend"]
		if bestFriend == nil {
			t.Fatal("bestFriend property should exist")
		}
		if !bestFriend.Circular() {
			t.Error("Person self-reference should be marked circular")
		}
	})
}
