package openapi20

import "github.com/apitrix/openapi-parser/parsers/shared"

// Thin wrappers delegating to shared.Get* / shared.Has* / shared.Parse* functions.
// This preserves the unexported API so all callers within the package work unchanged.

func getString(m map[string]interface{}, key string) string {
	return shared.GetString(m, key)
}

func getStringPtr(m map[string]interface{}, key string) *string {
	return shared.GetStringPtr(m, key)
}

func getBool(m map[string]interface{}, key string) bool {
	return shared.GetBool(m, key)
}

func getBoolPtr(m map[string]interface{}, key string) *bool {
	return shared.GetBoolPtr(m, key)
}

func getInt(m map[string]interface{}, key string) int {
	return shared.GetInt(m, key)
}

func getFloat64(m map[string]interface{}, key string) float64 {
	return shared.GetFloat64(m, key)
}

func getFloat64Ptr(m map[string]interface{}, key string) *float64 {
	return shared.GetFloat64Ptr(m, key)
}

func getUint64Ptr(m map[string]interface{}, key string) *uint64 {
	return shared.GetUint64Ptr(m, key)
}

func getMap(m map[string]interface{}, key string) map[string]interface{} {
	return shared.GetMap(m, key)
}

func getSlice(m map[string]interface{}, key string) []interface{} {
	return shared.GetSlice(m, key)
}

func getStringSlice(m map[string]interface{}, key string) []string {
	return shared.GetStringSlice(m, key)
}

func getAny(m map[string]interface{}, key string) interface{} {
	return shared.GetAny(m, key)
}

func hasKey(m map[string]interface{}, key string) bool {
	return shared.HasKey(m, key)
}

func hasRef(m map[string]interface{}) bool {
	return shared.HasRef(m)
}

func getRef(m map[string]interface{}) string {
	return shared.GetRef(m)
}

func parseExtensions(m map[string]interface{}) map[string]interface{} {
	return shared.ParseExtensions(m)
}

func getStringMap(m map[string]interface{}, key string) map[string]string {
	return shared.GetStringMap(m, key)
}

func getInterfaceSlice(m map[string]interface{}, key string) []interface{} {
	return shared.GetInterfaceSlice(m, key)
}

func itoa(i int) string {
	return shared.Itoa(i)
}
