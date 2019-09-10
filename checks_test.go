package validator

import (
	"strconv"
	"testing"
)

func TestChecks_checkRequired(t *testing.T) {
	var testdata = []struct {
		key         string
		in          map[string]string
		noErrorFlag bool
	}{
		{"a",
			map[string]string{
				"a": "A",
			},
			true},
		{"c",
			map[string]string{
				"b": "B",
			},
			false},
	}

	for i, td := range testdata {
		t.Run("TestCheckRequired_"+strconv.Itoa(i), func(t *testing.T) {
			err := checkRequired(td.key, td.in)
			if td.noErrorFlag && err != nil {
				t.Error()
			} else if !td.noErrorFlag && err == nil {
				t.Error()
			}
		})
	}
}

func TestChecks_checkInt(t *testing.T) {
	var testdata = []struct {
		key         string
		in          map[string]string
		noErrorFlag bool
	}{
		{"a",
			map[string]string{
				"a": "1234",
			},
			true},
		{"a",
			map[string]string{
				"a": "asdf",
			},
			false},
		{"a",
			map[string]string{
				"a": "123wer",
			},
			false},
		{"a",
			map[string]string{
				"a": "-1235",
			},
			true},
		{"a",
			map[string]string{
				"a": "3.05",
			},
			false},
	}

	for i, td := range testdata {
		t.Run("TestCheckInt_"+strconv.Itoa(i), func(t *testing.T) {
			err := checkInt(td.key, td.in)
			if td.noErrorFlag && err != nil {
				t.Error()
			} else if !td.noErrorFlag && err == nil {
				t.Error()
			}
		})
	}
}

func TestChecks_checkUnsigned(t *testing.T) {
	var testdata = []struct {
		key         string
		in          map[string]string
		noErrorFlag bool
	}{
		{"a",
			map[string]string{
				"a": "1234",
			},
			true},
		{"a",
			map[string]string{
				"a": "-1234",
			},
			false},
		{"a",
			map[string]string{
				"a": "1234asdas",
			},
			false},
		{"a",
			map[string]string{
				"a": "dasdaw",
			},
			false},
		{"a",
			map[string]string{
				"a": "9.03",
			},
			false},
		{"a",
			map[string]string{
				"a": "-9.035",
			},
			false},
	}

	for i, td := range testdata {
		t.Run("TestCheckUnsigned_"+strconv.Itoa(i), func(t *testing.T) {
			err := checkUnsigned(td.key, td.in)
			if td.noErrorFlag && err != nil {
				t.Error()
			} else if !td.noErrorFlag && err == nil {
				t.Error()
			}
		})
	}
}

func TestChecks_checkTime(t *testing.T) {
	var testdata = []struct {
		key         string
		in          map[string]string
		noErrorFlag bool
	}{
		{"a",
			map[string]string{
				"a": "2019-08-21T09:00:00Z",
			},
			true},
		{"a",
			map[string]string{
				"a": "201900:00Z",
			},
			false},
		{"a",
			map[string]string{
				"a": "2019-08-21T09:00:00",
			},
			false},
		{"a",
			map[string]string{
				"a": "2019-20-21T09:00:00",
			},
			false},
		{"a",
			map[string]string{
				"a": "2019-08-21T30:00:00",
			},
			false},
		{"a",
			map[string]string{
				"a": "2019-08-21 09:00:00",
			},
			false},
	}

	for i, td := range testdata {
		t.Run("TestCheckTime_"+strconv.Itoa(i), func(t *testing.T) {
			err := checkTime(td.key, td.in)
			if td.noErrorFlag && err != nil {
				t.Error()
			} else if !td.noErrorFlag && err == nil {
				t.Error()
			}
		})
	}
}
