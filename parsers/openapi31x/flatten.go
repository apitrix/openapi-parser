package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
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

	collectNodeErrors(&doc.Node, &result)
	visited[doc] = true

	// Info
	if info := doc.Info(); info != nil {
		collectNodeErrors(&info.Node, &result)
		if c := info.Contact(); c != nil {
			collectNodeErrors(&c.Node, &result)
		}
		if l := info.License(); l != nil {
			collectNodeErrors(&l.Node, &result)
		}
	}

	// Servers
	for _, s := range doc.Servers() {
		flattenServer(s, &result)
	}

	// Paths
	if paths := doc.Paths(); paths != nil {
		collectNodeErrors(&paths.Node, &result)
		for _, pi := range paths.Items() {
			flattenPathItem(pi, &result, visited)
		}
	}

	// Webhooks
	for _, ref := range doc.Webhooks() {
		if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
			flattenPathItem(ref.Value(), &result, visited)
		}
	}

	// Components
	if comp := doc.Components(); comp != nil {
		collectNodeErrors(&comp.Node, &result)

		for _, ref := range comp.Schemas() {
			if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
				flattenSchema(ref.Value(), &result, visited)
			}
		}
		for _, ref := range comp.Parameters() {
			if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
				flattenParameter(ref.Value(), &result, visited)
			}
		}
		for _, ref := range comp.Headers() {
			if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
				flattenHeader(ref.Value(), &result, visited)
			}
		}
		for _, ref := range comp.RequestBodies() {
			if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
				flattenRequestBody(ref.Value(), &result, visited)
			}
		}
		for _, ref := range comp.Responses() {
			if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
				flattenResponse(ref.Value(), &result, visited)
			}
		}
		for _, ref := range comp.SecuritySchemes() {
			if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
				flattenSecurityScheme(ref.Value(), &result)
			}
		}
		for _, ref := range comp.Links() {
			if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
				flattenLink(ref.Value(), &result)
			}
		}
		for _, ref := range comp.Callbacks() {
			if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
				flattenCallback(ref.Value(), &result, visited)
			}
		}
		for _, ref := range comp.Examples() {
			if ref != nil && ref.Value() != nil {
				collectNodeErrors(&ref.Value().Node, &result)
			}
		}
		for _, ref := range comp.PathItems() {
			if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
				flattenPathItem(ref.Value(), &result, visited)
			}
		}
	}

	// Tags
	for _, tag := range doc.Tags() {
		if tag != nil {
			collectNodeErrors(&tag.Node, &result)
			if ed := tag.ExternalDocs(); ed != nil {
				collectNodeErrors(&ed.Node, &result)
			}
		}
	}

	// ExternalDocs
	if ed := doc.ExternalDocs(); ed != nil {
		collectNodeErrors(&ed.Node, &result)
	}

	return result
}

func collectNodeErrors(node *openapi31models.Node, result *[]*shared.ParseError) {
	for _, e := range node.Trix.Errors {
		*result = append(*result, modelParseErrorToShared(e))
	}
}

func flattenServer(s *openapi31models.Server, result *[]*shared.ParseError) {
	if s == nil {
		return
	}
	collectNodeErrors(&s.Node, result)
	for _, sv := range s.Variables() {
		if sv != nil {
			collectNodeErrors(&sv.Node, result)
		}
	}
}

