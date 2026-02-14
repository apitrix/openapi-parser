package openapi20

import (
	"errors"
	"testing"
)

func TestSwagger_SetSwaggerVersion_WithoutHook(t *testing.T) {
	s := &Swagger{}
	s.swagger = "old"
	err := s.SetSwaggerVersion("new")
	if err != nil {
		t.Fatalf("SetSwaggerVersion without hook should succeed, got %v", err)
	}
	if s.SwaggerVersion() != "new" {
		t.Errorf("SwaggerVersion() = %q, want %q", s.SwaggerVersion(), "new")
	}
}

func TestSwagger_SetSwaggerVersion_WithHook_Rejects(t *testing.T) {
	s := &Swagger{}
	s.swagger = "old"
	rejectErr := errors.New("rejected")
	s.Trix.OnSet("swagger", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := s.SetSwaggerVersion("new")
	if err != rejectErr {
		t.Errorf("SetSwaggerVersion with rejecting hook should return error, got %v", err)
	}
	if s.SwaggerVersion() != "old" {
		t.Errorf("SwaggerVersion should be unchanged after rejection, got %q", s.SwaggerVersion())
	}
}

func TestSwagger_SetSwaggerVersion_WithHook_Passes(t *testing.T) {
	s := &Swagger{}
	s.swagger = "old"
	s.Trix.OnSet("swagger", func(field string, oldVal, newVal interface{}) error {
		return nil
	})
	err := s.SetSwaggerVersion("new")
	if err != nil {
		t.Fatalf("SetSwaggerVersion with passing hook should succeed, got %v", err)
	}
	if s.SwaggerVersion() != "new" {
		t.Errorf("SwaggerVersion() = %q, want %q", s.SwaggerVersion(), "new")
	}
}

func TestSwagger_SetInfo_WithoutHook(t *testing.T) {
	s := &Swagger{}
	info := NewInfo("API", "", "", "1.0", nil, nil)
	err := s.SetInfo(info)
	if err != nil {
		t.Fatalf("SetInfo without hook should succeed, got %v", err)
	}
	if s.Info() != info {
		t.Errorf("Info() = %v, want %v", s.Info(), info)
	}
}

func TestSwagger_SetHost_WithoutHook(t *testing.T) {
	s := &Swagger{}
	err := s.SetHost("api.example.com")
	if err != nil {
		t.Fatalf("SetHost without hook should succeed, got %v", err)
	}
	if s.Host() != "api.example.com" {
		t.Errorf("Host() = %q, want api.example.com", s.Host())
	}
}

func TestSwagger_SetBasePath_WithoutHook(t *testing.T) {
	s := &Swagger{}
	err := s.SetBasePath("/v1")
	if err != nil {
		t.Fatalf("SetBasePath without hook should succeed, got %v", err)
	}
	if s.BasePath() != "/v1" {
		t.Errorf("BasePath() = %q, want /v1", s.BasePath())
	}
}

func TestSwagger_SetSchemes_WithoutHook(t *testing.T) {
	s := &Swagger{}
	err := s.SetSchemes([]string{"https"})
	if err != nil {
		t.Fatalf("SetSchemes without hook should succeed, got %v", err)
	}
	if len(s.Schemes()) != 1 || s.Schemes()[0] != "https" {
		t.Errorf("Schemes() = %v, want [https]", s.Schemes())
	}
}

func TestSwagger_SetPaths_WithoutHook(t *testing.T) {
	s := &Swagger{}
	paths := NewPaths(map[string]*PathItem{"/pets": NewPathItem("", nil, nil, nil, nil, nil, nil, nil, nil)})
	err := s.SetPaths(paths)
	if err != nil {
		t.Fatalf("SetPaths without hook should succeed, got %v", err)
	}
	if s.Paths() != paths {
		t.Errorf("Paths() = %v, want %v", s.Paths(), paths)
	}
}

func TestSwagger_SetPaths_WithHook_Rejects(t *testing.T) {
	s := &Swagger{}
	rejectErr := errors.New("rejected")
	s.Trix.OnSet("paths", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	paths := NewPaths(map[string]*PathItem{"/pets": NewPathItem("", nil, nil, nil, nil, nil, nil, nil, nil)})
	err := s.SetPaths(paths)
	if err != rejectErr {
		t.Errorf("SetPaths with rejecting hook should return error, got %v", err)
	}
	if s.Paths() != nil {
		t.Errorf("Paths should be unchanged after rejection")
	}
}

func TestSwagger_SetTags_WithoutHook(t *testing.T) {
	s := &Swagger{}
	tags := []*Tag{NewTag("pets", "", nil)}
	err := s.SetTags(tags)
	if err != nil {
		t.Fatalf("SetTags without hook should succeed, got %v", err)
	}
	if len(s.Tags()) != 1 || s.Tags()[0].Name() != "pets" {
		t.Errorf("Tags() = %v, want [pets]", s.Tags())
	}
}

func TestSwagger_SetExternalDocs_WithoutHook(t *testing.T) {
	s := &Swagger{}
	ed := NewExternalDocs("docs", "http://docs.com")
	err := s.SetExternalDocs(ed)
	if err != nil {
		t.Fatalf("SetExternalDocs without hook should succeed, got %v", err)
	}
	if s.ExternalDocs() != ed {
		t.Errorf("ExternalDocs() = %v, want %v", s.ExternalDocs(), ed)
	}
}
