//This file is the main implementation of the validator
//Here you can find the public APIs as well as the constants used
//
//Below you can find an example on how to define a struct and a map in order to
//use the validator:
//
//		type InnerStruct struct {
//			C int `datakey:"c" validate:"required,nothing,int"`
//			D string `datakey:"d" validate:"nothing"`
//		}
//
//		type MyStruct struct {
//			A time.Time `validate:"required,nothing,time" datakey:"a"`
//			B int `validate:"required,int" datakey:"b"`
//			IS InnerStruct
//		}
//
//		s := MyStruct{}
//		m := map[string]string{
//					"a": "2019-08-21T09:00:00Z",
//					"b": "1",
//					"c": "1",
//					"d": "qwerty",
//				}
//
//		v := validator.New()
//		err := v.ValidateAndInit(m, &s)

package validator

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

const (
	//public

	//private
	//struct's tag keys
	tagMapKey   string = "datakey"
	tagValidate string = "validate"

	//rule names
	ruleRequired string = "required"
	ruleInt      string = "int"
	ruleUnsigned string = "unsigned"
	ruleTime     string = "time"
	ruleBool     string = "bool"

	//converter types
	convertInt    string = "int"
	convertString string = "string"
	convertTime   string = "time.Time"
	convertBool   string = "bool"
)

type (
	//The definition of the validator
	//Rule mappings is a map that connects a string rule name to an implementation/validator function
	//Rule mapping example: "ruleA" -> CheckA()
	//ConverterMappings is a map that connects a string type name to a converter function; changes the string value to
	//the desired type
	//Converter example: "MyStruct" -> ConvertToMyStruct()
	validator struct {
		//public

		//private
		ruleMappings      map[string]func(mapKey string, m map[string]string, params ...string) error
		converterMappings map[string]func(value string, params ...string) (interface{}, error)
		isInit            bool
	}
)

//Creates a new instance of the validator type
//Also it initializes the validator with the default configuration with before returning it
func New() *validator {
	validator := validator{}
	validator.initValidator()

	return &validator
}

//Initializes the validator with the default configuration and mappings
//Creates the "ruleMappings" and "converterMappings" maps and adds teh build in functions
//Before finishing it will set the "isInit" flag to true, which signifies that the validator is ready to be used
func (v *validator) initValidator() {
	v.ruleMappings = make(map[string]func(mapKey string, m map[string]string, params ...string) error)
	v.converterMappings = make(map[string]func(value string, params ...string) (interface{}, error))

	v.ruleMappings[ruleRequired] = checkRequired
	v.ruleMappings[ruleInt] = checkInt
	v.ruleMappings[ruleUnsigned] = checkUnsigned
	v.ruleMappings[ruleTime] = checkTime
	v.ruleMappings[ruleBool] = checkBool

	v.converterMappings[convertInt] = convertToInt
	v.converterMappings[convertString] = convertToString
	v.converterMappings[convertTime] = convertToTime
	v.converterMappings[convertBool] = convertToBool

	v.isInit = true
}

//Main function of the validator that is responsible for both validating the provided
//data based on the rules applied on the struct's tags, but also initializes the
//provided empty interface with the data
//
//The m parameter is a map with string keys and string values, while i parameter is an
//interface (normally a pointer to an empty struct)
func (v *validator) ValidateAndInit(m map[string]string, i interface{}) error {

	//If the i parameter is not a pointer, return an exception
	if reflect.ValueOf(i).Kind() != reflect.Ptr {
		return fmt.Errorf("please provide a pointer to the interface")
	}

	//If the i parameter is not a pointer to a struct, return an exception
	if reflect.Indirect(reflect.ValueOf(i)).Type().Kind() != reflect.Struct {
		return fmt.Errorf("please provide a pointer to the struct")
	}

	//If the validator is not initialized, return an error
	if !v.isInit {
		return fmt.Errorf("validator not initialized: call New()")
	}

	//Using reflection get the concrete value of the pointer i through the Elem()
	t := reflect.ValueOf(i).Elem()

	//Validation step
	//Check if the map values respect the rules defined on the struct's fields
	//If the validation fails, return an error
	err := v.checkRules(m, t)
	if err != nil {
		return errors.Wrap(err, "error validation map values based on rules")
	}

	//Struct initialization step
	//Initialize the struct with the values from the map
	//If the data initialization fails, return an error
	err = v.initData(m, t)
	if err != nil {
		return errors.Wrap(err, "error initializing struct with values")
	}

	return nil
}

