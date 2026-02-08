package openapi31x

import (
	"regexp"

	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// httpStatusCodePattern matches valid HTTP status codes in responses
var httpStatusCodePattern = regexp.MustCompile(`^[1-5][0-9][0-9]$|^[1-5]XX$`)

// parseSharedResponses parses a Responses object from a yaml.Node.
// OpenAPI 3.0.3 spec: https://spec.openapis.org/oas/v3.0.3#responses-object
func parseSharedResponses(node *yaml.Node, ctx *ParseContext) (*openapi31models.Responses, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "responses must be an object")
	}

	responses := &openapi31models.Responses{}
	responses.Codes = make(map[string]*openapi31models.ResponseRef)

	for key, valueNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		if key == "default" {
			defaultResp, err := parseResponseRef(valueNode, ctx.push("default"))
			if err != nil {
				return nil, err
			}
			responses.Default = defaultResp
		} else if httpStatusCodePattern.MatchString(key) {
			resp, err := parseResponseRef(valueNode, ctx.push(key))
			if err != nil {
				return nil, err
			}
			responses.Codes[key] = resp
		}
	}

	responses.Extensions = parseNodeExtensions(node)
	responses.NodeSource = ctx.nodeSource(node)

	return responses, nil
}
