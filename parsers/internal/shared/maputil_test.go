package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetString(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]interface{}
		key  string
		want string
	}{
		{"found", map[string]interface{}{"k": "v"}, "k", "v"},
		{"missing key", map[string]interface{}{"k": "v"}, "other", ""},
		{"wrong type", map[string]interface{}{"k": 42}, "k", ""},
		{"nil map", nil, "k", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetString(tt.m, tt.key))
		})
	}
}

func TestGetStringPtr(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		got := GetStringPtr(map[string]interface{}{"k": "hello"}, "k")
		require.NotNil(t, got)
		assert.Equal(t, "hello", *got)
	})
	t.Run("missing", func(t *testing.T) {
		assert.Nil(t, GetStringPtr(map[string]interface{}{}, "k"))
	})
	t.Run("wrong type", func(t *testing.T) {
		assert.Nil(t, GetStringPtr(map[string]interface{}{"k": 42}, "k"))
	})
}

func TestGetBool(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]interface{}
		want bool
	}{
		{"true", map[string]interface{}{"k": true}, true},
		{"false", map[string]interface{}{"k": false}, false},
		{"missing", map[string]interface{}{}, false},
		{"wrong type", map[string]interface{}{"k": "true"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetBool(tt.m, "k"))
		})
	}
}

func TestGetBoolPtr(t *testing.T) {
	t.Run("found true", func(t *testing.T) {
		got := GetBoolPtr(map[string]interface{}{"k": true}, "k")
		require.NotNil(t, got)
		assert.True(t, *got)
	})
	t.Run("found false", func(t *testing.T) {
		got := GetBoolPtr(map[string]interface{}{"k": false}, "k")
		require.NotNil(t, got)
		assert.False(t, *got)
	})
	t.Run("missing", func(t *testing.T) {
		assert.Nil(t, GetBoolPtr(map[string]interface{}{}, "k"))
	})
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]interface{}
		want int
	}{
		{"int", map[string]interface{}{"k": 42}, 42},
		{"int64", map[string]interface{}{"k": int64(100)}, 100},
		{"float64", map[string]interface{}{"k": float64(7)}, 7},
		{"missing", map[string]interface{}{}, 0},
		{"wrong type", map[string]interface{}{"k": "not a number"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetInt(tt.m, "k"))
		})
	}
}

func TestGetFloat64(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]interface{}
		want float64
	}{
		{"float64", map[string]interface{}{"k": 3.14}, 3.14},
		{"int", map[string]interface{}{"k": 5}, float64(5)},
		{"int64", map[string]interface{}{"k": int64(99)}, float64(99)},
		{"missing", map[string]interface{}{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.InDelta(t, tt.want, GetFloat64(tt.m, "k"), 0.001)
		})
	}
}

func TestGetFloat64Ptr(t *testing.T) {
	t.Run("float64", func(t *testing.T) {
		got := GetFloat64Ptr(map[string]interface{}{"k": 2.5}, "k")
		require.NotNil(t, got)
		assert.InDelta(t, 2.5, *got, 0.001)
	})
	t.Run("int", func(t *testing.T) {
		got := GetFloat64Ptr(map[string]interface{}{"k": 10}, "k")
		require.NotNil(t, got)
		assert.InDelta(t, 10.0, *got, 0.001)
	})
	t.Run("int64", func(t *testing.T) {
		got := GetFloat64Ptr(map[string]interface{}{"k": int64(20)}, "k")
		require.NotNil(t, got)
		assert.InDelta(t, 20.0, *got, 0.001)
	})
	t.Run("wrong type", func(t *testing.T) {
		assert.Nil(t, GetFloat64Ptr(map[string]interface{}{"k": "text"}, "k"))
	})
	t.Run("missing", func(t *testing.T) {
		assert.Nil(t, GetFloat64Ptr(map[string]interface{}{}, "k"))
	})
}

func TestGetUint64Ptr(t *testing.T) {
	t.Run("float64 positive", func(t *testing.T) {
		got := GetUint64Ptr(map[string]interface{}{"k": float64(42)}, "k")
		require.NotNil(t, got)
		assert.Equal(t, uint64(42), *got)
	})
	t.Run("float64 negative", func(t *testing.T) {
		assert.Nil(t, GetUint64Ptr(map[string]interface{}{"k": float64(-1)}, "k"))
	})
	t.Run("int positive", func(t *testing.T) {
		got := GetUint64Ptr(map[string]interface{}{"k": 5}, "k")
		require.NotNil(t, got)
		assert.Equal(t, uint64(5), *got)
	})
	t.Run("int negative", func(t *testing.T) {
		assert.Nil(t, GetUint64Ptr(map[string]interface{}{"k": -3}, "k"))
	})
	t.Run("int64 positive", func(t *testing.T) {
		got := GetUint64Ptr(map[string]interface{}{"k": int64(99)}, "k")
		require.NotNil(t, got)
		assert.Equal(t, uint64(99), *got)
	})
	t.Run("int64 negative", func(t *testing.T) {
		assert.Nil(t, GetUint64Ptr(map[string]interface{}{"k": int64(-7)}, "k"))
	})
	t.Run("uint64", func(t *testing.T) {
		got := GetUint64Ptr(map[string]interface{}{"k": uint64(1000)}, "k")
		require.NotNil(t, got)
		assert.Equal(t, uint64(1000), *got)
	})
	t.Run("wrong type", func(t *testing.T) {
		assert.Nil(t, GetUint64Ptr(map[string]interface{}{"k": "text"}, "k"))
	})
	t.Run("missing", func(t *testing.T) {
		assert.Nil(t, GetUint64Ptr(map[string]interface{}{}, "k"))
	})
}

