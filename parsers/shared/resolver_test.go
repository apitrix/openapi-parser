package shared

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// =============================================================================
// SplitRef tests
// =============================================================================

func TestSplitRef_LocalPointer(t *testing.T) {
	file, pointer := SplitRef("#/components/schemas/Pet")
	assert.Equal(t, "", file)
	assert.Equal(t, "/components/schemas/Pet", pointer)
}

func TestSplitRef_FileOnly(t *testing.T) {
	file, pointer := SplitRef("./schemas/pet.yaml")
	assert.Equal(t, "./schemas/pet.yaml", file)
	assert.Equal(t, "", pointer)
}

func TestSplitRef_FileWithPointer(t *testing.T) {
	file, pointer := SplitRef("./common.yaml#/definitions/Error")
	assert.Equal(t, "./common.yaml", file)
	assert.Equal(t, "/definitions/Error", pointer)
}

func TestSplitRef_BareName(t *testing.T) {
	file, pointer := SplitRef("Pet")
	assert.Equal(t, "Pet", file)
	assert.Equal(t, "", pointer)
}

func TestSplitRef_EmptyRef(t *testing.T) {
	file, pointer := SplitRef("")
	assert.Equal(t, "", file)
	assert.Equal(t, "", pointer)
}

func TestSplitRef_RemoteURL(t *testing.T) {
	file, pointer := SplitRef("https://example.com/schemas/pet.yaml#/definitions/Pet")
	assert.Equal(t, "https://example.com/schemas/pet.yaml", file)
	assert.Equal(t, "/definitions/Pet", pointer)
}

func TestSplitRef_RemoteURLNoPointer(t *testing.T) {
	file, pointer := SplitRef("https://example.com/schemas/pet.yaml")
	assert.Equal(t, "https://example.com/schemas/pet.yaml", file)
	assert.Equal(t, "", pointer)
}

func TestIsExternalRef(t *testing.T) {
	assert.True(t, IsExternalRef("./schemas/pet.yaml"))
	assert.True(t, IsExternalRef("./common.yaml#/definitions/Error"))
	assert.True(t, IsExternalRef("https://example.com/pet.yaml"))
	assert.False(t, IsExternalRef("#/components/schemas/Pet"))
	assert.False(t, IsExternalRef(""))
}

func TestIsLocalRef(t *testing.T) {
	assert.True(t, IsLocalRef("#/components/schemas/Pet"))
	assert.False(t, IsLocalRef("./schemas/pet.yaml"))
	assert.False(t, IsLocalRef("https://example.com/pet.yaml"))
	assert.False(t, IsLocalRef(""))
}

func TestIsRemoteRef(t *testing.T) {
	assert.True(t, IsRemoteRef("https://example.com/pet.yaml"))
	assert.True(t, IsRemoteRef("http://example.com/pet.yaml"))
	assert.True(t, IsRemoteRef("https://example.com/common.yaml#/definitions/Error"))
	assert.False(t, IsRemoteRef("./schemas/pet.yaml"))
	assert.False(t, IsRemoteRef("#/components/schemas/Pet"))
	assert.False(t, IsRemoteRef(""))
}

// =============================================================================
// ResolveJSONPointer tests
// =============================================================================

func parseYAML(t *testing.T, data string) *yaml.Node {
	t.Helper()
	var node yaml.Node
	require.NoError(t, yaml.Unmarshal([]byte(data), &node))
	return &node
}

func TestResolveJSONPointer_Simple(t *testing.T) {
	root := parseYAML(t, `
components:
  schemas:
    Pet:
      type: object
`)
	node, err := ResolveJSONPointer(root, "/components/schemas/Pet")
	require.NoError(t, err)
	require.NotNil(t, node)
	assert.Equal(t, yaml.MappingNode, node.Kind)
}

func TestResolveJSONPointer_Nested(t *testing.T) {
	root := parseYAML(t, `
a:
  b:
    c: hello
`)
	node, err := ResolveJSONPointer(root, "/a/b/c")
	require.NoError(t, err)
	require.NotNil(t, node)
	assert.Equal(t, "hello", node.Value)
}

func TestResolveJSONPointer_Root(t *testing.T) {
	root := parseYAML(t, `type: object`)
	node, err := ResolveJSONPointer(root, "")
	require.NoError(t, err)
	require.NotNil(t, node)
}

