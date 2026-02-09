package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for operation_servers.go - parseOperationServers function
// =============================================================================

// --- Single Server ---

func TestParseOperationServers_Single(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      servers:
        - url: https://pets.example.com
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	servers := result.Document.Paths.Items["/pets"].Get.Servers
	assert.Len(t, servers, 1)
	assert.Equal(t, "https://pets.example.com", servers[0].URL)
}

// --- Multiple Servers ---

func TestParseOperationServers_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      servers:
        - url: https://primary.example.com
          description: "Primary"
        - url: https://backup.example.com
          description: "Backup"
        - url: https://fallback.example.com
          description: "Fallback"
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	servers := result.Document.Paths.Items["/pets"].Get.Servers
	assert.Len(t, servers, 3)
}

// --- No Operation Servers ---

func TestParseOperationServers_None(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	// Uses path-level or global servers
	assert.Empty(t, result.Document.Paths.Items["/pets"].Get.Servers)
}

// --- With Variables ---

func TestParseOperationServers_WithVariables(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      servers:
        - url: https://{env}.example.com
          variables:
            env:
              default: prod
              enum:
                - prod
                - staging
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	servers := result.Document.Paths.Items["/pets"].Get.Servers
	require.Len(t, servers, 1)
	assert.NotNil(t, servers[0].Variables)
	assert.Equal(t, "prod", servers[0].Variables["env"].Default)
}