//Used when user needs to add a custom rule
//The parameter "ruleName" is the name of teh rule as specified int he validation tag of the struct's field
//The second parameter is a function that needs to respect the required definition:
//* "mapKey" string parameter which will be the map key to which the struct field will be linked to
//* "m" a map with string keys and string values which will contains the data provided at the ValidateAndInit function
//* "params" a list of optional arguments
func (v *validator) RegisterRule(ruleName string, rule func(mapKey string, m map[string]string, params ...string) error) error { //If the validator is not initialized, return an error
	if !v.isInit {
		return fmt.Errorf("validator not initialized: call New()")
	}
	if ruleName == "" {
		return fmt.Errorf("empty rule name provided")
	}

	v.ruleMappings[ruleName] = rule

	return nil
}

//Used when the user needs to add custom converter from string value to a new data type
//The parameter "toType" is the name of the converter and must be the name of the new data type (e.g. MyStruct)
//The second parameter is a function that needs to respect the required definition:
//* "value" string parameter is the string value that needs to be converted to the "toType" data type
//* "params" is a list of optional arguments
func (v *validator) RegisterConverter(toType string, converter func(value string, params ...string) (interface{}, error)) error {
	if !v.isInit {
		return fmt.Errorf("validator not initialized: call New()")
	}
	if toType == "" {
		return fmt.Errorf("empty rule name provided")
	}
	v.converterMappings[toType] = converter

	return nil
}

//Validates the map data based on the rules defined on the struct's tags
//If the provided struct contains several sub struct as fields, they too will be taken into account recursively
func (v *validator) checkRules(m map[string]string, t reflect.Value) error {
	//Iterate over the list of struct fields
	for index := 0; index < t.Type().NumField(); index++ {
		//Get the current field from the struct
		currField := t.Type().Field(index)
		//If the current field is a sub struct, call the validation function recursively
		if currField.Type.Kind() == reflect.Struct {
			err := v.checkRules(m, t.Field(index))
			if err != nil {
				msg := fmt.Sprintf("validation of sub-struct '%s' failed", currField.Name)
				return errors.Wrap(err, msg)
			}
		}

		//Extract the list of rules from the tag "validate"
		validationRules, isValidationKey := currField.Tag.Lookup(tagValidate)
		//If the validation tag is present in the field tags apply the checks for each validation rule
		if isValidationKey {
			rules := strings.Split(validationRules, ",")
			if len(rules) == 1 && rules[0] == "" {
				continue
			}
			for _, ruleName := range rules {
				stripedRuleName := strings.TrimSpace(ruleName)
				//Extract the mapped function for the current rule and call it using the map data
				//If the rule name is not mapped in the validator, return an error
				if ruleImpl, ok := v.ruleMappings[stripedRuleName]; ok {
					if currField.Tag.Get(tagMapKey) == "" {
						return fmt.Errorf("tag '%s' not present struct tags", tagMapKey)
					}
					err := ruleImpl(currField.Tag.Get(tagMapKey), m)
					if err != nil {
						msg := fmt.Sprintf("validation failed at rule '%s'", ruleName)
						return errors.Wrap(err, msg)
					}
				} else {
					return fmt.Errorf("validation rule '%s' has no implementation. "+
						"please use 'RegisterRule' to provide one", ruleName)
				}
			}
		}
	}
	return nil
}

//Initializes the struct with the values form the map
//If the provided struct contains several sub struct as fields, they too will be taken into account recursively
func (v *validator) initData(m map[string]string, t reflect.Value) error {
	//Iterate over the list of struct fields
	for index := 0; index < t.Type().NumField(); index++ {
		//Get the current field from the struct
		currField := t.Type().Field(index)
		//If the current field is a sub struct, call the initialization function recursively
		if currField.Type.Kind() == reflect.Struct {
			err := v.initData(m, t.Field(index))
			if err != nil {
				msg := fmt.Sprintf("error initializing sub-struct '%s'", currField.Name)
				return errors.Wrap(err, msg)
			}
		}

		//Get the map value associated with the current field via the "datakey" tag
		if mapValue, ok := m[currField.Tag.Get(tagMapKey)]; ok {
			structFieldValue := t.FieldByName(currField.Name)
			//If the builtin data time is more complex (e.g. time.Time) it will build the type
			//of the struct using the package path and the name of the type
			structFieldType := structFieldValue.Type()
			converterName := ""
			if structFieldType.PkgPath() != "" {
				converterName = structFieldType.PkgPath() + "." + structFieldType.Name()
			} else {
				converterName = structFieldType.Name()
			}

			//Get the designated converter function for the current type from the "converterMappings" and call the
			//converter function with the map value
			//If the type is not mapped to a converter it will return an error
			if converter, ok := v.converterMappings[converterName]; ok {
				result, err := converter(mapValue, converterName)
				if err != nil {
					msg := fmt.Sprintf("error converting to type '%s'", structFieldType.Name())
					return errors.Wrap(err, msg)
				}

				//Set the computed value the field
				structFieldValue.Set(reflect.ValueOf(result))
			} else {
				return fmt.Errorf("conversion to '%s' is not defined, please use RegisterConverter", structFieldType)
			}
		}
	}
	return nil
}
