package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	modelsShared "openapi-parser/models/shared"
	"openapi-parser/parsers/shared"
)

// flattenErrors walks the parsed document tree via getter methods and collects
// all Trix.Errors from every node into a flat []*shared.ParseError slice.
// Because model fields are now private (readonly pattern), we must use getters
// rather than reflection to traverse the tree.
func flattenErrors(doc *openapi31models.OpenAPI) []*shared.ParseError {
	if doc == nil {
		return nil
	}
	var result []*shared.ParseError
	visited := make(map[interface{}]bool)

	collectNodeErrors(&doc.ElementBase, &result)
	visited[doc] = true

	// Info
	if info := doc.Info(); info != nil {
		collectNodeErrors(&info.ElementBase, &result)
		if c := info.Contact(); c != nil {
			collectNodeErrors(&c.ElementBase, &result)
		}
		if l := info.License(); l != nil {
			collectNodeErrors(&l.ElementBase, &result)
		}
	}

	// Servers
	for _, s := range doc.Servers() {
		flattenServer(s, &result)
	}

	// Paths
	if paths := doc.Paths(); paths != nil {
		collectNodeErrors(&paths.ElementBase, &result)
		for _, pi := range paths.Items() {
			flattenPathItem(pi, &result, visited)
		}
	}

	// Webhooks
	for _, ref := range doc.Webhooks() {
		if ref != nil {
			collectRefErrors(ref, &result)
			if ref.Value() != nil && !visited[ref.Value()] {
				flattenPathItem(ref.Value(), &result, visited)
			}
		}
	}

	// Components
	if comp := doc.Components(); comp != nil {
		collectNodeErrors(&comp.ElementBase, &result)

		for _, ref := range comp.Schemas() {
			if ref != nil {
				collectRefErrors(ref, &result)
				if ref.Value() != nil && !visited[ref.Value()] {
					flattenSchema(ref.Value(), &result, visited)
				}
			}
		}
		for _, ref := range comp.Parameters() {
			if ref != nil {
				collectRefErrors(ref, &result)
				if ref.Value() != nil && !visited[ref.Value()] {
					flattenParameter(ref.Value(), &result, visited)
				}
			}
		}
		for _, ref := range comp.Headers() {
			if ref != nil {
				collectRefErrors(ref, &result)
				if ref.Value() != nil && !visited[ref.Value()] {
					flattenHeader(ref.Value(), &result, visited)
				}
			}
		}
		for _, ref := range comp.RequestBodies() {
			if ref != nil {
				collectRefErrors(ref, &result)
				if ref.Value() != nil && !visited[ref.Value()] {
					flattenRequestBody(ref.Value(), &result, visited)
				}
			}
		}
		for _, ref := range comp.Responses() {
			if ref != nil {
				collectRefErrors(ref, &result)
				if ref.Value() != nil && !visited[ref.Value()] {
					flattenResponse(ref.Value(), &result, visited)
				}
			}
		}
		for _, ref := range comp.SecuritySchemes() {
			if ref != nil {
				collectRefErrors(ref, &result)
				if ref.Value() != nil && !visited[ref.Value()] {
					flattenSecurityScheme(ref.Value(), &result)
				}
			}
		}
		for _, ref := range comp.Links() {
			if ref != nil {
				collectRefErrors(ref, &result)
				if ref.Value() != nil && !visited[ref.Value()] {
					flattenLink(ref.Value(), &result)
				}
			}
		}
		for _, ref := range comp.Callbacks() {
			if ref != nil {
				collectRefErrors(ref, &result)
				if ref.Value() != nil && !visited[ref.Value()] {
					flattenCallback(ref.Value(), &result, visited)
				}
			}
		}
		for _, ref := range comp.Examples() {
			if ref != nil {
				collectRefErrors(ref, &result)
				if ref.Value() != nil {
					collectNodeErrors(&ref.Value().ElementBase, &result)
				}
			}
		}
		for _, ref := range comp.PathItems() {
			if ref != nil {
				collectRefErrors(ref, &result)
				if ref.Value() != nil && !visited[ref.Value()] {
					flattenPathItem(ref.Value(), &result, visited)
				}
			}
		}
	}

	// Tags
	for _, tag := range doc.Tags() {
		if tag != nil {
			collectNodeErrors(&tag.ElementBase, &result)
			if ed := tag.ExternalDocs(); ed != nil {
				collectNodeErrors(&ed.ElementBase, &result)
			}
		}
	}

	// ExternalDocs
	if ed := doc.ExternalDocs(); ed != nil {
		collectNodeErrors(&ed.ElementBase, &result)
	}

	return result
}

func collectNodeErrors(node *openapi31models.ElementBase, result *[]*shared.ParseError) {
	for _, e := range node.Trix.Errors {
		*result = append(*result, modelParseErrorToShared(e))
	}
}

