//This file contains all the built in checks/validators for teh built in rules like
//"required", "time", "int", "unsigned int"

package validator

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Validates if, the given key "mapKey" for a given map "m" is present in m
//First it checks if the mapKey is an empty string and if no it will return an error
func checkRequired(mapKey string, m map[string]string, params ...string) error {
	if _, ok := m[mapKey]; !ok {
		return fmt.Errorf("requred field '%s' is not preset in map", mapKey)
	}
	return nil
}

//Validates if, for a given key "mapKey" and a given map "m", the value of the key can be converted to an integer
//If not it will return an error
//First it checks if the mapKey is an empty string and if no it will return an error
func checkInt(mapKey string, m map[string]string, params ...string) error {
	if mapValue, ok := m[mapKey]; ok {
		_, err := strconv.Atoi(mapValue)
		if err != nil {
			return fmt.Errorf("failed to convert string to int for key '%s'", mapKey)
		}
	}

	return nil
}

//Validates if, for a given key "mapKey" and a given map "m", the value starts with a dash ("-")ß
//First it checks if the mapKey is an empty string and if no it will return an error
func checkUnsigned(mapKey string, m map[string]string, params ...string) error {
	if mapValue, ok := m[mapKey]; ok {
		if strings.HasPrefix(mapValue, "-"){
			return fmt.Errorf("map key '%s' does not match constraint '%s'", mapKey, ruleUnsigned)
		}
	}
	err := checkInt(mapKey, m)
	if err != nil {
		return err
	}
	return nil
}

//Validates if, for a given key "mapKey" and a given map "m", the value can be converted to a time.Time instance
//of the following format: 	RFC3339 = "2006-01-02T15:04:05Z07:00"
//First it checks if the mapKey is an empty string and if no it will return an error
func checkTime(mapKey string, m map[string]string, params ...string) error {
	if mapValue, ok := m[mapKey]; ok {
		_, err := time.Parse(time.RFC3339, mapValue)
		if err != nil {
			return fmt.Errorf("error checking time string")
		}
	}
	return nil
}

//Validates if, for a given key "mapKey" and a given map "m", the value has a boolean meaning
//Checks if the value is either "true" or "false" CASES-SENSITIVE
func checkBool(mapKey string, m map[string]string, params ...string) error {
	if mapValue, ok := m[mapKey]; ok {
		if mapValue != "true" && mapValue != "false" {
			return fmt.Errorf("error checking bool string")
		}
	}
	return nil
}