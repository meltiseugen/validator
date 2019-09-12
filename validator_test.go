package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestValidator_New(t *testing.T) {
	v := New()
	if v == nil {
		t.Error()
	} else {
		if len(v.converterMappings) != 4 {
			t.Error()
		}
		if v.isInit != true {
			t.Error()
		}
	}
}

func TestValidator_RegisterRule2(t *testing.T) {
	v := Validator{}
	err := v.RegisterRule("", func(mapKey string, m map[string]string, params ...string) error {
		return nil
	})
	if err == nil {
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

func TestValidator_RegisterConverter2(t *testing.T) {
	v := Validator{}
	err := v.RegisterConverter("", func(value string, params ...string) (i interface{}, e error) {
		return nil, nil
	})
	if err == nil {
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

func TestValidator_checkRules2(t *testing.T) {
	type InnerStruct struct {
		C int `validate:"int" datakey:"c"`
	}
	type MyStruct struct {
		A string `validate:"required" datakey:"a"`
		B int    `validate:"int" datakey:"b"`
		IS InnerStruct
	}
	testdata := []struct {
		m           map[string]string
		noErrorFlag bool
	}{
		{
			map[string]string{
				"a": "123",
				"b": "1",
				"c": "2",
			},
			true,
		},
		{
			map[string]string{
				"a": "123",
				"b": "1",
				"c": "asdf",
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

func TestValidator_checkRules3(t *testing.T) {
	type MyStruct struct {
		A string `validate:"required"`
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

func TestValidator_checkRules4(t *testing.T) {
	type MyStruct struct {
		A string `validate:"required, newRule" datakey:"a"`
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

func TestValidator_initData2(t *testing.T) {
	type InnerStruct struct {
		C int `validate:"int" datakey:"c"`
	}
	type MyStruct struct {
		A string `validate:"required" datakey:"a"`
		B int    `validate:"int" datakey:"b"`
		IS InnerStruct
	}
	testdata := []struct {
		m           map[string]string
		noErrorFlag bool
	}{
		{
			map[string]string{
				"a": "123",
				"b": "1",
				"c": "2",
			},
			true,
		},
		{
			map[string]string{
				"a": "123",
				"b": "1",
				"c": "asdf",
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

func TestValidator_initData3(t *testing.T) {
	type MyStruct struct {
		A bool `validate:"required" datakey:"a"`
		B int    `validate:"int" datakey:"b"`
	}
	testdata := []struct {
		m           map[string]string
		noErrorFlag bool
	}{
		{
			map[string]string{
				"a": "TRUE",
				"b": "1",
			},
			true,
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

func TestValidator_initData4(t *testing.T) {
	type NewType struct {
		C string
	}
	type MyStruct struct {
		A bool `validate:"required" datakey:"a"`
		B int    `validate:"int" datakey:"b"`
		N NewType `datakey:"c"`
	}
	testdata := []struct {
		m           map[string]string
		noErrorFlag bool
	}{
		{
			map[string]string{
				"a": "TRUE",
				"b": "1",
				"c": "asdaf",
			},
			false,
		},
	}

	v := New()
	for i, td := range testdata {
		t.Run("TestCheckTime_"+strconv.Itoa(i), func(t *testing.T) {
			err := v.initData(td.m, reflect.ValueOf(&MyStruct{}).Elem())
			if td.noErrorFlag && err != nil {
				fmt.Println(err)
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

func TestValidator_ValidateAndInit2(t *testing.T) {
	type MyStruct struct {
		A string `validate:"required" datakey:"a"`
		B int    `validate:"int" datakey:"b"`
	}
	m := map[string]string{
		"a": "asdfsa",
		"b": "asdf",
	}

	s := MyStruct{}
	i := 1
	v := New()

	err := v.ValidateAndInit(m, s)
	if err == nil {
		t.Error()
	}

	err = v.ValidateAndInit(m, &i)
	if err == nil {
		t.Error()
	}
}

func TestValidator_ValidateAndInit3(t *testing.T) {
	type MyStruct struct {
		A string `validate:"required" datakey:"a"`
		B int    `validate:"int" datakey:"b"`
	}
	m := map[string]string{
		"a": "asdfsa",
		"b": "asdf",
	}

	s := MyStruct{}
	v := Validator{}

	err := v.ValidateAndInit(m, &s)
	if err == nil {
		t.Error()
	}
}

func TestValidator_ValidateAndInit4(t *testing.T) {
	type MyStruct struct {
		A string `validate:"required" datakey:"a"`
		B bool    `validate:"" datakey:"b"`
	}
	testdata := []struct {
		m           map[string]string
		noErrorFlag bool
	}{
		{
			map[string]string{
				"a": "123",
				"b": "sdasD",
			},
			false,
		},
	}

	v := New()
	for i, td := range testdata {
		t.Run("TestCheckTime_"+strconv.Itoa(i), func(t *testing.T) {
			s := MyStruct{}
			err := v.ValidateAndInit(td.m, &s)
			if td.noErrorFlag && err != nil {
				t.Error()
			} else if !td.noErrorFlag && err == nil {
				t.Error()
			}
		})
	}
}