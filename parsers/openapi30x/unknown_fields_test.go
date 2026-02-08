package openapi30x

import (
	"testing"
)

func TestUnknownFieldDetection(t *testing.T) {
	tests := []struct {
		name           string
		data           string
		wantUnknown    int
		wantUnknownKey string
		wantPath       string
	}{
		{
			name: "unknown field at root level",
			data: `
openapi: "3.0.3"
info:
  title: "Test API"
  version: "1.0.0"
unknownRoot: "this should be detected"
paths: {}
`,
			wantUnknown:    1,
			wantUnknownKey: "unknownRoot",
			wantPath:       "",
		},
		{
			name: "unknown field in info",
			data: `
openapi: "3.0.3"
info:
  title: "Test API"
  version: "1.0.0"
  unknownInfo: "oops"
paths: {}
`,
			wantUnknown:    1,
			wantUnknownKey: "unknownInfo",
			wantPath:       "info",
		},
		{
			name: "extension is not flagged",
			data: `
openapi: "3.0.3"
info:
  title: "Test API"
  version: "1.0.0"
  x-custom-extension: "this is valid"
paths: {}
x-another-extension: "also valid"
`,
			wantUnknown: 0,
		},
		{
			name: "multiple unknown fields",
			data: `
openapi: "3.0.3"
info:
  title: "Test API"
  version: "1.0.0"
  typo1: "unknown"
  typo2: "also unknown"
paths: {}
`,
			wantUnknown: 2,
		},
		{
			name: "unknown field in operation",
			data: `
openapi: "3.0.3"
info:
  title: "Test API"
  version: "1.0.0"
paths:
  /test:
    get:
      summary: "Test endpoint"
      unknownOperation: "detected"
      responses:
        "200":
          description: "OK"
`,
			wantUnknown:    1,
			wantUnknownKey: "unknownOperation",
			wantPath:       "paths./test.get",
		},
		{
			name: "unknown field in schema",
			data: `
openapi: "3.0.3"
info:
  title: "Test API"
  version: "1.0.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      unknownSchemaField: true
`,
			wantUnknown:    1,
			wantUnknownKey: "unknownSchemaField",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseWithUnknownFields([]byte(tt.data))
			if err != nil {
				t.Fatalf("ParseWithUnknownFields failed: %v", err)
			}

			if len(result.UnknownFields) != tt.wantUnknown {
				t.Errorf("got %d unknown fields, want %d", len(result.UnknownFields), tt.wantUnknown)
				for _, f := range result.UnknownFields {
					t.Logf("  unknown: %s at %s (line %d)", f.Key, f.Path, f.Line)
				}
			}

			if tt.wantUnknownKey != "" && len(result.UnknownFields) > 0 {
				found := false
				for _, f := range result.UnknownFields {
					if f.Key == tt.wantUnknownKey {
						found = true
						if tt.wantPath != "" && f.Path != tt.wantPath {
							t.Errorf("unknown field %q has path %q, want %q", f.Key, f.Path, tt.wantPath)
						}
						// Verify line number is set
						if f.Line == 0 {
							t.Errorf("unknown field %q has no line number", f.Key)
						}
					}
				}
				if !found {
					t.Errorf("expected unknown field %q not found", tt.wantUnknownKey)
				}
			}
		})
	}
}

func TestParseWithUnknownFieldsBasic(t *testing.T) {
	// Test that ParseWithUnknownFields returns the document correctly
	data := []byte(`
openapi: "3.0.3"
info:
  title: "Test API"
  version: "1.0.0"
paths: {}
`)
	result, err := ParseWithUnknownFields(data)
	if err != nil {
		t.Fatalf("ParseWithUnknownFields failed: %v", err)
	}

	if result.Document == nil {
		t.Fatal("expected Document to be non-nil")
	}

	if result.Document.OpenAPI != "3.0.3" {
		t.Errorf("expected openapi 3.0.3, got %s", result.Document.OpenAPI)
	}

	if len(result.UnknownFields) != 0 {
		t.Errorf("expected no unknown fields for valid doc, got %d", len(result.UnknownFields))
	}
}

func TestUnknownFieldError(t *testing.T) {
	// Test the error formatting
	fields := []UnknownField{
		{Path: "info", Key: "typo", Line: 5, Column: 3},
	}
	err := &UnknownFieldError{Fields: fields}
	errStr := err.Error()

	if errStr == "" {
		t.Error("expected non-empty error string")
	}

	if !contains(errStr, "typo") || !contains(errStr, "info") {
		t.Errorf("error message should contain field name and path, got: %s", errStr)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
