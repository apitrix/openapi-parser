package openapi30x

import (
	"fmt"

	openapi30models "openapi-parser/models/openapi30"
	"openapi-parser/parsers/shared"

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
	if err := resolveDocument(doc, r, resolving); err != nil {
		return err
	}
	// Post-resolve: wire up operationRefs now that all operations are parsed
	resolveOperationRefs(doc)
	return nil
}

func resolveDocument(doc *openapi30models.OpenAPI, r *shared.RefResolver, resolving map[string]bool) error {
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

	if doc.Components() != nil {
		if err := resolveComponents(doc.Components(), r, resolving); err != nil {
			return err
		}
	}

	return nil
}

func resolveComponents(c *openapi30models.Components, r *shared.RefResolver, resolving map[string]bool) error {
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
	return nil
}

func resolvePathItem(pi *openapi30models.PathItem, r *shared.RefResolver, resolving map[string]bool) error {
	if pi == nil {
		return nil
	}

	for _, op := range []*openapi30models.Operation{
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

func resolveOperation(op *openapi30models.Operation, r *shared.RefResolver, resolving map[string]bool) error {
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

func resolveSchemaRef(ref *openapi30models.SchemaRef, r *shared.RefResolver, resolving map[string]bool) error {
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

func resolveSchema(schema *openapi30models.Schema, r *shared.RefResolver, resolving map[string]bool) error {
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

	// Resolve discriminator.mapping values
	if schema.Discriminator() != nil && len(schema.Discriminator().Mapping()) > 0 {
		resolved := make(map[string]*openapi30models.SchemaRef)
		for key, val := range schema.Discriminator().Mapping() {
			mapResult, mapErr := r.ResolveMapping(val)
			if mapErr == nil {
				ctx := newParseContext(mapResult.Node, shared.All())
				s, parseErr := parseSharedSchema(mapResult.Node, ctx)
				if parseErr == nil {
					ref := openapi30models.NewSchemaRef(val)
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

func resolveResponseRef(ref *openapi30models.ResponseRef, r *shared.RefResolver, resolving map[string]bool) error {
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

func resolveParameterRef(ref *openapi30models.ParameterRef, r *shared.RefResolver, resolving map[string]bool) error {
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

func resolveExampleRef(ref *openapi30models.ExampleRef, r *shared.RefResolver, resolving map[string]bool) error {
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

func resolveRequestBodyRef(ref *openapi30models.RequestBodyRef, r *shared.RefResolver, resolving map[string]bool) error {
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

func resolveHeaderRef(ref *openapi30models.HeaderRef, r *shared.RefResolver, resolving map[string]bool) error {
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

func resolveSecuritySchemeRef(ref *openapi30models.SecuritySchemeRef, r *shared.RefResolver, resolving map[string]bool) error {
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

func resolveLinkRef(ref *openapi30models.LinkRef, r *shared.RefResolver, resolving map[string]bool) error {
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

func resolveCallbackRef(ref *openapi30models.CallbackRef, r *shared.RefResolver, resolving map[string]bool) error {
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

func resolveMediaType(mt *openapi30models.MediaType, r *shared.RefResolver, resolving map[string]bool) error {
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

func resolveOperationRefs(doc *openapi30models.OpenAPI) {
	if doc == nil {
		return
	}
	ops := collectOperations(doc)
	walkLinksForOperationRef(doc, ops)
}

func collectOperations(doc *openapi30models.OpenAPI) map[string]*openapi30models.Operation {
	ops := make(map[string]*openapi30models.Operation)
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

func pathItemOps(pi *openapi30models.PathItem) map[string]*openapi30models.Operation {
	return map[string]*openapi30models.Operation{
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

func walkLinksForOperationRef(doc *openapi30models.OpenAPI, ops map[string]*openapi30models.Operation) {
	resolveLinksInPaths := func(pi *openapi30models.PathItem) {
		if pi == nil {
			return
		}
		for _, op := range []*openapi30models.Operation{
			pi.Get(), pi.Put(), pi.Post(), pi.Delete(),
			pi.Options(), pi.Head(), pi.Patch(), pi.Trace(),
		} {
			if op == nil || op.Responses() == nil {
				continue
			}
			if op.Responses().Default() != nil && op.Responses().Default().RawValue() != nil {
				resolveLinksInResponse(op.Responses().Default().RawValue(), ops)
			}
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

	if doc.Components() != nil {
		for _, linkRef := range doc.Components().Links() {
			if linkRef != nil && linkRef.RawValue() != nil {
				resolveSingleLinkOperationRef(linkRef.RawValue(), ops)
			}
		}
	}
}

func resolveLinksInResponse(resp *openapi30models.Response, ops map[string]*openapi30models.Operation) {
	for _, linkRef := range resp.Links() {
		if linkRef != nil && linkRef.RawValue() != nil {
			resolveSingleLinkOperationRef(linkRef.RawValue(), ops)
		}
	}
}

func resolveSingleLinkOperationRef(link *openapi30models.Link, ops map[string]*openapi30models.Operation) {
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
func initRefDoneChannels(doc *openapi30models.OpenAPI) {
	if doc == nil {
		return
	}

	if doc.Paths() != nil {
		for _, pi := range doc.Paths().Items() {
			initPathItemDone(pi)
		}
	}

	if doc.Components() != nil {
		initComponentsDone(doc.Components())
	}
}

func initComponentsDone(c *openapi30models.Components) {
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
}

func initPathItemDone(pi *openapi30models.PathItem) {
	if pi == nil {
		return
	}
	for _, op := range []*openapi30models.Operation{
		pi.Get(), pi.Put(), pi.Post(), pi.Delete(),
		pi.Options(), pi.Head(), pi.Patch(), pi.Trace(),
	} {
		initOperationDone(op)
	}
	for _, ref := range pi.Parameters() {
		initParameterRefDone(ref)
	}
}

func initOperationDone(op *openapi30models.Operation) {
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

func initSchemaRefDone(ref *openapi30models.SchemaRef) {
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

func initSchemaDone(s *openapi30models.Schema) {
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
}

func initResponseRefDone(ref *openapi30models.ResponseRef) {
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

func initParameterRefDone(ref *openapi30models.ParameterRef) {
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

func initExampleRefDone(ref *openapi30models.ExampleRef) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
	}
}

func initRequestBodyRefDone(ref *openapi30models.RequestBodyRef) {
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

func initHeaderRefDone(ref *openapi30models.HeaderRef) {
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

func initSecuritySchemeRefDone(ref *openapi30models.SecuritySchemeRef) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
	}
}

func initLinkRefDone(ref *openapi30models.LinkRef) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
	}
}

func initCallbackRefDone(ref *openapi30models.CallbackRef) {
	if ref == nil {
		return
	}
	if ref.Ref != "" {
		ref.InitDone()
		return
	}
	if ref.RawValue() != nil {
		for _, pi := range ref.RawValue().Paths() {
			initPathItemDone(pi)
		}
	}
}

func initMediaTypeDone(mt *openapi30models.MediaType) {
	if mt == nil {
		return
	}
	initSchemaRefDone(mt.Schema())
	for _, ref := range mt.Examples() {
		initExampleRefDone(ref)
	}
	for _, enc := range mt.Encoding() {
		if enc != nil {
			for _, hRef := range enc.Headers() {
				initHeaderRefDone(hRef)
			}
		}
	}
}
