// Package openapi20 provides parsing functionality for OpenAPI 2.0 (Swagger) specifications.
//
// This package parses YAML or JSON documents conforming to the Swagger 2.0 specification
// (https://swagger.io/specification/v2/) into strongly-typed Go models defined in
// the openapi-parser/models/openapi20 package.
//
// # Usage
//
// Basic parsing:
//
//	doc, err := openapi20.Parse(yamlData)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(doc.Info.Title)
//
// Parsing with unknown field detection:
//
//	result, err := openapi20.ParseWithUnknownFields(yamlData)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, field := range result.UnknownFields {
//	    fmt.Printf("Unknown field at %s: %s\n", field.Path, field.Name)
//	}
//
// # Features
//
//   - Parses both JSON and YAML formats
//   - Preserves source location information (line/column)
//   - Detects unknown/unrecognized fields
//   - Handles x-* extension fields
//   - Resolves $ref references within the document
package openapi20
