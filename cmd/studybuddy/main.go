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
package main

import (
	"log"

	"github.com/kozigh01/go_yt_DivRhino/cmd/studybuddy/cmd"
	_ "github.com/kozigh01/go_yt_DivRhino/cmd/studybuddy/cmd/note"
	"github.com/kozigh01/go_yt_DivRhino/cmd/studybuddy/data"
)

func main() {
	if err := data.OpenDatabase(); err != nil {
		log.Printf("there was an issue opening the db: %v\n", err)
	}
	cmd.Execute()
}
