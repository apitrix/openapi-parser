package openapi30

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
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.Len(t, doc.Tags, 1)
	require.NotNil(t, doc.Tags[0].ExternalDocs)
	assert.Equal(t, "Pet docs", doc.Tags[0].ExternalDocs.Description)
	assert.Equal(t, "https://example.com/pets", doc.Tags[0].ExternalDocs.URL)
}
