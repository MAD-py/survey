package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/validators"
)

// the questions to ask
var simpleQs = []*survey.Question{
	{
		Name: "letter",
		Prompt: &survey.Select{
			Message: "Choose a letter:",
			Options: []string{
				"a",
				"b",
				"c",
				"d",
				"e",
				"f",
				"g",
				"h",
				"i",
				"j",
			},
		},
		Validate: validators.Required,
	},
}

func main() {
	answers := struct {
		Letter string
	}{}

	// ask the question
	err := survey.Ask(simpleQs, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the answers
	fmt.Printf("you chose %s.\n", answers.Letter)
}
