package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type callbackParser struct{}

// defaultCallbackParser is the singleton instance used by parsing functions.
var defaultCallbackParser = &callbackParser{}

// parseSharedCallback parses a Callback object from a yaml.Node.
func parseSharedCallback(node *yaml.Node, ctx *ParseContext) (*openapi30models.Callback, error) {
	return defaultCallbackParser.Parse(node, ctx)
}

// Parse parses a Callback object.
func (p *callbackParser) Parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Callback, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "callback must be an object")
	}

	callback := &openapi30models.Callback{}
	var err error

	// Callbacks are maps of expression -> PathItem
	callback.Paths, err = p.ParsePaths(node, ctx)
	if err != nil {
		return nil, err
	}

	callback.Extensions = parseNodeExtensions(node)
	callback.NodeSource = ctx.nodeSource(node)

	return callback, nil
}

// ParsePaths parses all path items in a Callback.
func (p *callbackParser) ParsePaths(node *yaml.Node, c *ParseContext) (map[string]*openapi30models.PathItem, error) {
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	paths := make(map[string]*openapi30models.PathItem)
	ctx := c
	for _, expr := range nodeKeys(node) {
		// Skip extensions
		if len(expr) > 2 && expr[0] == 'x' && expr[1] == '-' {
			continue
		}

		pathItemNode := nodeGetValue(node, expr)
		pathItem, err := parseOpenAPIPathsPathItem(pathItemNode, ctx.push(expr))
		if err != nil {
			return nil, err
		}
		paths[expr] = pathItem
	}
	return paths, nil
}
