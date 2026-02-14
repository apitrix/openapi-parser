package openapi31

import (
	"errors"
	"testing"
)

func TestInfo_SetTitle_WithoutHook(t *testing.T) {
	i := NewInfo("old", "", "", "", "1.0", nil, nil)
	err := i.SetTitle("new")
	if err != nil {
		t.Fatalf("SetTitle without hook should succeed, got %v", err)
	}
	if i.Title() != "new" {
		t.Errorf("Title() = %q, want %q", i.Title(), "new")
	}
}

func TestInfo_SetTitle_WithHook_Rejects(t *testing.T) {
	i := NewInfo("old", "", "", "", "1.0", nil, nil)
	rejectErr := errors.New("rejected")
	i.Trix.OnSet("title", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := i.SetTitle("new")
	if err != rejectErr {
		t.Errorf("SetTitle with rejecting hook should return error, got %v", err)
	}
	if i.Title() != "old" {
		t.Errorf("Title should be unchanged after rejection, got %q", i.Title())
	}
}

func TestInfo_SetSummary_WithoutHook(t *testing.T) {
	i := NewInfo("", "old", "", "", "1.0", nil, nil)
	err := i.SetSummary("new")
	if err != nil {
		t.Fatalf("SetSummary without hook should succeed, got %v", err)
	}
	if i.Summary() != "new" {
		t.Errorf("Summary() = %q, want %q", i.Summary(), "new")
	}
}

func TestInfo_SetDescription_WithoutHook(t *testing.T) {
	i := NewInfo("", "", "old", "", "1.0", nil, nil)
	err := i.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if i.Description() != "new" {
		t.Errorf("Description() = %q, want %q", i.Description(), "new")
	}
}

func TestInfo_SetVersion_WithoutHook(t *testing.T) {
	i := NewInfo("", "", "", "", "old", nil, nil)
	err := i.SetVersion("new")
	if err != nil {
		t.Fatalf("SetVersion without hook should succeed, got %v", err)
	}
	if i.Version() != "new" {
		t.Errorf("Version() = %q, want %q", i.Version(), "new")
	}
}

func TestInfo_SetContact_WithoutHook(t *testing.T) {
	i := NewInfo("", "", "", "", "1.0", nil, nil)
	c := NewContact("x", "http://x.com", "x@x.com")
	err := i.SetContact(c)
	if err != nil {
		t.Fatalf("SetContact without hook should succeed, got %v", err)
	}
	if i.Contact() != c {
		t.Errorf("Contact() = %v, want %v", i.Contact(), c)
	}
}

func TestInfo_SetLicense_WithoutHook(t *testing.T) {
	i := NewInfo("", "", "", "", "1.0", nil, nil)
	lic := NewLicense("MIT", "MIT", "http://license.com")
	err := i.SetLicense(lic)
	if err != nil {
		t.Fatalf("SetLicense without hook should succeed, got %v", err)
	}
	if i.License() != lic {
		t.Errorf("License() = %v, want %v", i.License(), lic)
	}
}
