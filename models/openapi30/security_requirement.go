package openapi30

// SecurityRequirement lists required security schemes to execute an operation.
// https://spec.openapis.org/oas/v3.0.3#security-requirement-object
type SecurityRequirement map[string][]string
