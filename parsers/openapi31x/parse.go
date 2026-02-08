// package openapi31x provides parsing functionality for OpenAPI 3.1/3.2 specifications.
package openapi31x

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/parsers/internal/shared"

	"gopkg.in/yaml.v3"
)

// ParseResult contains the parsed OpenAPI document along with any unknown
// fields that were detected during parsing.
type ParseResult struct {
	// Document is the parsed OpenAPI specification.
	Document *openapi31models.OpenAPI

	// UnknownFields contains all fields that were not recognized as valid
	// OpenAPI fields during parsing. Extensions (x-*) are NOT included here
	// as they are valid per the OpenAPI specification.
	UnknownFields []UnknownField
}

// Parse parses OpenAPI 3.1/3.2 specification from bytes (JSON or YAML).
// Uses yaml.Node for lossless parsing with line/column preservation.
// This function does not detect unknown fields; use ParseWithUnknownFields for that.
func Parse(data []byte) (*openapi31models.OpenAPI, error) {
	result, err := ParseWithUnknownFields(data)
	if err != nil {
		return nil, err
	}
	return result.Document, nil
}

// ParseWithUnknownFields parses OpenAPI 3.1/3.2 specification from bytes (JSON or YAML)
// and detects any unknown fields in the document. Unknown fields are fields that
// are not part of the OpenAPI 3.1 specification and are not extensions (x-*).
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
		return nil, fmt.Errorf("OpenAPI document must be an object")
	}

	ctx := newParseContext(docNode)
	doc, err := parseOpenAPI(docNode, ctx)
	if err != nil {
		return nil, err
	}

	return &ParseResult{
		Document:      doc,
		UnknownFields: ctx.UnknownFieldsResult(),
	}, nil
}

// ParseReader parses OpenAPI 3.1/3.2 specification from an io.Reader.
func ParseReader(r io.Reader) (*openapi31models.OpenAPI, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}
	return Parse(data)
}

// ParseReaderWithUnknownFields parses OpenAPI 3.1/3.2 specification from an io.Reader
// and detects any unknown fields in the document.
func ParseReaderWithUnknownFields(r io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}
	return ParseWithUnknownFields(data)
}

// ParseFile parses an OpenAPI 3.1/3.2 specification from a file path or HTTP/HTTPS URL,
// resolving all $ref references relative to the source location.
// It auto-detects whether the input is a URL or a local file path.
func ParseFile(pathOrURL string) (*openapi31models.OpenAPI, error) {
	var data []byte
	var basePath string

	if shared.IsRemoteRef(pathOrURL) {
		// Remote URL
		var err error
		data, basePath, err = shared.FetchURL(pathOrURL)
		if err != nil {
			return nil, err
		}
	} else {
		// Local file
		absPath, err := filepath.Abs(pathOrURL)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve absolute path: %w", err)
		}
		data, err = os.ReadFile(absPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}
		basePath = filepath.Dir(absPath)
	}

	return parseAndResolve(data, basePath)
}

// parseAndResolve unmarshals YAML data, parses the OpenAPI document, and resolves all $ref references.
func parseAndResolve(data []byte, basePath string) (*openapi31models.OpenAPI, error) {
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
		return nil, fmt.Errorf("OpenAPI document must be an object")
	}

	ctx := newParseContext(docNode)
	doc, err := parseOpenAPI(docNode, ctx)
	if err != nil {
		return nil, err
	}

	if err := Resolve(doc, docNode, basePath); err != nil {
		return nil, fmt.Errorf("failed to resolve references: %w", err)
	}

	return doc, nil
}
