package json2

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

type Model struct {
	Value      string
	IntValue   int
	FloatValue float64
	ArrayValue []int
}

type ModelWithStruct struct {
	Value       string
	WithModel   Model
	StructArray []Model
}

func TestSimpleEncodeDecode(t *testing.T) {
	model := Model{
		Value:"UpperCase",
	}

	buf, err := Marshal(model);
	if !assert.Nil(t, err) {
		fmt.Printf("Marshal Error%s\n", err.Error());
		return;
	}

	fmt.Printf("JSON = %s\n", string(buf))

	decode := Model{};

	err = Unmarshal(buf, &decode);
	if !assert.Nil(t, err) {
		fmt.Printf("Unmarshal Error%s\n", err.Error());
		return;
	}

	assert.Equal(t, model.Value, decode.Value);
}

func TestStructEncodeDecode(t *testing.T) {
	model := ModelWithStruct{
		Value:"UpperCase",
		WithModel:Model{
			Value:"InStruct",
		},
	}

	buf, err := Marshal(model);
	if !assert.Nil(t, err) {
		fmt.Printf("Marshal Error%s\n", err.Error());
		return;
	}

	fmt.Printf("JSON = %s\n", string(buf))

	decode := ModelWithStruct{};

	err = Unmarshal(buf, &decode);
	if !assert.Nil(t, err) {
		fmt.Printf("Unmarshal Error%s\n", err.Error());
		return;
	}

	assert.Equal(t, model.Value, decode.Value);
}
