package openapi20

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// slowSchemaServer returns an httptest.Server that waits `delay` before
// responding with the given YAML body.
func slowSchemaServer(delay time.Duration, yamlBody string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.Header().Set("Content-Type", "application/x-yaml")
		w.Write([]byte(yamlBody))
	}))
}

// TestResolve_BackgroundParseReturnsImmediately verifies that parseAndResolve
// returns well before the simulated remote response arrives.
func TestResolve_BackgroundParseReturnsImmediately(t *testing.T) {
	// Arrange
	const delay = 500 * time.Millisecond
	srv := slowSchemaServer(delay, `type: object
properties:
  name:
    type: string`)
	defer srv.Close()

	spec := fmt.Sprintf(`swagger: "2.0"
info:
  title: Test
  version: "1.0"
paths: {}
definitions:
  Pet:
    $ref: '%s/pet.yaml'`, srv.URL)

	// Act
	start := time.Now()
	result, err := parseAndResolve([]byte(spec), "/tmp", All())
	elapsed := time.Since(start)

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Less(t, elapsed, delay/2, "parseAndResolve should return before remote response arrives")
}

// TestResolve_PropertyAccessBlocksUntilResolved verifies that accessing a
// property through a ref blocks until background resolution completes.
func TestResolve_PropertyAccessBlocksUntilResolved(t *testing.T) {
	// Arrange
	const delay = 500 * time.Millisecond
	srv := slowSchemaServer(delay, `type: object
properties:
  name:
    type: string`)
	defer srv.Close()

	spec := fmt.Sprintf(`swagger: "2.0"
info:
  title: Test
  version: "1.0"
paths: {}
definitions:
  Pet:
    $ref: '%s/pet.yaml'`, srv.URL)

	result, err := parseAndResolve([]byte(spec), "/tmp", All())
	require.NoError(t, err)

	ref := result.Document.Definitions()["Pet"]
	require.NotNil(t, ref, "Pet definition ref should exist in parsed document")

	// Act — access a property through the ref WITHOUT calling Wait() first.
	// This must block until the background resolver completes (~500ms).
	start := time.Now()
	val := ref.Value()
	elapsed := time.Since(start)

	// Assert
	require.NotNil(t, val, "ref.Value() should be resolved (check ref.ResolveErr(): %v)", ref.ResolveErr())
	assert.GreaterOrEqual(t, elapsed, delay/2, "ref.Value() should block until resolved")
	assert.Equal(t, "object", val.Type())
	assert.NotNil(t, val.Properties()["name"], "resolved schema should have 'name' property")
}

// TestResolve_WaitThenPropertyAccessIsInstant verifies that after Wait()
// returns, property access is instant (non-blocking).
func TestResolve_WaitThenPropertyAccessIsInstant(t *testing.T) {
	// Arrange
	const delay = 500 * time.Millisecond
	srv := slowSchemaServer(delay, `type: string
description: a name`)
	defer srv.Close()

	spec := fmt.Sprintf(`swagger: "2.0"
info:
  title: Test
  version: "1.0"
paths: {}
definitions:
  Name:
    $ref: '%s/name.yaml'`, srv.URL)

	result, err := parseAndResolve([]byte(spec), "/tmp", All())
	require.NoError(t, err)
	require.NoError(t, result.Wait())

	ref := result.Document.Definitions()["Name"]
	require.NotNil(t, ref)

	// Act — access property after Wait() has returned
	start := time.Now()
	val := ref.Value()
	elapsed := time.Since(start)

	// Assert
	require.NotNil(t, val)
	assert.Less(t, elapsed, 10*time.Millisecond, "ref.Value() should be instant after Wait()")
	assert.Equal(t, "string", val.Type())
}
