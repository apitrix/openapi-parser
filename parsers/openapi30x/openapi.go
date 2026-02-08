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

	openapi := &openapi30models.OpenAPI{}
	var err error

	// Simple property - version (inline)
	openapi.OpenAPI, err = parseOpenAPIVersion(node, ctx)
	if err != nil {
		return nil, err // version is fatal — can't proceed without it
	}

	// Complex properties - delegated (errors collected on node)
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

	paths := &openapi30models.Paths{}
	paths.Items = make(map[string]*openapi30models.PathItem)

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
