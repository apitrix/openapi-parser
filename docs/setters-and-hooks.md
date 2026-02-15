# Setters and Hooks

This document describes the mutable API for OpenAPI model types: `SetXxx()` methods and the hook infrastructure in `Trix` for attaching validators, loggers, and other middleware.

## Overview

All model types across OpenAPI 2.0, 3.0, and 3.1 expose:

- **Getters** — read-only access to fields (e.g., `Contact.Name()`)
- **Setters** — mutation with optional validation (e.g., `Contact.SetName(name)`)

Setters run a hook pipeline before applying changes. Hooks can validate, log, transform, or reject updates.

## Basic Usage

### Setting a field (no hooks)

```go
contact := openapi20.NewContact("Alice", "https://example.com", "alice@example.com")
err := contact.SetName("Bob")
if err != nil {
    // handle error
}
// contact.Name() == "Bob"
```

### Attaching a validator

```go
contact := openapi20.NewContact("Alice", "https://example.com", "alice@example.com")

contact.Trix.OnSet("name", func(field string, oldVal, newVal interface{}) error {
    name := newVal.(string)
    if len(name) > 100 {
        return fmt.Errorf("name too long: %d chars", len(name))
    }
    return nil
})

err := contact.SetName("Bob")       // succeeds
err = contact.SetName(strings.Repeat("x", 101))  // returns error, name unchanged
```

### Attaching a logger

```go
op.Trix.OnSet("tags", func(field string, oldVal, newVal interface{}) error {
    log.Printf("tags changed: %v -> %v", oldVal, newVal)
    return nil
})
op.SetTags([]string{"pets", "store"})
```

## Hook Infrastructure

Hooks live in `Trix`, the per-node library metadata container embedded in every model via `ElementBase`.

| Method | Description |
|--------|-------------|
| `OnSet(field string, fn HookFunc)` | Registers a hook that runs before the field is set |
| `RunHooks(field, oldVal, newVal)` | Called by setters; runs all registered hooks for the field |

`HookFunc` signature:

```go
type HookFunc func(field string, oldVal, newVal interface{}) error
```

- Return `nil` to allow the change
- Return a non-nil `error` to reject the change (the setter will not update the field)

### Field names

Use the JSON/YAML property name for the `field` argument (e.g., `"operationId"`, `"externalDocs"`, `"$ref"`).

### Zero cost when unused

- `hooks` is lazy-initialized — no allocation until the first `OnSet` call
- `RunHooks` is a no-op when no hooks are registered — no overhead for models that don't use hooks

## Naming Convention

Setters follow the pattern `Set` + getter name:

| Getter | Setter |
|--------|--------|
| `Name()` | `SetName(name string)` |
| `OperationID()` | `SetOperationID(operationID string)` |
| `ExternalDocs()` | `SetExternalDocs(externalDocs *ExternalDocs)` |

## Scope

Setters are available on all model types in:

- `github.com/apitrix/openapi-parser/models/openapi20` (Swagger 2.0)
- `github.com/apitrix/openapi-parser/models/openapi30` (OpenAPI 3.0.x)
- `github.com/apitrix/openapi-parser/models/openapi31` (OpenAPI 3.1.x)

Each unexported field with a getter has a corresponding setter that runs hooks before applying the change.
