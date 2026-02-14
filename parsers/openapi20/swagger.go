package openapi20

import (
	"openapi-parser/parsers/shared"
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseSwagger parses a full Swagger 2.0 document from a yaml.Node.
// Swagger 2.0 spec: https://swagger.io/specification/v2/#swagger-object
func parseSwagger(node *yaml.Node, ctx *ParseContext) (*openapi20models.Swagger, error) {
	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "Swagger document must be an object")
	}

	swagger := &openapi20models.Swagger{}
	var err error

	// Simple property - version (inline)
	version, err := parseSwaggerVersion(node, ctx)
	if err != nil {
		return nil, err // version is fatal — can't proceed without it
	}
	swagger.SetProperty("swagger", version)

	// Simple properties - inline
	swagger.SetProperty("host", nodeGetString(node, "host"))
	basePath := nodeGetString(node, "basePath")
	if basePath == "" && shared.ApplySpecDefaults(ctx.config) {
		basePath = shared.DefaultBasePath
	}
	swagger.SetProperty("basePath", basePath)
	if schemes := nodeGetStringSlice(node, "schemes"); schemes != nil {
		swagger.SetProperty("schemes", schemes)
	}
	if consumes := nodeGetStringSlice(node, "consumes"); consumes != nil {
		swagger.SetProperty("consumes", consumes)
	}
	if produces := nodeGetStringSlice(node, "produces"); produces != nil {
		swagger.SetProperty("produces", produces)
	}

	// Complex property - Info (delegated)
	infoNode := nodeGetValue(node, "info")
	info, err := parseSwaggerInfo(infoNode, ctx.push("info"))
	if err != nil {
		swagger.Trix.Errors = append(swagger.Trix.Errors, toParseError(err))
	}
	if info != nil {
		swagger.SetProperty("info", info)
	}

	// Complex property - Paths (delegated)
	pathsNode := nodeGetValue(node, "paths")
	paths, err := parseSwaggerPaths(pathsNode, ctx.push("paths"))
	if err != nil {
		swagger.Trix.Errors = append(swagger.Trix.Errors, toParseError(err))
	}
	if paths != nil {
		swagger.SetProperty("paths", paths)
	}

	// Complex properties - definitions, parameters, responses, securityDefinitions
	if defsNode := nodeGetValue(node, "definitions"); defsNode != nil {
		defs, err := parseSwaggerDefinitions(defsNode, ctx.push("definitions"))
		if err != nil {
			swagger.Trix.Errors = append(swagger.Trix.Errors, toParseError(err))
		}
		if defs != nil {
			swagger.SetProperty("definitions", defs)
		}
	}

	if paramsNode := nodeGetValue(node, "parameters"); paramsNode != nil {
		params, err := parseSwaggerParameters(paramsNode, ctx.push("parameters"))
		if err != nil {
			swagger.Trix.Errors = append(swagger.Trix.Errors, toParseError(err))
		}
		if params != nil {
			swagger.SetProperty("parameters", params)
		}
	}

	if respNode := nodeGetValue(node, "responses"); respNode != nil {
		resps, err := parseSwaggerResponses(respNode, ctx.push("responses"))
		if err != nil {
			swagger.Trix.Errors = append(swagger.Trix.Errors, toParseError(err))
		}
		if resps != nil {
			swagger.SetProperty("responses", resps)
		}
	}

	if secDefsNode := nodeGetValue(node, "securityDefinitions"); secDefsNode != nil {
		secDefs, err := parseSwaggerSecurityDefinitions(secDefsNode, ctx.push("securityDefinitions"))
		if err != nil {
			swagger.Trix.Errors = append(swagger.Trix.Errors, toParseError(err))
		}
		if secDefs != nil {
			swagger.SetProperty("securityDefinitions", secDefs)
		}
	}

	// Complex property - Security
	if secNode := nodeGetValue(node, "security"); secNode != nil {
		sec, err := parseSecurityRequirements(secNode, ctx.push("security"))
		if err != nil {
			swagger.Trix.Errors = append(swagger.Trix.Errors, toParseError(err))
		}
		if sec != nil {
			swagger.SetProperty("security", sec)
		}
	}

	// Complex property - Tags
	if tagsNode := nodeGetValue(node, "tags"); tagsNode != nil {
		tags, err := parseTags(tagsNode, ctx.push("tags"))
		if err != nil {
			swagger.Trix.Errors = append(swagger.Trix.Errors, toParseError(err))
		}
		if tags != nil {
			swagger.SetProperty("tags", tags)
		}
	}

	// Complex property - ExternalDocs
	if edNode := nodeGetValue(node, "externalDocs"); edNode != nil {
		ed, err := parseExternalDocs(edNode, ctx.push("externalDocs"))
		if err != nil {
			swagger.Trix.Errors = append(swagger.Trix.Errors, toParseError(err))
		}
		if ed != nil {
			swagger.SetProperty("externalDocs", ed)
		}
	}

	swagger.VendorExtensions = parseNodeExtensions(node)
	swagger.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields at root level
	swagger.Trix.Errors = append(swagger.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, swaggerKnownFieldsSet))...)

	return swagger, nil
}

