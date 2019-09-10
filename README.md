# Map data Validator and struct initializer based on field tags

Validates and initialises a struct with values from a map, based on rules defined inside tags on the struct's fields
In order to link the field to a map key, use the tag `datakey` (e.g. `datakey:"aKey"`)

The builtin rules are the following:
* `required`: checks if the `datakey` is present inside the values map
* `int`: checks if the map value is convertible to integer
* `time`: checks if the map value is convertible to `time.Time`

The initialization of the struct is based on the type of the field. This means that, after the validation 
has passed, the values from the map will be converted to the type of the designated field

The values map is a map of type `map[string]string`

# Quick example
            type InnerStruct struct {
			C int `datakey:"c" validate:"required,int"`
			D string `datakey:"d"`
		}

		type MyStruct struct {
			A time.Time `validate:"required,time" datakey:"a"`
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
		err := v.ValidateAndInit(m, &s)

# Defining custom rules
In order to add a new rule you must register it inside the validator by using `RegisterRule` like this:

```   
func MyRule(mapKey string, m map[string]string, params ...string) error {
    return nil
}
v.RegisterRule("myRule", MyRule)
```

Now you can use the rule inside the `validate` tag along side the builtin ones

# Defining custom type converters
Type Converters are useful when you have to convert a string value to a more complex data type (e.g. Mongo primitive ObjectId)
In order to add a new converter you must register it inside the validator by using `RegisterConverter` like this:

```
func MyConverter(value string, params ...string) (i interface{}, e error) {
    return nil
}
v.RegisterConverter("MyType", MyConverter)
```
