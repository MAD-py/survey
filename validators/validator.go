package validators

import (
	"errors"
	"reflect"
)

// Validator is a function passed to a Question after a user has provided a response.
// If the function returns an error, then the user will be prompted again for another
// response.
type Validator func(ans interface{}) error

// Required does not allow an empty value
func Required(val interface{}) error {
	// the reflect value of the result
	value := reflect.ValueOf(val)

	// if the value passed in is the zero value of the appropriate type
	if IsZero(value) && value.Kind() != reflect.Bool {
		return errors.New("Value is required")
	}
	return nil
}

// ComposeValidators is a variadic function used to create one validator from many.
func ComposeValidators(validators ...Validator) Validator {
	// return a validator that calls each one sequentially
	return func(val interface{}) error {
		// execute each validator
		for _, validator := range validators {
			// if the answer's value is not valid
			if err := validator(val); err != nil {
				// return the error
				return err
			}
		}
		// we passed all validators, the answer is valid
		return nil
	}
}

// isZero returns true if the passed value is the zero object
func IsZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		return v.Len() == 0
	}

	// compare the types directly with more general coverage
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
