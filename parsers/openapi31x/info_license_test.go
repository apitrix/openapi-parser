package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for info_license.go - parseInfoLicense function
// =============================================================================

// --- Complete License ---

func TestParseInfoLicense_Complete(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  license:
    name: "MIT"
    url: "https://opensource.org/licenses/MIT"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.License)
	assert.Equal(t, "MIT", doc.Info.License.Name)
	assert.Equal(t, "https://opensource.org/licenses/MIT", doc.Info.License.URL)
}

// --- Name Only (URL optional) ---

func TestParseInfoLicense_NameOnly(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  license:
    name: "Apache 2.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.License)
	assert.Equal(t, "Apache 2.0", doc.Info.License.Name)
	assert.Empty(t, doc.Info.License.URL)
}

// --- Missing License ---

func TestParseInfoLicense_Missing(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	// License is nil when not provided
	assert.Nil(t, doc.Info.License)
}

// --- Different License Types ---

func TestParseInfoLicense_Apache(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  license:
    name: "Apache 2.0"
    url: "https://www.apache.org/licenses/LICENSE-2.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.License)
	assert.Equal(t, "Apache 2.0", doc.Info.License.Name)
}

func TestParseInfoLicense_GPL(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  license:
    name: "GPL-3.0"
    url: "https://www.gnu.org/licenses/gpl-3.0.html"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.License)
	assert.Equal(t, "GPL-3.0", doc.Info.License.Name)
}

func TestParseInfoLicense_Proprietary(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  license:
    name: "Proprietary"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.License)
	assert.Equal(t, "Proprietary", doc.Info.License.Name)
}

// --- Extensions ---

func TestParseInfoLicense_Extensions(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  license:
    name: "MIT"
    x-spdx-id: "MIT"
    x-osi-approved: true
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.License)
	require.NotNil(t, doc.Info.License.Extensions)
	assert.Equal(t, "MIT", doc.Info.License.Extensions["x-spdx-id"])
	assert.Equal(t, true, doc.Info.License.Extensions["x-osi-approved"])
}

// --- Empty Name (should still parse) ---

func TestParseInfoLicense_EmptyName(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  license:
    name: ""
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.License)
	assert.Empty(t, doc.Info.License.Name)
}

// --- Special Characters in Name ---

func TestParseInfoLicense_SpecialCharacters(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
  license:
    name: "CC-BY-NC-SA 4.0 (Creative Commons)"
    url: "https://creativecommons.org/licenses/by-nc-sa/4.0/"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Info)
	require.NotNil(t, doc.Info.License)
	assert.Equal(t, "CC-BY-NC-SA 4.0 (Creative Commons)", doc.Info.License.Name)
}
