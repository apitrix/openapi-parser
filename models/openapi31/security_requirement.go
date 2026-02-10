package openapi31

// SecurityRequirement lists required security schemes to execute an operation.
// https://spec.openapis.org/oas/v3.1.0#security-requirement-object
type SecurityRequirement map[string][]string
