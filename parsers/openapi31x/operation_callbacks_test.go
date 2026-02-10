package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for operation_callbacks.go - parseOperationCallbacks function
// =============================================================================

// --- Single Callback ---

func TestParseOperationCallbacks_Single(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /subscribe:
    post:
      callbacks:
        onEvent:
          '{$url}':
            post:
              responses:
                "200":
                  description: "OK"
      responses:
        "201":
          description: "Created"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	callbacks := result.Document.Paths().Items()["/subscribe"].Post().Callbacks()
	assert.Len(t, callbacks, 1)
	assert.Contains(t, callbacks, "onEvent")
}

// --- Multiple Callbacks ---

func TestParseOperationCallbacks_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /subscribe:
    post:
      callbacks:
        onSuccess:
          '{$url}':
            post:
              responses:
                "200":
                  description: "OK"
        onFailure:
          '{$url}':
            post:
              responses:
                "200":
                  description: "OK"
        onProgress:
          '{$url}':
            post:
              responses:
                "200":
                  description: "OK"
      responses:
        "201":
          description: "Created"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Paths().Items()["/subscribe"].Post().Callbacks(), 3)
}

// --- No Callbacks ---

func TestParseOperationCallbacks_None(t *testing.T) {
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
	assert.Nil(t, result.Document.Paths().Items()["/pets"].Get().Callbacks())
}
