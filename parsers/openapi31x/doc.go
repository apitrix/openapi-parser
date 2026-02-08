// Package openapi31x implements a parser for OpenAPI 3.1.x and 3.2.x specifications.
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
// # Key differences from openapi30x
//
// - Accepts OpenAPI versions 3.1.x and 3.2.x (3.2 is backwards-compatible with 3.1)
// - Schema uses JSON Schema Draft 2020-12 (type arrays, new keywords, no nullable)
// - Reference objects support summary and description alongside $ref
// - Root document has webhooks and jsonSchemaDialect
// - Info has summary, License has identifier
// - Components has pathItems
//
// # Usage
//
//	data, _ := os.ReadFile("openapi.yaml")
//	doc, err := openapi31x.Parse(data)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// doc is now a fully parsed *openapi31models.OpenAPI
package openapi31x
