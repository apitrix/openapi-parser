package openapi30x

import (
	"fmt"

	openapi30models "openapi-parser/models/openapi30"
	"openapi-parser/parsers/internal/shared"

	"gopkg.in/yaml.v3"
)

// Resolve resolves all $ref references in a parsed OpenAPI 3.0 document.
// basePath is the directory containing the root document (for relative file refs).
// root is the yaml.Node tree (for local JSON pointer refs).
//
// This function walks the entire document tree and populates the Value field
// on all *Ref types whose Ref field is non-empty and Value is nil.
// Circular references are detected and marked with Circular=true.
func Resolve(doc *openapi30models.OpenAPI, root *yaml.Node, basePath string) error {
	r := shared.NewRefResolver(basePath, root)
	resolving := make(map[string]bool)
	return resolveDocument(doc, r, resolving)
}

func resolveDocument(doc *openapi30models.OpenAPI, r *shared.RefResolver, resolving map[string]bool) error {
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

	if doc.Components != nil {
		if err := resolveComponents(doc.Components, r, resolving); err != nil {
			return err
		}
	}

	return nil
}

func resolveComponents(c *openapi30models.Components, r *shared.RefResolver, resolving map[string]bool) error {
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
	return nil
}

func resolvePathItem(pi *openapi30models.PathItem, r *shared.RefResolver, resolving map[string]bool) error {
	if pi == nil {
		return nil
	}

	for _, op := range []*openapi30models.Operation{
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

func resolveOperation(op *openapi30models.Operation, r *shared.RefResolver, resolving map[string]bool) error {
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

func resolveSchemaRef(ref *openapi30models.SchemaRef, r *shared.RefResolver, resolving map[string]bool) error {
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
		ctx := newParseContext(result.Node)
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

func resolveSchema(schema *openapi30models.Schema, r *shared.RefResolver, resolving map[string]bool) error {
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
	return resolveSchemaRef(schema.AdditionalProperties, r, resolving)
}

func resolveResponseRef(ref *openapi30models.ResponseRef, r *shared.RefResolver, resolving map[string]bool) error {
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
		ctx := newParseContext(result.Node)
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

func resolveParameterRef(ref *openapi30models.ParameterRef, r *shared.RefResolver, resolving map[string]bool) error {
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
		ctx := newParseContext(result.Node)
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

func resolveExampleRef(ref *openapi30models.ExampleRef, r *shared.RefResolver, resolving map[string]bool) error {
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
		ctx := newParseContext(result.Node)
		val, err := parseSharedExample(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved example ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	return nil
}

func resolveRequestBodyRef(ref *openapi30models.RequestBodyRef, r *shared.RefResolver, resolving map[string]bool) error {
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
		ctx := newParseContext(result.Node)
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

func resolveHeaderRef(ref *openapi30models.HeaderRef, r *shared.RefResolver, resolving map[string]bool) error {
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
		ctx := newParseContext(result.Node)
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

func resolveSecuritySchemeRef(ref *openapi30models.SecuritySchemeRef, r *shared.RefResolver, resolving map[string]bool) error {
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
		ctx := newParseContext(result.Node)
		val, err := parseSharedSecurityScheme(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved securityScheme ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	return nil
}

func resolveLinkRef(ref *openapi30models.LinkRef, r *shared.RefResolver, resolving map[string]bool) error {
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
		ctx := newParseContext(result.Node)
		val, err := parseSharedLink(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved link ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	return nil
}

func resolveCallbackRef(ref *openapi30models.CallbackRef, r *shared.RefResolver, resolving map[string]bool) error {
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
		ctx := newParseContext(result.Node)
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

func resolveMediaType(mt *openapi30models.MediaType, r *shared.RefResolver, resolving map[string]bool) error {
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
