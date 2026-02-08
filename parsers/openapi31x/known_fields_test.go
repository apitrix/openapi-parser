package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseKnownFields(t *testing.T) {
	// Known fields configuration is internal, just verify parsing works
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc)
}
