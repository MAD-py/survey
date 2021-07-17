package validators

import (
	"testing"
)

func TestMaxLength(t *testing.T) {
	// the string to test
	testStr := randString(150)
	// validate the string
	if err := MaxLength(140)(testStr); err == nil {
		t.Error("No error returned with input greater than 150 characters.")
	}

	// emoji test
	emojiStr := "I😍Golang"
	// validate visible length with Maxlength
	if err := MaxLength(10)(emojiStr); err != nil {
		t.Errorf("Error returned with emoji containing 8 characters long input.")
	}
}

func TestMinLength(t *testing.T) {
	// validate the string
	if err := MinLength(12)(randString(10)); err == nil {
		t.Error("No error returned with input less than 12 characters.")
	}

	// emoji test
	emojiStr := "I😍Golang"
	// validate visibly 8 characters long string with MinLength
	if err := MinLength(10)(emojiStr); err == nil {
		t.Error("No error returned with emoji containing input less than 10 characters.")
	}
}

func TestMinLength_onInt(t *testing.T) {
	// validate the string
	if err := MinLength(12)(1); err == nil {
		t.Error("No error returned when enforcing length on int.")
	}
}

func TestMaxLength_onInt(t *testing.T) {
	// validate the string
	if err := MaxLength(12)(1); err == nil {
		t.Error("No error returned when enforcing length on int.")
	}
}

func TestEmail(t *testing.T) {
	// the string to test
	testStr := "test.1234gmail.com"
	// validate the string
	if err := Email(testStr); err == nil {
		t.Error("No error returned when input has not email structure")
	}
}
