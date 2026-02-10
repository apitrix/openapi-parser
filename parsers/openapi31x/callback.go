package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type callbackParser struct{}

// defaultCallbackParser is the singleton instance used by parsing functions.
var defaultCallbackParser = &callbackParser{}

// parseSharedCallback parses a Callback object from a yaml.Node.
func parseSharedCallback(node *yaml.Node, ctx *ParseContext) (*openapi31models.Callback, error) {
	return defaultCallbackParser.parse(node, ctx)
}

// Parse parses a Callback object.
func (p *callbackParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.Callback, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "callback must be an object")
	}

	var errs []openapi31models.ParseError

	// Callbacks are maps of expression -> PathItem
	paths, err := p.ParsePaths(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	callback := openapi31models.NewCallback(paths)

	callback.VendorExtensions = parseNodeExtensions(node)
	callback.Trix.Source = ctx.nodeSource(node)
	callback.Trix.Errors = append(callback.Trix.Errors, errs...)

	return callback, nil
}

// ParsePaths parses all path items in a Callback.
func (p *callbackParser) ParsePaths(node *yaml.Node, c *ParseContext) (map[string]*openapi31models.PathItem, error) {
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	paths := make(map[string]*openapi31models.PathItem)
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
