package validator

import (
	"strconv"
	"testing"
	"time"
)

func TestConverters_convertToString(t *testing.T) {
	testdata := []struct {
		in          string
		out         interface{}
		noErrorFlag bool
	}{
		{
			"asdf",
			"asdf",
			true,
		},
		{
			"",
			"",
			true,
		},
		{
			"234",
			"234",
			true,
		},
		{
			"wer",
			"sadf",
			true,
		},
	}

	for i, td := range testdata {
		t.Run("TestConvertToString_"+strconv.Itoa(i), func(t *testing.T) {
			result, err := convertToString(td.in)
			if td.noErrorFlag && err != nil && result != td.out {
				t.Error()
			} else if !td.noErrorFlag && err == nil {
				t.Error()
			}
		})
	}
}

func TestConverters_convertToInt(t *testing.T) {
	testdata := []struct {
		in          string
		out         interface{}
		case_		string
		noErrorFlag bool
	}{
		{
			"123",
			123,
			"int",
			true,
		},
		{
			"-123",
			-123,
			"int",
			true,
		},
		{
			"123",
			123,
			"uint",
			true,
		},
		{
			"-123",
			-123,
			"uint",
			true,
		},
		{
			"qewrq",
			123,
			"uint",
			false,
		},
		{
			"123",
			123,
			"int64",
			true,
		},
	}

	for i, td := range testdata {
		t.Run("TestConvertToInt_"+strconv.Itoa(i), func(t *testing.T) {
			result, err := convertToInt(td.in, td.case_)
			if td.noErrorFlag && err != nil && result != td.out {
				t.Error()
			} else if !td.noErrorFlag && err == nil {
				t.Error()
			}
		})
	}
}

func TestConverters_convertToTime(t *testing.T) {
	testdata := []struct {
		in          string
		noErrorFlag bool
	}{
		{
			"2019-08-21T09:00:00Z",
			true,
		},
		{
			"2019-000:00Z",
			false,
		},
		{
			"",
			false,
		},
		{
			"2019-08-21T41:00:00Z",
			false,
		},
		{
			"2019-32-21T41:00:00Z",
			false,
		},
	}

	for i, td := range testdata {
		t.Run("TestConvertToTime_"+strconv.Itoa(i), func(t *testing.T) {
			result, err := convertToTime(td.in)
			if td.noErrorFlag && err != nil {
				expected, _ := time.Parse(time.RFC3339, td.in)
				if !expected.Equal(result.(time.Time)) {
					t.Error()
				}
			} else if !td.noErrorFlag && err == nil {
				t.Error()
			}
		})
	}
}