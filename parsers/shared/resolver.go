package shared

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

// RefResolver resolves $ref references — local (JSON pointer), external (file), and remote (HTTP/HTTPS URL).
// It caches loaded documents and detects circular references.
type RefResolver struct {
	// BasePath is the directory of the root document, used to resolve relative file paths.
	BasePath string

	// Root is the yaml.Node of the root document.
	Root *yaml.Node

	// Fs is the filesystem to use for reading external files.
	// Defaults to afero.OsFs if nil.
	Fs afero.Fs

	// HTTPClient is the HTTP client used for fetching remote $ref URLs.
	// Defaults to a client with a 30-second timeout.
	HTTPClient *http.Client

	// fileCache caches loaded external documents (files and URLs) to avoid re-fetching.
	fileCache map[string]*yaml.Node

	// visiting tracks refs currently being resolved for cycle detection.
	// key is the canonical ref string (e.g. "schemas/pet.yaml#/Pet").
	visiting map[string]bool
}

// defaultHTTPClient is the default HTTP client with a 30-second timeout.
var defaultHTTPClient = &http.Client{Timeout: 30 * time.Second}

// NewRefResolver creates a new RefResolver using the real OS filesystem.
// basePath is the directory containing the root document.
// root is the parsed yaml.Node of the root document.
func NewRefResolver(basePath string, root *yaml.Node) *RefResolver {
	return NewRefResolverWithFs(basePath, root, afero.NewOsFs())
}

// NewRefResolverWithFs creates a new RefResolver with a custom filesystem.
// This is useful for testing with in-memory filesystems.
func NewRefResolverWithFs(basePath string, root *yaml.Node, fs afero.Fs) *RefResolver {
	return &RefResolver{
		BasePath:   basePath,
		Root:       root,
		Fs:         fs,
		HTTPClient: defaultHTTPClient,
		fileCache:  make(map[string]*yaml.Node),
		visiting:   make(map[string]bool),
	}
}

// SplitRef splits a $ref string into a file path and a JSON pointer.
// Examples:
//
//	"#/components/schemas/Pet"          → ("", "/components/schemas/Pet")
//	"./schemas/pet.yaml"                → ("./schemas/pet.yaml", "")
//	"./common.yaml#/definitions/Error"  → ("./common.yaml", "/definitions/Error")
//	"Pet"                               → ("", "Pet")  (bare name, e.g. Swagger 2.0)
func SplitRef(ref string) (filePath, pointer string) {
	if idx := strings.Index(ref, "#"); idx >= 0 {
		return ref[:idx], ref[idx+1:]
	}
	return ref, ""
}

// IsExternalRef returns true if the ref points to an external file or URL.
func IsExternalRef(ref string) bool {
	filePath, _ := SplitRef(ref)
	return filePath != ""
}

// IsLocalRef returns true if the ref is a local JSON pointer (starts with #).
func IsLocalRef(ref string) bool {
	return strings.HasPrefix(ref, "#")
}

// IsRemoteRef returns true if the ref points to a remote URL (http:// or https://).
func IsRemoteRef(ref string) bool {
	filePath, _ := SplitRef(ref)
	return strings.HasPrefix(filePath, "http://") || strings.HasPrefix(filePath, "https://")
}

// ResolveResult contains the result of resolving a $ref.
type ResolveResult struct {
	// Node is the resolved yaml.Node.
	Node *yaml.Node

	// Circular is true if a circular reference was detected.
	Circular bool
}

// Resolve resolves a $ref string to a yaml.Node.
// For local refs (#/path/to/thing), resolves within the root document.
// For external refs (file.yaml or file.yaml#/pointer), loads the file and resolves.
// Returns ResolveResult with Circular=true if a circular reference is detected.
func (r *RefResolver) Resolve(ref string) (*ResolveResult, error) {
	// Canonicalize the ref for cycle detection
	canonicalRef := r.canonicalize(ref)

	// Check for circular reference
	if r.visiting[canonicalRef] {
		return &ResolveResult{Circular: true}, nil
	}
	r.visiting[canonicalRef] = true
	defer func() { delete(r.visiting, canonicalRef) }()

	filePath, pointer := SplitRef(ref)

	var targetRoot *yaml.Node
	if filePath == "" {
		// Local reference — resolve within root document
		targetRoot = r.Root
	} else if isRemoteURL(filePath) {
		// Remote URL reference — fetch and cache
		node, err := r.loadURL(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve remote ref %q: %w", ref, err)
		}
		targetRoot = node
	} else {
		// External file reference — load and cache the file
		node, err := r.loadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve external ref %q: %w", ref, err)
		}
		targetRoot = node
	}

	// If there's no pointer, return the document root
	if pointer == "" {
		return &ResolveResult{Node: targetRoot}, nil
	}

	// Resolve JSON pointer within the target document
	node, err := ResolveJSONPointer(targetRoot, pointer)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve ref %q: %w", ref, err)
	}

	return &ResolveResult{Node: node}, nil
}

