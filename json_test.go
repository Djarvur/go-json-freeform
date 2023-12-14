package freeform_test

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	freeform "github.com/Djarvur/go-json-freeform"
)

//go:embed testdata/good.json
var good []byte

func TestRaw(t *testing.T) {
	var raw freeform.Raw

	err := json.Unmarshal(good, &raw)
	require.NoError(t, err)

	var b bytes.Buffer
	c := json.NewEncoder(&b)
	c.SetIndent("", " ")

	err = c.Encode(raw)
	require.NoError(t, err)
	require.Equal(t, good, b.Bytes())

	require.Equal(t, "some string", raw.AsMap().Get("mess").AsMap().Get("0").AsList().Get(3).AsList().Get(0).AsString())
	require.Nil(t, raw.AsList())
	require.Nil(t, raw.AsMap().Get("number").AsMap())
	require.Equal(t, float64(0), raw.AsNumber())
	require.Equal(t, float64(1000), raw.AsMap().Get("number").AsNumber())
	require.Equal(t, freeform.Raw{}, raw.AsMap().Get("mess").AsMap().Get("0").AsList().Get(3).AsList().Get(10))
	require.Equal(t, freeform.Raw{}, raw.AsMap().Get("mess").AsMap().Get("nonexistent"))
	require.Nil(t, raw.AsMap().Get("mess").AsMap().Get("nonexistent").AsMap())
	require.Equal(t, freeform.Raw{}, raw.AsMap().Get("mess").AsMap().Get("nonexistent").AsMap().Get(""))
	require.Equal(t, "", raw.AsString())
}
