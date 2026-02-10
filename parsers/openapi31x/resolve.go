package openapi31x

import (
	"fmt"

	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/parsers/shared"

	"gopkg.in/yaml.v3"
)

// Resolve resolves all $ref references in a parsed OpenAPI 3.1 document.
// basePath is the directory containing the root document (for relative file refs).
// root is the yaml.Node tree (for local JSON pointer refs).
func Resolve(doc *openapi31models.OpenAPI, root *yaml.Node, basePath string) error {
	r := shared.NewRefResolver(basePath, root)
	resolving := make(map[string]bool)
	return resolveDocument(doc, r, resolving)
}

func resolveDocument(doc *openapi31models.OpenAPI, r *shared.RefResolver, resolving map[string]bool) error {
	if doc == nil {
		return nil
	}

	if doc.Paths != nil {
		for _, pathItem := range doc.Paths.Items {
			if err := resolvePathItem(pathItem, r, resolving); err != nil {
				return err
			}
		}
	}

	// Resolve webhooks (new in 3.1)
	for _, ref := range doc.Webhooks {
		if err := resolvePathItemRef(ref, r, resolving); err != nil {
			return err
		}
	}

	if doc.Components != nil {
		if err := resolveComponents(doc.Components, r, resolving); err != nil {
			return err
		}
	}

	return nil
}

