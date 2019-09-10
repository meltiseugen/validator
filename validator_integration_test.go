package validator_test

import (
	"testing"
	"time"
	"validator"
)

func TestValidator_ValidateAndInit2(t *testing.T) {
	type MyStruct struct {
		A string `validate:"required" datakey:"a"`
		B int    `validate:"int" datakey:"b"`
	}
	m := map[string]string{
		"a": "123",
		"b": "1",
	}

	s := MyStruct{}
	v := validator.New()
	err := v.ValidateAndInit(m, &s)
	if err != nil {
		t.Error()
	}
	if s.A != "123" || s.B != 1 {
		t.Error()
	}
}

func TestValidator_ValidateAndInit3(t *testing.T) {
	type InnerStruct struct {
		C int `datakey:"c" validate:"required,nothing,int"`
		D string `datakey:"d" validate:"nothing"`
	}
	type MyStruct struct {
		A time.Time `validate:"required,nothing,time" datakey:"a"`
		B int `validate:"required,int" datakey:"b"`
		IS InnerStruct
	}

	s := MyStruct{}
	m := map[string]string{
		"a": "2019-08-21T09:00:00Z",
		"b": "1",
		"c": "1",
		"d": "qwerty",
	}

	v := validator.New()
	_ = v.RegisterRule("nothing", func(mapKey string, m map[string]string, params ...string) error {
		return nil
	})
	err := v.ValidateAndInit(m, &s)
	if err != nil {
		t.Error()
	}

	time_, _ := time.Parse(time.RFC3339, "2019-08-21T09:00:00Z")
	if !s.A.Equal(time_) || s.B != 1 || s.IS.C != 1 || s.IS.D != "qwerty" {
		t.Error()
	}
}