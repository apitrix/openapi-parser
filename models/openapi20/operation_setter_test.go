package openapi20

import (
	"errors"
	"testing"
)

func TestOperation_SetTags_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	err := op.SetTags([]string{"pets", "store"})
	if err != nil {
		t.Fatalf("SetTags without hook should succeed, got %v", err)
	}
	if len(op.Tags()) != 2 || op.Tags()[0] != "pets" || op.Tags()[1] != "store" {
		t.Errorf("Tags() = %v, want [pets store]", op.Tags())
	}
}

func TestOperation_SetTags_WithHook_Rejects(t *testing.T) {
	op := NewOperation([]string{"old"}, "", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	rejectErr := errors.New("rejected")
	op.Trix.OnSet("tags", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := op.SetTags([]string{"new"})
	if err != rejectErr {
		t.Errorf("SetTags with rejecting hook should return error, got %v", err)
	}
	if len(op.Tags()) != 1 || op.Tags()[0] != "old" {
		t.Errorf("Tags should be unchanged after rejection, got %v", op.Tags())
	}
}

func TestOperation_SetTags_WithHook_Passes(t *testing.T) {
	op := NewOperation([]string{"old"}, "", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	op.Trix.OnSet("tags", func(field string, oldVal, newVal interface{}) error {
		return nil
	})
	err := op.SetTags([]string{"new"})
	if err != nil {
		t.Fatalf("SetTags with passing hook should succeed, got %v", err)
	}
	if len(op.Tags()) != 1 || op.Tags()[0] != "new" {
		t.Errorf("Tags() = %v, want [new]", op.Tags())
	}
}

func TestOperation_SetSummary_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "old", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	err := op.SetSummary("new")
	if err != nil {
		t.Fatalf("SetSummary without hook should succeed, got %v", err)
	}
	if op.Summary() != "new" {
		t.Errorf("Summary() = %q, want %q", op.Summary(), "new")
	}
}

func TestOperation_SetSummary_WithHook_Rejects(t *testing.T) {
	op := NewOperation(nil, "old", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	rejectErr := errors.New("rejected")
	op.Trix.OnSet("summary", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := op.SetSummary("new")
	if err != rejectErr {
		t.Errorf("SetSummary with rejecting hook should return error, got %v", err)
	}
	if op.Summary() != "old" {
		t.Errorf("Summary should be unchanged after rejection, got %q", op.Summary())
	}
}

func TestOperation_SetSummary_WithHook_Passes(t *testing.T) {
	op := NewOperation(nil, "old", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	op.Trix.OnSet("summary", func(field string, oldVal, newVal interface{}) error {
		return nil
	})
	err := op.SetSummary("new")
	if err != nil {
		t.Fatalf("SetSummary with passing hook should succeed, got %v", err)
	}
	if op.Summary() != "new" {
		t.Errorf("Summary() = %q, want %q", op.Summary(), "new")
	}
}

func TestOperation_SetDescription_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "old", nil, "", nil, nil, nil, nil, nil, false, nil)
	err := op.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if op.Description() != "new" {
		t.Errorf("Description() = %q, want %q", op.Description(), "new")
	}
}

func TestOperation_SetOperationID_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "old", nil, nil, nil, nil, nil, false, nil)
	err := op.SetOperationID("new")
	if err != nil {
		t.Fatalf("SetOperationID without hook should succeed, got %v", err)
	}
	if op.OperationID() != "new" {
		t.Errorf("OperationID() = %q, want %q", op.OperationID(), "new")
	}
}

func TestOperation_SetConsumes_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "", []string{"old"}, nil, nil, nil, nil, false, nil)
	err := op.SetConsumes([]string{"new"})
	if err != nil {
		t.Fatalf("SetConsumes without hook should succeed, got %v", err)
	}
	if len(op.Consumes()) != 1 || op.Consumes()[0] != "new" {
		t.Errorf("Consumes() = %v, want [new]", op.Consumes())
	}
}

func TestOperation_SetProduces_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "", nil, []string{"old"}, nil, nil, nil, false, nil)
	err := op.SetProduces([]string{"new"})
	if err != nil {
		t.Fatalf("SetProduces without hook should succeed, got %v", err)
	}
	if len(op.Produces()) != 1 || op.Produces()[0] != "new" {
		t.Errorf("Produces() = %v, want [new]", op.Produces())
	}
}

func TestOperation_SetSchemes_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, nil, []string{"old"}, false, nil)
	err := op.SetSchemes([]string{"new"})
	if err != nil {
		t.Fatalf("SetSchemes without hook should succeed, got %v", err)
	}
	if len(op.Schemes()) != 1 || op.Schemes()[0] != "new" {
		t.Errorf("Schemes() = %v, want [new]", op.Schemes())
	}
}

func TestOperation_SetDeprecated_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	err := op.SetDeprecated(true)
	if err != nil {
		t.Fatalf("SetDeprecated without hook should succeed, got %v", err)
	}
	if !op.Deprecated() {
		t.Errorf("Deprecated() = false, want true")
	}
}

func TestOperation_SetDeprecated_WithHook_Rejects(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	rejectErr := errors.New("rejected")
	op.Trix.OnSet("deprecated", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := op.SetDeprecated(true)
	if err != rejectErr {
		t.Errorf("SetDeprecated with rejecting hook should return error, got %v", err)
	}
	if op.Deprecated() {
		t.Errorf("Deprecated should be unchanged after rejection")
	}
}

func TestOperation_SetExternalDocs_WithoutHook(t *testing.T) {
	oldEd := NewExternalDocs("old", "http://old.com")
	newEd := NewExternalDocs("new", "http://new.com")
	op := NewOperation(nil, "", "", oldEd, "", nil, nil, nil, nil, nil, false, nil)
	err := op.SetExternalDocs(newEd)
	if err != nil {
		t.Fatalf("SetExternalDocs without hook should succeed, got %v", err)
	}
	if op.ExternalDocs() != newEd {
		t.Errorf("ExternalDocs() = %v, want %v", op.ExternalDocs(), newEd)
	}
}

func TestOperation_SetParameters_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	params := []*RefParameter{NewRefParameter("#/params/1")}
	err := op.SetParameters(params)
	if err != nil {
		t.Fatalf("SetParameters without hook should succeed, got %v", err)
	}
	if len(op.Parameters()) != 1 {
		t.Errorf("Parameters() len = %d, want 1", len(op.Parameters()))
	}
}

func TestOperation_SetResponses_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	resp := NewResponses(nil, map[string]*RefResponse{"200": NewRefResponse("#/responses/ok")})
	err := op.SetResponses(resp)
	if err != nil {
		t.Fatalf("SetResponses without hook should succeed, got %v", err)
	}
	if op.Responses() != resp {
		t.Errorf("Responses() = %v, want %v", op.Responses(), resp)
	}
}

func TestOperation_SetSecurity_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	sec := []SecurityRequirement{{"apiKey": []string{}}}
	err := op.SetSecurity(sec)
	if err != nil {
		t.Fatalf("SetSecurity without hook should succeed, got %v", err)
	}
	if len(op.Security()) != 1 {
		t.Errorf("Security() len = %d, want 1", len(op.Security()))
	}
}
