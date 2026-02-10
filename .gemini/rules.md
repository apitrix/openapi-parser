# Project Rules for openapi-parser

## Project Context
This is a Go-based OpenAPI parser library that handles parsing and validation of OpenAPI 2.0, 3.0, and 3.1 specifications.

## Code Style
- Follow standard Go conventions and idioms
- Use descriptive variable and function names
- Add comments for exported functions and types
- Maintain consistency with existing code patterns in the repository
- Fix linter errors as soon as possible

## Testing
- Write tests for all new functionality
- Run tests after every change
- Follow the existing test patterns in `*_test.go` files.
- Every `.go` source file that contains logic (not only struct definitions) MUST have a corresponding `_test.go` file in the same directory (e.g. `schema.go` → `schema_test.go`). The only exception is `doc.go` files that only contain package documentation.
- Use table-driven tests where appropriate
- Ensure schema conformance tests pass after model changes
- Use Arrange Act Assert(AAA) pattern for tests


## Project Structure
- `models/` - Contains OpenAPI specification model definitions
  - `openapi20/` - OpenAPI 2.0 (Swagger) models
  - `openapi30/` - OpenAPI 3.0 models
  - `openapi31/` - OpenAPI 3.1 models
- `parsers/` - Contains parsing logic for different OpenAPI versions
  - `openapi20/` - OpenAPI 2.0 parser
  - `openapi30x/` - OpenAPI 3.0.x parser
  - `openapi31x/` - OpenAPI 3.1.x / 3.2.x parser

## Guidelines
- When modifying models, ensure they conform to the official OpenAPI JSON schemas
- Use reflection-based tests to validate struct field mappings
- Clean temporary files that you create!
