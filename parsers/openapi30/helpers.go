package openapi30

// getString retrieves a string value from a map, returning empty string if not found or wrong type.
func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// getStringPtr retrieves a string pointer from a map, returning nil if not found.
func getStringPtr(m map[string]interface{}, key string) *string {
	if v, ok := m[key]; ok {
		if s, ok := v.(string); ok {
			return &s
		}
	}
	return nil
}

// getBool retrieves a bool value from a map, returning false if not found or wrong type.
func getBool(m map[string]interface{}, key string) bool {
	if v, ok := m[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}

// getBoolPtr retrieves a bool pointer from a map, returning nil if not found.
func getBoolPtr(m map[string]interface{}, key string) *bool {
	if v, ok := m[key]; ok {
		if b, ok := v.(bool); ok {
			return &b
		}
	}
	return nil
}

// getInt retrieves an int value from a map, returning 0 if not found or wrong type.
func getInt(m map[string]interface{}, key string) int {
	if v, ok := m[key]; ok {
		switch n := v.(type) {
		case int:
			return n
		case int64:
			return int(n)
		case float64:
			return int(n)
		}
	}
	return 0
}

// getFloat64 retrieves a float64 value from a map, returning 0 if not found or wrong type.
func getFloat64(m map[string]interface{}, key string) float64 {
	if v, ok := m[key]; ok {
		switch n := v.(type) {
		case float64:
			return n
		case int:
			return float64(n)
		case int64:
			return float64(n)
		}
	}
	return 0
}

// getFloat64Ptr retrieves a float64 pointer from a map, returning nil if not found.
func getFloat64Ptr(m map[string]interface{}, key string) *float64 {
	if v, ok := m[key]; ok {
		var f float64
		switch n := v.(type) {
		case float64:
			f = n
		case int:
			f = float64(n)
		case int64:
			f = float64(n)
		default:
			return nil
		}
		return &f
	}
	return nil
}

// getUint64Ptr retrieves a uint64 pointer from a map, returning nil if not found.
func getUint64Ptr(m map[string]interface{}, key string) *uint64 {
	if v, ok := m[key]; ok {
		var u uint64
		switch n := v.(type) {
		case float64:
			if n >= 0 {
				u = uint64(n)
			} else {
				return nil
			}
		case int:
			if n >= 0 {
				u = uint64(n)
			} else {
				return nil
			}
		case int64:
			if n >= 0 {
				u = uint64(n)
			} else {
				return nil
			}
		case uint64:
			u = n
		default:
			return nil
		}
		return &u
	}
	return nil
}

// getMap retrieves a map value from a map, returning nil if not found or wrong type.
func getMap(m map[string]interface{}, key string) map[string]interface{} {
	if v, ok := m[key]; ok {
		if sub, ok := v.(map[string]interface{}); ok {
			return sub
		}
	}
	return nil
}

// getSlice retrieves a slice value from a map, returning nil if not found or wrong type.
func getSlice(m map[string]interface{}, key string) []interface{} {
	if v, ok := m[key]; ok {
		if s, ok := v.([]interface{}); ok {
			return s
		}
	}
	return nil
}

// getStringSlice retrieves a string slice from a map, returning nil if not found.
func getStringSlice(m map[string]interface{}, key string) []string {
	slice := getSlice(m, key)
	if slice == nil {
		return nil
	}

	result := make([]string, 0, len(slice))
	for _, v := range slice {
		if s, ok := v.(string); ok {
			result = append(result, s)
		}
	}
	return result
}

// getAny retrieves any value from a map, returning nil if not found.
func getAny(m map[string]interface{}, key string) interface{} {
	return m[key]
}

// hasKey checks if a key exists in the map.
func hasKey(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}

// hasRef checks if the map contains a $ref key.
func hasRef(m map[string]interface{}) bool {
	_, ok := m["$ref"]
	return ok
}

// getRef retrieves the $ref value from a map, returning empty string if not found.
func getRef(m map[string]interface{}) string {
	return getString(m, "$ref")
}

// parseExtensions extracts extension fields (x-*) from a map.
func parseExtensions(m map[string]interface{}) map[string]interface{} {
	var extensions map[string]interface{}

	for key, value := range m {
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			if extensions == nil {
				extensions = make(map[string]interface{})
			}
			extensions[key] = value
		}
	}

	return extensions
}

// getStringMap retrieves a map[string]string from a map, returning nil if not found.
func getStringMap(m map[string]interface{}, key string) map[string]string {
	sub := getMap(m, key)
	if sub == nil {
		return nil
	}

	result := make(map[string]string)
	for k, v := range sub {
		if s, ok := v.(string); ok {
			result[k] = s
		}
	}
	return result
}

// getInterfaceSlice retrieves a slice of interface{} from a map, returning nil if not found.
func getInterfaceSlice(m map[string]interface{}, key string) []interface{} {
	return getSlice(m, key)
}

// itoa converts an int to a string (simple implementation for indices).
func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	if i < 0 {
		return "-" + itoa(-i)
	}
	var digits []byte
	for i > 0 {
		digits = append([]byte{byte('0' + i%10)}, digits...)
		i /= 10
	}
	return string(digits)
}