// parseSwaggerVersion parses and validates the Swagger version.
func parseSwaggerVersion(node *yaml.Node, ctx *ParseContext) (string, error) {
	version := nodeGetString(node, "swagger")
	if version == "" {
		return "", ctx.errorAt(nodeGetKeyNode(node, "swagger"), "swagger field is required")
	}

	// Validate version is "2.0"
	if version != "2.0" {
		return "", ctx.errorAt(nodeGetValue(node, "swagger"), "unsupported Swagger version: %s (expected 2.0)", version)
	}

	return version, nil
}

// parseSwaggerPaths parses the Swagger.Paths field.
func parseSwaggerPaths(node *yaml.Node, ctx *ParseContext) (*openapi20models.Paths, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "paths must be an object")
	}

	items := make(map[string]*openapi20models.PathItem)

	for key, pathItemNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		pathItem, err := parsePathItem(pathItemNode, ctx.push(key))
		if err != nil {
			return nil, err
		}
		items[key] = pathItem
	}

	paths := openapi20models.NewPaths(items)

	paths.VendorExtensions = parseNodeExtensions(node)
	paths.Trix.Source = ctx.nodeSource(node)

	return paths, nil
}

// parseSwaggerDefinitions parses the Swagger.Definitions field.
func parseSwaggerDefinitions(node *yaml.Node, ctx *ParseContext) (map[string]*openapi20models.RefSchema, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "definitions must be an object")
	}

	definitions := make(map[string]*openapi20models.RefSchema)

	for key, schemaNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		schemaRef, err := parseSchemaRef(schemaNode, ctx.push(key))
		if err != nil {
			return nil, err
		}
		definitions[key] = schemaRef
	}

	return definitions, nil
}

// parseSwaggerParameters parses the Swagger.Parameters field.
func parseSwaggerParameters(node *yaml.Node, ctx *ParseContext) (map[string]*openapi20models.RefParameter, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "parameters must be an object")
	}

	parameters := make(map[string]*openapi20models.RefParameter)

	for key, paramNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		paramRef, err := parseParameterRef(paramNode, ctx.push(key))
		if err != nil {
			return nil, err
		}
		parameters[key] = paramRef
	}

	return parameters, nil
}

// parseSwaggerResponses parses the Swagger.Responses field.
func parseSwaggerResponses(node *yaml.Node, ctx *ParseContext) (map[string]*openapi20models.RefResponse, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "responses must be an object")
	}

	responses := make(map[string]*openapi20models.RefResponse)

	for key, respNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		respRef, err := parseResponseRef(respNode, ctx.push(key))
		if err != nil {
			return nil, err
		}
		responses[key] = respRef
	}

	return responses, nil
}

// parseSwaggerSecurityDefinitions parses the Swagger.SecurityDefinitions field.
func parseSwaggerSecurityDefinitions(node *yaml.Node, ctx *ParseContext) (map[string]*openapi20models.SecurityScheme, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "securityDefinitions must be an object")
	}

	secDefs := make(map[string]*openapi20models.SecurityScheme)

	for key, schemeNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		scheme, err := parseSecurityScheme(schemeNode, ctx.push(key))
		if err != nil {
			return nil, err
		}
		secDefs[key] = scheme
	}

	return secDefs, nil
}
