package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseTagExternalDocs(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
tags:
  - name: pets
    externalDocs:
      description: "Pet docs"
      url: "https://example.com/pets"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.Len(t, result.Document.Tags, 1)
	require.NotNil(t, result.Document.Tags[0].ExternalDocs)
	assert.Equal(t, "Pet docs", result.Document.Tags[0].ExternalDocs.Description)
	assert.Equal(t, "https://example.com/pets", result.Document.Tags[0].ExternalDocs.URL)
}
