package openapi20

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	openapi20models "openapi-parser/models/openapi20"
	"openapi-parser/parsers/internal/shared"

	"gopkg.in/yaml.v3"
)

// ParseResult contains the parsed Swagger document along with any unknown
// fields that were detected during parsing.
type ParseResult struct {
	// Document is the parsed Swagger specification.
	Document *openapi20models.Swagger

	// UnknownFields contains all fields that were not recognized as valid
	// Swagger fields during parsing. Extensions (x-*) are NOT included here
	// as they are valid per the Swagger specification.
	UnknownFields []UnknownField
}

// Parse parses Swagger 2.0 specification from bytes (JSON or YAML).
// Uses yaml.Node for lossless parsing with line/column preservation.
// This function does not detect unknown fields; use ParseWithUnknownFields for that.
func Parse(data []byte) (*openapi20models.Swagger, error) {
	result, err := ParseWithUnknownFields(data)
	if err != nil {
		return nil, err
	}
	return result.Document, nil
}

// ParseWithUnknownFields parses Swagger 2.0 specification from bytes (JSON or YAML)
// and detects any unknown fields in the document. Unknown fields are fields that
// are not part of the Swagger 2.0 specification and are not extensions (x-*).
func ParseWithUnknownFields(data []byte) (*ParseResult, error) {
	var rootNode yaml.Node
	if err := yaml.Unmarshal(data, &rootNode); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	// Handle document node wrapper
	var docNode *yaml.Node
	if rootNode.Kind == yaml.DocumentNode && len(rootNode.Content) > 0 {
		docNode = rootNode.Content[0]
	} else {
		docNode = &rootNode
	}

	if docNode.Kind != yaml.MappingNode {
		return nil, fmt.Errorf("Swagger document must be an object")
	}

	ctx := newParseContext(docNode)
	doc, err := parseSwagger(docNode, ctx)
	if err != nil {
		return nil, err
	}

	return &ParseResult{
		Document:      doc,
		UnknownFields: ctx.UnknownFieldsResult(),
	}, nil
}

// ParseReader parses Swagger 2.0 specification from an io.Reader.
func ParseReader(r io.Reader) (*openapi20models.Swagger, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}
	return Parse(data)
}

// ParseReaderWithUnknownFields parses Swagger 2.0 specification from an io.Reader
// and detects any unknown fields in the document.
func ParseReaderWithUnknownFields(r io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}
	return ParseWithUnknownFields(data)
}

// ParseFile parses a Swagger 2.0 specification from a file path,
// resolving all external $ref references relative to the file's directory.
func ParseFile(filePath string) (*openapi20models.Swagger, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var rootNode yaml.Node
	if err := yaml.Unmarshal(data, &rootNode); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	var docNode *yaml.Node
	if rootNode.Kind == yaml.DocumentNode && len(rootNode.Content) > 0 {
		docNode = rootNode.Content[0]
	} else {
		docNode = &rootNode
	}

	if docNode.Kind != yaml.MappingNode {
		return nil, fmt.Errorf("Swagger document must be an object")
	}

	ctx := newParseContext(docNode)
	doc, err := parseSwagger(docNode, ctx)
	if err != nil {
		return nil, err
	}

	basePath := filepath.Dir(absPath)
	if err := Resolve(doc, docNode, basePath); err != nil {
		return nil, fmt.Errorf("failed to resolve references: %w", err)
	}

	return doc, nil
}

// ParseURL parses a Swagger 2.0 specification from an HTTP/HTTPS URL,
// resolving all $ref references relative to the URL's base path.
func ParseURL(rawURL string) (*openapi20models.Swagger, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL %q: %w", rawURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch URL %q: HTTP %d", rawURL, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body from %q: %w", rawURL, err)
	}

	var rootNode yaml.Node
	if err := yaml.Unmarshal(data, &rootNode); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data from %q: %w", rawURL, err)
	}

	var docNode *yaml.Node
	if rootNode.Kind == yaml.DocumentNode && len(rootNode.Content) > 0 {
		docNode = rootNode.Content[0]
	} else {
		docNode = &rootNode
	}

	if docNode.Kind != yaml.MappingNode {
		return nil, fmt.Errorf("Swagger document must be an object")
	}

	ctx := newParseContext(docNode)
	doc, err := parseSwagger(docNode, ctx)
	if err != nil {
		return nil, err
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL %q: %w", rawURL, err)
	}
	basePath := resolveURLBase(parsedURL)

	r := shared.NewRefResolver(basePath, docNode)
	resolving := make(map[string]bool)
	if err := resolveDocument(doc, r, resolving); err != nil {
		return nil, fmt.Errorf("failed to resolve references: %w", err)
	}

	return doc, nil
}

// resolveURLBase extracts the base URL (without the filename) for resolving
// relative $ref references in a remote document.
func resolveURLBase(u *url.URL) string {
	base := *u
	if i := len(base.Path) - 1; i >= 0 {
		for i > 0 && base.Path[i] != '/' {
			i--
		}
		base.Path = base.Path[:i+1]
	}
	return base.String()
}
