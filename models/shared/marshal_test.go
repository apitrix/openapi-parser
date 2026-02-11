package shared

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

// --- MarshalFieldsJSON ---

func TestMarshalFieldsJSON_BasicFields(t *testing.T) {
	fields := []Field{
		{Key: "name", Value: "Alice"},
		{Key: "age", Value: 30},
	}
	got, err := MarshalFieldsJSON(fields)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"Alice","age":30}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMarshalFieldsJSON_PreservesOrder(t *testing.T) {
	fields := []Field{
		{Key: "z", Value: 1},
		{Key: "a", Value: 2},
		{Key: "m", Value: 3},
	}
	got, err := MarshalFieldsJSON(fields)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"z":1,"a":2,"m":3}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMarshalFieldsJSON_OmitsZeroValues(t *testing.T) {
	fields := []Field{
		{Key: "title", Value: "hello"},
		{Key: "empty_string", Value: ""},
		{Key: "nil_value", Value: nil},
		{Key: "false_bool", Value: false},
		{Key: "present", Value: "yes"},
	}
	got, err := MarshalFieldsJSON(fields)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"title":"hello","present":"yes"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMarshalFieldsJSON_OmitsNilPointer(t *testing.T) {
	var p *string
	fields := []Field{
		{Key: "name", Value: "ok"},
		{Key: "ptr", Value: p},
	}
	got, err := MarshalFieldsJSON(fields)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"ok"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMarshalFieldsJSON_IncludesNonNilPointer(t *testing.T) {
	v := "hello"
	fields := []Field{
		{Key: "ptr", Value: &v},
	}
	got, err := MarshalFieldsJSON(fields)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"ptr":"hello"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMarshalFieldsJSON_OmitsNilAndEmptySlice(t *testing.T) {
	var nilSlice []string
	emptySlice := []string{}
	fields := []Field{
		{Key: "nil_slice", Value: nilSlice},
		{Key: "empty_slice", Value: emptySlice},
		{Key: "present", Value: []string{"a"}},
	}
	got, err := MarshalFieldsJSON(fields)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"present":["a"]}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMarshalFieldsJSON_OmitsNilAndEmptyMap(t *testing.T) {
	var nilMap map[string]string
	emptyMap := map[string]string{}
	fields := []Field{
		{Key: "nil_map", Value: nilMap},
		{Key: "empty_map", Value: emptyMap},
		{Key: "present", Value: map[string]string{"k": "v"}},
	}
	got, err := MarshalFieldsJSON(fields)
	if err != nil {
		t.Fatal(err)
	}
	// Unmarshal to compare since map key order is not guaranteed
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if len(result) != 1 {
		t.Errorf("expected 1 key, got %d: %s", len(result), got)
	}
	if _, ok := result["present"]; !ok {
		t.Errorf("expected 'present' key in output: %s", got)
	}
}

func TestMarshalFieldsJSON_EmptyFields(t *testing.T) {
	got, err := MarshalFieldsJSON(nil)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMarshalFieldsJSON_AllOmitted(t *testing.T) {
	fields := []Field{
		{Key: "a", Value: ""},
		{Key: "b", Value: nil},
		{Key: "c", Value: false},
	}
	got, err := MarshalFieldsJSON(fields)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMarshalFieldsJSON_NestedObject(t *testing.T) {
	inner := map[string]interface{}{"nested": true}
	fields := []Field{
		{Key: "outer", Value: inner},
	}
	got, err := MarshalFieldsJSON(fields)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"outer":{"nested":true}}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMarshalFieldsJSON_TrueBoolIncluded(t *testing.T) {
	fields := []Field{
		{Key: "enabled", Value: true},
	}
	got, err := MarshalFieldsJSON(fields)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"enabled":true}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMarshalFieldsJSON_IntZeroNotOmitted(t *testing.T) {
	fields := []Field{
		{Key: "count", Value: 0},
	}
	got, err := MarshalFieldsJSON(fields)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"count":0}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// --- MarshalFieldsYAML ---

func TestMarshalFieldsYAML_BasicFields(t *testing.T) {
	fields := []Field{
		{Key: "name", Value: "Alice"},
		{Key: "age", Value: 30},
	}
	node, err := MarshalFieldsYAML(fields)
	if err != nil {
		t.Fatal(err)
	}
	if node.Kind != yaml.MappingNode {
		t.Fatalf("expected MappingNode, got %d", node.Kind)
	}
	// Content is key/value pairs: [key, val, key, val, ...]
	if len(node.Content) != 4 {
		t.Fatalf("expected 4 content nodes (2 pairs), got %d", len(node.Content))
	}
	if node.Content[0].Value != "name" {
		t.Errorf("expected first key 'name', got %q", node.Content[0].Value)
	}
	if node.Content[2].Value != "age" {
		t.Errorf("expected second key 'age', got %q", node.Content[2].Value)
	}
}

func TestMarshalFieldsYAML_OmitsZeroValues(t *testing.T) {
	fields := []Field{
		{Key: "title", Value: "hello"},
		{Key: "empty", Value: ""},
		{Key: "nil_val", Value: nil},
	}
	node, err := MarshalFieldsYAML(fields)
	if err != nil {
		t.Fatal(err)
	}
	if len(node.Content) != 2 { // 1 pair
		t.Fatalf("expected 2 content nodes (1 pair), got %d", len(node.Content))
	}
	if node.Content[0].Value != "title" {
		t.Errorf("expected key 'title', got %q", node.Content[0].Value)
	}
}

func TestMarshalFieldsYAML_EmptyFields(t *testing.T) {
	node, err := MarshalFieldsYAML(nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(node.Content) != 0 {
		t.Errorf("expected 0 content nodes, got %d", len(node.Content))
	}
}

func TestMarshalFieldsYAML_RoundTrip(t *testing.T) {
	fields := []Field{
		{Key: "openapi", Value: "3.0.3"},
		{Key: "title", Value: "Pet Store"},
	}
	node, err := MarshalFieldsYAML(fields)
	if err != nil {
		t.Fatal(err)
	}
	// Marshal to YAML bytes via a document node
	doc := &yaml.Node{Kind: yaml.DocumentNode, Content: []*yaml.Node{node}}
	out, err := yaml.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	// Parse back
	var result map[string]string
	if err := yaml.Unmarshal(out, &result); err != nil {
		t.Fatal(err)
	}
	if result["openapi"] != "3.0.3" {
		t.Errorf("openapi: got %q, want %q", result["openapi"], "3.0.3")
	}
	if result["title"] != "Pet Store" {
		t.Errorf("title: got %q, want %q", result["title"], "Pet Store")
	}
}

// --- AppendExtensions ---

func TestAppendExtensions_NilMap(t *testing.T) {
	fields := []Field{{Key: "a", Value: 1}}
	got := AppendExtensions(fields, nil)
	if len(got) != 1 {
		t.Errorf("expected 1 field, got %d", len(got))
	}
}

func TestAppendExtensions_EmptyMap(t *testing.T) {
	fields := []Field{{Key: "a", Value: 1}}
	got := AppendExtensions(fields, map[string]interface{}{})
	if len(got) != 1 {
		t.Errorf("expected 1 field, got %d", len(got))
	}
}

func TestAppendExtensions_SortedOrder(t *testing.T) {
	fields := []Field{{Key: "spec", Value: "val"}}
	extensions := map[string]interface{}{
		"x-zebra": "z",
		"x-alpha": "a",
		"x-mid":   "m",
	}
	got := AppendExtensions(fields, extensions)
	if len(got) != 4 {
		t.Fatalf("expected 4 fields, got %d", len(got))
	}
	// Extensions should be sorted alphabetically after spec fields
	if got[1].Key != "x-alpha" {
		t.Errorf("got[1].Key = %q, want %q", got[1].Key, "x-alpha")
	}
	if got[2].Key != "x-mid" {
		t.Errorf("got[2].Key = %q, want %q", got[2].Key, "x-mid")
	}
	if got[3].Key != "x-zebra" {
		t.Errorf("got[3].Key = %q, want %q", got[3].Key, "x-zebra")
	}
}

// --- isOmittable ---

func TestIsOmittable(t *testing.T) {
	var nilPtr *string
	nonNilStr := "hello"

	tests := []struct {
		name string
		val  interface{}
		want bool
	}{
		{"nil", nil, true},
		{"empty string", "", true},
		{"non-empty string", "hi", false},
		{"false bool", false, true},
		{"true bool", true, false},
		{"nil pointer", nilPtr, true},
		{"non-nil pointer", &nonNilStr, false},
		{"nil slice", ([]string)(nil), true},
		{"empty slice", []string{}, true},
		{"populated slice", []string{"a"}, false},
		{"nil map", (map[string]string)(nil), true},
		{"empty map", map[string]string{}, true},
		{"populated map", map[string]string{"k": "v"}, false},
		{"zero int", 0, false},
		{"non-zero int", 42, false},
		{"float64", 3.14, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isOmittable(tt.val)
			if got != tt.want {
				t.Errorf("isOmittable(%v) = %v, want %v", tt.val, got, tt.want)
			}
		})
	}
}
