package openapi30x

import (
	"strings"

	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseOpenAPI parses a full OpenAPI 3.0.x document from a yaml.Node.
// OpenAPI 3.0.3 spec: https://spec.openapis.org/oas/v3.0.3#openapi-object
func parseOpenAPI(node *yaml.Node, ctx *ParseContext) (*openapi30models.OpenAPI, error) {
	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "OpenAPI document must be an object")
	}

	var err error

	// Simple property - version (inline)
	version, err := parseOpenAPIVersion(node, ctx)
	if err != nil {
		return nil, err // version is fatal — can't proceed without it
	}

	// Complex properties - delegated (errors collected on node)
	infoNode := nodeGetValue(node, "info")
	info, infoErr := parseOpenAPIInfo(infoNode, ctx.push("info"))

	var servers []*openapi30models.Server
	if serversNode := nodeGetValue(node, "servers"); serversNode != nil {
		servers, err = parseSharedServers(serversNode, ctx.push("servers"))
		if err != nil {
			infoErr = err // will collect below
		}
	}

	pathsNode := nodeGetValue(node, "paths")
	paths, pathsErr := parseOpenAPIPaths(pathsNode, ctx.push("paths"))

	componentsNode := nodeGetValue(node, "components")
	components, componentsErr := parseOpenAPIComponents(componentsNode, ctx.push("components"))

	var security []openapi30models.SecurityRequirement
	if securityNode := nodeGetValue(node, "security"); securityNode != nil {
		security, err = parseSharedSecurityRequirements(securityNode, ctx.push("security"))
		if err != nil {
			// collect below
			_ = err
		}
	}

	var tags []*openapi30models.Tag
	if tagsNode := nodeGetValue(node, "tags"); tagsNode != nil {
		tags, err = parseSharedTags(tagsNode, ctx.push("tags"))
		if err != nil {
			_ = err
		}
	}

	var externalDocs *openapi30models.ExternalDocumentation
	if edNode := nodeGetValue(node, "externalDocs"); edNode != nil {
		externalDocs, err = parseSharedExternalDocs(edNode, ctx.push("externalDocs"))
		if err != nil {
			_ = err
		}
	}

	// Create the OpenAPI document using constructor + SetProperty for incremental fields
	openapi := openapi30models.NewOpenAPI(version, info)
	openapi.SetProperty("servers", servers)
	openapi.SetProperty("paths", paths)
	openapi.SetProperty("components", components)
	openapi.SetProperty("security", security)
	openapi.SetProperty("tags", tags)
	openapi.SetProperty("externalDocs", externalDocs)

	// Collect errors
	if infoErr != nil {
		openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(infoErr))
	}
	if pathsErr != nil {
		openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(pathsErr))
	}
	if componentsErr != nil {
		openapi.Trix.Errors = append(openapi.Trix.Errors, toParseError(componentsErr))
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

	// Validate version is 3.0.x
	if !strings.HasPrefix(version, "3.0.") {
		return "", ctx.errorAt(nodeGetValue(node, "openapi"), "unsupported OpenAPI version: %s (expected 3.0.x)", version)
	}

	return version, nil
}

// parseOpenAPIPaths parses the OpenAPI.Paths field.
func parseOpenAPIPaths(node *yaml.Node, ctx *ParseContext) (*openapi30models.Paths, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "paths must be an object")
	}

	items := make(map[string]*openapi30models.PathItem)

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

	paths := openapi30models.NewPaths(items)
	paths.VendorExtensions = parseNodeExtensions(node)
	paths.Trix.Source = ctx.nodeSource(node)

	return paths, nil
}
