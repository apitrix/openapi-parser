package openapi30x

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// slowSchemaServer starts an httptest server that serves a YAML schema
// after the given delay. Returns the server (caller must defer Close).
func slowSchemaServer(delay time.Duration, yamlBody string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.Header().Set("Content-Type", "application/yaml")
		fmt.Fprintln(w, yamlBody)
	}))
}

// TestResolve_BackgroundParseReturnsImmediately verifies that parseAndResolve
// returns well before the remote $ref has resolved.
func TestResolve_BackgroundParseReturnsImmediately(t *testing.T) {
	// Arrange
	const delay = 500 * time.Millisecond
	srv := slowSchemaServer(delay, `type: object
properties:
  id:
    type: integer`)
	defer srv.Close()

	spec := fmt.Sprintf(`openapi: "3.0.3"
info:
  title: Test
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      $ref: '%s/pet.yaml'`, srv.URL)

	// Act
	start := time.Now()
	result, err := parseAndResolve([]byte(spec), "/tmp", All())
	elapsed := time.Since(start)

	// Assert
	require.NoError(t, err)
	assert.Less(t, elapsed, delay, "parseAndResolve() should return before resolution finishes")

	// Clean up: wait for background goroutine so the httptest server isn't closed early
	result.Wait()
}

// TestResolve_PropertyAccessBlocksUntilResolved verifies that accessing a
// property through a resolved ref blocks until background resolution completes.
//
// This is the core background-resolution contract: calling ref.Value().Type()
// on an as-yet-unresolved ref must block until the resolver populates it.
func TestResolve_PropertyAccessBlocksUntilResolved(t *testing.T) {
	// Arrange
	const delay = 500 * time.Millisecond
	srv := slowSchemaServer(delay, `type: object
properties:
  name:
    type: string`)
	defer srv.Close()

	spec := fmt.Sprintf(`openapi: "3.0.3"
info:
  title: Test
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      $ref: '%s/pet.yaml'`, srv.URL)

	result, err := parseAndResolve([]byte(spec), "/tmp", All())
	require.NoError(t, err)

	ref := result.Document.Components().Schemas()["Pet"]
	require.NotNil(t, ref, "Pet schema ref should exist in parsed document")

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
// returns, accessing properties through refs is instant (non-blocking).
func TestResolve_WaitThenPropertyAccessIsInstant(t *testing.T) {
	// Arrange
	const delay = 500 * time.Millisecond
	srv := slowSchemaServer(delay, `type: string
description: a name`)
	defer srv.Close()

	spec := fmt.Sprintf(`openapi: "3.0.3"
info:
  title: Test
  version: "1.0"
paths: {}
components:
  schemas:
    Name:
      $ref: '%s/name.yaml'`, srv.URL)

	result, err := parseAndResolve([]byte(spec), "/tmp", All())
	require.NoError(t, err)
	require.NoError(t, result.Wait())

	ref := result.Document.Components().Schemas()["Name"]
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
