// package openapi31x provides parsing functionality for OpenAPI 3.1/3.2 specifications.
package openapi31x

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/parsers/shared"

	"gopkg.in/yaml.v3"
)

// ParseResult contains the parsed OpenAPI document along with any
// errors and unknown fields detected during parsing.
type ParseResult struct {
	// Document is the parsed OpenAPI specification.
	Document *openapi31models.OpenAPI

	// Errors contains all flattened errors (parse errors + unknown field errors)
	// collected from across the entire document tree.
	Errors []*shared.ParseError

	// Config is the ParseConfig that was used for this parse.
	Config *shared.ParseConfig

	// done is closed when background reference resolution is complete.
	done chan struct{}

	// resolveErr holds any error from background reference resolution.
	resolveErr error
}

// Wait blocks until all background reference resolution is complete.
// Returns any error that occurred during resolution.
// It is safe to call Wait multiple times.
func (pr *ParseResult) Wait() error {
	if pr.done != nil {
		<-pr.done
	}
	return pr.resolveErr
}

// Parse parses OpenAPI 3.1/3.2 specification from bytes (JSON or YAML).
// Uses yaml.Node for lossless parsing with line/column preservation.
// An optional ParseConfig controls which features are enabled (nil = All).
func Parse(data []byte, cfgs ...*shared.ParseConfig) (*ParseResult, error) {
	cfg := shared.FirstConfig(cfgs)

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

	ctx := newParseContext(docNode, cfg)
	doc, err := parseOpenAPI(docNode, ctx)
	if err != nil {
		return nil, err
	}

	return &ParseResult{
		Document: doc,
		Errors:   flattenErrors(doc),
		Config:   cfg,
	}, nil
}

// ParseReader parses OpenAPI 3.1/3.2 specification from an io.Reader.
// An optional ParseConfig controls which features are enabled (nil = All).
func ParseReader(r io.Reader, cfgs ...*shared.ParseConfig) (*ParseResult, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}
	return Parse(data, cfgs...)
}

// ParseFile parses an OpenAPI 3.1/3.2 specification from a file path or HTTP/HTTPS URL,
// resolving all $ref references relative to the source location.
// It auto-detects whether the input is a URL or a local file path.
// An optional ParseConfig controls which features are enabled (nil = All).
func ParseFile(pathOrURL string, cfgs ...*shared.ParseConfig) (*ParseResult, error) {
	cfg := shared.FirstConfig(cfgs)

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

	return parseAndResolve(data, basePath, cfg)
}

// parseAndResolve unmarshals YAML data, parses the OpenAPI document, and optionally
// resolves all $ref references in the background based on the config.
func parseAndResolve(data []byte, basePath string, cfg *shared.ParseConfig) (*ParseResult, error) {
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

	ctx := newParseContext(docNode, cfg)
	doc, err := parseOpenAPI(docNode, ctx)
	if err != nil {
		return nil, err
	}

	pr := &ParseResult{
		Document: doc,
		Errors:   flattenErrors(doc),
		Config:   cfg,
	}

	// Resolve $ref references in background if enabled
	if cfg.ResolveInternalRefs || cfg.ResolveExternalRefs {
		initRefDoneChannels(doc)
		pr.done = make(chan struct{})
		go func() {
			pr.resolveErr = Resolve(doc, docNode, basePath)
			close(pr.done)
		}()
	}

	return pr, nil
}
