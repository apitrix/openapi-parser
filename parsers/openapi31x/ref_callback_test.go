package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for ref_callback.go - callback reference parsing
// =============================================================================

// --- Basic Reference ---

func TestParseRefCallback_Basic(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /subscribe:
    post:
      callbacks:
        onEvent:
          $ref: '#/components/callbacks/WebhookCallback'
      responses:
        "201":
          description: "Subscribed"
components:
  callbacks:
    WebhookCallback:
      '{$url}':
        post:
          responses:
            "200":
              description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := doc.Paths.Items["/subscribe"].Post.Callbacks["onEvent"]
	assert.Equal(t, "#/components/callbacks/WebhookCallback", ref.Ref)
}

// --- Multiple References ---

func TestParseRefCallback_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /subscribe:
    post:
      callbacks:
        onSuccess:
          $ref: '#/components/callbacks/SuccessCallback'
        onFailure:
          $ref: '#/components/callbacks/FailureCallback'
      responses:
        "201":
          description: "Subscribed"
components:
  callbacks:
    SuccessCallback:
      '{$url}':
        post:
          responses:
            "200":
              description: "OK"
    FailureCallback:
      '{$url}':
        post:
          responses:
            "200":
              description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	callbacks := doc.Paths.Items["/subscribe"].Post.Callbacks
	assert.Equal(t, "#/components/callbacks/SuccessCallback", callbacks["onSuccess"].Ref)
	assert.Equal(t, "#/components/callbacks/FailureCallback", callbacks["onFailure"].Ref)
}

// --- Mixed Inline and Reference ---

func TestParseRefCallback_Mixed(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /subscribe:
    post:
      callbacks:
        onEvent:
          $ref: '#/components/callbacks/WebhookCallback'
        onProgress:
          '{$url}':
            post:
              responses:
                "200":
                  description: "OK"
      responses:
        "201":
          description: "Subscribed"
components:
  callbacks:
    WebhookCallback:
      '{$url}':
        post:
          responses:
            "200":
              description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	callbacks := doc.Paths.Items["/subscribe"].Post.Callbacks
	assert.Equal(t, "#/components/callbacks/WebhookCallback", callbacks["onEvent"].Ref)
	assert.NotNil(t, callbacks["onProgress"].Value)
}