func resolveComponents(c *openapi31models.Components, r *shared.RefResolver, resolving map[string]bool) error {
	for name, ref := range c.Schemas {
		canonicalRef := "#/components/schemas/" + name
		resolving[canonicalRef] = true
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.Responses {
		canonicalRef := "#/components/responses/" + name
		resolving[canonicalRef] = true
		if err := resolveResponseRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.Parameters {
		canonicalRef := "#/components/parameters/" + name
		resolving[canonicalRef] = true
		if err := resolveParameterRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.Examples {
		canonicalRef := "#/components/examples/" + name
		resolving[canonicalRef] = true
		if err := resolveExampleRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.RequestBodies {
		canonicalRef := "#/components/requestBodies/" + name
		resolving[canonicalRef] = true
		if err := resolveRequestBodyRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.Headers {
		canonicalRef := "#/components/headers/" + name
		resolving[canonicalRef] = true
		if err := resolveHeaderRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.SecuritySchemes {
		canonicalRef := "#/components/securitySchemes/" + name
		resolving[canonicalRef] = true
		if err := resolveSecuritySchemeRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.Links {
		canonicalRef := "#/components/links/" + name
		resolving[canonicalRef] = true
		if err := resolveLinkRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.Callbacks {
		canonicalRef := "#/components/callbacks/" + name
		resolving[canonicalRef] = true
		if err := resolveCallbackRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	// PathItems in components (new in 3.1)
	for name, ref := range c.PathItems {
		canonicalRef := "#/components/pathItems/" + name
		resolving[canonicalRef] = true
		if err := resolvePathItemRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	return nil
}

func resolvePathItem(pi *openapi31models.PathItem, r *shared.RefResolver, resolving map[string]bool) error {
	if pi == nil {
		return nil
	}

	for _, op := range []*openapi31models.Operation{
		pi.Get, pi.Put, pi.Post, pi.Delete,
		pi.Options, pi.Head, pi.Patch, pi.Trace,
	} {
		if err := resolveOperation(op, r, resolving); err != nil {
			return err
		}
	}

	for _, ref := range pi.Parameters {
		if err := resolveParameterRef(ref, r, resolving); err != nil {
			return err
		}
	}

	return nil
}

func resolvePathItemRef(ref *openapi31models.PathItemRef, r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.Circular {
		return nil
	}

	if ref.Ref != "" && ref.Value == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			return fmt.Errorf("resolving pathItem ref %q: %w", ref.Ref, err)
		}
		if result.Circular {
			ref.Circular = true
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseOpenAPIPathsPathItem(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved pathItem ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	if ref.Value != nil {
		return resolvePathItem(ref.Value, r, resolving)
	}

	return nil
}

func resolveOperation(op *openapi31models.Operation, r *shared.RefResolver, resolving map[string]bool) error {
	if op == nil {
		return nil
	}

	for _, ref := range op.Parameters {
		if err := resolveParameterRef(ref, r, resolving); err != nil {
			return err
		}
	}

	if err := resolveRequestBodyRef(op.RequestBody, r, resolving); err != nil {
		return err
	}

	if op.Responses != nil {
		if err := resolveResponseRef(op.Responses.Default, r, resolving); err != nil {
			return err
		}
		for _, ref := range op.Responses.Codes {
			if err := resolveResponseRef(ref, r, resolving); err != nil {
				return err
			}
		}
	}

	for _, ref := range op.Callbacks {
		if err := resolveCallbackRef(ref, r, resolving); err != nil {
			return err
		}
	}

	return nil
}

// =============================================================================
// Individual ref type resolvers
// =============================================================================

func resolveSchemaRef(ref *openapi31models.SchemaRef, r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.Circular {
		return nil
	}

	if ref.Ref != "" && ref.Value == nil {
		// Model-level cycle detection
		if resolving[ref.Ref] {
			ref.Circular = true
			return nil
		}
		resolving[ref.Ref] = true
		defer func() { delete(resolving, ref.Ref) }()

		result, err := r.Resolve(ref.Ref)
		if err != nil {
			return fmt.Errorf("resolving schema ref %q: %w", ref.Ref, err)
		}
		if result.Circular {
			ref.Circular = true
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		schema, err := parseSharedSchema(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved schema ref %q: %w", ref.Ref, err)
		}
		ref.Value = schema
	}

	if ref.Value != nil {
		return resolveSchema(ref.Value, r, resolving)
	}

	return nil
}

func resolveSchema(schema *openapi31models.Schema, r *shared.RefResolver, resolving map[string]bool) error {
	if schema == nil {
		return nil
	}

	for _, ref := range schema.AllOf {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
	}
	for _, ref := range schema.OneOf {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
	}
	for _, ref := range schema.AnyOf {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
	}
	if err := resolveSchemaRef(schema.Not, r, resolving); err != nil {
		return err
	}
	if err := resolveSchemaRef(schema.Items, r, resolving); err != nil {
		return err
	}
	for _, ref := range schema.Properties {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
	}
	if err := resolveSchemaRef(schema.AdditionalProperties, r, resolving); err != nil {
		return err
	}

	// OpenAPI 3.1 / JSON Schema 2020-12 additions
	if err := resolveSchemaRef(schema.If, r, resolving); err != nil {
		return err
	}
	if err := resolveSchemaRef(schema.Then, r, resolving); err != nil {
		return err
	}
	if err := resolveSchemaRef(schema.Else, r, resolving); err != nil {
		return err
	}
	for _, ref := range schema.DependentSchemas {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
	}
	for _, ref := range schema.PrefixItems {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
	}
	if err := resolveSchemaRef(schema.ContentSchema, r, resolving); err != nil {
		return err
	}
	if err := resolveSchemaRef(schema.UnevaluatedItems, r, resolving); err != nil {
		return err
	}
	return resolveSchemaRef(schema.UnevaluatedProperties, r, resolving)
}

func resolveResponseRef(ref *openapi31models.ResponseRef, r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.Circular {
		return nil
	}

	if ref.Ref != "" && ref.Value == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			return fmt.Errorf("resolving response ref %q: %w", ref.Ref, err)
		}
		if result.Circular {
			ref.Circular = true
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedResponse(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved response ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	if ref.Value != nil {
		for _, hRef := range ref.Value.Headers {
			if err := resolveHeaderRef(hRef, r, resolving); err != nil {
				return err
			}
		}
		for _, mt := range ref.Value.Content {
			if err := resolveMediaType(mt, r, resolving); err != nil {
				return err
			}
		}
		for _, lRef := range ref.Value.Links {
			if err := resolveLinkRef(lRef, r, resolving); err != nil {
				return err
			}
		}
	}

	return nil
}

func resolveParameterRef(ref *openapi31models.ParameterRef, r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.Circular {
		return nil
	}

	if ref.Ref != "" && ref.Value == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			return fmt.Errorf("resolving parameter ref %q: %w", ref.Ref, err)
		}
		if result.Circular {
			ref.Circular = true
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedParameter(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved parameter ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	if ref.Value != nil {
		if err := resolveSchemaRef(ref.Value.Schema, r, resolving); err != nil {
			return err
		}
		for _, eRef := range ref.Value.Examples {
			if err := resolveExampleRef(eRef, r, resolving); err != nil {
				return err
			}
		}
	}

	return nil
}

func resolveExampleRef(ref *openapi31models.ExampleRef, r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.Circular {
		return nil
	}

	if ref.Ref != "" && ref.Value == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			return fmt.Errorf("resolving example ref %q: %w", ref.Ref, err)
		}
		if result.Circular {
			ref.Circular = true
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedExample(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved example ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	return nil
}

func resolveRequestBodyRef(ref *openapi31models.RequestBodyRef, r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.Circular {
		return nil
	}

	if ref.Ref != "" && ref.Value == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			return fmt.Errorf("resolving requestBody ref %q: %w", ref.Ref, err)
		}
		if result.Circular {
			ref.Circular = true
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedRequestBody(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved requestBody ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	if ref.Value != nil {
		for _, mt := range ref.Value.Content {
			if err := resolveMediaType(mt, r, resolving); err != nil {
				return err
			}
		}
	}

	return nil
}

func resolveHeaderRef(ref *openapi31models.HeaderRef, r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.Circular {
		return nil
	}

	if ref.Ref != "" && ref.Value == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			return fmt.Errorf("resolving header ref %q: %w", ref.Ref, err)
		}
		if result.Circular {
			ref.Circular = true
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedHeader(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved header ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	if ref.Value != nil {
		if err := resolveSchemaRef(ref.Value.Schema, r, resolving); err != nil {
			return err
		}
		for _, eRef := range ref.Value.Examples {
			if err := resolveExampleRef(eRef, r, resolving); err != nil {
				return err
			}
		}
	}

	return nil
}

func resolveSecuritySchemeRef(ref *openapi31models.SecuritySchemeRef, r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.Circular {
		return nil
	}

	if ref.Ref != "" && ref.Value == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			return fmt.Errorf("resolving securityScheme ref %q: %w", ref.Ref, err)
		}
		if result.Circular {
			ref.Circular = true
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedSecurityScheme(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved securityScheme ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	return nil
}

func resolveLinkRef(ref *openapi31models.LinkRef, r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.Circular {
		return nil
	}

	if ref.Ref != "" && ref.Value == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			return fmt.Errorf("resolving link ref %q: %w", ref.Ref, err)
		}
		if result.Circular {
			ref.Circular = true
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedLink(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved link ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	return nil
}

func resolveCallbackRef(ref *openapi31models.CallbackRef, r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.Circular {
		return nil
	}

	if ref.Ref != "" && ref.Value == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			return fmt.Errorf("resolving callback ref %q: %w", ref.Ref, err)
		}
		if result.Circular {
			ref.Circular = true
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedCallback(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved callback ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	if ref.Value != nil {
		for _, pathItem := range ref.Value.Paths {
			if err := resolvePathItem(pathItem, r, resolving); err != nil {
				return err
			}
		}
	}

	return nil
}

func resolveMediaType(mt *openapi31models.MediaType, r *shared.RefResolver, resolving map[string]bool) error {
	if mt == nil {
		return nil
	}

	if err := resolveSchemaRef(mt.Schema, r, resolving); err != nil {
		return err
	}

	for _, ref := range mt.Examples {
		if err := resolveExampleRef(ref, r, resolving); err != nil {
			return err
		}
	}

	for _, enc := range mt.Encoding {
		if enc != nil {
			for _, hRef := range enc.Headers {
				if err := resolveHeaderRef(hRef, r, resolving); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
