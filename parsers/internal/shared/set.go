package shared

// ToSet converts a slice of field names to a set for O(1) lookup.
func ToSet(fields []string) map[string]struct{} {
	m := make(map[string]struct{}, len(fields))
	for _, f := range fields {
		m[f] = struct{}{}
	}
	return m
}
