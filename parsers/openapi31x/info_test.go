package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for info.go - parseInfo function
// =============================================================================

// --- Required Fields ---

func TestParseInfo_RequiredFields(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "My API"
  version: "1.0.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "My API", doc.Info.Title)
	assert.Equal(t, "1.0.0", doc.Info.Version)
}

// --- Description ---

func TestParseInfo_Description(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "API"
  version: "1.0"
  description: |
    This is a multi-line
    description of the API.
    
    It can have multiple paragraphs.
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Contains(t, doc.Info.Description, "multi-line")
	assert.Contains(t, doc.Info.Description, "multiple paragraphs")
}

// --- Terms of Service ---

func TestParseInfo_TermsOfService(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "API"
  version: "1.0"
  termsOfService: "https://example.com/terms"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "https://example.com/terms", doc.Info.TermsOfService)
}

// --- Contact ---

func TestParseInfo_Contact(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "API"
  version: "1.0"
  contact:
    name: "API Support"
    url: "https://example.com/support"
    email: "support@example.com"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info.Contact)
	assert.Equal(t, "API Support", doc.Info.Contact.Name)
	assert.Equal(t, "https://example.com/support", doc.Info.Contact.URL)
	assert.Equal(t, "support@example.com", doc.Info.Contact.Email)
}

func TestParseInfo_ContactPartial(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "API"
  version: "1.0"
  contact:
    email: "support@example.com"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info.Contact)
	assert.Equal(t, "support@example.com", doc.Info.Contact.Email)
	assert.Empty(t, doc.Info.Contact.Name)
}

// --- License ---

func TestParseInfo_License(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "API"
  version: "1.0"
  license:
    name: "Apache 2.0"
    url: "https://www.apache.org/licenses/LICENSE-2.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info.License)
	assert.Equal(t, "Apache 2.0", doc.Info.License.Name)
	assert.Equal(t, "https://www.apache.org/licenses/LICENSE-2.0", doc.Info.License.URL)
}

func TestParseInfo_LicenseNameOnly(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "API"
  version: "1.0"
  license:
    name: "MIT"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info.License)
	assert.Equal(t, "MIT", doc.Info.License.Name)
	assert.Empty(t, doc.Info.License.URL)
}

// --- Complete Info ---

func TestParseInfo_Complete(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Pet Store API"
  version: "1.2.3"
  description: "A sample pet store API"
  termsOfService: "https://example.com/terms"
  contact:
    name: "Support Team"
    url: "https://example.com/support"
    email: "support@example.com"
  license:
    name: "Apache 2.0"
    url: "https://www.apache.org/licenses/LICENSE-2.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "Pet Store API", doc.Info.Title)
	assert.Equal(t, "1.2.3", doc.Info.Version)
	assert.NotEmpty(t, doc.Info.Description)
	assert.NotEmpty(t, doc.Info.TermsOfService)
	require.NotNil(t, doc.Info.Contact)
	require.NotNil(t, doc.Info.License)
}

// --- Extensions ---

func TestParseInfo_Extensions(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "API"
  version: "1.0"
  x-logo:
    url: "https://example.com/logo.png"
  x-internal: true
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info.Extensions)
	assert.Equal(t, true, doc.Info.Extensions["x-internal"])
}

// --- Missing Optional Fields ---

func TestParseInfo_MinimalFields(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "API"
  version: "1.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, doc.Info.Description)
	assert.Empty(t, doc.Info.TermsOfService)
	assert.Nil(t, doc.Info.Contact)
	assert.Nil(t, doc.Info.License)
}
