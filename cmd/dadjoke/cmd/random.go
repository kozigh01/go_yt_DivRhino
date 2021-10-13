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
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

const (
	baseUrl = "https://icanhazdadjoke.com/"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random  dad joke",
	Long:  `This command fetches a random dad joke from the icanhazdadjoke api`,
	Run: func(cmd *cobra.Command, args []string) {
		jokeTerm, _ := cmd.Flags().GetString("term")
		if jokeTerm != "" {
			getRandomJokeWithTerm(jokeTerm)
		} else {
			getRandomJoke()
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")
	randomCmd.PersistentFlags().String("term", "", "search term for a dad joke.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

type SearchResult struct {
	Results    json.RawMessage `json:"results"`
	SearchTerm string          `json:"search_term"`
	Status     int             `json:"status"`
	TotalJokes int             `json:"total_jokes"`
}

func getRandomJokeWithTerm(jokeTerm string) {
	total_jokes, results := getJokeDataWithTerm(jokeTerm)
	if total_jokes <= 0 {
		fmt.Printf("No jokes found for the term '%s'\n", jokeTerm)
	} else {
		fmt.Println(randomiseJokeList(results).Joke)
	}
}

func randomiseJokeList(jokes []Joke) Joke {
	// rand := rand.New(rand.NewSource(time.Now().UnixMicro()))
	rand.Seed(time.Now().UnixMilli())
	indx := rand.Intn(len(jokes))
	return jokes[indx]
}

func getJokeDataWithTerm(jokeTerm string) (totalJokes int, jokeList []Joke) {
	url := fmt.Sprintf("%s/search?term=%s", baseUrl, jokeTerm)
	respBytes := getJokeData(url)
	jokeListRaw := SearchResult{}

	if err := json.Unmarshal(respBytes, &jokeListRaw); err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	jokes := []Joke{}
	if err := json.Unmarshal(jokeListRaw.Results, &jokes); err != nil {
		log.Printf("could not unmarshal jokeListRaw results - %v", err)
	}

	return jokeListRaw.TotalJokes, jokes
}

func getRandomJoke() {
	joke := new(Joke)
	jokeData := getJokeData(baseUrl)

	if err := json.Unmarshal(jokeData, &joke); err != nil {
		log.Fatalf("Had a problem unmarshalling the dad joke - %v", err)
	}

	fmt.Println(joke.Joke)
}

func getJokeData(baseUrl string) []byte {
	req, err := http.NewRequest(
		http.MethodGet,
		baseUrl,
		nil,
	)
	if err != nil {
		log.Fatalf("Could not request a dadjoke - %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Dadjoke CLI (github.com/kozigh01/go_yt_DivRhino)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Had problems calling the dad joke endpoint - %v", err)
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Had a problem reading the body of the response = %v", err)
	}
	return respBytes
}
