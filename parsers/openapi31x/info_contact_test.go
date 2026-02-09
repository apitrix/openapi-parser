package openapi31x

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
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  contact:
    name: "Support Team"
    url: "https://example.com/support"
    email: "support@example.com"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info)
	require.NotNil(t, result.Document.Info.Contact)
	assert.Equal(t, "Support Team", result.Document.Info.Contact.Name)
	assert.Equal(t, "https://example.com/support", result.Document.Info.Contact.URL)
	assert.Equal(t, "support@example.com", result.Document.Info.Contact.Email)
}

// --- Missing Contact ---

func TestParseInfoContact_Missing(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info)
	// Contact is nil when not provided
	assert.Nil(t, result.Document.Info.Contact)
}

// --- Partial Fields ---

func TestParseInfoContact_NameOnly(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  contact:
    name: "John Doe"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info)
	require.NotNil(t, result.Document.Info.Contact)
	assert.Equal(t, "John Doe", result.Document.Info.Contact.Name)
	assert.Empty(t, result.Document.Info.Contact.URL)
	assert.Empty(t, result.Document.Info.Contact.Email)
}

func TestParseInfoContact_EmailOnly(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  contact:
    email: "api@example.com"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info)
	require.NotNil(t, result.Document.Info.Contact)
	assert.Empty(t, result.Document.Info.Contact.Name)
	assert.Empty(t, result.Document.Info.Contact.URL)
	assert.Equal(t, "api@example.com", result.Document.Info.Contact.Email)
}

func TestParseInfoContact_URLOnly(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  contact:
    url: "https://example.com/contact"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info)
	require.NotNil(t, result.Document.Info.Contact)
	assert.Empty(t, result.Document.Info.Contact.Name)
	assert.Equal(t, "https://example.com/contact", result.Document.Info.Contact.URL)
	assert.Empty(t, result.Document.Info.Contact.Email)
}

func TestParseInfoContact_NameAndEmail(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  contact:
    name: "API Support"
    email: "support@example.com"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info)
	require.NotNil(t, result.Document.Info.Contact)
	assert.Equal(t, "API Support", result.Document.Info.Contact.Name)
	assert.Empty(t, result.Document.Info.Contact.URL)
	assert.Equal(t, "support@example.com", result.Document.Info.Contact.Email)
}

// --- Extensions ---

func TestParseInfoContact_Extensions(t *testing.T) {
	yaml := `openapi: "3.1.0"
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info)
	require.NotNil(t, result.Document.Info.Contact)
	require.NotNil(t, result.Document.Info.Contact.VendorExtensions)
	assert.Equal(t, "#api-support", result.Document.Info.Contact.VendorExtensions["x-slack"])
	assert.Equal(t, "+1-555-0100", result.Document.Info.Contact.VendorExtensions["x-phone"])
}

// --- Empty Fields ---

func TestParseInfoContact_EmptyFields(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  contact:
    name: ""
    url: ""
    email: ""
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info)
	require.NotNil(t, result.Document.Info.Contact)
	assert.Empty(t, result.Document.Info.Contact.Name)
	assert.Empty(t, result.Document.Info.Contact.URL)
	assert.Empty(t, result.Document.Info.Contact.Email)
}

// --- Special Characters ---

func TestParseInfoContact_SpecialCharacters(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  contact:
    name: "José García (API Team)"
    email: "jose+api@example.com"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info)
	require.NotNil(t, result.Document.Info.Contact)
	assert.Equal(t, "José García (API Team)", result.Document.Info.Contact.Name)
	assert.Equal(t, "jose+api@example.com", result.Document.Info.Contact.Email)
}
