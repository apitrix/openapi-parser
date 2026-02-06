package openapi30

// Known fields for each OpenAPI 3.0.x object type.
// These are used for detecting unknown/unrecognized fields during parsing.
// Extensions (x-*) are always allowed and handled separately.
//
// Reference: https://spec.openapis.org/oas/v3.0.3

// OpenAPI Object - root level
var openAPIKnownFields = []string{
	"openapi",
	"info",
	"servers",
	"paths",
	"components",
	"security",
	"tags",
	"externalDocs",
}

// Info Object
var infoKnownFields = []string{
	"title",
	"description",
	"termsOfService",
	"contact",
	"license",
	"version",
}

// Contact Object
var contactKnownFields = []string{
	"name",
	"url",
	"email",
}

// License Object
var licenseKnownFields = []string{
	"name",
	"url",
}

// Server Object
var serverKnownFields = []string{
	"url",
	"description",
	"variables",
}

// ServerVariable Object
var serverVariableKnownFields = []string{
	"enum",
	"default",
	"description",
}

// Components Object
var componentsKnownFields = []string{
	"schemas",
	"responses",
	"parameters",
	"examples",
	"requestBodies",
	"headers",
	"securitySchemes",
	"links",
	"callbacks",
}

// PathItem Object
var pathItemKnownFields = []string{
	"$ref",
	"summary",
	"description",
	"get",
	"put",
	"post",
	"delete",
	"options",
	"head",
	"patch",
	"trace",
	"servers",
	"parameters",
}

// Operation Object
var operationKnownFields = []string{
	"tags",
	"summary",
	"description",
	"externalDocs",
	"operationId",
	"parameters",
	"requestBody",
	"responses",
	"callbacks",
	"deprecated",
	"security",
	"servers",
}

// ExternalDocs Object
var externalDocsKnownFields = []string{
	"description",
	"url",
}

// Parameter Object
var parameterKnownFields = []string{
	"name",
	"in",
	"description",
	"required",
	"deprecated",
	"allowEmptyValue",
	"style",
	"explode",
	"allowReserved",
	"schema",
	"example",
	"examples",
	"content",
}

// RequestBody Object
var requestBodyKnownFields = []string{
	"description",
	"content",
	"required",
}

// MediaType Object
var mediaTypeKnownFields = []string{
	"schema",
	"example",
	"examples",
	"encoding",
}

// Encoding Object
var encodingKnownFields = []string{
	"contentType",
	"headers",
	"style",
	"explode",
	"allowReserved",
}

// Response Object
var responseKnownFields = []string{
	"description",
	"headers",
	"content",
	"links",
}

// Header Object (same as Parameter, but without name/in)
var headerKnownFields = []string{
	"description",
	"required",
	"deprecated",
	"allowEmptyValue",
	"style",
	"explode",
	"allowReserved",
	"schema",
	"example",
	"examples",
	"content",
}

// Link Object
var linkKnownFields = []string{
	"operationRef",
	"operationId",
	"parameters",
	"requestBody",
	"description",
	"server",
}

// Tag Object
var tagKnownFields = []string{
	"name",
	"description",
	"externalDocs",
}

// Example Object
var exampleKnownFields = []string{
	"summary",
	"description",
	"value",
	"externalValue",
}

// Schema Object (JSON Schema subset for OpenAPI 3.0)
var schemaKnownFields = []string{
	// Basic metadata
	"title",
	"description",
	// Type constraints
	"type",
	"format",
	"nullable",
	// String validation
	"pattern",
	"maxLength",
	"minLength",
	// Numeric validation
	"multipleOf",
	"maximum",
	"exclusiveMaximum",
	"minimum",
	"exclusiveMinimum",
	// Array validation
	"items",
	"maxItems",
	"minItems",
	"uniqueItems",
	// Object validation
	"properties",
	"additionalProperties",
	"required",
	"maxProperties",
	"minProperties",
	// Composition
	"allOf",
	"oneOf",
	"anyOf",
	"not",
	// Enumerations
	"enum",
	"default",
	// Documentation
	"example",
	"deprecated",
	"readOnly",
	"writeOnly",
	// OpenAPI extensions
	"discriminator",
	"xml",
	"externalDocs",
}

// Discriminator Object
var discriminatorKnownFields = []string{
	"propertyName",
	"mapping",
}

// XML Object
var xmlKnownFields = []string{
	"name",
	"namespace",
	"prefix",
	"attribute",
	"wrapped",
}

