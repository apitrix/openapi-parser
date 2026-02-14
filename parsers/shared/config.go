package shared

// ParseConfig controls which features are enabled during parsing.
// A nil config is treated as All() — all features enabled.
type ParseConfig struct {
	// DetectUnknownFields enables detection and reporting of fields not
	// recognized as valid OpenAPI fields for their context.
	DetectUnknownFields bool

	// ResolveInternalRefs enables resolution of local JSON pointer $refs (#/...).
	ResolveInternalRefs bool

	// ResolveExternalRefs enables resolution of external file and URL $refs.
	ResolveExternalRefs bool

	// ApplySpecDefaults fills in OpenAPI-specified default values when optional
	// fields are absent (e.g., servers absent → [{ url: "/" }]).
	ApplySpecDefaults bool
}

// All returns a config with all features enabled. This is the default
// when no config (nil) is passed to a Parse function.
func All() *ParseConfig {
	return &ParseConfig{
		DetectUnknownFields: true,
		ResolveInternalRefs: true,
		ResolveExternalRefs: true,
		ApplySpecDefaults:   true,
	}
}

// None returns a config with all features disabled.
// Only basic parsing is performed — no unknown field detection,
// no $ref resolution, and no error collection on Trix.Errors.
func None() *ParseConfig {
	return &ParseConfig{}
}

// Defaults returns the provided config if non-nil, otherwise All().
func Defaults(cfg *ParseConfig) *ParseConfig {
	if cfg == nil {
		return All()
	}
	return cfg
}

// FirstConfig extracts the first non-nil config from a variadic list,
// falling back to All() if none provided.
func FirstConfig(cfgs []*ParseConfig) *ParseConfig {
	for _, cfg := range cfgs {
		if cfg != nil {
			return cfg
		}
	}
	return All()
}
