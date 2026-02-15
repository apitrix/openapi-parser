package openapi30x

import (
	"reflect"

	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"
	"github.com/apitrix/openapi-parser/parsers/shared"
)

// flattenErrors recursively walks the parsed document tree and collects
// all Trix.Errors from every node into a flat []*shared.ParseError slice.
func flattenErrors(doc *openapi30models.OpenAPI) []*shared.ParseError {
	if doc == nil {
		return nil
	}
	var result []*shared.ParseError
	collectTrixErrors(reflect.ValueOf(doc), make(map[uintptr]bool), &result)
	return result
}

// collectTrixErrors recursively inspects a value for Trix.Errors fields.
// It handles both exported (public) and unexported (private/getter-only) fields.
// For unexported fields, it attempts to call a corresponding getter method.
func collectTrixErrors(v reflect.Value, visited map[uintptr]bool, result *[]*shared.ParseError) {
	// Dereference pointers and interfaces
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return
		}
		if v.Kind() == reflect.Ptr {
			ptr := v.Pointer()
			if visited[ptr] {
				return // avoid infinite recursion on circular refs
			}
			visited[ptr] = true
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()

		// First, check for Trix field and collect its errors (Trix is public/embedded)
		trixField := v.FieldByName("Trix")
		if trixField.IsValid() && trixField.Kind() == reflect.Struct {
			errorsField := trixField.FieldByName("Errors")
			if errorsField.IsValid() && errorsField.Kind() == reflect.Slice {
				for j := 0; j < errorsField.Len(); j++ {
					modelErr := errorsField.Index(j).Interface().(openapi30models.ParseError)
					*result = append(*result, modelParseErrorToShared(modelErr))
				}
			}
		}

		// Recurse into all struct fields, using getter methods for unexported fields
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := t.Field(i)

			// Skip the Trix and ElementBase fields (already handled above or not needed)
			if fieldType.Name == "Trix" || fieldType.Name == "ElementBase" {
				continue
			}

			// Skip VendorExtensions (not a model type)
			if fieldType.Name == "VendorExtensions" {
				continue
			}

			if field.CanInterface() {
				// Exported field - use directly
				collectTrixErrors(field, visited, result)
			} else {
				// Unexported field - try getter method on the pointer receiver
				// Construct the method name: capitalize first letter of field name
				// Go convention: unexported field "foo" has getter "Foo()"
				methodName := exportedName(fieldType.Name)
				if methodName == "" {
					continue
				}

				// We need the pointer to the struct to call pointer-receiver methods
				var ptrVal reflect.Value
				if v.CanAddr() {
					ptrVal = v.Addr()
				} else {
					// If we can't get a pointer, create a new one
					newPtr := reflect.New(v.Type())
					newPtr.Elem().Set(v)
					ptrVal = newPtr
				}

				method := ptrVal.MethodByName(methodName)
				if !method.IsValid() {
					continue
				}

				// Call the getter (no arguments)
				results := method.Call(nil)
				if len(results) > 0 {
					val := results[0]
					if val.IsValid() {
						collectTrixErrors(val, visited, result)
					}
				}
			}
		}

	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			collectTrixErrors(v.Index(i), visited, result)
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			collectTrixErrors(v.MapIndex(key), visited, result)
		}
	}
}

// exportedName converts an unexported field name to its expected getter method name.
// e.g., "info" → "Info", "openAPI" → "OpenAPIVersion" (special case)
func exportedName(fieldName string) string {
	// Special cases where field name doesn't match getter name
	switch fieldName {
	case "openAPI":
		return "OpenAPIVersion"
	case "default":
		return "Default"
	case "delete":
		return "Delete"
	}

	if len(fieldName) == 0 {
		return ""
	}

	// Standard capitalization: first letter uppercase
	runes := []rune(fieldName)
	if runes[0] >= 'a' && runes[0] <= 'z' {
		runes[0] = runes[0] - 'a' + 'A'
	}
	return string(runes)
}

// modelParseErrorToShared converts a model-level ParseError to a shared.ParseError.
func modelParseErrorToShared(e openapi30models.ParseError) *shared.ParseError {
	return &shared.ParseError{
		Path:    e.Path,
		Message: e.Message,
		Kind:    e.Kind,
	}
}
