/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package note

import (
	"errors"
	"fmt"
	"log"

	"github.com/kozigh01/go_yt_DivRhino/cmd/studybuddy/data"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "creates a new studybuddy note",
	Long:  `creates a new studybuddy note`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	noteCmd.AddCommand(newCmd)
}

func promptGetInput(pc promptContent) string {
	prompt := promptui.Prompt{
		Label: pc.label,
		Templates: &promptui.PromptTemplates{
			Prompt:  "{{ . }}",
			Valid:   "{{ . | green }}",
			Invalid: "{{ . | red }}",
			Success: "{{ . | bold }}",
		},
		Validate: func(input string) error {
			if len(input) <= 0 {
				return errors.New(pc.errorMsg)
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	log.Printf("Input: %s\n", result)

	return result
}

func promptGetSelect(pc promptContent) string {
	items := []string{"animal", "food", "person", "object"}
	index := -1

	var result string
	var err error

	prompt := promptui.SelectWithAdd{
		Label:    pc.label,
		Items:    items,
		AddLabel: "Other",
	}
	
	for index < 0 {
		index, result, err = prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}

		if index == -1 {
			prompt.Items = append(prompt.Items, result)
		}
	}

	fmt.Printf("Input: %s\n", result)
	return result
}

func createNewNote() {
	wordPromptContent := promptContent{
		errorMsg: "Please provide a word",
		label:    "What word would you like to make a note of? ",
	}
	word := promptGetInput(wordPromptContent)

	definitionPromptContent := promptContent{
		errorMsg: "Please provide a definition",
		label:    fmt.Sprintf("What is the definition of %s? ", word),
	}
	definition := promptGetInput(definitionPromptContent)

	categoryPromptContent := promptContent{
		errorMsg: "Please provide a category",
		label:    fmt.Sprintf("What category does %s belong to? ", word),
	}
	category := promptGetSelect(categoryPromptContent)

	data.InsertNote(word, definition, category)
}