func TestResolveJSONPointer_NotFound(t *testing.T) {
	root := parseYAML(t, `a: 1`)
	_, err := ResolveJSONPointer(root, "/b")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestResolveJSONPointer_EscapedTilde(t *testing.T) {
	root := parseYAML(t, `
"a/b":
  "c~d": value
`)
	// ~1 = /    ~0 = ~
	node, err := ResolveJSONPointer(root, "/a~1b/c~0d")
	require.NoError(t, err)
	require.NotNil(t, node)
	assert.Equal(t, "value", node.Value)
}

func TestResolveJSONPointer_SequenceIndex(t *testing.T) {
	root := parseYAML(t, `
items:
  - first
  - second
  - third
`)
	node, err := ResolveJSONPointer(root, "/items/1")
	require.NoError(t, err)
	assert.Equal(t, "second", node.Value)
}

func TestResolveJSONPointer_SequenceOutOfRange(t *testing.T) {
	root := parseYAML(t, `
items:
  - first
`)
	_, err := ResolveJSONPointer(root, "/items/5")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "out of range")
}

// =============================================================================
// RefResolver tests with afero in-memory filesystem
// =============================================================================

func TestRefResolver_LocalRef(t *testing.T) {
	root := parseYAML(t, `
components:
  schemas:
    Pet:
      type: object
      properties:
        name:
          type: string
`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())
	result, err := resolver.Resolve("#/components/schemas/Pet")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Circular)
	assert.Equal(t, yaml.MappingNode, result.Node.Kind)
}

