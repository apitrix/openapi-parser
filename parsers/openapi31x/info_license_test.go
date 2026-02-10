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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info())
	require.NotNil(t, result.Document.Info().License())
	assert.Equal(t, "MIT", result.Document.Info().License().Name())
	assert.Equal(t, "https://opensource.org/licenses/MIT", result.Document.Info().License().URL())
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info())
	require.NotNil(t, result.Document.Info().License())
	assert.Equal(t, "Apache 2.0", result.Document.Info().License().Name())
	assert.Empty(t, result.Document.Info().License().URL())
}

// --- Missing License ---

func TestParseInfoLicense_Missing(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info())
	// License is nil when not provided
	assert.Nil(t, result.Document.Info().License())
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info())
	require.NotNil(t, result.Document.Info().License())
	assert.Equal(t, "Apache 2.0", result.Document.Info().License().Name())
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info())
	require.NotNil(t, result.Document.Info().License())
	assert.Equal(t, "GPL-3.0", result.Document.Info().License().Name())
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info())
	require.NotNil(t, result.Document.Info().License())
	assert.Equal(t, "Proprietary", result.Document.Info().License().Name())
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info())
	require.NotNil(t, result.Document.Info().License())
	require.NotNil(t, result.Document.Info().License().VendorExtensions)
	assert.Equal(t, "MIT", result.Document.Info().License().VendorExtensions["x-spdx-id"])
	assert.Equal(t, true, result.Document.Info().License().VendorExtensions["x-osi-approved"])
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info())
	require.NotNil(t, result.Document.Info().License())
	assert.Empty(t, result.Document.Info().License().Name())
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
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Info())
	require.NotNil(t, result.Document.Info().License())
	assert.Equal(t, "CC-BY-NC-SA 4.0 (Creative Commons)", result.Document.Info().License().Name())
}
