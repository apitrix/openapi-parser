package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestApplySpecDefaults(t *testing.T) {
	assert.False(t, ApplySpecDefaults(nil))
	assert.False(t, ApplySpecDefaults(&ParseConfig{ApplySpecDefaults: false}))
	assert.True(t, ApplySpecDefaults(&ParseConfig{ApplySpecDefaults: true}))
}

func TestServersAbsentOrEmpty(t *testing.T) {
	assert.True(t, ServersAbsentOrEmpty(nil))

	emptySeq := &yaml.Node{Kind: yaml.SequenceNode, Content: []*yaml.Node{}}
	assert.True(t, ServersAbsentOrEmpty(emptySeq))

	nonEmptySeq := &yaml.Node{Kind: yaml.SequenceNode, Content: []*yaml.Node{{Kind: yaml.MappingNode}}}
	assert.False(t, ServersAbsentOrEmpty(nonEmptySeq))

	scalarNode := &yaml.Node{Kind: yaml.ScalarNode}
	assert.True(t, ServersAbsentOrEmpty(scalarNode))
}

func TestDefaultConstants(t *testing.T) {
	assert.Equal(t, "/", DefaultServersURL)
	assert.Equal(t, "/", DefaultBasePath)
}