// ResolveJSONPointer resolves a JSON pointer (RFC 6901) within a yaml.Node tree.
// The pointer should start with "/" (e.g. "/components/schemas/Pet").
// Handles both mapping nodes (by key lookup) and sequence nodes (by index).
func ResolveJSONPointer(root *yaml.Node, pointer string) (*yaml.Node, error) {
	if pointer == "" || pointer == "/" {
		return root, nil
	}

	// Unwrap document node
	node := root
	if node.Kind == yaml.DocumentNode && len(node.Content) > 0 {
		node = node.Content[0]
	}

	// Remove leading "/"
	if !strings.HasPrefix(pointer, "/") {
		return nil, fmt.Errorf("JSON pointer must start with '/': %q", pointer)
	}
	pointer = pointer[1:]

	// Walk each segment
	segments := strings.Split(pointer, "/")
	for _, rawSegment := range segments {
		// Unescape JSON pointer tokens (RFC 6901)
		segment := unescapeJSONPointer(rawSegment)

		switch node.Kind {
		case yaml.MappingNode:
			found := false
			for i := 0; i < len(node.Content)-1; i += 2 {
				if node.Content[i].Value == segment {
					node = node.Content[i+1]
					found = true
					break
				}
			}
			if !found {
				return nil, fmt.Errorf("JSON pointer segment %q not found", segment)
			}

		case yaml.SequenceNode:
			idx := 0
			if _, err := fmt.Sscanf(segment, "%d", &idx); err != nil {
				return nil, fmt.Errorf("JSON pointer segment %q is not a valid array index", segment)
			}
			if idx < 0 || idx >= len(node.Content) {
				return nil, fmt.Errorf("JSON pointer index %d out of range (length %d)", idx, len(node.Content))
			}
			node = node.Content[idx]

		default:
			return nil, fmt.Errorf("cannot traverse JSON pointer through node of kind %d at segment %q", node.Kind, segment)
		}
	}

	return node, nil
}

// loadFile loads and caches a YAML/JSON file relative to BasePath.
func (r *RefResolver) loadFile(filePath string) (*yaml.Node, error) {
	// Resolve to absolute path
	absPath := filePath
	if !filepath.IsAbs(filePath) {
		absPath = filepath.Join(r.BasePath, filePath)
	}
	absPath = filepath.Clean(absPath)

	// Check cache
	if cached, ok := r.fileCache[absPath]; ok {
		return cached, nil
	}

	// Read file using the configured filesystem
	data, err := afero.ReadFile(r.Fs, absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %q: %w", absPath, err)
	}

	// Parse YAML
	var rootNode yaml.Node
	if err := yaml.Unmarshal(data, &rootNode); err != nil {
		return nil, fmt.Errorf("failed to parse file %q: %w", absPath, err)
	}

	// Unwrap document node
	var docNode *yaml.Node
	if rootNode.Kind == yaml.DocumentNode && len(rootNode.Content) > 0 {
		docNode = rootNode.Content[0]
	} else {
		docNode = &rootNode
	}

	// Cache the result
	r.fileCache[absPath] = docNode

	return docNode, nil
}

// canonicalize creates a canonical key for a ref for cycle detection.
func (r *RefResolver) canonicalize(ref string) string {
	filePath, pointer := SplitRef(ref)
	if filePath == "" {
		return "#" + pointer
	}

	// Remote URLs are already globally unique — use as-is
	if isRemoteURL(filePath) {
		if pointer != "" {
			return filePath + "#" + pointer
		}
		return filePath
	}

	// Resolve to absolute path for consistent keys
	absPath := filePath
	if !filepath.IsAbs(filePath) {
		absPath = filepath.Join(r.BasePath, filePath)
	}
	absPath = filepath.Clean(absPath)

	if pointer != "" {
		return absPath + "#" + pointer
	}
	return absPath
}

// loadURL fetches and caches a remote YAML/JSON document from an HTTP/HTTPS URL.
func (r *RefResolver) loadURL(rawURL string) (*yaml.Node, error) {
	// Check cache
	if cached, ok := r.fileCache[rawURL]; ok {
		return cached, nil
	}

	resp, err := r.HTTPClient.Get(rawURL)
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

	// Parse YAML
	var rootNode yaml.Node
	if err := yaml.Unmarshal(data, &rootNode); err != nil {
		return nil, fmt.Errorf("failed to parse response from %q: %w", rawURL, err)
	}

	// Unwrap document node
	var docNode *yaml.Node
	if rootNode.Kind == yaml.DocumentNode && len(rootNode.Content) > 0 {
		docNode = rootNode.Content[0]
	} else {
		docNode = &rootNode
	}

	// Cache the result
	r.fileCache[rawURL] = docNode

	return docNode, nil
}

// isRemoteURL returns true if the path starts with http:// or https://.
func isRemoteURL(path string) bool {
	return strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://")
}

// unescapeJSONPointer unescapes a JSON pointer token per RFC 6901.
// ~1 → /    ~0 → ~
func unescapeJSONPointer(s string) string {
	s = strings.ReplaceAll(s, "~1", "/")
	s = strings.ReplaceAll(s, "~0", "~")
	// Also handle URL-encoded characters
	if strings.Contains(s, "%") {
		if unescaped, err := url.PathUnescape(s); err == nil {
			s = unescaped
		}
	}
	return s
}