// SecurityScheme Object
var securitySchemeKnownFields = []string{
	"type",
	"description",
	"name",
	"in",
	"scheme",
	"bearerFormat",
	"flows",
	"openIdConnectUrl",
}

// OAuthFlows Object
var oauthFlowsKnownFields = []string{
	"implicit",
	"password",
	"clientCredentials",
	"authorizationCode",
}

// OAuthFlow Object
var oauthFlowKnownFields = []string{
	"authorizationUrl",
	"tokenUrl",
	"refreshUrl",
	"scopes",
}

// Callback Object has dynamic keys (expressions), no specific fields
// (all non-extension keys are path expressions pointing to PathItem)

// Reference Object fields ($ref handling)
var referenceKnownFields = []string{
	"$ref",
}

// knownFieldsMap provides a lookup for known fields by object type name.
// This enables runtime lookup for dynamic detection scenarios.
var knownFieldsMap = map[string][]string{
	"OpenAPI":        openAPIKnownFields,
	"Info":           infoKnownFields,
	"Contact":        contactKnownFields,
	"License":        licenseKnownFields,
	"Server":         serverKnownFields,
	"ServerVariable": serverVariableKnownFields,
	"Components":     componentsKnownFields,
	"PathItem":       pathItemKnownFields,
	"Operation":      operationKnownFields,
	"ExternalDocs":   externalDocsKnownFields,
	"Parameter":      parameterKnownFields,
	"RequestBody":    requestBodyKnownFields,
	"MediaType":      mediaTypeKnownFields,
	"Encoding":       encodingKnownFields,
	"Response":       responseKnownFields,
	"Header":         headerKnownFields,
	"Link":           linkKnownFields,
	"Tag":            tagKnownFields,
	"Example":        exampleKnownFields,
	"Schema":         schemaKnownFields,
	"Discriminator":  discriminatorKnownFields,
	"XML":            xmlKnownFields,
	"SecurityScheme": securitySchemeKnownFields,
	"OAuthFlows":     oauthFlowsKnownFields,
	"OAuthFlow":      oauthFlowKnownFields,
	"Reference":      referenceKnownFields,
}

// toSet converts a slice of field names to a set for O(1) lookup.
func toSet(fields []string) map[string]struct{} {
	m := make(map[string]struct{}, len(fields))
	for _, f := range fields {
		m[f] = struct{}{}
	}
	return m
}

// Precomputed sets for O(1) lookup during unknown field detection.
// These are built once at init time to avoid repeated slice-to-map conversion.
var (
	openAPIKnownFieldsSet        = toSet(openAPIKnownFields)
	infoKnownFieldsSet           = toSet(infoKnownFields)
	contactKnownFieldsSet        = toSet(contactKnownFields)
	licenseKnownFieldsSet        = toSet(licenseKnownFields)
	serverKnownFieldsSet         = toSet(serverKnownFields)
	serverVariableKnownFieldsSet = toSet(serverVariableKnownFields)
	componentsKnownFieldsSet     = toSet(componentsKnownFields)
	pathItemKnownFieldsSet       = toSet(pathItemKnownFields)
	operationKnownFieldsSet      = toSet(operationKnownFields)
	externalDocsKnownFieldsSet   = toSet(externalDocsKnownFields)
	parameterKnownFieldsSet      = toSet(parameterKnownFields)
	requestBodyKnownFieldsSet    = toSet(requestBodyKnownFields)
	mediaTypeKnownFieldsSet      = toSet(mediaTypeKnownFields)
	encodingKnownFieldsSet       = toSet(encodingKnownFields)
	responseKnownFieldsSet       = toSet(responseKnownFields)
	headerKnownFieldsSet         = toSet(headerKnownFields)
	linkKnownFieldsSet           = toSet(linkKnownFields)
	tagKnownFieldsSet            = toSet(tagKnownFields)
	exampleKnownFieldsSet        = toSet(exampleKnownFields)
	schemaKnownFieldsSet         = toSet(schemaKnownFields)
	discriminatorKnownFieldsSet  = toSet(discriminatorKnownFields)
	xmlKnownFieldsSet            = toSet(xmlKnownFields)
	securitySchemeKnownFieldsSet = toSet(securitySchemeKnownFields)
	oauthFlowsKnownFieldsSet     = toSet(oauthFlowsKnownFields)
	oauthFlowKnownFieldsSet      = toSet(oauthFlowKnownFields)
	referenceKnownFieldsSet      = toSet(referenceKnownFields)
)
