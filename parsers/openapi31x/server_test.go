package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for server.go - parseServer function
// =============================================================================

// --- Basic Server ---

func TestParseServer_Basic(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
servers:
  - url: https://api.example.com
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.Len(t, result.Document.Servers(), 1)
	assert.Equal(t, "https://api.example.com", result.Document.Servers()[0].URL())
}

// --- With Description ---

func TestParseServer_WithDescription(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
servers:
  - url: https://api.example.com
    description: "Production server"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "Production server", result.Document.Servers()[0].Description())
}

// --- Multiple Servers ---

func TestParseServer_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
servers:
  - url: https://api.example.com
    description: "Production"
  - url: https://staging-api.example.com
    description: "Staging"
  - url: https://dev-api.example.com
    description: "Development"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Servers(), 3)
}

// --- Variables ---

func TestParseServer_Variables(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
servers:
  - url: https://{environment}.example.com:{port}
    variables:
      environment:
        default: api
        enum:
          - api
          - staging
          - dev
      port:
        default: "443"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	vars := result.Document.Servers()[0].Variables()
	require.NotNil(t, vars)
	assert.Len(t, vars, 2)
	assert.Equal(t, "api", vars["environment"].Default())
	assert.Len(t, vars["environment"].Enum(), 3)
}

// --- Variable with Description ---

func TestParseServer_VariableWithDescription(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
servers:
  - url: https://{env}.example.com
    variables:
      env:
        default: prod
        description: "Environment to connect to"
        enum:
          - prod
          - staging
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	envVar := result.Document.Servers()[0].Variables()["env"]
	assert.Equal(t, "Environment to connect to", envVar.Description())
}

// --- Path-Level Servers ---

func TestParseServer_PathLevel(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
servers:
  - url: https://api.example.com
paths:
  /special:
    servers:
      - url: https://special-api.example.com
    get:
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Servers(), 1)
	assert.Len(t, result.Document.Paths().Items()["/special"].Servers(), 1)
	assert.Equal(t, "https://special-api.example.com", result.Document.Paths().Items()["/special"].Servers()[0].URL())
}

// --- Operation-Level Servers ---

func TestParseServer_OperationLevel(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
servers:
  - url: https://api.example.com
paths:
  /pets:
    get:
      servers:
        - url: https://pets-api.example.com
        - url: https://backup-pets-api.example.com
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	opServers := result.Document.Paths().Items()["/pets"].Get().Servers()
	assert.Len(t, opServers, 2)
}

// --- Extensions ---

func TestParseServer_Extensions(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
servers:
  - url: https://api.example.com
    x-internal: true
    x-region: "us-east-1"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ext := result.Document.Servers()[0].VendorExtensions
	require.NotNil(t, ext)
	assert.Equal(t, true, ext["x-internal"])
	assert.Equal(t, "us-east-1", ext["x-region"])
}

// --- Empty Servers ---

func TestParseServer_Empty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Servers())
}

// --- Relative URL ---

func TestParseServer_RelativeURL(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
servers:
  - url: /v1
    description: "Relative URL"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "/v1", result.Document.Servers()[0].URL())
}
