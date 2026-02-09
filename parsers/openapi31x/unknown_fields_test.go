package openapi31x

import (
	"strings"
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
openapi: "3.1.0"
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
openapi: "3.1.0"
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
openapi: "3.1.0"
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
openapi: "3.1.0"
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
openapi: "3.1.0"
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
openapi: "3.1.0"
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
			result, err := Parse([]byte(tt.data))
			if err != nil {
				t.Fatalf("Parse failed: %v", err)
			}

			unknownErrors := filterUnknownFieldErrors(result)
			if len(unknownErrors) != tt.wantUnknown {
				t.Errorf("got %d unknown field errors, want %d", len(unknownErrors), tt.wantUnknown)
				for _, e := range unknownErrors {
					t.Logf("  unknown: %s at %s", e.Message, strings.Join(e.Path, "."))
				}
			}

			if tt.wantUnknownKey != "" && len(unknownErrors) > 0 {
				found := false
				for _, e := range unknownErrors {
					if strings.Contains(e.Message, tt.wantUnknownKey) {
						found = true
						errPath := strings.Join(e.Path, ".")
						if tt.wantPath != "" && errPath != tt.wantPath {
							t.Errorf("unknown field error for %q has path %q, want %q", tt.wantUnknownKey, errPath, tt.wantPath)
						}
					}
				}
				if !found {
					t.Errorf("expected unknown field error containing %q not found", tt.wantUnknownKey)
				}
			}
		})
	}
}

func TestParseBasic(t *testing.T) {
	// Test that Parse returns the document correctly
	data := []byte(`
openapi: "3.1.0"
info:
  title: "Test API"
  version: "1.0.0"
paths: {}
`)
	result, err := Parse(data)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if result.Document == nil {
		t.Fatal("expected Document to be non-nil")
	}

	if result.Document.OpenAPI != "3.1.0" {
		t.Errorf("expected openapi 3.1.0, got %s", result.Document.OpenAPI)
	}

	unknownErrors := filterUnknownFieldErrors(result)
	if len(unknownErrors) != 0 {
		t.Errorf("expected no unknown field errors for valid doc, got %d", len(unknownErrors))
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

	if !strings.Contains(errStr, "typo") || !strings.Contains(errStr, "info") {
		t.Errorf("error message should contain field name and path, got: %s", errStr)
	}
}

// filterUnknownFieldErrors returns only errors with Kind "unknown_field" from a ParseResult.
func filterUnknownFieldErrors(result *ParseResult) []*ParseError {
	var filtered []*ParseError
	for _, e := range result.Errors {
		if e.Kind == "unknown_field" {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
