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
	if mapKey == "" {
		message := fmt.Sprintf("tag '%s' not present struct tags", tagMapKey)
		return fmt.Errorf(message)
	}

	if _, ok := m[mapKey]; !ok {
		msg := fmt.Sprintf("requred field '%s' is not preset in map", mapKey)
		return fmt.Errorf(msg)
	}
	return nil
}

//Validates if, for a given key "mapKey" and a given map "m", the value of the key can be converted to an integer
//If not it will return an error
//First it checks if the mapKey is an empty string and if no it will return an error
func checkInt(mapKey string, m map[string]string, params ...string) error {
	if mapKey == "" {
		message := fmt.Sprintf("tag '%s' not present struct tags", tagMapKey)
		return fmt.Errorf(message)
	}

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
	if mapKey == "" {
		message := fmt.Sprintf("tag '%s' not present struct tags", tagMapKey)
		return fmt.Errorf(message)
	}

	if mapValue, ok := m[mapKey]; ok {
		if strings.HasPrefix(mapValue, "-"){
			msg := fmt.Sprintf("map key '%s' does not match constraint '%s'", mapKey, ruleUnsigned)
			return fmt.Errorf(msg)
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
	if mapKey == "" {
		message := fmt.Sprintf("tag '%s' not present struct tags", tagMapKey)
		return fmt.Errorf(message)
	}

	if mapValue, ok := m[mapKey]; ok {
		_, err := time.Parse(time.RFC3339, mapValue)
		if err != nil {
			return fmt.Errorf("error parsing time string")
		}
	}
	return nil
}
