package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"
	"gopkg.in/yaml.v3"
)

type serverParser struct{}

// defaultServerParser is the singleton instance used by parsing functions.
var defaultServerParser = &serverParser{}

// parseSharedServer parses a Server object from a yaml.Node.
func parseSharedServer(node *yaml.Node, ctx *ParseContext) (*openapi30models.Server, error) {
	return defaultServerParser.parse(node, ctx)
}

// parseSharedServers parses an array of Server objects.
func parseSharedServers(node *yaml.Node, ctx *ParseContext) ([]*openapi30models.Server, error) {
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	servers := make([]*openapi30models.Server, 0, len(node.Content))
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
func (p *serverParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Server, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "server must be an object")
	}

	server := &openapi30models.Server{}
	var err error

	// Simple properties - inline
	server.URL = p.ParseURL(node)
	server.Description = p.ParseDescription(node)

	// Complex properties - delegated to dedicated files
	server.Variables, err = p.ParseVariables(node, ctx)
	if err != nil {
		return nil, err
	}

	server.Extensions = parseNodeExtensions(node)
	server.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, serverKnownFields)

	return server, nil
}

func (p *serverParser) ParseURL(node *yaml.Node) string {
	return nodeGetString(node, "url")
}

func (p *serverParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}
