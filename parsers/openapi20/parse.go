package openapi20

import (
	"fmt"
	"io"
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

// ParseFile parses a Swagger 2.0 specification from a file path or HTTP/HTTPS URL,
// resolving all $ref references relative to the source location.
// It auto-detects whether the input is a URL or a local file path.
func ParseFile(pathOrURL string) (*openapi20models.Swagger, error) {
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

// parseAndResolve unmarshals YAML data, parses the Swagger document, and resolves all $ref references.
func parseAndResolve(data []byte, basePath string) (*openapi20models.Swagger, error) {
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

	if err := Resolve(doc, docNode, basePath); err != nil {
		return nil, fmt.Errorf("failed to resolve references: %w", err)
	}

	return doc, nil
}
