package main

import (
	"fmt"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/validators"
)

func suggestFiles(toComplete string) []string {
	files, _ := filepath.Glob(toComplete + "*")
	return files
}

// the questions to ask
var q = []*survey.Question{
	{
		Name: "file",
		Prompt: &survey.Input{
			Message: "Which file should be read?",
			Suggest: suggestFiles,
			Help:    "Any file; do not need to exist yet",
		},
		Validate: validators.Required,
	},
}

func main() {
	answers := struct {
		File string
	}{}

	// ask the question
	err := survey.Ask(q, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the answers
	fmt.Printf("File chosen %s.\n", answers.File)
}
