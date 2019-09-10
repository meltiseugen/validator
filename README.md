# Validator
Validates and initialises a struct with values from a map, based on rules defined inside tags on the struct's fields

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