func flattenPathItem(pi *openapi31models.PathItem, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if pi == nil || visited[pi] {
		return
	}
	visited[pi] = true
	collectNodeErrors(&pi.Node, result)

	for _, op := range []*openapi31models.Operation{
		pi.Get(), pi.Put(), pi.Post(), pi.Delete(),
		pi.Options(), pi.Head(), pi.Patch(), pi.Trace(),
	} {
		flattenOperation(op, result, visited)
	}

	for _, ref := range pi.Parameters() {
		if ref != nil && ref.Value() != nil {
			flattenParameter(ref.Value(), result, visited)
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
	collectNodeErrors(&op.Node, result)

	for _, ref := range op.Parameters() {
		if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
			flattenParameter(ref.Value(), result, visited)
		}
	}

	if rb := op.RequestBody(); rb != nil && rb.Value() != nil && !visited[rb.Value()] {
		flattenRequestBody(rb.Value(), result, visited)
	}

	if resp := op.Responses(); resp != nil {
		collectNodeErrors(&resp.Node, result)
		for _, ref := range resp.Codes() {
			if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
				flattenResponse(ref.Value(), result, visited)
			}
		}
	}

	for _, ref := range op.Callbacks() {
		if ref != nil && ref.Value() != nil && !visited[ref.Value()] {
			flattenCallback(ref.Value(), result, visited)
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
	collectNodeErrors(&p.Node, result)

	if s := p.Schema(); s != nil && s.Value() != nil {
		flattenSchema(s.Value(), result, visited)
	}
	for _, ref := range p.Examples() {
		if ref != nil && ref.Value() != nil {
			collectNodeErrors(&ref.Value().Node, result)
		}
	}
	flattenContent(p.Content(), result, visited)
}

func flattenHeader(h *openapi31models.Header, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if h == nil || visited[h] {
		return
	}
	visited[h] = true
	collectNodeErrors(&h.Node, result)

	if s := h.Schema(); s != nil && s.Value() != nil {
		flattenSchema(s.Value(), result, visited)
	}
	for _, ref := range h.Examples() {
		if ref != nil && ref.Value() != nil {
			collectNodeErrors(&ref.Value().Node, result)
		}
	}
	flattenContent(h.Content(), result, visited)
}

func flattenRequestBody(rb *openapi31models.RequestBody, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if rb == nil || visited[rb] {
		return
	}
	visited[rb] = true
	collectNodeErrors(&rb.Node, result)
	flattenContent(rb.Content(), result, visited)
}

func flattenResponse(resp *openapi31models.Response, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if resp == nil || visited[resp] {
		return
	}
	visited[resp] = true
	collectNodeErrors(&resp.Node, result)

	for _, ref := range resp.Headers() {
		if ref != nil && ref.Value() != nil {
			flattenHeader(ref.Value(), result, visited)
		}
	}
	flattenContent(resp.Content(), result, visited)

	for _, ref := range resp.Links() {
		if ref != nil && ref.Value() != nil {
			flattenLink(ref.Value(), result)
		}
	}
}

func flattenContent(content map[string]*openapi31models.MediaType, result *[]*shared.ParseError, visited map[interface{}]bool) {
	for _, mt := range content {
		if mt == nil {
			continue
		}
		collectNodeErrors(&mt.Node, result)

		if s := mt.Schema(); s != nil && s.Value() != nil {
			flattenSchema(s.Value(), result, visited)
		}
		for _, ref := range mt.Examples() {
			if ref != nil && ref.Value() != nil {
				collectNodeErrors(&ref.Value().Node, result)
			}
		}
		for _, enc := range mt.Encoding() {
			if enc != nil {
				collectNodeErrors(&enc.Node, result)
			}
		}
	}
}

func flattenSchema(s *openapi31models.Schema, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if s == nil || visited[s] {
		return
	}
	visited[s] = true
	collectNodeErrors(&s.Node, result)

	if d := s.Discriminator(); d != nil {
		collectNodeErrors(&d.Node, result)
	}
	if x := s.XML(); x != nil {
		collectNodeErrors(&x.Node, result)
	}
	if ed := s.ExternalDocs(); ed != nil {
		collectNodeErrors(&ed.Node, result)
	}

	// Composition keywords
	for _, ref := range s.AllOf() {
		if ref != nil && ref.Value() != nil {
			flattenSchema(ref.Value(), result, visited)
		}
	}
	for _, ref := range s.OneOf() {
		if ref != nil && ref.Value() != nil {
			flattenSchema(ref.Value(), result, visited)
		}
	}
	for _, ref := range s.AnyOf() {
		if ref != nil && ref.Value() != nil {
			flattenSchema(ref.Value(), result, visited)
		}
	}
	if n := s.Not(); n != nil && n.Value() != nil {
		flattenSchema(n.Value(), result, visited)
	}
	if itm := s.Items(); itm != nil && itm.Value() != nil {
		flattenSchema(itm.Value(), result, visited)
	}
	for _, ref := range s.Properties() {
		if ref != nil && ref.Value() != nil {
			flattenSchema(ref.Value(), result, visited)
		}
	}
	if ap := s.AdditionalProperties(); ap != nil && ap.Value() != nil {
		flattenSchema(ap.Value(), result, visited)
	}

	// 3.1 additions
	if ifRef := s.If(); ifRef != nil && ifRef.Value() != nil {
		flattenSchema(ifRef.Value(), result, visited)
	}
	if thenRef := s.Then(); thenRef != nil && thenRef.Value() != nil {
		flattenSchema(thenRef.Value(), result, visited)
	}
	if elseRef := s.Else(); elseRef != nil && elseRef.Value() != nil {
		flattenSchema(elseRef.Value(), result, visited)
	}
	for _, ref := range s.PrefixItems() {
		if ref != nil && ref.Value() != nil {
			flattenSchema(ref.Value(), result, visited)
		}
	}
	for _, ref := range s.DependentSchemas() {
		if ref != nil && ref.Value() != nil {
			flattenSchema(ref.Value(), result, visited)
		}
	}
	if cs := s.ContentSchema(); cs != nil && cs.Value() != nil {
		flattenSchema(cs.Value(), result, visited)
	}
	if ui := s.UnevaluatedItems(); ui != nil && ui.Value() != nil {
		flattenSchema(ui.Value(), result, visited)
	}
	if up := s.UnevaluatedProperties(); up != nil && up.Value() != nil {
		flattenSchema(up.Value(), result, visited)
	}
}

func flattenSecurityScheme(ss *openapi31models.SecurityScheme, result *[]*shared.ParseError) {
	if ss == nil {
		return
	}
	collectNodeErrors(&ss.Node, result)

	if flows := ss.Flows(); flows != nil {
		collectNodeErrors(&flows.Node, result)
		for _, flow := range []*openapi31models.OAuthFlow{
			flows.Implicit(), flows.Password(),
			flows.ClientCredentials(), flows.AuthorizationCode(),
		} {
			if flow != nil {
				collectNodeErrors(&flow.Node, result)
			}
		}
	}
}

func flattenLink(l *openapi31models.Link, result *[]*shared.ParseError) {
	if l == nil {
		return
	}
	collectNodeErrors(&l.Node, result)
	if s := l.Server(); s != nil {
		flattenServer(s, result)
	}
}

func flattenCallback(cb *openapi31models.Callback, result *[]*shared.ParseError, visited map[interface{}]bool) {
	if cb == nil || visited[cb] {
		return
	}
	visited[cb] = true
	collectNodeErrors(&cb.Node, result)
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
