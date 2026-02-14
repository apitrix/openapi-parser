package openapi30

import (
	"errors"
	"testing"
)

func TestComponents_SetSchemas_WithoutHook(t *testing.T) {
	c := NewComponents(nil, nil, nil, nil, nil, nil, nil, nil, nil)
	schemas := map[string]*RefSchema{"Pet": NewRefSchema("#/components/schemas/Pet")}
	err := c.SetSchemas(schemas)
	if err != nil {
		t.Fatalf("SetSchemas without hook should succeed, got %v", err)
	}
	if c.Schemas()["Pet"] != schemas["Pet"] {
		t.Errorf("Schemas() = %v, want %v", c.Schemas(), schemas)
	}
}

func TestComponents_SetSchemas_WithHook_Rejects(t *testing.T) {
	c := NewComponents(nil, nil, nil, nil, nil, nil, nil, nil, nil)
	rejectErr := errors.New("rejected")
	c.Trix.OnSet("schemas", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	schemas := map[string]*RefSchema{"Pet": NewRefSchema("#/components/schemas/Pet")}
	err := c.SetSchemas(schemas)
	if err != rejectErr {
		t.Errorf("SetSchemas with rejecting hook should return error, got %v", err)
	}
	if c.Schemas() != nil {
		t.Errorf("Schemas should be unchanged after rejection")
	}
}

func TestComponents_SetResponses_WithoutHook(t *testing.T) {
	c := NewComponents(nil, nil, nil, nil, nil, nil, nil, nil, nil)
	resp := map[string]*RefResponse{"Ok": NewRefResponse("#/components/responses/Ok")}
	err := c.SetResponses(resp)
	if err != nil {
		t.Fatalf("SetResponses without hook should succeed, got %v", err)
	}
	if c.Responses()["Ok"] != resp["Ok"] {
		t.Errorf("Responses() = %v, want %v", c.Responses(), resp)
	}
}

func TestComponents_SetParameters_WithoutHook(t *testing.T) {
	c := NewComponents(nil, nil, nil, nil, nil, nil, nil, nil, nil)
	params := map[string]*RefParameter{"limit": NewRefParameter("#/components/parameters/limit")}
	err := c.SetParameters(params)
	if err != nil {
		t.Fatalf("SetParameters without hook should succeed, got %v", err)
	}
	if c.Parameters()["limit"] != params["limit"] {
		t.Errorf("Parameters() = %v, want %v", c.Parameters(), params)
	}
}

func TestComponents_SetHeaders_WithoutHook(t *testing.T) {
	c := NewComponents(nil, nil, nil, nil, nil, nil, nil, nil, nil)
	headers := map[string]*RefHeader{"X-Rate": NewRefHeader("#/components/headers/X-Rate")}
	err := c.SetHeaders(headers)
	if err != nil {
		t.Fatalf("SetHeaders without hook should succeed, got %v", err)
	}
	if c.Headers()["X-Rate"] != headers["X-Rate"] {
		t.Errorf("Headers() = %v, want %v", c.Headers(), headers)
	}
}

func TestComponents_SetSecuritySchemes_WithoutHook(t *testing.T) {
	c := NewComponents(nil, nil, nil, nil, nil, nil, nil, nil, nil)
	ss := map[string]*RefSecurityScheme{"apiKey": NewRefSecurityScheme("#/components/securitySchemes/apiKey")}
	err := c.SetSecuritySchemes(ss)
	if err != nil {
		t.Fatalf("SetSecuritySchemes without hook should succeed, got %v", err)
	}
	if c.SecuritySchemes()["apiKey"] != ss["apiKey"] {
		t.Errorf("SecuritySchemes() = %v, want %v", c.SecuritySchemes(), ss)
	}
}

func TestComponents_SetLinks_WithoutHook(t *testing.T) {
	c := NewComponents(nil, nil, nil, nil, nil, nil, nil, nil, nil)
	links := map[string]*RefLink{"next": NewRefLink("#/components/links/next")}
	err := c.SetLinks(links)
	if err != nil {
		t.Fatalf("SetLinks without hook should succeed, got %v", err)
	}
	if c.Links()["next"] != links["next"] {
		t.Errorf("Links() = %v, want %v", c.Links(), links)
	}
}

func TestComponents_SetCallbacks_WithoutHook(t *testing.T) {
	c := NewComponents(nil, nil, nil, nil, nil, nil, nil, nil, nil)
	cb := map[string]*RefCallback{"ev": NewRefCallback("#/components/callbacks/ev")}
	err := c.SetCallbacks(cb)
	if err != nil {
		t.Fatalf("SetCallbacks without hook should succeed, got %v", err)
	}
	if c.Callbacks()["ev"] != cb["ev"] {
		t.Errorf("Callbacks() = %v, want %v", c.Callbacks(), cb)
	}
}
