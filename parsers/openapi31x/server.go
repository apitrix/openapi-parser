package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type serverParser struct{}

// defaultServerParser is the singleton instance used by parsing functions.
var defaultServerParser = &serverParser{}

// parseSharedServer parses a Server object from a yaml.Node.
func parseSharedServer(node *yaml.Node, ctx *ParseContext) (*openapi31models.Server, error) {
	return defaultServerParser.parse(node, ctx)
}

// parseSharedServers parses an array of Server objects.
func parseSharedServers(node *yaml.Node, ctx *ParseContext) ([]*openapi31models.Server, error) {
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	servers := make([]*openapi31models.Server, 0, len(node.Content))
	for i, serverNode := range node.Content {
		server, err := parseSharedServer(serverNode, ctx.push(itoa(i)))
		if err != nil {
			return nil, err
		}
		servers = append(servers, server)
	}
	return servers, nil
}

// Parse parses a Server object.
func (p *serverParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.Server, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "server must be an object")
	}

	var errs []openapi31models.ParseError

	// Complex properties - delegated to dedicated files
	variables, err := p.ParseVariables(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor
	server := openapi31models.NewServer(
		p.ParseURL(node),
		p.ParseDescription(node),
		variables,
	)

	server.VendorExtensions = parseNodeExtensions(node)
	server.Trix.Source = ctx.nodeSource(node)
	server.Trix.Errors = append(server.Trix.Errors, errs...)

	// Detect unknown fields
	server.Trix.Errors = append(server.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, serverKnownFieldsSet))...)

	return server, nil
}

func (p *serverParser) ParseURL(node *yaml.Node) string {
	return nodeGetString(node, "url")
}

func (p *serverParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}
