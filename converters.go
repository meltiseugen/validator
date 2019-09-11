//This file is used to define al the builtin type converters (from string to interface{}) of the validators
//The current converters are: convertToInt, convertToTime, convertToString

package validator

import (
	"fmt"
	"strconv"
	"time"
)

//Converts a string numeric value to and int type value
//The optional parameter is used to provide the type of int return value: int, uint, int64
func convertToInt(value string, params ...string) (interface{}, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return nil, fmt.Errorf("error converting '%s' to int", value)
	}

	switch params[0] {
	case "int":
		return intValue, nil
	case "uint":
		return uint(intValue), nil
	case "int64":
		return int64(intValue), nil
	default:
		return intValue, nil
	}
}

//Converts a string time value to a time.Time value
func convertToTime(value string, params ...string) (interface{}, error) {
	time_, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return nil, fmt.Errorf("error parsing time string")
	}

	return time_, nil
}

//Converts a string to a string value
//Basically it just returns the value
//Defined in order to have consistency and to work well with the overall converter mechanism
func convertToString(value string, params ...string) (interface{}, error) {
	return value, nil
}

//Converts a string bool value to a bool type value
func convertToBool(value string, params ...string) (interface{}, error) {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return nil, fmt.Errorf("error checking bool string")
	}

	return b, nil
}