package openapi30

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Server represents a server.
// https://spec.openapis.org/oas/v3.0.3#server-object
type Server struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	url         string
	description string
	variables   map[string]*ServerVariable
}

func (s *Server) URL() string                           { return s.url }
func (s *Server) Description() string                   { return s.description }
func (s *Server) Variables() map[string]*ServerVariable { return s.variables }

func (s *Server) SetURL(url string) error {
	if err := s.Trix.RunHooks("url", s.url, url); err != nil {
		return err
	}
	s.url = url
	return nil
}
func (s *Server) SetDescription(description string) error {
	if err := s.Trix.RunHooks("description", s.description, description); err != nil {
		return err
	}
	s.description = description
	return nil
}
func (s *Server) SetVariables(variables map[string]*ServerVariable) error {
	if err := s.Trix.RunHooks("variables", s.variables, variables); err != nil {
		return err
	}
	s.variables = variables
	return nil
}

// NewServer creates a new Server instance.
func NewServer(url, description string, variables map[string]*ServerVariable) *Server {
	return &Server{url: url, description: description, variables: variables}
}

func (s *Server) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "url", Value: s.url},
		{Key: "description", Value: s.description},
		{Key: "variables", Value: s.variables},
	}
	return shared.AppendExtensions(fields, s.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (s *Server) MarshalFields() []shared.Field { return s.marshalFields() }

func (s *Server) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(s.marshalFields())
}

func (s *Server) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(s.marshalFields())
}

var _ yaml.Marshaler = (*Server)(nil)
