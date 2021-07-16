package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/validators"
)

var simpleQs = []*survey.Question{
	{
		Name: "color",
		Prompt: &survey.Select{
			Message: "select1:",
			Options: []string{"red", "blue", "green"},
		},
		Validate: validators.Required,
	},
	{
		Name: "color2",
		Prompt: &survey.Select{
			Message: "select2:",
			Options: []string{"red", "blue", "green"},
		},
		Validate: validators.Required,
	},
}

func main() {
	answers := struct {
		Color  string
		Color2 string
	}{}
	// ask the question
	err := survey.Ask(simpleQs, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the answers
	fmt.Printf("%s and %s.\n", answers.Color, answers.Color2)
}
