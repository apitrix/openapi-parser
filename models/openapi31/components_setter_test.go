package openapi31

import (
	"errors"
	"testing"

	"openapi-parser/models/shared"
)

func TestComponents_SetSchemas_WithoutHook(t *testing.T) {
	c := NewComponents()
	schemas := map[string]*shared.RefWithMeta[Schema]{"Pet": shared.NewRefWithMeta[Schema]("#/components/schemas/Pet")}
	err := c.SetSchemas(schemas)
	if err != nil {
		t.Fatalf("SetSchemas without hook should succeed, got %v", err)
	}
	if c.Schemas()["Pet"] != schemas["Pet"] {
		t.Errorf("Schemas() = %v, want %v", c.Schemas(), schemas)
	}
}

func TestComponents_SetSchemas_WithHook_Rejects(t *testing.T) {
	c := NewComponents()
	rejectErr := errors.New("rejected")
	c.Trix.OnSet("schemas", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	schemas := map[string]*shared.RefWithMeta[Schema]{"Pet": shared.NewRefWithMeta[Schema]("#/components/schemas/Pet")}
	err := c.SetSchemas(schemas)
	if err != rejectErr {
		t.Errorf("SetSchemas with rejecting hook should return error, got %v", err)
	}
	if c.Schemas() != nil {
		t.Errorf("Schemas should be unchanged after rejection")
	}
}

func TestComponents_SetResponses_WithoutHook(t *testing.T) {
	c := NewComponents()
	resp := map[string]*shared.RefWithMeta[Response]{"Ok": shared.NewRefWithMeta[Response]("#/components/responses/Ok")}
	err := c.SetResponses(resp)
	if err != nil {
		t.Fatalf("SetResponses without hook should succeed, got %v", err)
	}
	if c.Responses()["Ok"] != resp["Ok"] {
		t.Errorf("Responses() = %v, want %v", c.Responses(), resp)
	}
}

func TestComponents_SetParameters_WithoutHook(t *testing.T) {
	c := NewComponents()
	params := map[string]*shared.RefWithMeta[Parameter]{"limit": shared.NewRefWithMeta[Parameter]("#/components/parameters/limit")}
	err := c.SetParameters(params)
	if err != nil {
		t.Fatalf("SetParameters without hook should succeed, got %v", err)
	}
	if c.Parameters()["limit"] != params["limit"] {
		t.Errorf("Parameters() = %v, want %v", c.Parameters(), params)
	}
}

func TestComponents_SetPathItems_WithoutHook(t *testing.T) {
	c := NewComponents()
	items := map[string]*shared.RefWithMeta[PathItem]{"PetPath": shared.NewRefWithMeta[PathItem]("#/components/pathItems/PetPath")}
	err := c.SetPathItems(items)
	if err != nil {
		t.Fatalf("SetPathItems without hook should succeed, got %v", err)
	}
	if c.PathItems()["PetPath"] != items["PetPath"] {
		t.Errorf("PathItems() = %v, want %v", c.PathItems(), items)
	}
}
