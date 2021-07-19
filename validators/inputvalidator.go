package validators

import (
	"fmt"
	"net/mail"
	"reflect"
	"time"
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

// Email requires that the string has a structure type email
func Email(val interface{}) error {
	if str, ok := val.(string); ok {
		// if the string does not have an email structure
		if _, err := mail.ParseAddress(str); err != nil {
			// yell loudly
			return fmt.Errorf("the answer has no email structure")
		}
	} else {
		// otherwise we cannot convert the value into a string and cannot enforce email validation
		return fmt.Errorf("cannot enforce email validation on response of type %v", reflect.TypeOf(val).Name())
	}

	// the input is fine
	return nil
}

// Time requires that the string has a structure type time
func Time(layout string) Validator {
	// if the string is not valid for time.parse
	if _, err := time.Parse(layout, layout); err != nil {
		panic("The layout is invalid, the valid layouts are the same of time.Parse")
	}
	// return a validator that checks the structure of the date and/or time
	return func(val interface{}) error {
		if str, ok := val.(string); ok {
			// if the string has not structure of date and/or time
			if _, err := time.Parse(layout, str); err != nil {
				// yell loudly
				return fmt.Errorf("the answer has no the expected structure (%v)", layout)
			}
		} else {
			// otherwise we cannot convert the value into a string and cannot enforce time validation
			return fmt.Errorf("cannot enforce time validation on response of type %v", reflect.TypeOf(val).Name())
		}

		// the input is fine
		return nil
	}
}
