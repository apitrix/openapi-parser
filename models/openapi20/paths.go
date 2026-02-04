package openapi20

// Paths holds the relative paths to individual endpoints.
// https://swagger.io/specification/v2/#paths-object
type Paths struct {
	Node // embedded - provides NodeSource and Extensions

	// Items maps paths to their definitions.
	Items map[string]*PathItem `json:"-" yaml:"-"`
}

// PathItem describes operations available on a single path.
// https://swagger.io/specification/v2/#path-item-object
type PathItem struct {
	Node // embedded - provides NodeSource and Extensions

	Ref        string          `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Get        *Operation      `json:"get,omitempty" yaml:"get,omitempty"`
	Put        *Operation      `json:"put,omitempty" yaml:"put,omitempty"`
	Post       *Operation      `json:"post,omitempty" yaml:"post,omitempty"`
	Delete     *Operation      `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options    *Operation      `json:"options,omitempty" yaml:"options,omitempty"`
	Head       *Operation      `json:"head,omitempty" yaml:"head,omitempty"`
	Patch      *Operation      `json:"patch,omitempty" yaml:"patch,omitempty"`
	Parameters []*ParameterRef `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
