package openapi20

import (
	"errors"
	"testing"
)

func TestSchema_SetTitle_WithoutHook(t *testing.T) {
	s := NewSchema(SchemaFields{Title: "old"})
	err := s.SetTitle("new")
	if err != nil {
		t.Fatalf("SetTitle without hook should succeed, got %v", err)
	}
	if s.Title() != "new" {
		t.Errorf("Title() = %q, want %q", s.Title(), "new")
	}
}

func TestSchema_SetTitle_WithHook_Rejects(t *testing.T) {
	s := NewSchema(SchemaFields{Title: "old"})
	rejectErr := errors.New("rejected")
	s.Trix.OnSet("title", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := s.SetTitle("new")
	if err != rejectErr {
		t.Errorf("SetTitle with rejecting hook should return error, got %v", err)
	}
	if s.Title() != "old" {
		t.Errorf("Title should be unchanged after rejection, got %q", s.Title())
	}
}

func TestSchema_SetTitle_WithHook_Passes(t *testing.T) {
	s := NewSchema(SchemaFields{Title: "old"})
	s.Trix.OnSet("title", func(field string, oldVal, newVal interface{}) error {
		return nil
	})
	err := s.SetTitle("new")
	if err != nil {
		t.Fatalf("SetTitle with passing hook should succeed, got %v", err)
	}
	if s.Title() != "new" {
		t.Errorf("Title() = %q, want %q", s.Title(), "new")
	}
}

func TestSchema_SetDescription_WithoutHook(t *testing.T) {
	s := NewSchema(SchemaFields{Description: "old"})
	err := s.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if s.Description() != "new" {
		t.Errorf("Description() = %q, want %q", s.Description(), "new")
	}
}

func TestSchema_SetType_WithoutHook(t *testing.T) {
	s := NewSchema(SchemaFields{Type: "old"})
	err := s.SetType("new")
	if err != nil {
		t.Fatalf("SetType without hook should succeed, got %v", err)
	}
	if s.Type() != "new" {
		t.Errorf("Type() = %q, want %q", s.Type(), "new")
	}
}

func TestSchema_SetFormat_WithoutHook(t *testing.T) {
	s := NewSchema(SchemaFields{Format: "old"})
	err := s.SetFormat("new")
	if err != nil {
		t.Fatalf("SetFormat without hook should succeed, got %v", err)
	}
	if s.Format() != "new" {
		t.Errorf("Format() = %q, want %q", s.Format(), "new")
	}
}

func TestSchema_SetPattern_WithoutHook(t *testing.T) {
	s := NewSchema(SchemaFields{Pattern: "old"})
	err := s.SetPattern("new")
	if err != nil {
		t.Fatalf("SetPattern without hook should succeed, got %v", err)
	}
	if s.Pattern() != "new" {
		t.Errorf("Pattern() = %q, want %q", s.Pattern(), "new")
	}
}

func TestSchema_SetReadOnly_WithoutHook(t *testing.T) {
	s := NewSchema(SchemaFields{})
	err := s.SetReadOnly(true)
	if err != nil {
		t.Fatalf("SetReadOnly without hook should succeed, got %v", err)
	}
	if !s.ReadOnly() {
		t.Errorf("ReadOnly() = false, want true")
	}
}

func TestSchema_SetDiscriminator_WithoutHook(t *testing.T) {
	s := NewSchema(SchemaFields{Discriminator: "old"})
	err := s.SetDiscriminator("new")
	if err != nil {
		t.Fatalf("SetDiscriminator without hook should succeed, got %v", err)
	}
	if s.Discriminator() != "new" {
		t.Errorf("Discriminator() = %q, want %q", s.Discriminator(), "new")
	}
}

func TestSchema_SetRequired_WithoutHook(t *testing.T) {
	s := NewSchema(SchemaFields{})
	err := s.SetRequired([]string{"id", "name"})
	if err != nil {
		t.Fatalf("SetRequired without hook should succeed, got %v", err)
	}
	if len(s.Required()) != 2 || s.Required()[0] != "id" {
		t.Errorf("Required() = %v, want [id name]", s.Required())
	}
}

func TestSchema_SetProperties_WithoutHook(t *testing.T) {
	s := NewSchema(SchemaFields{})
	props := map[string]*RefSchema{"name": NewRefSchema("#/definitions/name")}
	err := s.SetProperties(props)
	if err != nil {
		t.Fatalf("SetProperties without hook should succeed, got %v", err)
	}
	if s.Properties()["name"] != props["name"] {
		t.Errorf("Properties() = %v, want %v", s.Properties(), props)
	}
}

func TestSchema_SetItems_WithoutHook(t *testing.T) {
	s := NewSchema(SchemaFields{})
	ref := NewRefSchema("#/definitions/item")
	err := s.SetItems(ref)
	if err != nil {
		t.Fatalf("SetItems without hook should succeed, got %v", err)
	}
	if s.Items() != ref {
		t.Errorf("Items() = %v, want %v", s.Items(), ref)
	}
}

func TestSchema_SetXML_WithoutHook(t *testing.T) {
	s := NewSchema(SchemaFields{})
	xml := NewXML("Pet", "", "", false, false)
	err := s.SetXML(xml)
	if err != nil {
		t.Fatalf("SetXML without hook should succeed, got %v", err)
	}
	if s.XML() != xml {
		t.Errorf("XML() = %v, want %v", s.XML(), xml)
	}
}

func TestSchema_SetExternalDocs_WithoutHook(t *testing.T) {
	s := NewSchema(SchemaFields{})
	ed := NewExternalDocs("desc", "http://x.com")
	err := s.SetExternalDocs(ed)
	if err != nil {
		t.Fatalf("SetExternalDocs without hook should succeed, got %v", err)
	}
	if s.ExternalDocs() != ed {
		t.Errorf("ExternalDocs() = %v, want %v", s.ExternalDocs(), ed)
	}
}
