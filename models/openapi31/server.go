package openapi31

// Server represents a server.
// https://spec.openapis.org/oas/v3.1.0#server-object
type Server struct {
	Node // embedded - provides VendorExtensions and Trix

	URL         string                     `json:"url" yaml:"url"`
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

// NewServer creates a new Server instance.
func NewServer(url string) *Server {
	return &Server{URL: url}
}
