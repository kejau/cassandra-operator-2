package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"encoding/json"
)

func Test_use_number_for_unmarshal(t *testing.T) {
	should := require.New(t)
	api := jsoniter.Config{UseNumber: true}.Froze()
	var obj interface{}
	should.Nil(api.UnmarshalFromString("123", &obj))
	should.Equal(json.Number("123"), obj)
}

func Test_customize_float_marshal(t *testing.T) {
	should := require.New(t)
	json := jsoniter.Config{MarshalFloatWith6Digits: true}.Froze()
	str, err := json.MarshalToString(float32(1.23456789))
	should.Nil(err)
	should.Equal("1.234568", str)
}


func Test_customize_tag_key(t *testing.T) {

	type TestObject struct {
		Field string `orm:"field"`
	}

	should := require.New(t)
	json := jsoniter.Config{TagKey: "orm"}.Froze()
	str, err := json.MarshalToString(TestObject{"hello"})
	should.Nil(err)
	should.Equal(`{"field":"hello"}`, str)
}