package shared

import (
	"bytes"
	"encoding/json"
	"reflect"
	"sort"

	"gopkg.in/yaml.v3"
)

// Field is a key-value pair for ordered marshalling.
type Field struct {
	Key   string
	Value interface{}
}

// MarshalFieldsJSON marshals an ordered list of fields as a JSON object.
// Fields whose values are considered empty (nil, zero-length, zero-value)
// are omitted.
func MarshalFieldsJSON(fields []Field) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	first := true
	for _, f := range fields {
		if isOmittable(f.Value) {
			continue
		}
		if !first {
			buf.WriteByte(',')
		}
		first = false
		key, err := json.Marshal(f.Key)
		if err != nil {
			return nil, err
		}
		buf.Write(key)
		buf.WriteByte(':')
		val, err := json.Marshal(f.Value)
		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// MarshalFieldsYAML builds a *yaml.Node mapping from an ordered list of fields.
// Fields whose values are considered empty are omitted.
func MarshalFieldsYAML(fields []Field) (*yaml.Node, error) {
	mapping := &yaml.Node{
		Kind: yaml.MappingNode,
		Tag:  "!!map",
	}
	for _, f := range fields {
		if isOmittable(f.Value) {
			continue
		}
		keyNode := &yaml.Node{
			Kind:  yaml.ScalarNode,
			Tag:   "!!str",
			Value: f.Key,
		}
		valNode := &yaml.Node{}
		if err := valNode.Encode(f.Value); err != nil {
			return nil, err
		}
		mapping.Content = append(mapping.Content, keyNode, valNode)
	}
	return mapping, nil
}

// AppendExtensions appends sorted vendor extension entries to a field list.
func AppendExtensions(fields []Field, extensions map[string]interface{}) []Field {
	if len(extensions) == 0 {
		return fields
	}
	keys := make([]string, 0, len(extensions))
	for k := range extensions {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fields = append(fields, Field{Key: k, Value: extensions[k]})
	}
	return fields
}

// isOmittable returns true if the value should be omitted from output.
func isOmittable(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface:
		return rv.IsNil()
	case reflect.String:
		return rv.Len() == 0
	case reflect.Slice, reflect.Map:
		return rv.IsNil() || rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	default:
		return false
	}
}
