package openapi31x

import (
	"strings"

	"openapi-parser/models/shared"
	parsersshared "openapi-parser/parsers/shared"
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseOpenAPI parses a full OpenAPI 3.1.x/3.2.x document from a yaml.Node.
// OpenAPI 3.1.0 spec: https://spec.openapis.org/oas/v3.1.0#openapi-object
func parseOpenAPI(node *yaml.Node, ctx *ParseContext) (*openapi31models.OpenAPI, error) {
	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "OpenAPI document must be an object")
	}

	// Parse version first (required)
	version, err := parseOpenAPIVersion(node, ctx)
	if err != nil {
		return nil, err // version is fatal — can't proceed without it
	}

	// Parse info (required sub-object)
	infoNode := nodeGetValue(node, "info")
	info, err := parseOpenAPIInfo(infoNode, ctx.push("info"))

	openapi := openapi31models.NewOpenAPI(version, info)
	if err != nil {
		openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
	}

	// Simple property - jsonSchemaDialect
	if dialect := nodeGetString(node, "jsonSchemaDialect"); dialect != "" {
		openapi.SetProperty("jsonSchemaDialect", dialect)
	}

	// Complex properties - delegated
	serversNode := nodeGetValue(node, "servers")
	if serversNode != nil {
		servers, err := parseSharedServers(serversNode, ctx.push("servers"))
		if err != nil {
			openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
		}
		if servers != nil {
			openapi.SetProperty("servers", servers)
		}
	}
	if parsersshared.ServersAbsentOrEmpty(serversNode) && parsersshared.ApplySpecDefaults(ctx.config) {
		openapi.SetProperty("servers", []*openapi31models.Server{openapi31models.NewServer(parsersshared.DefaultServersURL, "", nil)})
	}

	pathsNode := nodeGetValue(node, "paths")
	paths, err := parseOpenAPIPaths(pathsNode, ctx.push("paths"))
	if err != nil {
		openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
	}
	if paths != nil {
		openapi.SetProperty("paths", paths)
	}

	// Webhooks - new in 3.1
	if webhooksNode := nodeGetValue(node, "webhooks"); webhooksNode != nil {
		webhooks, err := parseOpenAPIWebhooks(webhooksNode, ctx.push("webhooks"))
		if err != nil {
			openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
		}
		if webhooks != nil {
			openapi.SetProperty("webhooks", webhooks)
		}
	}

	componentsNode := nodeGetValue(node, "components")
	components, err := parseOpenAPIComponents(componentsNode, ctx.push("components"))
	if err != nil {
		openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
	}
	if components != nil {
		openapi.SetProperty("components", components)
	}

	if securityNode := nodeGetValue(node, "security"); securityNode != nil {
		security, err := parseSharedSecurityRequirements(securityNode, ctx.push("security"))
		if err != nil {
			openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
		}
		if security != nil {
			openapi.SetProperty("security", security)
		}
	}

	if tagsNode := nodeGetValue(node, "tags"); tagsNode != nil {
		tags, err := parseSharedTags(tagsNode, ctx.push("tags"))
		if err != nil {
			openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
		}
		if tags != nil {
			openapi.SetProperty("tags", tags)
		}
	}

	if edNode := nodeGetValue(node, "externalDocs"); edNode != nil {
		externalDocs, err := parseSharedExternalDocs(edNode, ctx.push("externalDocs"))
		if err != nil {
			openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(err))
		}
		if externalDocs != nil {
			openapi.SetProperty("externalDocs", externalDocs)
		}
	}

	openapi.VendorExtensions = parseNodeExtensions(node)
	openapi.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields at root level
	openapi.Trix.Errors = append(openapi.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, openAPIKnownFieldsSet))...)

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

	items := make(map[string]*openapi31models.PathItem)

	for key, pathItemNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		pathItem, err := parseOpenAPIPathsPathItem(pathItemNode, ctx.push(key))
		if err != nil {
			return nil, err
		}
		items[key] = pathItem
	}

	// Create via constructor
	paths := openapi31models.NewPaths(items)

	paths.VendorExtensions = parseNodeExtensions(node)
	paths.Trix.Source = ctx.nodeSource(node)

	return paths, nil
}

// parseOpenAPIWebhooks parses the OpenAPI.Webhooks field (new in 3.1).
func parseOpenAPIWebhooks(node *yaml.Node, ctx *ParseContext) (map[string]*shared.RefWithMeta[openapi31models.PathItem], error) {
	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "webhooks must be an object")
	}

	webhooks := make(map[string]*shared.RefWithMeta[openapi31models.PathItem])

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
