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
		t.Error("No error returned when input has not email structure.")
	}
}

func TestTime(t *testing.T) {
	// the layout and date to test
	layout := "2006-Jan-02"
	testTime := "2014-07-04"

	// validate the date
	if err := Time(layout)(testTime); err == nil {
		t.Error("No error returned when input has not expected structure.")
	}
}

func TestTimeErrorLayout(t *testing.T) {
	// the layout to test
	layout := "206-11-11"
	defer func() {
		if err := recover(); err == nil {
			t.Error("No error returned when layout is not valid.")
		}
	}()
	Time(layout)
}

func TestInvalidChars(t *testing.T) {
	chars := "*@#%.,/\\=-"
	if err := InvalidChars(chars)("I love gol@ng"); err == nil {
		t.Error("No error returned when enforcing the invalid chars.")
	}
}
