package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for callback.go - parseCallback function
// =============================================================================

// --- Basic Callback ---

func TestParseCallback_Basic(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /subscribe:
    post:
      callbacks:
        onEvent:
          '{$request.body#/callbackUrl}':
            post:
              responses:
                "200":
                  description: "OK"
      responses:
        "201":
          description: "Subscribed"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	callbacks := doc.Paths.Items["/subscribe"].Post.Callbacks
	require.NotNil(t, callbacks)
	assert.Contains(t, callbacks, "onEvent")
}

// --- Multiple Callbacks ---

func TestParseCallback_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /subscribe:
    post:
      callbacks:
        onSuccess:
          '{$request.body#/successUrl}':
            post:
              responses:
                "200":
                  description: "OK"
        onFailure:
          '{$request.body#/failureUrl}':
            post:
              responses:
                "200":
                  description: "OK"
        onProgress:
          '{$request.body#/progressUrl}':
            post:
              responses:
                "200":
                  description: "OK"
      responses:
        "201":
          description: "Subscribed"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, doc.Paths.Items["/subscribe"].Post.Callbacks, 3)
}

// --- Expression Path ---

func TestParseCallback_ExpressionPath(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /subscribe:
    post:
      callbacks:
        onEvent:
          'http://example.com/{$request.body#/id}':
            post:
              responses:
                "200":
                  description: "OK"
      responses:
        "201":
          description: "Subscribed"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	callback := doc.Paths.Items["/subscribe"].Post.Callbacks["onEvent"].Value
	assert.Contains(t, callback.Paths, "http://example.com/{$request.body#/id}")
}

// --- Multiple URLs ---

func TestParseCallback_MultipleURLs(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /subscribe:
    post:
      callbacks:
        onEvent:
          '{$request.body#/primaryUrl}':
            post:
              responses:
                "200":
                  description: "OK"
          '{$request.body#/backupUrl}':
            post:
              responses:
                "200":
                  description: "OK"
      responses:
        "201":
          description: "Subscribed"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	callback := doc.Paths.Items["/subscribe"].Post.Callbacks["onEvent"].Value
	assert.Len(t, callback.Paths, 2)
}

// --- In Components ---

func TestParseCallback_InComponents(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  callbacks:
    WebhookCallback:
      '{$url}':
        post:
          requestBody:
            content:
              application/json:
                schema:
                  type: object
          responses:
            "200":
              description: "OK"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Contains(t, doc.Components.Callbacks, "WebhookCallback")
}

// --- Reference ---

func TestParseCallback_Reference(t *testing.T) {
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
	callbackRef := doc.Paths.Items["/subscribe"].Post.Callbacks["onEvent"]
	assert.Equal(t, "#/components/callbacks/WebhookCallback", callbackRef.Ref)
}

// --- Extensions ---

func TestParseCallback_Extensions(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /subscribe:
    post:
      callbacks:
        onEvent:
          x-custom: "value"
          '{$url}':
            post:
              responses:
                "200":
                  description: "OK"
      responses:
        "201":
          description: "Subscribed"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	callback := doc.Paths.Items["/subscribe"].Post.Callbacks["onEvent"].Value
	require.NotNil(t, callback.VendorExtensions)
	assert.Equal(t, "value", callback.VendorExtensions["x-custom"])
}

// --- Full Callback Path ---

func TestParseCallback_FullPathItem(t *testing.T) {
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
            summary: "Webhook endpoint"
            post:
              summary: "Receive webhook"
              requestBody:
                content:
                  application/json:
                    schema:
                      type: object
              responses:
                "200":
                  description: "OK"
            get:
              responses:
                "200":
                  description: "Check status"
      responses:
        "201":
          description: "Subscribed"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	pathItem := doc.Paths.Items["/subscribe"].Post.Callbacks["onEvent"].Value.Paths["{$url}"]
	require.NotNil(t, pathItem)
	assert.NotNil(t, pathItem.Post)
	assert.NotNil(t, pathItem.Get)
}