func flattenServer(s *openapi31models.Server, result *[]*shared.ParseError) {
	if s == nil {
		return
	}
	collectNodeErrors(&s.ElementBase, result)
	for _, sv := range s.Variables() {
		if sv != nil {
			collectNodeErrors(&sv.ElementBase, result)
		}
	}
}

func flattenPathItem(pi *openapi31models.PathItem, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if pi == nil || visited[pi] {
		return
	}
	visited[pi] = true
	collectNodeErrors(&pi.ElementBase, result)

	for _, op := range []*openapi31models.Operation{
		pi.Get(), pi.Put(), pi.Post(), pi.Delete(),
		pi.Options(), pi.Head(), pi.Patch(), pi.Trace(),
	} {
		flattenOperation(op, result, visited)
	}

	for _, ref := range pi.Parameters() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil {
				flattenParameter(ref.Value(), result, visited)
			}
		}
	}

	for _, s := range pi.Servers() {
		flattenServer(s, result)
	}
}

func flattenOperation(op *openapi31models.Operation, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if op == nil {
		return
	}
	collectNodeErrors(&op.ElementBase, result)

	for _, ref := range op.Parameters() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil && !visited[ref.Value()] {
				flattenParameter(ref.Value(), result, visited)
			}
		}
	}

	if rb := op.RequestBody(); rb != nil {
		collectRefErrors(rb, result)
		if rb.Value() != nil && !visited[rb.Value()] {
			flattenRequestBody(rb.Value(), result, visited)
		}
	}

	if resp := op.Responses(); resp != nil {
		collectNodeErrors(&resp.ElementBase, result)
		for _, ref := range resp.Codes() {
			if ref != nil {
				collectRefErrors(ref, result)
				if ref.Value() != nil && !visited[ref.Value()] {
					flattenResponse(ref.Value(), result, visited)
				}
			}
		}
	}

	for _, ref := range op.Callbacks() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil && !visited[ref.Value()] {
				flattenCallback(ref.Value(), result, visited)
			}
		}
	}

	for _, s := range op.Servers() {
		flattenServer(s, result)
	}
}

func flattenParameter(p *openapi31models.Parameter, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if p == nil || visited[p] {
		return
	}
	visited[p] = true
	collectNodeErrors(&p.ElementBase, result)

	if s := p.Schema(); s != nil {
		collectRefErrors(s, result)
		if s.Value() != nil {
			flattenSchema(s.Value(), result, visited)
		}
	}
	for _, ref := range p.Examples() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil {
				collectNodeErrors(&ref.Value().ElementBase, result)
			}
		}
	}
	flattenContent(p.Content(), result, visited)
}

func flattenHeader(h *openapi31models.Header, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if h == nil || visited[h] {
		return
	}
	visited[h] = true
	collectNodeErrors(&h.ElementBase, result)

	if s := h.Schema(); s != nil {
		collectRefErrors(s, result)
		if s.Value() != nil {
			flattenSchema(s.Value(), result, visited)
		}
	}
	for _, ref := range h.Examples() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil {
				collectNodeErrors(&ref.Value().ElementBase, result)
			}
		}
	}
	flattenContent(h.Content(), result, visited)
}

func flattenRequestBody(rb *openapi31models.RequestBody, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if rb == nil || visited[rb] {
		return
	}
	visited[rb] = true
	collectNodeErrors(&rb.ElementBase, result)
	flattenContent(rb.Content(), result, visited)
}

func flattenResponse(resp *openapi31models.Response, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if resp == nil || visited[resp] {
		return
	}
	visited[resp] = true
	collectNodeErrors(&resp.ElementBase, result)

	for _, ref := range resp.Headers() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil {
				flattenHeader(ref.Value(), result, visited)
			}
		}
	}
	flattenContent(resp.Content(), result, visited)

	for _, ref := range resp.Links() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil {
				flattenLink(ref.Value(), result)
			}
		}
	}
}

func flattenContent(content map[string]*openapi31models.MediaType, result *[]*shared.ParseError, visited map[interface{}]bool) {
	for _, mt := range content {
		if mt == nil {
			continue
		}
		collectNodeErrors(&mt.ElementBase, result)

		if s := mt.Schema(); s != nil {
			collectRefErrors(s, result)
			if s.Value() != nil {
				flattenSchema(s.Value(), result, visited)
			}
		}
		for _, ref := range mt.Examples() {
			if ref != nil {
				collectRefErrors(ref, result)
				if ref.Value() != nil {
					collectNodeErrors(&ref.Value().ElementBase, result)
				}
			}
		}
		for _, enc := range mt.Encoding() {
			if enc != nil {
				collectNodeErrors(&enc.ElementBase, result)
			}
		}
	}
}

