package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_contentschema.go - contentSchema/contentEncoding/contentMediaType
// =============================================================================

func TestParseSchema_ContentSchema_HappyPath(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    EncodedPayload:
      type: string
      contentEncoding: base64
      contentMediaType: application/json
      contentSchema:
        type: object
        properties:
          name:
            type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["EncodedPayload"].Value
	assert.Equal(t, "base64", schema.ContentEncoding())
	assert.Equal(t, "application/json", schema.ContentMediaType())
	require.NotNil(t, schema.ContentSchema())
	assert.Equal(t, "object", schema.ContentSchema().Value.Type().Single)
}

func TestParseSchema_ContentEncodingOnly(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    BinaryData:
      type: string
      contentEncoding: base64
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["BinaryData"].Value
	assert.Equal(t, "base64", schema.ContentEncoding())
	assert.Empty(t, schema.ContentMediaType())
	assert.Nil(t, schema.ContentSchema())
}

func TestParseSchema_NoContentFields(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Plain:
      type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	schema := result.Document.Components().Schemas()["Plain"].Value
	assert.Empty(t, schema.ContentEncoding())
	assert.Empty(t, schema.ContentMediaType())
	assert.Nil(t, schema.ContentSchema())
}
