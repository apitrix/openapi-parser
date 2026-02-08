// Package openapi30x implements a parser for OpenAPI 3.0.x specifications.
//
// # Architecture
//
// The parser uses an imperative approach with yaml.Node for lossless parsing:
//   - Simple properties (string, bool, int) are inline in main type files
//   - Complex properties (nested objects, arrays, maps) have dedicated files
//   - Source location (line/column) is preserved for all nodes
//
// This design enables:
//   - Clear separation of simple vs complex property handling
//   - Easy navigation to complex property logic via dedicated files
//   - Detection of unknown/unrecognized fields
//
// # File Naming Conventions
//
// Files follow these naming patterns:
//
//   - {type}.go - Main type file with orchestrator + simple properties inline
//   - {type}_{property}.go - Dedicated file for complex property
//   - openapi_{path}.go - Document hierarchy files
//   - ref_{type}.go - Reference handling files for $ref resolution
//
// # Example: Schema Type
//
//	schema.go                      - Main file: orchestrator + 25 simple properties inline
//	schema_items.go                - Complex: recursive SchemaRef
//	schema_allof.go                - Complex: []SchemaRef
//	schema_properties.go           - Complex: map[string]SchemaRef
//	schema_additionalproperties.go - Complex: polymorphic (bool OR SchemaRef)
//	schema_discriminator.go        - Complex: nested Discriminator object
//
// # Usage
//
//	data, _ := os.ReadFile("openapi.yaml")
//	doc, err := openapi30x.Parse(data)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// doc is now a fully parsed *openapi30models.OpenAPI
package openapi30x
