# Project Rules for openapi-parser

## Project Context
This is a Go-based OpenAPI parser library that handles parsing and validation of OpenAPI 2.0 and 3.0 specifications.

## Code Style
- Follow standard Go conventions and idioms
- Use descriptive variable and function names
- Add comments for exported functions and types
- Maintain consistency with existing code patterns in the repository
- Fix linter errors as soon as possible

## Testing
- Write tests for all new functionality
- Follow the existing test patterns in `*_test.go` files.
- All `.go` files should have a corresponding `.test.go` file.
- Use table-driven tests where appropriate
- Ensure schema conformance tests pass after model changes
- Use Arrange Act Assert(AAA) pattern for tests


## Project Structure
- `models/` - Contains OpenAPI specification model definitions
  - `openapi20/` - OpenAPI 2.0 (Swagger) models
  - `openapi30/` - OpenAPI 3.0 models
- `parsers/` - Contains parsing logic for different OpenAPI versions
  - `openapi20/` - OpenAPI 2.0 parser
  - `openapi30/` - OpenAPI 3.0 parser

## Guidelines
- When modifying models, ensure they conform to the official OpenAPI JSON schemas
- Use reflection-based tests to validate struct field mappings
- Handle JSON tags correctly (`json:"fieldName,omitempty"`)
- Extensions (x-*) should use `map[string]any` types
