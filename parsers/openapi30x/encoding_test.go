package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for encoding.go - parseEncoding function
// =============================================================================

// --- Basic Encoding ---

func TestParseEncoding_Basic(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /upload:
    post:
      requestBody:
        content:
          multipart/form-data:
            schema:
              type: object
            encoding:
              file:
                contentType: application/octet-stream
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	enc := result.Document.Paths.Items["/upload"].Post.RequestBody.Value.Content["multipart/form-data"].Encoding["file"]
	require.NotNil(t, enc)
	assert.Equal(t, "application/octet-stream", enc.ContentType)
}

// --- Multiple Encodings ---

func TestParseEncoding_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /upload:
    post:
      requestBody:
        content:
          multipart/form-data:
            schema:
              type: object
            encoding:
              file:
                contentType: application/octet-stream
              metadata:
                contentType: application/json
              tags:
                contentType: text/csv
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	enc := result.Document.Paths.Items["/upload"].Post.RequestBody.Value.Content["multipart/form-data"].Encoding
	assert.Len(t, enc, 3)
}

// --- Headers ---

func TestParseEncoding_Headers(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /upload:
    post:
      requestBody:
        content:
          multipart/form-data:
            schema:
              type: object
            encoding:
              file:
                contentType: application/octet-stream
                headers:
                  X-Custom-Header:
                    schema:
                      type: string
                  X-Another-Header:
                    schema:
                      type: integer
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	enc := result.Document.Paths.Items["/upload"].Post.RequestBody.Value.Content["multipart/form-data"].Encoding["file"]
	assert.Len(t, enc.Headers, 2)
}

// --- Style and Explode ---

func TestParseEncoding_StyleExplode(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /upload:
    post:
      requestBody:
        content:
          application/x-www-form-urlencoded:
            schema:
              type: object
            encoding:
              tags:
                style: form
                explode: true
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	enc := result.Document.Paths.Items["/upload"].Post.RequestBody.Value.Content["application/x-www-form-urlencoded"].Encoding["tags"]
	assert.Equal(t, "form", enc.Style)
	require.NotNil(t, enc.Explode)
	assert.True(t, *enc.Explode)
}

// --- AllowReserved ---

func TestParseEncoding_AllowReserved(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /upload:
    post:
      requestBody:
        content:
          application/x-www-form-urlencoded:
            schema:
              type: object
            encoding:
              path:
                allowReserved: true
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	enc := result.Document.Paths.Items["/upload"].Post.RequestBody.Value.Content["application/x-www-form-urlencoded"].Encoding["path"]
	assert.True(t, enc.AllowReserved)
}

// --- Extensions ---

func TestParseEncoding_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /upload:
    post:
      requestBody:
        content:
          multipart/form-data:
            schema:
              type: object
            encoding:
              file:
                contentType: application/octet-stream
                x-custom: "value"
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	enc := result.Document.Paths.Items["/upload"].Post.RequestBody.Value.Content["multipart/form-data"].Encoding["file"]
	require.NotNil(t, enc.VendorExtensions)
	assert.Equal(t, "value", enc.VendorExtensions["x-custom"])
}

// --- Empty Encoding ---

func TestParseEncoding_Empty(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	enc := result.Document.Paths.Items["/pets"].Post.RequestBody.Value.Content["application/json"].Encoding
	assert.Empty(t, enc)
}
