package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for components_callbacks.go - parseComponentsCallbacks function
// =============================================================================

// --- Single Callback ---

func TestParseComponentsCallbacks_Single(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  callbacks:
    Webhook:
      '{$url}':
        post:
          responses:
            "200":
              description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Callbacks, 1)
	assert.Contains(t, doc.Components.Callbacks, "Webhook")
}

// --- Multiple Callbacks ---

func TestParseComponentsCallbacks_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  callbacks:
    OnSuccess:
      '{$url}':
        post:
          responses:
            "200":
              description: "OK"
    OnFailure:
      '{$url}':
        post:
          responses:
            "200":
              description: "OK"
    OnProgress:
      '{$url}':
        post:
          responses:
            "200":
              description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Components.Callbacks, 3)
}

// --- Empty ---

func TestParseComponentsCallbacks_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  callbacks: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, doc.Components.Callbacks)
}

// --- Missing ---

func TestParseComponentsCallbacks_Missing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
components: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Nil(t, doc.Components.Callbacks)
}