func flattenSchema(s *openapi31models.Schema, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if s == nil || visited[s] {
		return
	}
	visited[s] = true
	collectNodeErrors(&s.ElementBase, result)

	if d := s.Discriminator(); d != nil {
		collectNodeErrors(&d.ElementBase, result)
	}
	if x := s.XML(); x != nil {
		collectNodeErrors(&x.ElementBase, result)
	}
	if ed := s.ExternalDocs(); ed != nil {
		collectNodeErrors(&ed.ElementBase, result)
	}

	// Composition keywords
	for _, ref := range s.AllOf() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil {
				flattenSchema(ref.Value(), result, visited)
			}
		}
	}
	for _, ref := range s.OneOf() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil {
				flattenSchema(ref.Value(), result, visited)
			}
		}
	}
	for _, ref := range s.AnyOf() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil {
				flattenSchema(ref.Value(), result, visited)
			}
		}
	}
	if n := s.Not(); n != nil {
		collectRefErrors(n, result)
		if n.Value() != nil {
			flattenSchema(n.Value(), result, visited)
		}
	}
	if itm := s.Items(); itm != nil {
		collectRefErrors(itm, result)
		if itm.Value() != nil {
			flattenSchema(itm.Value(), result, visited)
		}
	}
	for _, ref := range s.Properties() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil {
				flattenSchema(ref.Value(), result, visited)
			}
		}
	}
	if ap := s.AdditionalProperties(); ap != nil {
		collectRefErrors(ap, result)
		if ap.Value() != nil {
			flattenSchema(ap.Value(), result, visited)
		}
	}

	// 3.1 additions
	if ifRef := s.If(); ifRef != nil {
		collectRefErrors(ifRef, result)
		if ifRef.Value() != nil {
			flattenSchema(ifRef.Value(), result, visited)
		}
	}
	if thenRef := s.Then(); thenRef != nil {
		collectRefErrors(thenRef, result)
		if thenRef.Value() != nil {
			flattenSchema(thenRef.Value(), result, visited)
		}
	}
	if elseRef := s.Else(); elseRef != nil {
		collectRefErrors(elseRef, result)
		if elseRef.Value() != nil {
			flattenSchema(elseRef.Value(), result, visited)
		}
	}
	for _, ref := range s.PrefixItems() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil {
				flattenSchema(ref.Value(), result, visited)
			}
		}
	}
	for _, ref := range s.DependentSchemas() {
		if ref != nil {
			collectRefErrors(ref, result)
			if ref.Value() != nil {
				flattenSchema(ref.Value(), result, visited)
			}
		}
	}
	if cs := s.ContentSchema(); cs != nil {
		collectRefErrors(cs, result)
		if cs.Value() != nil {
			flattenSchema(cs.Value(), result, visited)
		}
	}
	if ui := s.UnevaluatedItems(); ui != nil {
		collectRefErrors(ui, result)
		if ui.Value() != nil {
			flattenSchema(ui.Value(), result, visited)
		}
	}
	if up := s.UnevaluatedProperties(); up != nil {
		collectRefErrors(up, result)
		if up.Value() != nil {
			flattenSchema(up.Value(), result, visited)
		}
	}
}

func flattenSecurityScheme(ss *openapi31models.SecurityScheme, result *[]*shared.ParseError) {
	if ss == nil {
		return
	}
	collectNodeErrors(&ss.ElementBase, result)

	if flows := ss.Flows(); flows != nil {
		collectNodeErrors(&flows.ElementBase, result)
		for _, flow := range []*openapi31models.OAuthFlow{
			flows.Implicit(), flows.Password(),
			flows.ClientCredentials(), flows.AuthorizationCode(),
		} {
			if flow != nil {
				collectNodeErrors(&flow.ElementBase, result)
			}
		}
	}
}

func flattenLink(l *openapi31models.Link, result *[]*shared.ParseError) {
	if l == nil {
		return
	}
	collectNodeErrors(&l.ElementBase, result)
	if s := l.Server(); s != nil {
		flattenServer(s, result)
	}
}

func flattenCallback(cb *openapi31models.Callback, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if cb == nil || visited[cb] {
		return
	}
	visited[cb] = true
	collectNodeErrors(&cb.ElementBase, result)
	for _, pi := range cb.Paths() {
		if pi != nil {
			flattenPathItem(pi, result, visited)
		}
	}
}

// modelParseErrorToShared converts a model-level ParseError to a shared.ParseError.
func modelParseErrorToShared(e openapi31models.ParseError) *shared.ParseError {
	return &shared.ParseError{
		Path:    e.Path,
		Message: e.Message,
		Kind:    e.Kind,
	}
}

// collectRefErrors collects Trix.Errors from a ref (including resolution errors).
func collectRefErrors[T any](ref *modelsShared.RefWithMeta[T], result *[]*shared.ParseError) {
	if ref == nil {
		return
	}
	for _, e := range ref.Trix.Errors {
		*result = append(*result, modelParseErrorToShared(e))
	}
}
