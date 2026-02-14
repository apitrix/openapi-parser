package openapi30

import (
	"errors"
	"testing"

	"openapi-parser/models/shared"
)

func TestResponse_SetDescription_WithoutHook(t *testing.T) {
	r := NewResponse("old", nil, nil, nil)
	err := r.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if r.Description() != "new" {
		t.Errorf("Description() = %q, want %q", r.Description(), "new")
	}
}

func TestResponse_SetDescription_WithHook_Rejects(t *testing.T) {
	r := NewResponse("old", nil, nil, nil)
	rejectErr := errors.New("rejected")
	r.Trix.OnSet("description", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := r.SetDescription("new")
	if err != rejectErr {
		t.Errorf("SetDescription with rejecting hook should return error, got %v", err)
	}
	if r.Description() != "old" {
		t.Errorf("Description should be unchanged after rejection, got %q", r.Description())
	}
}

func TestResponse_SetHeaders_WithoutHook(t *testing.T) {
	r := NewResponse("", nil, nil, nil)
	headers := map[string]*shared.Ref[Header]{"X-Rate": shared.NewRef[Header]("#/components/headers/X-Rate")}
	err := r.SetHeaders(headers)
	if err != nil {
		t.Fatalf("SetHeaders without hook should succeed, got %v", err)
	}
	if r.Headers()["X-Rate"] != headers["X-Rate"] {
		t.Errorf("Headers() = %v, want %v", r.Headers(), headers)
	}
}

func TestResponse_SetContent_WithoutHook(t *testing.T) {
	r := NewResponse("", nil, nil, nil)
	content := map[string]*MediaType{"application/json": NewMediaType(nil, nil, nil, nil)}
	err := r.SetContent(content)
	if err != nil {
		t.Fatalf("SetContent without hook should succeed, got %v", err)
	}
	if r.Content()["application/json"] != content["application/json"] {
		t.Errorf("Content() = %v, want %v", r.Content(), content)
	}
}

func TestResponse_SetLinks_WithoutHook(t *testing.T) {
	r := NewResponse("", nil, nil, nil)
	links := map[string]*shared.Ref[Link]{"next": shared.NewRef[Link]("#/components/links/next")}
	err := r.SetLinks(links)
	if err != nil {
		t.Fatalf("SetLinks without hook should succeed, got %v", err)
	}
	if r.Links()["next"] != links["next"] {
		t.Errorf("Links() = %v, want %v", r.Links(), links)
	}
}
