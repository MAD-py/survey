package validators

import (
	"fmt"
	"reflect"
)

// MaxLength requires that the string is no longer than the specified value
func MaxLength(length int) Validator {
	// return a validator that checks the length of the string
	return func(val interface{}) error {
		if str, ok := val.(string); ok {
			// if the string is longer than the given value
			if len([]rune(str)) > length {
				// yell loudly
				return fmt.Errorf("value is too long. Max length is %v", length)
			}
		} else {
			// otherwise we cannot convert the value into a string and cannot enforce length
			return fmt.Errorf("cannot enforce length on response of type %v", reflect.TypeOf(val).Name())
		}

		// the input is fine
		return nil
	}
}

// MinLength requires that the string is longer or equal in length to the specified value
func MinLength(length int) Validator {
	// return a validator that checks the length of the string
	return func(val interface{}) error {
		if str, ok := val.(string); ok {
			// if the string is shorter than the given value
			if len([]rune(str)) < length {
				// yell loudly
				return fmt.Errorf("value is too short. Min length is %v", length)
			}
		} else {
			// otherwise we cannot convert the value into a string and cannot enforce length
			return fmt.Errorf("cannot enforce length on response of type %v", reflect.TypeOf(val).Name())
		}

		// the input is fine
		return nil
	}
}
