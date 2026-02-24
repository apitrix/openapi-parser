package openapi31x

import "github.com/apitrix/openapi-parser/parsers/shared"

// ============================================================================
// Known field lists for OpenAPI 3.1 objects.
// These define the valid fields for each object type, used for unknown field detection.
// ============================================================================

// OpenAPI Object
var openAPIKnownFields = []string{
	"openapi", "$self", "info", "jsonSchemaDialect", "servers",
	"paths", "webhooks", "components", "security", "tags", "externalDocs",
}

var openAPIKnownFieldsSet = shared.ToSet(openAPIKnownFields)

// Info Object
var infoKnownFields = []string{
	"title", "summary", "description", "termsOfService",
	"contact", "license", "version",
}

var infoKnownFieldsSet = shared.ToSet(infoKnownFields)

// Contact Object
var contactKnownFields = []string{
	"name", "url", "email",
}

var contactKnownFieldsSet = shared.ToSet(contactKnownFields)

// License Object
var licenseKnownFields = []string{
	"name", "identifier", "url",
}

var licenseKnownFieldsSet = shared.ToSet(licenseKnownFields)

// Server Object
var serverKnownFields = []string{
	"url", "description", "name", "variables",
}

var serverKnownFieldsSet = shared.ToSet(serverKnownFields)

// Server Variable Object
var serverVariableKnownFields = []string{
	"enum", "default", "description",
}

var serverVariableKnownFieldsSet = shared.ToSet(serverVariableKnownFields)

// Components Object
var componentsKnownFields = []string{
	"schemas", "responses", "parameters", "examples",
	"requestBodies", "headers", "securitySchemes", "links",
	"callbacks", "pathItems", "mediaTypes",
}

var componentsKnownFieldsSet = shared.ToSet(componentsKnownFields)

// Path Item Object
var pathItemKnownFields = []string{
	"$ref", "summary", "description",
	"get", "put", "post", "delete", "options", "head", "patch", "trace",
	"servers", "parameters", "additionalOperations", "query",
}

var pathItemKnownFieldsSet = shared.ToSet(pathItemKnownFields)

// Operation Object
var operationKnownFields = []string{
	"tags", "summary", "description", "externalDocs", "operationId",
	"parameters", "requestBody", "responses", "callbacks",
	"deprecated", "security", "servers",
}

var operationKnownFieldsSet = shared.ToSet(operationKnownFields)

// External Documentation Object
var externalDocsKnownFields = []string{
	"description", "url",
}

var externalDocsKnownFieldsSet = shared.ToSet(externalDocsKnownFields)

// Parameter Object
var parameterKnownFields = []string{
	"name", "in", "description", "required", "deprecated", "allowEmptyValue",
	"style", "explode", "allowReserved", "schema", "example", "examples",
	"content",
}

var parameterKnownFieldsSet = shared.ToSet(parameterKnownFields)

// Header Object
var headerKnownFields = []string{
	"description", "required", "deprecated", "allowEmptyValue",
	"style", "explode", "allowReserved", "schema", "example", "examples",
	"content",
}

var headerKnownFieldsSet = shared.ToSet(headerKnownFields)

// Request Body Object
var requestBodyKnownFields = []string{
	"description", "content", "required",
}

var requestBodyKnownFieldsSet = shared.ToSet(requestBodyKnownFields)

// Media Type Object
var mediaTypeKnownFields = []string{
	"description", "schema", "example", "examples", "encoding",
	"itemSchema", "prefixEncoding", "itemEncoding",
}

var mediaTypeKnownFieldsSet = shared.ToSet(mediaTypeKnownFields)

// Encoding Object
var encodingKnownFields = []string{
	"contentType", "headers", "style", "explode", "allowReserved",
	"encoding", "prefixEncoding", "itemEncoding",
}

var encodingKnownFieldsSet = shared.ToSet(encodingKnownFields)

// Responses Object
var responsesKnownFields = []string{
	"default",
}

var responsesKnownFieldsSet = shared.ToSet(responsesKnownFields)

// Response Object
var responseKnownFields = []string{
	"summary", "description", "headers", "content", "links",
}

var responseKnownFieldsSet = shared.ToSet(responseKnownFields)

// Callback Object
// Note: callbacks are maps of runtime expressions to PathItem objects,
// so there are no fixed known fields - all keys are valid.
var callbackKnownFields = []string{}

var callbackKnownFieldsSet = shared.ToSet(callbackKnownFields)

// Example Object
var exampleKnownFields = []string{
	"summary", "description", "value", "externalValue",
	"dataValue", "serializedValue",
}

var exampleKnownFieldsSet = shared.ToSet(exampleKnownFields)

// Link Object
var linkKnownFields = []string{
	"operationRef", "operationId", "parameters", "requestBody",
	"description", "server",
}

var linkKnownFieldsSet = shared.ToSet(linkKnownFields)

// Tag Object
var tagKnownFields = []string{
	"name", "summary", "description", "externalDocs", "parent", "kind",
}

var tagKnownFieldsSet = shared.ToSet(tagKnownFields)

// Security Scheme Object
var securitySchemeKnownFields = []string{
	"type", "description", "name", "in", "scheme", "bearerFormat",
	"flows", "openIdConnectUrl", "oauth2MetadataUrl", "deprecated",
}

var securitySchemeKnownFieldsSet = shared.ToSet(securitySchemeKnownFields)

// OAuth Flows Object
var oauthFlowsKnownFields = []string{
	"implicit", "password", "clientCredentials", "authorizationCode",
	"deviceAuthorization",
}

var oauthFlowsKnownFieldsSet = shared.ToSet(oauthFlowsKnownFields)

// OAuth Flow Object
var oauthFlowKnownFields = []string{
	"authorizationUrl", "tokenUrl", "refreshUrl", "scopes",
	"deviceAuthorizationUrl",
}

var oauthFlowKnownFieldsSet = shared.ToSet(oauthFlowKnownFields)

// Discriminator Object
var discriminatorKnownFields = []string{
	"propertyName", "mapping", "defaultMapping",
}

var discriminatorKnownFieldsSet = shared.ToSet(discriminatorKnownFields)

// XML Object
var xmlKnownFields = []string{
	"name", "namespace", "prefix", "attribute", "wrapped",
}

var xmlKnownFieldsSet = shared.ToSet(xmlKnownFields)

// Schema Object (OpenAPI 3.1 / JSON Schema 2020-12)
var schemaKnownFields = []string{
	// JSON Schema core
	"title", "multipleOf", "maximum", "exclusiveMaximum",
	"minimum", "exclusiveMinimum", "maxLength", "minLength",
	"pattern", "maxItems", "minItems", "uniqueItems",
	"maxProperties", "minProperties", "required", "enum",
	"type", "allOf", "oneOf", "anyOf", "not",
	"items", "properties", "additionalProperties",
	"description", "format", "default",

	// OpenAPI extensions in schema
	"discriminator", "readOnly", "writeOnly", "xml",
	"externalDocs", "example", "deprecated",

	// JSON Schema 2020-12 new keywords
	"const",
	"if", "then", "else",
	"dependentSchemas",
	"prefixItems",
	"$anchor", "$dynamicRef", "$dynamicAnchor",
	"contentEncoding", "contentMediaType", "contentSchema",
	"unevaluatedItems", "unevaluatedProperties",
	"examples",
}

var schemaKnownFieldsSet = shared.ToSet(schemaKnownFields)
