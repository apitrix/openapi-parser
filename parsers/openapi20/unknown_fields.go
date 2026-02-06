package openapi20

import "gopkg.in/yaml.v3"

// UnknownField represents an unrecognized field found during parsing.
// Extension fields (x-*) are not considered unknown.
type UnknownField struct {
	Path   string // JSON path to the field (e.g., "paths./pets.get")
	Name   string // Field name
	Line   int    // Line number (1-based)
	Column int    // Column number (1-based)
}

// detectUnknownNodeFields checks a yaml.Node for fields not in the known set.
// Extensions (x-*) are excluded from the result.
func detectUnknownNodeFields(node *yaml.Node, knownFields map[string]struct{}, basePath string) []UnknownField {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}

	var unknown []UnknownField
	for i := 0; i < len(node.Content)-1; i += 2 {
		key := node.Content[i].Value

		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		// Check if known
		if _, ok := knownFields[key]; !ok {
			path := key
			if basePath != "" {
				path = basePath + "." + key
			}
			unknown = append(unknown, UnknownField{
				Path:   path,
				Name:   key,
				Line:   node.Content[i].Line,
				Column: node.Content[i].Column,
			})
		}
	}

	return unknown
}
