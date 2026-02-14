package openapi31x

import "openapi-parser/parsers/shared"

// ParseConfig controls which features are enabled during parsing.
// Use this alias instead of shared.ParseConfig when calling Parse, ParseReader, or ParseFile.
type ParseConfig = shared.ParseConfig

// All returns a config with all features enabled. This is the default
// when no config (nil) is passed to a Parse function.
func All() *ParseConfig { return shared.All() }

// None returns a config with all features disabled.
func None() *ParseConfig { return shared.None() }

// Defaults returns the provided config if non-nil, otherwise All().
func Defaults(cfg *ParseConfig) *ParseConfig { return shared.Defaults(cfg) }

// FirstConfig extracts the first non-nil config from a variadic list,
// falling back to All() if none provided.
func FirstConfig(cfgs []*ParseConfig) *ParseConfig { return shared.FirstConfig(cfgs) }
