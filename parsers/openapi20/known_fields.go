package openapi20

import "openapi-parser/parsers/shared"

// Known fields for each Swagger 2.0 object type.
// These are used for detecting unknown/unrecognized fields during parsing.
// Extensions (x-*) are always allowed and handled separately.
//
// Reference: https://swagger.io/specification/v2/

// Swagger Object - root level
var swaggerKnownFields = []string{
	"swagger",
	"info",
	"host",
	"basePath",
	"schemes",
	"consumes",
	"produces",
	"paths",
	"definitions",
	"parameters",
	"responses",
	"securityDefinitions",
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

// PathItem Object
var pathItemKnownFields = []string{
	"$ref",
	"get",
	"put",
	"post",
	"delete",
	"options",
	"head",
	"patch",
	"parameters",
}

// Operation Object
var operationKnownFields = []string{
	"tags",
	"summary",
	"description",
	"externalDocs",
	"operationId",
	"consumes",
	"produces",
	"parameters",
	"responses",
	"schemes",
	"deprecated",
	"security",
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
	"allowEmptyValue",
	// Body parameter
	"schema",
	// Non-body parameters
	"type",
	"format",
	"items",
	"collectionFormat",
	"default",
	"maximum",
	"exclusiveMaximum",
	"minimum",
	"exclusiveMinimum",
	"maxLength",
	"minLength",
	"pattern",
	"maxItems",
	"minItems",
	"uniqueItems",
	"enum",
	"multipleOf",
}

// Items Object
var itemsKnownFields = []string{
	"type",
	"format",
	"items",
	"collectionFormat",
	"default",
	"maximum",
	"exclusiveMaximum",
	"minimum",
	"exclusiveMinimum",
	"maxLength",
	"minLength",
	"pattern",
	"maxItems",
	"minItems",
	"uniqueItems",
	"enum",
	"multipleOf",
}

// Responses Object (container)
// Note: default and HTTP status codes are dynamic keys

// Response Object
var responseKnownFields = []string{
	"description",
	"schema",
	"headers",
	"examples",
}

// Header Object
var headerKnownFields = []string{
	"description",
	"type",
	"format",
	"items",
	"collectionFormat",
	"default",
	"maximum",
	"exclusiveMaximum",
	"minimum",
	"exclusiveMinimum",
	"maxLength",
	"minLength",
	"pattern",
	"maxItems",
	"minItems",
	"uniqueItems",
	"enum",
	"multipleOf",
}

// Tag Object
var tagKnownFields = []string{
	"name",
	"description",
	"externalDocs",
}

// Schema Object (JSON Schema subset for Swagger 2.0)
var schemaKnownFields = []string{
	// Basic metadata
	"title",
	"description",
	// Type constraints
	"type",
	"format",
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
	// Enumerations
	"enum",
	"default",
	// Swagger 2.0 specific
	"discriminator",
	"readOnly",
	"xml",
	"externalDocs",
	"example",
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
	"flow",
	"authorizationUrl",
	"tokenUrl",
	"scopes",
}

// Reference Object fields ($ref handling)
var referenceKnownFields = []string{
	"$ref",
}

// knownFieldsMap provides a lookup for known fields by object type name.
var knownFieldsMap = map[string][]string{
	"Swagger":        swaggerKnownFields,
	"Info":           infoKnownFields,
	"Contact":        contactKnownFields,
	"License":        licenseKnownFields,
	"PathItem":       pathItemKnownFields,
	"Operation":      operationKnownFields,
	"ExternalDocs":   externalDocsKnownFields,
	"Parameter":      parameterKnownFields,
	"Items":          itemsKnownFields,
	"Response":       responseKnownFields,
	"Header":         headerKnownFields,
	"Tag":            tagKnownFields,
	"Schema":         schemaKnownFields,
	"XML":            xmlKnownFields,
	"SecurityScheme": securitySchemeKnownFields,
	"Reference":      referenceKnownFields,
}

// toSet converts a slice of field names to a set for O(1) lookup.
func toSet(fields []string) map[string]struct{} {
	return shared.ToSet(fields)
}

// Precomputed sets for O(1) lookup during unknown field detection.
// These are built once at init time to avoid repeated slice-to-map conversion.
var (
	swaggerKnownFieldsSet        = toSet(swaggerKnownFields)
	infoKnownFieldsSet           = toSet(infoKnownFields)
	contactKnownFieldsSet        = toSet(contactKnownFields)
	licenseKnownFieldsSet        = toSet(licenseKnownFields)
	pathItemKnownFieldsSet       = toSet(pathItemKnownFields)
	operationKnownFieldsSet      = toSet(operationKnownFields)
	externalDocsKnownFieldsSet   = toSet(externalDocsKnownFields)
	parameterKnownFieldsSet      = toSet(parameterKnownFields)
	itemsKnownFieldsSet          = toSet(itemsKnownFields)
	responseKnownFieldsSet       = toSet(responseKnownFields)
	headerKnownFieldsSet         = toSet(headerKnownFields)
	tagKnownFieldsSet            = toSet(tagKnownFields)
	schemaKnownFieldsSet         = toSet(schemaKnownFields)
	xmlKnownFieldsSet            = toSet(xmlKnownFields)
	securitySchemeKnownFieldsSet = toSet(securitySchemeKnownFields)
	referenceKnownFieldsSet      = toSet(referenceKnownFields)
)
