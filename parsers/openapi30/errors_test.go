package openapi30

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseErrors_MalformedYAML(t *testing.T) {
	yaml := `{{{not valid yaml at all`
	_, err := Parse([]byte(yaml))
	assert.Error(t, err)
}

func TestParseErrors_MissingOpenAPI(t *testing.T) {
	yaml := `info:
  title: "Test"
  version: "1.0"
paths: {}
`
	_, err := Parse([]byte(yaml))
	assert.Error(t, err)
}
