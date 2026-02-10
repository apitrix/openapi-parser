package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type callbackParser struct{}

// defaultCallbackParser is the singleton instance used by parsing functions.
var defaultCallbackParser = &callbackParser{}

// parseSharedCallback parses a Callback object from a yaml.Node.
func parseSharedCallback(node *yaml.Node, ctx *ParseContext) (*openapi30models.Callback, error) {
	return defaultCallbackParser.parse(node, ctx)
}

// Parse parses a Callback object.
func (p *callbackParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Callback, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "callback must be an object")
	}

	var errors []openapi30models.ParseError

	// Callbacks are maps of expression -> PathItem
	paths, err := p.ParsePaths(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	// Create via constructor
	callback := openapi30models.NewCallback(paths)

	callback.VendorExtensions = parseNodeExtensions(node)
	callback.Trix.Source = ctx.nodeSource(node)
	callback.Trix.Errors = append(callback.Trix.Errors, errors...)

	return callback, nil
}

// ParsePaths parses all path items in a Callback.
func (p *callbackParser) ParsePaths(node *yaml.Node, c *ParseContext) (map[string]*openapi30models.PathItem, error) {
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	paths := make(map[string]*openapi30models.PathItem)
	ctx := c
	for expr, pathItemNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(expr) > 2 && expr[0] == 'x' && expr[1] == '-' {
			continue
		}

		pathItem, err := parseOpenAPIPathsPathItem(pathItemNode, ctx.push(expr))
		if err != nil {
			return nil, err
		}
		paths[expr] = pathItem
	}
	return paths, nil
}
