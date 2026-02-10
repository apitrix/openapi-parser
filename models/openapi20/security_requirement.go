package openapi20

// SecurityRequirement lists required security schemes to execute an operation.
// https://swagger.io/specification/v2/#security-requirement-object
type SecurityRequirement map[string][]string
