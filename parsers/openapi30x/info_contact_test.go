package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for info_contact.go - parseInfoContact function
// =============================================================================

// --- Complete Contact ---

func TestParseInfoContact_Complete(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
  contact:
    name: "Support Team"
    url: "https://example.com/support"
    email: "support@example.com"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.Contact)
	assert.Equal(t, "Support Team", doc.Info.Contact.Name)
	assert.Equal(t, "https://example.com/support", doc.Info.Contact.URL)
	assert.Equal(t, "support@example.com", doc.Info.Contact.Email)
}

// --- Missing Contact ---

func TestParseInfoContact_Missing(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	// Contact is nil when not provided
	assert.Nil(t, doc.Info.Contact)
}

// --- Partial Fields ---

func TestParseInfoContact_NameOnly(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
  contact:
    name: "John Doe"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.Contact)
	assert.Equal(t, "John Doe", doc.Info.Contact.Name)
	assert.Empty(t, doc.Info.Contact.URL)
	assert.Empty(t, doc.Info.Contact.Email)
}

func TestParseInfoContact_EmailOnly(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
  contact:
    email: "api@example.com"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.Contact)
	assert.Empty(t, doc.Info.Contact.Name)
	assert.Empty(t, doc.Info.Contact.URL)
	assert.Equal(t, "api@example.com", doc.Info.Contact.Email)
}

func TestParseInfoContact_URLOnly(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
  contact:
    url: "https://example.com/contact"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.Contact)
	assert.Empty(t, doc.Info.Contact.Name)
	assert.Equal(t, "https://example.com/contact", doc.Info.Contact.URL)
	assert.Empty(t, doc.Info.Contact.Email)
}

func TestParseInfoContact_NameAndEmail(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
  contact:
    name: "API Support"
    email: "support@example.com"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.Contact)
	assert.Equal(t, "API Support", doc.Info.Contact.Name)
	assert.Empty(t, doc.Info.Contact.URL)
	assert.Equal(t, "support@example.com", doc.Info.Contact.Email)
}

// --- Extensions ---

func TestParseInfoContact_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
  contact:
    name: "Support"
    x-slack: "#api-support"
    x-phone: "+1-555-0100"
    x-hours: "9am-5pm EST"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.Contact)
	require.NotNil(t, doc.Info.Contact.Extensions)
	assert.Equal(t, "#api-support", doc.Info.Contact.Extensions["x-slack"])
	assert.Equal(t, "+1-555-0100", doc.Info.Contact.Extensions["x-phone"])
}

// --- Empty Fields ---

func TestParseInfoContact_EmptyFields(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
  contact:
    name: ""
    url: ""
    email: ""
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.Contact)
	assert.Empty(t, doc.Info.Contact.Name)
	assert.Empty(t, doc.Info.Contact.URL)
	assert.Empty(t, doc.Info.Contact.Email)
}

// --- Special Characters ---

func TestParseInfoContact_SpecialCharacters(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
  contact:
    name: "José García (API Team)"
    email: "jose+api@example.com"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.Contact)
	assert.Equal(t, "José García (API Team)", doc.Info.Contact.Name)
	assert.Equal(t, "jose+api@example.com", doc.Info.Contact.Email)
}
