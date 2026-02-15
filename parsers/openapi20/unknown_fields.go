package openapi20

import (
	"github.com/apitrix/openapi-parser/parsers/shared"

	"gopkg.in/yaml.v3"
)

// UnknownField represents an unrecognized field found during parsing.
// Extension fields (x-*) are not considered unknown.
// This is a type alias for backward compatibility.
type UnknownField = shared.UnknownField

// detectUnknownNodeFields checks a yaml.Node for fields not in the known set.
// Extensions (x-*) are excluded from the result.
func detectUnknownNodeFields(node *yaml.Node, knownFields map[string]struct{}, basePath string) []UnknownField {
	return shared.DetectUnknownNodeFields(node, knownFields, basePath)
}
