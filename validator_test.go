package validator

import (
	"reflect"
	"strconv"
	"testing"
)

func TestValidator_New(t *testing.T) {
	v := New()
	if &v == nil {
		t.Error()
	}
	if len(v.ruleMappings) != 4 {
		t.Error()
	}
	if len(v.converterMappings) != 3 {
		t.Error()
	}
	if v.isInit != true {
		t.Error()
	}
}

func TestValidator_RegisterRule(t *testing.T) {
	v := New()
	err := v.RegisterRule("", func(mapKey string, m map[string]string, params ...string) error {
		return nil
	})
	if err == nil {
		t.Error()
	}

	err = v.RegisterRule("rule", func(mapKey string, m map[string]string, params ...string) error {
		return nil
	})
	if err != nil {
		t.Error()
	}
	if _, ok := v.ruleMappings["rule"]; !ok {
		t.Error()
	}
}

func TestValidator_RegisterConverter(t *testing.T) {
	v := New()
	err := v.RegisterConverter("", func(value string, params ...string) (i interface{}, e error) {
		return nil, nil
	})
	if err == nil {
		t.Error()
	}

	err = v.RegisterConverter("MyType", func(value string, params ...string) (i interface{}, e error) {
		return nil, nil
	})
	if err != nil {
		t.Error()
	}
	if _, ok := v.converterMappings["MyType"]; !ok {
		t.Error()
	}
}

func TestValidator_checkRules(t *testing.T) {
	type MyStruct struct {
		A string `validate:"required" datakey:"a"`
		B int    `validate:"int" datakey:"b"`
	}
	testdata := []struct {
		m           map[string]string
		noErrorFlag bool
	}{
		{
			map[string]string{
				"a": "123",
				"b": "1",
			},
			true,
		},
		{
			map[string]string{
				"a": "asdfsa",
				"b": "1",
			},
			true,
		},
		{
			map[string]string{
				"a": "asdfsa",
				"b": "asdf",
			},
			false,
		},
		{
			map[string]string{
				"b": "asdf",
			},
			false,
		},
	}

	v := New()
	for i, td := range testdata {
		t.Run("TestCheckTime_"+strconv.Itoa(i), func(t *testing.T) {
			err := v.checkRules(td.m, reflect.ValueOf(&MyStruct{}).Elem())
			if td.noErrorFlag && err != nil {
				t.Error()
			} else if !td.noErrorFlag && err == nil {
				t.Error()
			}
		})
	}
}

func TestValidator_initData(t *testing.T) {
	type MyStruct struct {
		A string `validate:"required" datakey:"a"`
		B int    `validate:"int" datakey:"b"`
	}
	testdata := []struct {
		m           map[string]string
		noErrorFlag bool
	}{
		{
			map[string]string{
				"a": "123",
				"b": "1",
			},
			true,
		},
		{
			map[string]string{
				"a": "asdfsa",
				"b": "1",
			},
			true,
		},
		{
			map[string]string{
				"a": "asdfsa",
				"b": "asdf",
			},
			false,
		},
		{
			map[string]string{
				"b": "asdf",
			},
			false,
		},
	}

	v := New()
	for i, td := range testdata {
		t.Run("TestCheckTime_"+strconv.Itoa(i), func(t *testing.T) {
			err := v.initData(td.m, reflect.ValueOf(&MyStruct{}).Elem())
			if td.noErrorFlag && err != nil {
				t.Error()
			} else if !td.noErrorFlag && err == nil {
				t.Error()
			}
		})
	}
}

func TestValidator_ValidateAndInit(t *testing.T) {
	type MyStruct struct {
		A string `validate:"required" datakey:"a"`
		B int    `validate:"int" datakey:"b"`
	}
	testdata := []struct {
		m           map[string]string
		noErrorFlag bool
	}{
		{
			map[string]string{
				"a": "123",
				"b": "1",
			},
			true,
		},
		{
			map[string]string{
				"a": "asdfsa",
				"b": "1",
			},
			true,
		},
		{
			map[string]string{
				"a": "asdfsa",
				"b": "asdf",
			},
			false,
		},
		{
			map[string]string{
				"b": "asdf",
			},
			false,
		},
	}

	v := New()
	for i, td := range testdata {
		t.Run("TestCheckTime_"+strconv.Itoa(i), func(t *testing.T) {
			err := v.ValidateAndInit(td.m, &MyStruct{})
			if td.noErrorFlag && err != nil {
				t.Error()
			} else if !td.noErrorFlag && err == nil {
				t.Error()
			}
		})
	}
}
