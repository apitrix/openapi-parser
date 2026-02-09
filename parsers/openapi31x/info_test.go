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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "My API", result.Document.Info.Title)
	assert.Equal(t, "1.0.0", result.Document.Info.Version)
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Contains(t, result.Document.Info.Description, "multi-line")
	assert.Contains(t, result.Document.Info.Description, "multiple paragraphs")
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "https://example.com/terms", result.Document.Info.TermsOfService)
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info.Contact)
	assert.Equal(t, "API Support", result.Document.Info.Contact.Name)
	assert.Equal(t, "https://example.com/support", result.Document.Info.Contact.URL)
	assert.Equal(t, "support@example.com", result.Document.Info.Contact.Email)
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info.Contact)
	assert.Equal(t, "support@example.com", result.Document.Info.Contact.Email)
	assert.Empty(t, result.Document.Info.Contact.Name)
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info.License)
	assert.Equal(t, "Apache 2.0", result.Document.Info.License.Name)
	assert.Equal(t, "https://www.apache.org/licenses/LICENSE-2.0", result.Document.Info.License.URL)
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info.License)
	assert.Equal(t, "MIT", result.Document.Info.License.Name)
	assert.Empty(t, result.Document.Info.License.URL)
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "Pet Store API", result.Document.Info.Title)
	assert.Equal(t, "1.2.3", result.Document.Info.Version)
	assert.NotEmpty(t, result.Document.Info.Description)
	assert.NotEmpty(t, result.Document.Info.TermsOfService)
	require.NotNil(t, result.Document.Info.Contact)
	require.NotNil(t, result.Document.Info.License)
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info.VendorExtensions)
	assert.Equal(t, true, result.Document.Info.VendorExtensions["x-internal"])
}

// --- Missing Optional Fields ---

func TestParseInfo_MinimalFields(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "API"
  version: "1.0"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Info.Description)
	assert.Empty(t, result.Document.Info.TermsOfService)
	assert.Nil(t, result.Document.Info.Contact)
	assert.Nil(t, result.Document.Info.License)
}
