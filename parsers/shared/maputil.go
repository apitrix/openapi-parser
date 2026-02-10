package shared

// GetString retrieves a string value from a map, returning empty string if not found or wrong type.
func GetString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// GetStringPtr retrieves a string pointer from a map, returning nil if not found.
func GetStringPtr(m map[string]interface{}, key string) *string {
	if v, ok := m[key]; ok {
		if s, ok := v.(string); ok {
			return &s
		}
	}
	return nil
}

// GetBool retrieves a bool value from a map, returning false if not found or wrong type.
func GetBool(m map[string]interface{}, key string) bool {
	if v, ok := m[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}

// GetBoolPtr retrieves a bool pointer from a map, returning nil if not found.
func GetBoolPtr(m map[string]interface{}, key string) *bool {
	if v, ok := m[key]; ok {
		if b, ok := v.(bool); ok {
			return &b
		}
	}
	return nil
}

// GetInt retrieves an int value from a map, returning 0 if not found or wrong type.
func GetInt(m map[string]interface{}, key string) int {
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

// GetFloat64 retrieves a float64 value from a map, returning 0 if not found or wrong type.
func GetFloat64(m map[string]interface{}, key string) float64 {
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

// GetFloat64Ptr retrieves a float64 pointer from a map, returning nil if not found.
func GetFloat64Ptr(m map[string]interface{}, key string) *float64 {
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

// GetUint64Ptr retrieves a uint64 pointer from a map, returning nil if not found.
func GetUint64Ptr(m map[string]interface{}, key string) *uint64 {
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

// GetMap retrieves a map value from a map, returning nil if not found or wrong type.
func GetMap(m map[string]interface{}, key string) map[string]interface{} {
	if v, ok := m[key]; ok {
		if sub, ok := v.(map[string]interface{}); ok {
			return sub
		}
	}
	return nil
}

// GetSlice retrieves a slice value from a map, returning nil if not found or wrong type.
func GetSlice(m map[string]interface{}, key string) []interface{} {
	if v, ok := m[key]; ok {
		if s, ok := v.([]interface{}); ok {
			return s
		}
	}
	return nil
}

// GetStringSlice retrieves a string slice from a map, returning nil if not found.
func GetStringSlice(m map[string]interface{}, key string) []string {
	slice := GetSlice(m, key)
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

// GetAny retrieves any value from a map, returning nil if not found.
func GetAny(m map[string]interface{}, key string) interface{} {
	return m[key]
}

// HasKey checks if a key exists in the map.
func HasKey(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}

// HasRef checks if the map contains a $ref key.
func HasRef(m map[string]interface{}) bool {
	_, ok := m["$ref"]
	return ok
}

// GetRef retrieves the $ref value from a map, returning empty string if not found.
func GetRef(m map[string]interface{}) string {
	return GetString(m, "$ref")
}

// ParseExtensions extracts extension fields (x-*) from a map.
func ParseExtensions(m map[string]interface{}) map[string]interface{} {
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

// GetStringMap retrieves a map[string]string from a map, returning nil if not found.
func GetStringMap(m map[string]interface{}, key string) map[string]string {
	sub := GetMap(m, key)
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

// GetInterfaceSlice retrieves a slice of interface{} from a map, returning nil if not found.
func GetInterfaceSlice(m map[string]interface{}, key string) []interface{} {
	return GetSlice(m, key)
}

// Itoa converts an int to a string (simple implementation for indices).
func Itoa(i int) string {
	if i == 0 {
		return "0"
	}
	if i < 0 {
		return "-" + Itoa(-i)
	}
	var digits []byte
	for i > 0 {
		digits = append([]byte{byte('0' + i%10)}, digits...)
		i /= 10
	}
	return string(digits)
}
