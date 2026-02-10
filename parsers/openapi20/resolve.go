package openapi20

import (
	"fmt"

	openapi20models "openapi-parser/models/openapi20"
	"openapi-parser/parsers/shared"

	"gopkg.in/yaml.v3"
)

// Resolve resolves all $ref references in a parsed Swagger 2.0 document.
// basePath is the directory containing the root document (for relative file refs).
// root is the yaml.Node tree (for local JSON pointer refs).
func Resolve(doc *openapi20models.Swagger, root *yaml.Node, basePath string) error {
	r := shared.NewRefResolver(basePath, root)
	resolving := make(map[string]bool)
	return resolveDocument(doc, r, resolving)
}

func resolveDocument(doc *openapi20models.Swagger, r *shared.RefResolver, resolving map[string]bool) error {
	if doc == nil {
		return nil
	}

	// Resolve top-level definitions — pre-register each definition's canonical
	// ref path so self-referencing schemas are immediately detected as circular.
	for name, ref := range doc.Definitions {
		canonicalRef := "#/definitions/" + name
		resolving[canonicalRef] = true
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range doc.Parameters {
		canonicalRef := "#/parameters/" + name
		resolving[canonicalRef] = true
		if err := resolveParameterRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range doc.Responses {
		canonicalRef := "#/responses/" + name
		resolving[canonicalRef] = true
		if err := resolveResponseRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}

	// Resolve paths
	if doc.Paths != nil {
		for _, pathItem := range doc.Paths.Items {
			if err := resolvePathItem(pathItem, r, resolving); err != nil {
				return err
			}
		}
	}

	return nil
}

func resolvePathItem(pi *openapi20models.PathItem, r *shared.RefResolver, resolving map[string]bool) error {
	if pi == nil {
		return nil
	}

	for _, op := range []*openapi20models.Operation{
		pi.Get, pi.Put, pi.Post, pi.Delete,
		pi.Options, pi.Head, pi.Patch,
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

func resolveOperation(op *openapi20models.Operation, r *shared.RefResolver, resolving map[string]bool) error {
	if op == nil {
		return nil
	}

	for _, ref := range op.Parameters {
		if err := resolveParameterRef(ref, r, resolving); err != nil {
			return err
		}
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

	return nil
}

// =============================================================================
// Individual ref type resolvers
// =============================================================================

func resolveSchemaRef(ref *openapi20models.SchemaRef, r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.Circular {
		return nil
	}

	if ref.Ref != "" && ref.Value == nil {
		// Check model-level cycle: if we're already resolving this $ref string,
		// it means we hit a circular reference through the parsed model tree.
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
		schema, err := parseSchema(result.Node, ctx)
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

func resolveSchema(schema *openapi20models.Schema, r *shared.RefResolver, resolving map[string]bool) error {
	if schema == nil {
		return nil
	}

	for _, ref := range schema.AllOf {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
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

func resolveResponseRef(ref *openapi20models.ResponseRef, r *shared.RefResolver, resolving map[string]bool) error {
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
		val, err := parseResponse(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved response ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	if ref.Value != nil {
		return resolveSchemaRef(ref.Value.Schema, r, resolving)
	}

	return nil
}

func resolveParameterRef(ref *openapi20models.ParameterRef, r *shared.RefResolver, resolving map[string]bool) error {
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
		val, err := parseParameter(result.Node, ctx)
		if err != nil {
			return fmt.Errorf("parsing resolved parameter ref %q: %w", ref.Ref, err)
		}
		ref.Value = val
	}

	if ref.Value != nil {
		return resolveSchemaRef(ref.Value.Schema, r, resolving)
	}

	return nil
}
