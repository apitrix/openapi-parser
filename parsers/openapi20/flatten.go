package openapi20

import (
	"reflect"

	openapi20models "openapi-parser/models/openapi20"
	"openapi-parser/parsers/shared"
)

// flattenErrors recursively walks the parsed document tree and collects
// all Trix.Errors from every node into a flat []*shared.ParseError slice.
func flattenErrors(doc *openapi20models.Swagger) []*shared.ParseError {
	if doc == nil {
		return nil
	}
	var result []*shared.ParseError
	collectTrixErrors(reflect.ValueOf(doc), make(map[uintptr]bool), &result)
	return result
}

// collectTrixErrors recursively inspects a value for Trix.Errors fields.
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
		// Recurse into all struct fields
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := t.Field(i)

			// If this field IS the Trix struct, collect its Errors directly
			// and do NOT recurse further into it (to avoid walking into Source, etc.).
			if fieldType.Name == "Trix" && field.Kind() == reflect.Struct {
				errorsField := field.FieldByName("Errors")
				if errorsField.IsValid() && errorsField.Kind() == reflect.Slice {
					for j := 0; j < errorsField.Len(); j++ {
						modelErr := errorsField.Index(j).Interface().(openapi20models.ParseError)
						*result = append(*result, modelParseErrorToShared(modelErr))
					}
				}
				continue
			}

			if field.CanInterface() {
				collectTrixErrors(field, visited, result)
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

// modelParseErrorToShared converts a model-level ParseError to a shared.ParseError.
func modelParseErrorToShared(e openapi20models.ParseError) *shared.ParseError {
	return &shared.ParseError{
		Path:    e.Path,
		Message: e.Message,
		Kind:    e.Kind,
	}
}
