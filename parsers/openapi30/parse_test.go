package openapi30

import (
	"os"
	"testing"
)

func TestParsePetstore(t *testing.T) {
	data, err := os.ReadFile("testdata/petstore.yaml")
	if err != nil {
		t.Fatalf("failed to read petstore.yaml: %v", err)
	}

	doc, err := Parse(data)
	if err != nil {
		t.Fatalf("failed to parse petstore.yaml: %v", err)
	}

	// Verify basic structure
	if doc.OpenAPI != "3.0.3" {
		t.Errorf("expected openapi version 3.0.3, got %s", doc.OpenAPI)
	}

	if doc.Info == nil {
		t.Fatal("expected info to be non-nil")
	}

	if doc.Info.Title != "Petstore API" {
		t.Errorf("expected title 'Petstore API', got %s", doc.Info.Title)
	}

	if doc.Info.Version != "1.0.0" {
		t.Errorf("expected version '1.0.0', got %s", doc.Info.Version)
	}

	if doc.Paths == nil {
		t.Fatal("expected paths to be non-nil")
	}

	if len(doc.Paths.Items) == 0 {
		t.Error("expected at least one path")
	}

	// Check /pets path exists
	if _, ok := doc.Paths.Items["/pets"]; !ok {
		t.Error("expected /pets path to exist")
	}

	// Check servers
	if len(doc.Servers) != 2 {
		t.Errorf("expected 2 servers, got %d", len(doc.Servers))
	}
}

func TestParseSimple(t *testing.T) {
	data, err := os.ReadFile("testdata/simple.json")
	if err != nil {
		t.Fatalf("failed to read simple.json: %v", err)
	}

	doc, err := Parse(data)
	if err != nil {
		t.Fatalf("failed to parse simple.json: %v", err)
	}

	if doc.OpenAPI != "3.0.3" {
		t.Errorf("expected openapi version 3.0.3, got %s", doc.OpenAPI)
	}

	if doc.Info.Title != "Simple API" {
		t.Errorf("expected title 'Simple API', got %s", doc.Info.Title)
	}
}

func TestParseInvalidVersion(t *testing.T) {
	data := []byte(`{"openapi": "2.0", "info": {"title": "Test", "version": "1.0"}}`)

	_, err := Parse(data)
	if err == nil {
		t.Error("expected error for invalid version")
	}
}

func TestParseMissingInfo(t *testing.T) {
	data := []byte(`{"openapi": "3.0.0"}`)

	_, err := Parse(data)
	if err == nil {
		t.Error("expected error for missing info")
	}
}

func TestParseLineColumnNumbers(t *testing.T) {
	// YAML document with known positions
	data := []byte(`openapi: "3.0.3"
info:
  title: "Test API"
  version: "1.0.0"
paths:
  /pets:
    get:
      summary: "Get pets"
      responses:
        "200":
          description: "Success"
`)

	doc, err := Parse(data)
	if err != nil {
		t.Fatalf("failed to parse YAML: %v", err)
	}

	// Verify line numbers are captured (non-zero)
	if doc.NodeSource.Start.Line == 0 {
		t.Error("expected root line number to be non-zero")
	}

	if doc.Info.NodeSource.Start.Line == 0 {
		t.Error("expected info line number to be non-zero")
	}

	if doc.Paths.NodeSource.Start.Line == 0 {
		t.Error("expected paths line number to be non-zero")
	}

	// /pets path item
	petsPath := doc.Paths.Items["/pets"]
	if petsPath == nil {
		t.Fatal("expected /pets path to exist")
	}
	if petsPath.NodeSource.Start.Line == 0 {
		t.Error("expected /pets line number to be non-zero")
	}

	// GET operation
	if petsPath.Get.NodeSource.Start.Line == 0 {
		t.Error("expected GET operation line number to be non-zero")
	}

	// Column numbers should also be captured
	if doc.NodeSource.Start.Column == 0 {
		t.Error("expected root column number to be non-zero")
	}

	t.Logf("Line/column tracking verified:")
	t.Logf("  root:     line %d, col %d", doc.NodeSource.Start.Line, doc.NodeSource.Start.Column)
	t.Logf("  info:     line %d, col %d", doc.Info.NodeSource.Start.Line, doc.Info.NodeSource.Start.Column)
	t.Logf("  paths:    line %d, col %d", doc.Paths.NodeSource.Start.Line, doc.Paths.NodeSource.Start.Column)
	t.Logf("  /pets:    line %d, col %d", petsPath.NodeSource.Start.Line, petsPath.NodeSource.Start.Column)
	t.Logf("  GET:      line %d, col %d", petsPath.Get.NodeSource.Start.Line, petsPath.Get.NodeSource.Start.Column)
}
