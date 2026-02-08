package openapi30x

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
	yaml := `openapi: "3.0.3"
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	callbacks := doc.Paths.Items["/subscribe"].Post.Callbacks
	assert.Len(t, callbacks, 1)
	assert.Contains(t, callbacks, "onEvent")
}

// --- Multiple Callbacks ---

func TestParseOperationCallbacks_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Paths.Items["/subscribe"].Post.Callbacks, 3)
}

// --- No Callbacks ---

func TestParseOperationCallbacks_None(t *testing.T) {
	yaml := `openapi: "3.0.3"
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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Nil(t, doc.Paths.Items["/pets"].Get.Callbacks)
}
