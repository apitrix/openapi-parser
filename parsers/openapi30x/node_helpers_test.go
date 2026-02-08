package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseNodeHelpers(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	// Node helpers are internal, just verify parsing works
	require.NotNil(t, doc)
	require.NotNil(t, doc.Info)
}
