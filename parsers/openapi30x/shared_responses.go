package openapi30x

import (
	"regexp"

	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// httpStatusCodePattern matches valid HTTP status codes in responses
var httpStatusCodePattern = regexp.MustCompile(`^[1-5][0-9][0-9]$|^[1-5]XX$`)

// parseSharedResponses parses a Responses object from a yaml.Node.
// OpenAPI 3.0.3 spec: https://spec.openapis.org/oas/v3.0.3#responses-object
func parseSharedResponses(node *yaml.Node, ctx *ParseContext) (*openapi30models.Responses, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "responses must be an object")
	}

	var defaultResp *openapi30models.RefResponse
	codes := make(map[string]*openapi30models.RefResponse)

	for key, valueNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		if key == "default" {
			resp, err := parseResponseRef(valueNode, ctx.push("default"))
			if err != nil {
				return nil, err
			}
			defaultResp = resp
		} else if httpStatusCodePattern.MatchString(key) {
			resp, err := parseResponseRef(valueNode, ctx.push(key))
			if err != nil {
				return nil, err
			}
			codes[key] = resp
		}
	}

	// Create via constructor
	responses := openapi30models.NewResponses(defaultResp, codes)

	responses.VendorExtensions = parseNodeExtensions(node)
	responses.Trix.Source = ctx.nodeSource(node)

	return responses, nil
}
