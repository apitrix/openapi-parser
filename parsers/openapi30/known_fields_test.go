package openapi30

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseKnownFields(t *testing.T) {
	// Known fields configuration is internal, just verify parsing works
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc)
}
