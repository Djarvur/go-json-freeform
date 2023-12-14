// Package freeform go_json_freeform is a simply and mostly useless package
// for (un)marshaling JSON docs and reading them without creating a struct to represent them.
package freeform

import (
	"encoding/json"
)

var (
	_ json.Marshaler   = (*Raw)(nil)
	_ json.Unmarshaler = (*Raw)(nil)
)

// Raw is a container for the free-form unmarshalled struct.
type Raw struct {
	payload interface{}
}

// UnmarshalJSON is required to unmarshal Raw properly.
func (r *Raw) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &r.payload) //nolint:wrapcheck
}

// MarshalJSON is required to marshal Raw properly.
func (r Raw) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.payload) //nolint:wrapcheck
}

// AsString returns a string contained in the Raw or empty string in case there is not a string.
func (r Raw) AsString() string {
	v, _ := r.payload.(string)

	return v
}

// AsNumber returns a float64 contained in the Raw or 0 in case there is not a number.
func (r Raw) AsNumber() float64 {
	v, _ := r.payload.(float64)

	return v
}

// AsBool returns a boolean contained in the Raw or false in case there is not a boolean.
func (r Raw) AsBool() bool {
	v, _ := r.payload.(bool)

	return v
}

// AsList returns an array (actually slice) contained in the Raw or nil in case there is not an array.
func (r Raw) AsList() List {
	l, ok := r.payload.([]interface{})
	if !ok {
		return nil
	}

	out := make(List, 0, len(l))

	for _, v := range l {
		out = append(out, Raw{payload: v})
	}

	return out
}

// List represents a json array.
type List []Raw

// Get returns a corresponding value from the list for the provided index
// or the empty Raw in case if invalid index or even the list is nil.
func (l List) Get(i int) Raw {
	if i >= len(l) || i < 0 {
		return Raw{} //nolint:exhaustruct
	}

	return l[i]
}

// AsMap returns an object (actually map) contained in the Raw
// or nil in case there is not an object/map.
func (r Raw) AsMap() Map {
	m, ok := r.payload.(map[string]interface{})
	if !ok {
		return nil
	}

	out := make(Map, len(m))

	for k, v := range m {
		out[k] = Raw{payload: v}
	}

	return out
}

// Map represents a json object.
type Map map[string]Raw

// Get returns a corresponding value from the map for the provided key
// or the empty Raw in case no such key exists or even the map is nil.
func (m Map) Get(key string) Raw {
	if m == nil {
		return Raw{} //nolint:exhaustruct
	}

	return m[key]
}