func TestRefResolver_ExternalFileRef(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/base/pet.yaml", []byte(`type: object
properties:
  name:
    type: string
`), 0644)

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs("/base", root, fs)

	result, err := resolver.Resolve("pet.yaml")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Circular)
	assert.Equal(t, yaml.MappingNode, result.Node.Kind)
}

func TestRefResolver_ExternalFileWithPointer(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/base/schemas/common.yaml", []byte(`definitions:
  Error:
    type: object
    properties:
      message:
        type: string
`), 0644)

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs("/base", root, fs)

	result, err := resolver.Resolve("schemas/common.yaml#/definitions/Error")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, yaml.MappingNode, result.Node.Kind)
}

func TestRefResolver_FileCaching(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/base/schema.yaml", []byte("type: string"), 0644)

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs("/base", root, fs)

	// Resolve twice — second should use cache
	r1, err := resolver.Resolve("schema.yaml")
	require.NoError(t, err)
	r2, err := resolver.Resolve("schema.yaml")
	require.NoError(t, err)

	// Same node pointer = cache hit
	assert.Equal(t, r1.Node, r2.Node)
}

func TestRefResolver_CircularDetection(t *testing.T) {
	root := parseYAML(t, `
definitions:
  Node:
    type: object
    properties:
      children:
        type: array
`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())

	// Manually mark as visiting to simulate recursion
	canonical := resolver.canonicalize("#/definitions/Node")
	resolver.visiting[canonical] = true

	result, err := resolver.Resolve("#/definitions/Node")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.Circular, "should detect circular reference")
	assert.Nil(t, result.Node)
}

func TestRefResolver_MissingFile(t *testing.T) {
	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())

	_, err := resolver.Resolve("nonexistent.yaml")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to resolve external ref")
}

func TestRefResolver_MissingPointer(t *testing.T) {
	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())

	_, err := resolver.Resolve("#/nonexistent/path")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

// =============================================================================
// Remote ref tests with httptest
// =============================================================================

func TestRefResolver_RemoteRef(t *testing.T) {
	// Arrange — start a test server serving a YAML schema
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		w.Write([]byte("type: object\nproperties:\n  name:\n    type: string\n"))
	}))
	defer srv.Close()

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())
	resolver.HTTPClient = srv.Client()

	// Act
	result, err := resolver.Resolve(srv.URL + "/schemas/pet.yaml")

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Circular)
	assert.Equal(t, yaml.MappingNode, result.Node.Kind)
}

func TestRefResolver_RemoteRefWithPointer(t *testing.T) {
	// Arrange
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		w.Write([]byte("definitions:\n  Error:\n    type: object\n    properties:\n      message:\n        type: string\n"))
	}))
	defer srv.Close()

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())
	resolver.HTTPClient = srv.Client()

	// Act
	result, err := resolver.Resolve(srv.URL + "/common.yaml#/definitions/Error")

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, yaml.MappingNode, result.Node.Kind)
}

func TestRefResolver_RemoteRefCaching(t *testing.T) {
	// Arrange — count requests to verify caching
	requestCount := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Write([]byte("type: string"))
	}))
	defer srv.Close()

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())
	resolver.HTTPClient = srv.Client()

	refURL := srv.URL + "/schema.yaml"

	// Act — resolve same URL twice
	r1, err := resolver.Resolve(refURL)
	require.NoError(t, err)
	r2, err := resolver.Resolve(refURL)
	require.NoError(t, err)

	// Assert — only one HTTP request was made
	assert.Equal(t, 1, requestCount, "second resolve should use cache")
	assert.Equal(t, r1.Node, r2.Node)
}

func TestRefResolver_RemoteRefNotFound(t *testing.T) {
	// Arrange — return 404
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer srv.Close()

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())
	resolver.HTTPClient = srv.Client()

	// Act
	_, err := resolver.Resolve(srv.URL + "/missing.yaml")

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "HTTP 404")
}

func TestRefResolver_RemoteRefInvalidYAML(t *testing.T) {
	// Arrange — return invalid YAML
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{{invalid yaml:::"))
	}))
	defer srv.Close()

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())
	resolver.HTTPClient = srv.Client()

	// Act
	_, err := resolver.Resolve(srv.URL + "/bad.yaml")

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse response")
}

// =============================================================================
// $anchor resolution tests
// =============================================================================

func TestBuildAnchorIndex_FindsAnchors(t *testing.T) {
	// Arrange
	root := parseYAML(t, `
components:
  schemas:
    Pet:
      $anchor: pet
      type: object
      properties:
        name:
          type: string
    Dog:
      $anchor: dog
      type: object
`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())

	// Act
	resolver.BuildAnchorIndex("", root)

	// Assert
	assert.Contains(t, resolver.anchorCache, "")
	assert.Contains(t, resolver.anchorCache[""], "pet")
	assert.Contains(t, resolver.anchorCache[""], "dog")
}

func TestResolve_AnchorRef(t *testing.T) {
	// Arrange
	root := parseYAML(t, `
components:
  schemas:
    Pet:
      $anchor: pet
      type: object
      properties:
        name:
          type: string
`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())
	resolver.BuildAnchorIndex("", root)

	// Act
	result, err := resolver.Resolve("#pet")

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Node)
	assert.False(t, result.Circular)
	// The resolved node should be the mapping node containing $anchor
	assert.Equal(t, yaml.MappingNode, result.Node.Kind)
}

func TestResolve_AnchorNotFound(t *testing.T) {
	// Arrange
	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())
	resolver.BuildAnchorIndex("", root)

	// Act
	_, err := resolver.Resolve("#nonexistent")

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "anchor \"nonexistent\" not found")
}

func TestResolve_AnchorInExternalFile(t *testing.T) {
	// Arrange
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/base/models.yaml", []byte(`
Tag:
  $anchor: tag
  type: object
  properties:
    id:
      type: integer
`), 0644)

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs("/base", root, fs)

	// Load the external file and build its anchor index
	extNode, err := resolver.loadFile("models.yaml")
	require.NoError(t, err)
	resolver.BuildAnchorIndex(resolver.anchorDocKey("models.yaml"), extNode)

	// Act
	result, err := resolver.Resolve("models.yaml#tag")

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Node)
	assert.Equal(t, yaml.MappingNode, result.Node.Kind)
}

// =============================================================================
// $dynamicRef / $dynamicAnchor tests
// =============================================================================

func TestBuildDynamicAnchorIndex_FindsDynamicAnchors(t *testing.T) {
	// Arrange
	root := parseYAML(t, `
components:
  schemas:
    Base:
      $dynamicAnchor: meta
      type: object
    Extension:
      $dynamicAnchor: ext
      type: string
`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())

	// Act
	resolver.BuildDynamicAnchorIndex(root)

	// Assert
	assert.Contains(t, resolver.dynamicAnchorCache, "meta")
	assert.Contains(t, resolver.dynamicAnchorCache, "ext")
}

func TestResolveDynamicRef(t *testing.T) {
	// Arrange
	root := parseYAML(t, `
components:
  schemas:
    Base:
      $dynamicAnchor: meta
      type: object
      properties:
        name:
          type: string
`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())
	resolver.BuildDynamicAnchorIndex(root)

	// Act
	result, err := resolver.ResolveDynamicRef("#meta")

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Node)
	assert.Equal(t, yaml.MappingNode, result.Node.Kind)
}

func TestResolveDynamicRef_NotFound(t *testing.T) {
	// Arrange
	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())
	resolver.BuildDynamicAnchorIndex(root)

	// Act
	_, err := resolver.ResolveDynamicRef("#nonexistent")

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "$dynamicAnchor \"nonexistent\" not found")
}

func TestResolveDynamicRef_Empty(t *testing.T) {
	// Arrange
	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())

	// Act
	_, err := resolver.ResolveDynamicRef("#")

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "empty $dynamicRef")
}

// =============================================================================
// discriminator.mapping resolution tests
// =============================================================================

func TestResolveMapping_BareSchemaName(t *testing.T) {
	// Arrange
	root := parseYAML(t, `
components:
  schemas:
    Dog:
      type: object
      properties:
        breed:
          type: string
`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())

	// Act — bare name "Dog" should resolve as #/components/schemas/Dog
	result, err := resolver.ResolveMapping("Dog")

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Node)
	assert.Equal(t, yaml.MappingNode, result.Node.Kind)
}

func TestResolveMapping_RefString(t *testing.T) {
	// Arrange
	root := parseYAML(t, `
components:
  schemas:
    Cat:
      type: object
`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())

	// Act — explicit ref string
	result, err := resolver.ResolveMapping("#/components/schemas/Cat")

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Node)
}

func TestResolveMapping_ExternalFile(t *testing.T) {
	// Arrange
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/base/models/dog.yaml", []byte(`
type: object
properties:
  breed:
    type: string
`), 0644)

	root := parseYAML(t, `type: object`)
	resolver := NewRefResolverWithFs("/base", root, fs)

	// Act — external file ref
	result, err := resolver.ResolveMapping("./models/dog.yaml")

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Node)
}

func TestResolveMapping_NotFound(t *testing.T) {
	// Arrange
	root := parseYAML(t, `
components:
  schemas:
    Cat:
      type: object
`)
	resolver := NewRefResolverWithFs(".", root, afero.NewMemMapFs())

	// Act
	_, err := resolver.ResolveMapping("NonExistent")

	// Assert
	require.Error(t, err)
}

// =============================================================================
// ParseOperationRef tests
// =============================================================================

func TestParseOperationRef_Simple(t *testing.T) {
	// Act
	path, method, err := ParseOperationRef("#/paths/~1users/get")

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "/users", path)
	assert.Equal(t, "get", method)
}

func TestParseOperationRef_WithPathParam(t *testing.T) {
	// Act
	path, method, err := ParseOperationRef("#/paths/~1users~1{id}/get")

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "/users/{id}", path)
	assert.Equal(t, "get", method)
}

func TestParseOperationRef_NestedPath(t *testing.T) {
	// Act
	path, method, err := ParseOperationRef("#/paths/~1api~1v1~1users~1{userId}~1orders/post")

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "/api/v1/users/{userId}/orders", path)
	assert.Equal(t, "post", method)
}

func TestParseOperationRef_AllMethods(t *testing.T) {
	methods := []string{"get", "put", "post", "delete", "options", "head", "patch", "trace"}
	for _, m := range methods {
		t.Run(m, func(t *testing.T) {
			path, method, err := ParseOperationRef("#/paths/~1test/" + m)
			require.NoError(t, err)
			assert.Equal(t, "/test", path)
			assert.Equal(t, m, method)
		})
	}
}

func TestParseOperationRef_InvalidMethod(t *testing.T) {
	_, _, err := ParseOperationRef("#/paths/~1users/foo")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid HTTP method")
}

func TestParseOperationRef_NoPointer(t *testing.T) {
	_, _, err := ParseOperationRef("Pet")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "has no JSON pointer")
}

func TestParseOperationRef_NotPathsFormat(t *testing.T) {
	_, _, err := ParseOperationRef("#/components/schemas/Pet")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "does not match")
}
