package openapi31x

import (
	"fmt"
	modelshared "openapi-parser/models/shared"

	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/parsers/shared"

	"gopkg.in/yaml.v3"
)

// Resolve resolves all $ref references in a parsed OpenAPI 3.1 document.
// basePath is the directory containing the root document (for relative file refs).
// root is the yaml.Node tree (for local JSON pointer refs).
func Resolve(doc *openapi31models.OpenAPI, root *yaml.Node, basePath string) error {
	r := shared.NewRefResolver(basePath, root)
	r.BuildAnchorIndex("", root)
	r.BuildDynamicAnchorIndex(root)
	resolving := make(map[string]bool)
	if err := resolveDocument(doc, r, resolving); err != nil {
		return err
	}
	// Post-resolve: wire up operationRefs now that all operations are parsed
	resolveOperationRefs(doc)
	return nil
}

func resolveDocument(doc *openapi31models.OpenAPI, r *shared.RefResolver, resolving map[string]bool) error {
	if doc == nil {
		return nil
	}

	if doc.Paths() != nil {
		for _, pathItem := range doc.Paths().Items() {
			if err := resolvePathItem(pathItem, r, resolving); err != nil {
				return err
			}
		}
	}

	// Resolve webhooks (new in 3.1)
	for _, ref := range doc.Webhooks() {
		if err := resolvePathItemRef(ref, r, resolving); err != nil {
			return err
		}
	}

	if doc.Components() != nil {
		if err := resolveComponents(doc.Components(), r, resolving); err != nil {
			return err
		}
	}

	return nil
}

