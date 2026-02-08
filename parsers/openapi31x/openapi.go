package openapi31x

import (
	"strings"

	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseOpenAPI parses a full OpenAPI 3.1.x/3.2.x document from a yaml.Node.
// OpenAPI 3.1.0 spec: https://spec.openapis.org/oas/v3.1.0#openapi-object
func parseOpenAPI(node *yaml.Node, ctx *ParseContext) (*openapi31models.OpenAPI, error) {
	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "OpenAPI document must be an object")
	}

	openapi := &openapi31models.OpenAPI{}
	var err error

	// Simple property - version (inline)
	openapi.OpenAPI, err = parseOpenAPIVersion(node, ctx)
	if err != nil {
		return nil, err // version is fatal — can't proceed without it
	}

	// Simple property - jsonSchemaDialect
	openapi.JsonSchemaDialect = nodeGetString(node, "jsonSchemaDialect")

	// Complex properties - delegated
	infoNode := nodeGetValue(node, "info")
	openapi.Info, err = parseOpenAPIInfo(infoNode, ctx.push("info"))
	if err != nil {
		openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
	}

	if serversNode := nodeGetValue(node, "servers"); serversNode != nil {
		openapi.Servers, err = parseSharedServers(serversNode, ctx.push("servers"))
		if err != nil {
			openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
		}
	}

	pathsNode := nodeGetValue(node, "paths")
	openapi.Paths, err = parseOpenAPIPaths(pathsNode, ctx.push("paths"))
	if err != nil {
		openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
	}

	// Webhooks - new in 3.1
	if webhooksNode := nodeGetValue(node, "webhooks"); webhooksNode != nil {
		openapi.Webhooks, err = parseOpenAPIWebhooks(webhooksNode, ctx.push("webhooks"))
		if err != nil {
			openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
		}
	}

	componentsNode := nodeGetValue(node, "components")
	openapi.Components, err = parseOpenAPIComponents(componentsNode, ctx.push("components"))
	if err != nil {
		openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
	}

	if securityNode := nodeGetValue(node, "security"); securityNode != nil {
		openapi.Security, err = parseSharedSecurityRequirements(securityNode, ctx.push("security"))
		if err != nil {
			openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
		}
	}

	if tagsNode := nodeGetValue(node, "tags"); tagsNode != nil {
		openapi.Tags, err = parseSharedTags(tagsNode, ctx.push("tags"))
		if err != nil {
			openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
		}
	}

	if edNode := nodeGetValue(node, "externalDocs"); edNode != nil {
		openapi.ExternalDocs, err = parseSharedExternalDocs(edNode, ctx.push("externalDocs"))
		if err != nil {
			openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
		}
	}

	openapi.VendorExtensions = parseNodeExtensions(node)
	openapi.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields at root level
	ctx.detectUnknown(node, openAPIKnownFieldsSet)

	return openapi, nil
}

// parseOpenAPIVersion parses and validates the OpenAPI version.
func parseOpenAPIVersion(node *yaml.Node, ctx *ParseContext) (string, error) {
	version := nodeGetString(node, "openapi")
	if version == "" {
		return "", ctx.errorAt(nodeGetKeyNode(node, "openapi"), "openapi field is required")
	}

	// Validate version is 3.1.x or 3.2.x
	if !strings.HasPrefix(version, "3.1.") && !strings.HasPrefix(version, "3.2.") {
		return "", ctx.errorAt(nodeGetValue(node, "openapi"), "unsupported OpenAPI version: %s (expected 3.1.x or 3.2.x)", version)
	}

	return version, nil
}

// parseOpenAPIPaths parses the OpenAPI.Paths field.
func parseOpenAPIPaths(node *yaml.Node, ctx *ParseContext) (*openapi31models.Paths, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "paths must be an object")
	}

	paths := &openapi31models.Paths{}
	paths.Items = make(map[string]*openapi31models.PathItem)

	for key, pathItemNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		pathItem, err := parseOpenAPIPathsPathItem(pathItemNode, ctx.push(key))
		if err != nil {
			return nil, err
		}
		paths.Items[key] = pathItem
	}

	paths.VendorExtensions = parseNodeExtensions(node)
	paths.Trix.Source = ctx.nodeSource(node)

	return paths, nil
}

// parseOpenAPIWebhooks parses the OpenAPI.Webhooks field (new in 3.1).
func parseOpenAPIWebhooks(node *yaml.Node, ctx *ParseContext) (map[string]*openapi31models.PathItemRef, error) {
	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "webhooks must be an object")
	}

	webhooks := make(map[string]*openapi31models.PathItemRef)

	for key, webhookNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		ref, err := parsePathItemRef(webhookNode, ctx.push(key))
		if err != nil {
			return nil, err
		}
		webhooks[key] = ref
	}

	return webhooks, nil
}
