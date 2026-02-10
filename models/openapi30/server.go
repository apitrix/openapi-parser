package openapi30

// Server represents a server.
// https://spec.openapis.org/oas/v3.0.3#server-object
type Server struct {
	Node // embedded - provides VendorExtensions and Trix

	url         string
	description string
	variables   map[string]*ServerVariable
}

func (s *Server) URL() string                           { return s.url }
func (s *Server) Description() string                   { return s.description }
func (s *Server) Variables() map[string]*ServerVariable { return s.variables }

// NewServer creates a new Server instance.
func NewServer(url, description string, variables map[string]*ServerVariable) *Server {
	return &Server{url: url, description: description, variables: variables}
}