func resolveComponents(c *openapi31models.Components, r *shared.RefResolver, resolving map[string]bool) error {
	for name, ref := range c.Schemas() {
		canonicalRef := "#/components/schemas/" + name
		resolving[canonicalRef] = true
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.Responses() {
		canonicalRef := "#/components/responses/" + name
		resolving[canonicalRef] = true
		if err := resolveResponseRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.Parameters() {
		canonicalRef := "#/components/parameters/" + name
		resolving[canonicalRef] = true
		if err := resolveParameterRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.Examples() {
		canonicalRef := "#/components/examples/" + name
		resolving[canonicalRef] = true
		if err := resolveExampleRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.RequestBodies() {
		canonicalRef := "#/components/requestBodies/" + name
		resolving[canonicalRef] = true
		if err := resolveRequestBodyRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.Headers() {
		canonicalRef := "#/components/headers/" + name
		resolving[canonicalRef] = true
		if err := resolveHeaderRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.SecuritySchemes() {
		canonicalRef := "#/components/securitySchemes/" + name
		resolving[canonicalRef] = true
		if err := resolveSecuritySchemeRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.Links() {
		canonicalRef := "#/components/links/" + name
		resolving[canonicalRef] = true
		if err := resolveLinkRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	for name, ref := range c.Callbacks() {
		canonicalRef := "#/components/callbacks/" + name
		resolving[canonicalRef] = true
		if err := resolveCallbackRef(ref, r, resolving); err != nil {
			return err
		}
		delete(resolving, canonicalRef)
	}
	// PathItems in components (new in 3.1)
	for name, ref := range c.PathItems() {
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
		pi.Get(), pi.Put(), pi.Post(), pi.Delete(),
		pi.Options(), pi.Head(), pi.Patch(), pi.Trace(),
	} {
		if err := resolveOperation(op, r, resolving); err != nil {
			return err
		}
	}

	for _, ref := range pi.Parameters() {
		if err := resolveParameterRef(ref, r, resolving); err != nil {
			return err
		}
	}

	return nil
}

func resolvePathItemRef(ref *modelshared.RefWithMeta[openapi31models.PathItem], r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.RawCircular() {
		return nil
	}

	if ref.Ref != "" && ref.RawValue() == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("resolving pathItem ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		if result.Circular {
			ref.SetCircular(true)
			ref.MarkDone()
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseOpenAPIPathsPathItem(result.Node, ctx)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("parsing resolved pathItem ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		ref.SetValue(val)
		ref.MarkDone()
	}

	if ref.RawValue() != nil {
		return resolvePathItem(ref.RawValue(), r, resolving)
	}

	return nil
}

func resolveOperation(op *openapi31models.Operation, r *shared.RefResolver, resolving map[string]bool) error {
	if op == nil {
		return nil
	}

	for _, ref := range op.Parameters() {
		if err := resolveParameterRef(ref, r, resolving); err != nil {
			return err
		}
	}

	if err := resolveRequestBodyRef(op.RequestBody(), r, resolving); err != nil {
		return err
	}

	if op.Responses() != nil {
		if err := resolveResponseRef(op.Responses().Default(), r, resolving); err != nil {
			return err
		}
		for _, ref := range op.Responses().Codes() {
			if err := resolveResponseRef(ref, r, resolving); err != nil {
				return err
			}
		}
	}

	for _, ref := range op.Callbacks() {
		if err := resolveCallbackRef(ref, r, resolving); err != nil {
			return err
		}
	}

	return nil
}

// =============================================================================
// Individual ref type resolvers
// =============================================================================

func resolveSchemaRef(ref *modelshared.RefWithMeta[openapi31models.Schema], r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.RawCircular() {
		return nil
	}

	if ref.Ref != "" && ref.RawValue() == nil {
		// Model-level cycle detection
		if resolving[ref.Ref] {
			ref.SetCircular(true)
			ref.MarkDone()
			return nil
		}
		resolving[ref.Ref] = true
		defer func() { delete(resolving, ref.Ref) }()

		result, err := r.Resolve(ref.Ref)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("resolving schema ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		if result.Circular {
			ref.SetCircular(true)
			ref.MarkDone()
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		schema, err := parseSharedSchema(result.Node, ctx)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("parsing resolved schema ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		ref.SetValue(schema)
		ref.MarkDone()
	}

	if ref.RawValue() != nil {
		return resolveSchema(ref.RawValue(), r, resolving)
	}

	return nil
}

func resolveSchema(schema *openapi31models.Schema, r *shared.RefResolver, resolving map[string]bool) error {
	if schema == nil {
		return nil
	}

	for _, ref := range schema.AllOf() {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
	}
	for _, ref := range schema.OneOf() {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
	}
	for _, ref := range schema.AnyOf() {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
	}
	if err := resolveSchemaRef(schema.Not(), r, resolving); err != nil {
		return err
	}
	if err := resolveSchemaRef(schema.Items(), r, resolving); err != nil {
		return err
	}
	for _, ref := range schema.Properties() {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
	}
	if err := resolveSchemaRef(schema.AdditionalProperties(), r, resolving); err != nil {
		return err
	}

	// OpenAPI 3.1 / JSON Schema 2020-12 additions
	if err := resolveSchemaRef(schema.If(), r, resolving); err != nil {
		return err
	}
	if err := resolveSchemaRef(schema.Then(), r, resolving); err != nil {
		return err
	}
	if err := resolveSchemaRef(schema.Else(), r, resolving); err != nil {
		return err
	}
	for _, ref := range schema.DependentSchemas() {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
	}
	for _, ref := range schema.PrefixItems() {
		if err := resolveSchemaRef(ref, r, resolving); err != nil {
			return err
		}
	}
	if err := resolveSchemaRef(schema.ContentSchema(), r, resolving); err != nil {
		return err
	}
	if err := resolveSchemaRef(schema.UnevaluatedItems(), r, resolving); err != nil {
		return err
	}
	if err := resolveSchemaRef(schema.UnevaluatedProperties(), r, resolving); err != nil {
		return err
	}

	// Resolve $dynamicRef → look up matching $dynamicAnchor
	if schema.DynamicRef() != "" {
		result, err := r.ResolveDynamicRef(schema.DynamicRef())
		if err == nil && !result.Circular {
			ctx := newParseContext(result.Node, shared.All())
			resolved, parseErr := parseSharedSchema(result.Node, ctx)
			if parseErr == nil {
				ref := modelshared.NewRefWithMeta[openapi31models.Schema]("")
				ref.SetValue(resolved)
				schema.Trix.ResolvedDynamicRef = ref
			}
		}
	}

	// Resolve discriminator.mapping values
	if schema.Discriminator() != nil && len(schema.Discriminator().Mapping()) > 0 {
		resolved := make(map[string]*modelshared.RefWithMeta[openapi31models.Schema])
		for key, val := range schema.Discriminator().Mapping() {
			mapResult, mapErr := r.ResolveMapping(val)
			if mapErr == nil {
				ctx := newParseContext(mapResult.Node, shared.All())
				s, parseErr := parseSharedSchema(mapResult.Node, ctx)
				if parseErr == nil {
					ref := modelshared.NewRefWithMeta[openapi31models.Schema](val)
					ref.SetValue(s)
					resolved[key] = ref
				}
			}
		}
		if len(resolved) > 0 {
			schema.Discriminator().Trix.ResolvedMapping = resolved
		}
	}

	return nil
}

func resolveResponseRef(ref *modelshared.RefWithMeta[openapi31models.Response], r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.RawCircular() {
		return nil
	}

	if ref.Ref != "" && ref.RawValue() == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("resolving response ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		if result.Circular {
			ref.SetCircular(true)
			ref.MarkDone()
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedResponse(result.Node, ctx)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("parsing resolved response ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		ref.SetValue(val)
		ref.MarkDone()
	}

	if ref.RawValue() != nil {
		for _, hRef := range ref.RawValue().Headers() {
			if err := resolveHeaderRef(hRef, r, resolving); err != nil {
				return err
			}
		}
		for _, mt := range ref.RawValue().Content() {
			if err := resolveMediaType(mt, r, resolving); err != nil {
				return err
			}
		}
		for _, lRef := range ref.RawValue().Links() {
			if err := resolveLinkRef(lRef, r, resolving); err != nil {
				return err
			}
		}
	}

	return nil
}

func resolveParameterRef(ref *modelshared.RefWithMeta[openapi31models.Parameter], r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.RawCircular() {
		return nil
	}

	if ref.Ref != "" && ref.RawValue() == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("resolving parameter ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		if result.Circular {
			ref.SetCircular(true)
			ref.MarkDone()
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedParameter(result.Node, ctx)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("parsing resolved parameter ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		ref.SetValue(val)
		ref.MarkDone()
	}

	if ref.RawValue() != nil {
		if err := resolveSchemaRef(ref.RawValue().Schema(), r, resolving); err != nil {
			return err
		}
		for _, eRef := range ref.RawValue().Examples() {
			if err := resolveExampleRef(eRef, r, resolving); err != nil {
				return err
			}
		}
	}

	return nil
}

func resolveExampleRef(ref *modelshared.RefWithMeta[openapi31models.Example], r *shared.RefResolver, _ map[string]bool) error {
	if ref == nil || ref.RawCircular() {
		return nil
	}

	if ref.Ref != "" && ref.RawValue() == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("resolving example ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		if result.Circular {
			ref.SetCircular(true)
			ref.MarkDone()
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedExample(result.Node, ctx)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("parsing resolved example ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		ref.SetValue(val)
		ref.MarkDone()
	}

	return nil
}

func resolveRequestBodyRef(ref *modelshared.RefWithMeta[openapi31models.RequestBody], r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.RawCircular() {
		return nil
	}

	if ref.Ref != "" && ref.RawValue() == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("resolving requestBody ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		if result.Circular {
			ref.SetCircular(true)
			ref.MarkDone()
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedRequestBody(result.Node, ctx)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("parsing resolved requestBody ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		ref.SetValue(val)
		ref.MarkDone()
	}

	if ref.RawValue() != nil {
		for _, mt := range ref.RawValue().Content() {
			if err := resolveMediaType(mt, r, resolving); err != nil {
				return err
			}
		}
	}

	return nil
}

func resolveHeaderRef(ref *modelshared.RefWithMeta[openapi31models.Header], r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.RawCircular() {
		return nil
	}

	if ref.Ref != "" && ref.RawValue() == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("resolving header ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		if result.Circular {
			ref.SetCircular(true)
			ref.MarkDone()
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedHeader(result.Node, ctx)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("parsing resolved header ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		ref.SetValue(val)
		ref.MarkDone()
	}

	if ref.RawValue() != nil {
		if err := resolveSchemaRef(ref.RawValue().Schema(), r, resolving); err != nil {
			return err
		}
		for _, eRef := range ref.RawValue().Examples() {
			if err := resolveExampleRef(eRef, r, resolving); err != nil {
				return err
			}
		}
	}

	return nil
}

func resolveSecuritySchemeRef(ref *modelshared.RefWithMeta[openapi31models.SecurityScheme], r *shared.RefResolver, _ map[string]bool) error {
	if ref == nil || ref.RawCircular() {
		return nil
	}

	if ref.Ref != "" && ref.RawValue() == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("resolving securityScheme ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		if result.Circular {
			ref.SetCircular(true)
			ref.MarkDone()
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedSecurityScheme(result.Node, ctx)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("parsing resolved securityScheme ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		ref.SetValue(val)
		ref.MarkDone()
	}

	return nil
}

func resolveLinkRef(ref *modelshared.RefWithMeta[openapi31models.Link], r *shared.RefResolver, _ map[string]bool) error {
	if ref == nil || ref.RawCircular() {
		return nil
	}

	if ref.Ref != "" && ref.RawValue() == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("resolving link ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		if result.Circular {
			ref.SetCircular(true)
			ref.MarkDone()
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedLink(result.Node, ctx)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("parsing resolved link ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		ref.SetValue(val)
		ref.MarkDone()
	}

	return nil
}

func resolveCallbackRef(ref *modelshared.RefWithMeta[openapi31models.Callback], r *shared.RefResolver, resolving map[string]bool) error {
	if ref == nil || ref.RawCircular() {
		return nil
	}

	if ref.Ref != "" && ref.RawValue() == nil {
		result, err := r.Resolve(ref.Ref)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("resolving callback ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		if result.Circular {
			ref.SetCircular(true)
			ref.MarkDone()
			return nil
		}
		ctx := newParseContext(result.Node, shared.All())
		val, err := parseSharedCallback(result.Node, ctx)
		if err != nil {
			ref.SetResolveErr(fmt.Errorf("parsing resolved callback ref %q: %w", ref.Ref, err))
			ref.MarkDone()
			return nil
		}
		ref.SetValue(val)
		ref.MarkDone()
	}

	if ref.RawValue() != nil {
		for _, pathItem := range ref.RawValue().Paths() {
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

	if err := resolveSchemaRef(mt.Schema(), r, resolving); err != nil {
		return err
	}

	for _, ref := range mt.Examples() {
		if err := resolveExampleRef(ref, r, resolving); err != nil {
			return err
		}
	}

	for _, enc := range mt.Encoding() {
		if enc != nil {
			for _, hRef := range enc.Headers() {
				if err := resolveHeaderRef(hRef, r, resolving); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// =============================================================================
// operationRef resolution — post-resolve step
// =============================================================================

// resolveOperationRefs walks all Link objects in the document and resolves
// operationRef values to the already-parsed Operation in the document tree.
// This runs after all $ref resolution is complete.
func resolveOperationRefs(doc *openapi31models.OpenAPI) {
	if doc == nil {
		return
	}

	// Collect all operations by path+method
	ops := collectOperations(doc)

	// Walk all responses to find links with operationRef
	walkLinksForOperationRef(doc, ops)
}

// collectOperations builds a map of "path|method" → *Operation from the document.
func collectOperations(doc *openapi31models.OpenAPI) map[string]*openapi31models.Operation {
	ops := make(map[string]*openapi31models.Operation)
	if doc.Paths() == nil {
		return ops
	}

	for path, pi := range doc.Paths().Items() {
		for method, op := range pathItemOps(pi) {
			if op != nil {
				ops[path+"|"+method] = op
			}
		}
	}
	return ops
}

// pathItemOps returns a map of method → operation for the given path item.
func pathItemOps(pi *openapi31models.PathItem) map[string]*openapi31models.Operation {
	return map[string]*openapi31models.Operation{
		"get":     pi.Get(),
		"put":     pi.Put(),
		"post":    pi.Post(),
		"delete":  pi.Delete(),
		"options": pi.Options(),
		"head":    pi.Head(),
		"patch":   pi.Patch(),
		"trace":   pi.Trace(),
	}
}

// walkLinksForOperationRef iterates over all response links and resolves operationRef.
func walkLinksForOperationRef(doc *openapi31models.OpenAPI, ops map[string]*openapi31models.Operation) {
	resolveLinksInPaths := func(pi *openapi31models.PathItem) {
		if pi == nil {
			return
		}
		for _, op := range []*openapi31models.Operation{
			pi.Get(), pi.Put(), pi.Post(), pi.Delete(),
			pi.Options(), pi.Head(), pi.Patch(), pi.Trace(),
		} {
			if op == nil || op.Responses() == nil {
				continue
			}
			// Check default response
			if op.Responses().Default() != nil && op.Responses().Default().RawValue() != nil {
				resolveLinksInResponse(op.Responses().Default().RawValue(), ops)
			}
			// Check status code responses
			for _, respRef := range op.Responses().Codes() {
				if respRef != nil && respRef.RawValue() != nil {
					resolveLinksInResponse(respRef.RawValue(), ops)
				}
			}
		}
	}

	if doc.Paths() != nil {
		for _, pi := range doc.Paths().Items() {
			resolveLinksInPaths(pi)
		}
	}

	// Also check components
	if doc.Components() != nil {
		for _, linkRef := range doc.Components().Links() {
			if linkRef != nil && linkRef.RawValue() != nil {
				resolveSingleLinkOperationRef(linkRef.RawValue(), ops)
			}
		}
	}
}

func resolveLinksInResponse(resp *openapi31models.Response, ops map[string]*openapi31models.Operation) {
	for _, linkRef := range resp.Links() {
		if linkRef != nil && linkRef.RawValue() != nil {
			resolveSingleLinkOperationRef(linkRef.RawValue(), ops)
		}
	}
}

func resolveSingleLinkOperationRef(link *openapi31models.Link, ops map[string]*openapi31models.Operation) {
	if link.OperationRef() == "" {
		return
	}
	path, method, err := shared.ParseOperationRef(link.OperationRef())
	if err != nil {
		return
	}
	if op, ok := ops[path+"|"+method]; ok {
		link.Trix.ResolvedOperation = op
	}
}

// =============================================================================
// Done channel initialization — called before background goroutine starts
// =============================================================================

// initRefDoneChannels walks the entire document and calls InitDone() on every
// ref node that has Ref != "". This MUST be called before the background
// goroutine that runs Resolve(), so that consumers calling Value() will
// correctly block until resolution completes.
func initRefDoneChannels(doc *openapi31models.OpenAPI) {
	if doc == nil {
		return
	}

	if doc.Paths() != nil {
		for _, pi := range doc.Paths().Items() {
			initPathItemDone(pi)
		}
	}

	for _, ref := range doc.Webhooks() {
		initPathItemRefDone(ref)
	}

	if doc.Components() != nil {
		initComponentsDone(doc.Components())
	}
}

func initComponentsDone(c *openapi31models.Components) {
	for _, ref := range c.Schemas() {
		initSchemaRefDone(ref)
	}
	for _, ref := range c.Responses() {
		initResponseRefDone(ref)
	}
	for _, ref := range c.Parameters() {
		initParameterRefDone(ref)
	}
	for _, ref := range c.Examples() {
		initExampleRefDone(ref)
	}
	for _, ref := range c.RequestBodies() {
		initRequestBodyRefDone(ref)
	}
	for _, ref := range c.Headers() {
		initHeaderRefDone(ref)
	}
	for _, ref := range c.SecuritySchemes() {
		initSecuritySchemeRefDone(ref)
	}
	for _, ref := range c.Links() {
		initLinkRefDone(ref)
	}
	for _, ref := range c.Callbacks() {
		initCallbackRefDone(ref)
	}
	for _, ref := range c.PathItems() {
		initPathItemRefDone(ref)
	}
}

func initPathItemDone(pi *openapi31models.PathItem) {
	if pi == nil {
		return
	}
	for _, op := range []*openapi31models.Operation{
		pi.Get(), pi.Put(), pi.Post(), pi.Delete(),
		pi.Options(), pi.Head(), pi.Patch(), pi.Trace(),
	} {
		initOperationDone(op)
	}
	for _, ref := range pi.Parameters() {
		initParameterRefDone(ref)
	}
}

func initPathItemRefDone(ref *modelshared.RefWithMeta[openapi31models.PathItem]) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
		return
	}
	if ref.RawValue() != nil {
		initPathItemDone(ref.RawValue())
	}
}

func initOperationDone(op *openapi31models.Operation) {
	if op == nil {
		return
	}
	for _, ref := range op.Parameters() {
		initParameterRefDone(ref)
	}
	if op.RequestBody() != nil {
		initRequestBodyRefDone(op.RequestBody())
	}
	if op.Responses() != nil {
		if op.Responses().Default() != nil {
			initResponseRefDone(op.Responses().Default())
		}
		for _, ref := range op.Responses().Codes() {
			initResponseRefDone(ref)
		}
	}
	for _, ref := range op.Callbacks() {
		initCallbackRefDone(ref)
	}
}

func initSchemaRefDone(ref *modelshared.RefWithMeta[openapi31models.Schema]) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
		return
	}
	if ref.RawValue() != nil {
		initSchemaDone(ref.RawValue())
	}
}

func initSchemaDone(s *openapi31models.Schema) {
	if s == nil {
		return
	}
	for _, ref := range s.AllOf() {
		initSchemaRefDone(ref)
	}
	for _, ref := range s.OneOf() {
		initSchemaRefDone(ref)
	}
	for _, ref := range s.AnyOf() {
		initSchemaRefDone(ref)
	}
	initSchemaRefDone(s.Not())
	initSchemaRefDone(s.Items())
	for _, ref := range s.Properties() {
		initSchemaRefDone(ref)
	}
	initSchemaRefDone(s.AdditionalProperties())

	// OpenAPI 3.1 / JSON Schema 2020-12 additions
	initSchemaRefDone(s.If())
	initSchemaRefDone(s.Then())
	initSchemaRefDone(s.Else())
	for _, ref := range s.DependentSchemas() {
		initSchemaRefDone(ref)
	}
	for _, ref := range s.PrefixItems() {
		initSchemaRefDone(ref)
	}
	initSchemaRefDone(s.ContentSchema())
	initSchemaRefDone(s.UnevaluatedItems())
	initSchemaRefDone(s.UnevaluatedProperties())
}

func initResponseRefDone(ref *modelshared.RefWithMeta[openapi31models.Response]) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
		return
	}
	if ref.RawValue() != nil {
		for _, hRef := range ref.RawValue().Headers() {
			initHeaderRefDone(hRef)
		}
		for _, mt := range ref.RawValue().Content() {
			initMediaTypeDone(mt)
		}
		for _, lRef := range ref.RawValue().Links() {
			initLinkRefDone(lRef)
		}
	}
}

func initParameterRefDone(ref *modelshared.RefWithMeta[openapi31models.Parameter]) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
		return
	}
	if ref.RawValue() != nil {
		initSchemaRefDone(ref.RawValue().Schema())
		for _, eRef := range ref.RawValue().Examples() {
			initExampleRefDone(eRef)
		}
	}
}

func initExampleRefDone(ref *modelshared.RefWithMeta[openapi31models.Example]) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
	}
}

func initRequestBodyRefDone(ref *modelshared.RefWithMeta[openapi31models.RequestBody]) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
		return
	}
	if ref.RawValue() != nil {
		for _, mt := range ref.RawValue().Content() {
			initMediaTypeDone(mt)
		}
	}
}

func initHeaderRefDone(ref *modelshared.RefWithMeta[openapi31models.Header]) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
		return
	}
	if ref.RawValue() != nil {
		initSchemaRefDone(ref.RawValue().Schema())
		for _, eRef := range ref.RawValue().Examples() {
			initExampleRefDone(eRef)
		}
	}
}

func initSecuritySchemeRefDone(ref *modelshared.RefWithMeta[openapi31models.SecurityScheme]) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
	}
}

func initLinkRefDone(ref *modelshared.RefWithMeta[openapi31models.Link]) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
	}
}

func initCallbackRefDone(ref *modelshared.RefWithMeta[openapi31models.Callback]) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
		return
	}
	if ref.RawValue() != nil {
		for _, pathItem := range ref.RawValue().Paths() {
			initPathItemDone(pathItem)
		}
	}
}

func initMediaTypeDone(mt *openapi31models.MediaType) {
	if mt == nil {
		return
	}
	initSchemaRefDone(mt.Schema())
	for _, eRef := range mt.Examples() {
		initExampleRefDone(eRef)
	}
	for _, enc := range mt.Encoding() {
		if enc != nil {
			for _, hRef := range enc.Headers() {
				initHeaderRefDone(hRef)
			}
		}
	}
}
