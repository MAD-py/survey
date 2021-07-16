package validators

import (
	"testing"

	"github.com/AlecAivazis/survey/v2/core"
)

func TestMaxItems(t *testing.T) {
	// the list to test
	testList := []core.OptionAnswer{
		core.OptionAnswer{Value: "a", Index: 0},
		core.OptionAnswer{Value: "b", Index: 1},
		core.OptionAnswer{Value: "c", Index: 2},
		core.OptionAnswer{Value: "d", Index: 3},
		core.OptionAnswer{Value: "e", Index: 4},
		core.OptionAnswer{Value: "f", Index: 5},
	}

	// validate the list
	if err := MaxItems(4)(testList); err == nil {
		t.Error("No error returned with input greater than 6 items.")
	}
}

func TestMinItems(t *testing.T) {
	// the list to test
	testList := []core.OptionAnswer{
		core.OptionAnswer{Value: "a", Index: 0},
		core.OptionAnswer{Value: "b", Index: 1},
		core.OptionAnswer{Value: "c", Index: 2},
		core.OptionAnswer{Value: "d", Index: 3},
		core.OptionAnswer{Value: "e", Index: 4},
		core.OptionAnswer{Value: "f", Index: 5},
	}

	// validate the list
	if err := MinItems(10)(testList); err == nil {
		t.Error("No error returned with input less than 10 items.")
	}
}