func TestGetMap(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		sub := map[string]interface{}{"nested": "val"}
		got := GetMap(map[string]interface{}{"k": sub}, "k")
		require.NotNil(t, got)
		assert.Equal(t, "val", got["nested"])
	})
	t.Run("missing", func(t *testing.T) {
		assert.Nil(t, GetMap(map[string]interface{}{}, "k"))
	})
	t.Run("wrong type", func(t *testing.T) {
		assert.Nil(t, GetMap(map[string]interface{}{"k": "string"}, "k"))
	})
}

func TestGetSlice(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		got := GetSlice(map[string]interface{}{"k": []interface{}{"a", "b"}}, "k")
		assert.Len(t, got, 2)
	})
	t.Run("missing", func(t *testing.T) {
		assert.Nil(t, GetSlice(map[string]interface{}{}, "k"))
	})
}

func TestGetStringSlice(t *testing.T) {
	t.Run("all strings", func(t *testing.T) {
		got := GetStringSlice(map[string]interface{}{"k": []interface{}{"a", "b", "c"}}, "k")
		assert.Equal(t, []string{"a", "b", "c"}, got)
	})
	t.Run("mixed types skips non-string", func(t *testing.T) {
		got := GetStringSlice(map[string]interface{}{"k": []interface{}{"a", 42, "c"}}, "k")
		assert.Equal(t, []string{"a", "c"}, got)
	})
	t.Run("missing", func(t *testing.T) {
		assert.Nil(t, GetStringSlice(map[string]interface{}{}, "k"))
	})
}

func TestGetAny(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		assert.Equal(t, 42, GetAny(map[string]interface{}{"k": 42}, "k"))
	})
	t.Run("missing", func(t *testing.T) {
		assert.Nil(t, GetAny(map[string]interface{}{}, "k"))
	})
}

func TestHasKey(t *testing.T) {
	m := map[string]interface{}{"present": true}
	assert.True(t, HasKey(m, "present"))
	assert.False(t, HasKey(m, "absent"))
}

func TestHasRef(t *testing.T) {
	assert.True(t, HasRef(map[string]interface{}{"$ref": "#/schemas/Pet"}))
	assert.False(t, HasRef(map[string]interface{}{"type": "object"}))
}

func TestGetRef(t *testing.T) {
	assert.Equal(t, "#/defs/Pet", GetRef(map[string]interface{}{"$ref": "#/defs/Pet"}))
	assert.Equal(t, "", GetRef(map[string]interface{}{}))
}

func TestParseExtensions(t *testing.T) {
	t.Run("has extensions", func(t *testing.T) {
		m := map[string]interface{}{"type": "object", "x-custom": "hello", "x-internal": true}
		got := ParseExtensions(m)
		assert.Len(t, got, 2)
		assert.Equal(t, "hello", got["x-custom"])
	})
	t.Run("no extensions", func(t *testing.T) {
		assert.Nil(t, ParseExtensions(map[string]interface{}{"type": "string"}))
	})
	t.Run("short x key not extension", func(t *testing.T) {
		assert.Nil(t, ParseExtensions(map[string]interface{}{"x": "not extension"}))
	})
}

func TestGetStringMap(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		m := map[string]interface{}{"k": map[string]interface{}{"a": "1", "b": "2"}}
		got := GetStringMap(m, "k")
		assert.Equal(t, map[string]string{"a": "1", "b": "2"}, got)
	})
	t.Run("missing", func(t *testing.T) {
		assert.Nil(t, GetStringMap(map[string]interface{}{}, "k"))
	})
	t.Run("mixed values keeps only strings", func(t *testing.T) {
		m := map[string]interface{}{"k": map[string]interface{}{"a": "str", "b": 42}}
		got := GetStringMap(m, "k")
		assert.Len(t, got, 1)
		assert.Equal(t, "str", got["a"])
	})
}

func TestGetInterfaceSlice(t *testing.T) {
	got := GetInterfaceSlice(map[string]interface{}{"k": []interface{}{1, "two", 3.0}}, "k")
	assert.Len(t, got, 3)
}

func TestItoa(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{0, "0"}, {1, "1"}, {42, "42"}, {100, "100"}, {-5, "-5"}, {-123, "-123"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			assert.Equal(t, tt.want, Itoa(tt.input))
		})
	}
}
