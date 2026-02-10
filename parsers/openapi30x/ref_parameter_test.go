package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for ref_parameter.go - parameter reference parsing
// =============================================================================

// --- Basic Reference ---

func TestParseRefParameter_Basic(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - $ref: '#/components/parameters/LimitParam'
      responses:
        "200":
          description: "OK"
components:
  parameters:
    LimitParam:
      name: limit
      in: query
      schema:
        type: integer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := result.Document.Paths().Items()["/pets"].Get().Parameters()[0]
	assert.Equal(t, "#/components/parameters/LimitParam", ref.Ref)
}

// --- Path-Level Parameter ---

func TestParseRefParameter_PathLevel(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets/{petId}:
    parameters:
      - $ref: '#/components/parameters/PetIdParam'
    get:
      responses:
        "200":
          description: "OK"
components:
  parameters:
    PetIdParam:
      name: petId
      in: path
      required: true
      schema:
        type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := result.Document.Paths().Items()["/pets/{petId}"].Parameters()[0]
	assert.Equal(t, "#/components/parameters/PetIdParam", ref.Ref)
}

// --- Multiple References ---

func TestParseRefParameter_Multiple(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - $ref: '#/components/parameters/LimitParam'
        - $ref: '#/components/parameters/OffsetParam'
        - $ref: '#/components/parameters/SortParam'
      responses:
        "200":
          description: "OK"
components:
  parameters:
    LimitParam:
      name: limit
      in: query
      schema:
        type: integer
    OffsetParam:
      name: offset
      in: query
      schema:
        type: integer
    SortParam:
      name: sort
      in: query
      schema:
        type: string
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	params := result.Document.Paths().Items()["/pets"].Get().Parameters()
	assert.Len(t, params, 3)
	assert.Equal(t, "#/components/parameters/LimitParam", params[0].Ref)
	assert.Equal(t, "#/components/parameters/OffsetParam", params[1].Ref)
	assert.Equal(t, "#/components/parameters/SortParam", params[2].Ref)
}

// --- Mixed Inline and Reference ---

func TestParseRefParameter_Mixed(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      parameters:
        - $ref: '#/components/parameters/LimitParam'
        - name: filter
          in: query
          schema:
            type: string
      responses:
        "200":
          description: "OK"
components:
  parameters:
    LimitParam:
      name: limit
      in: query
      schema:
        type: integer
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	params := result.Document.Paths().Items()["/pets"].Get().Parameters()
	assert.Len(t, params, 2)
	assert.Equal(t, "#/components/parameters/LimitParam", params[0].Ref)
	assert.Equal(t, "filter", params[1].Value.Name())
}
