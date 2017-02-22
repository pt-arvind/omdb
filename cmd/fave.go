// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/dghubble/sling"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Movie struct {
	Title string `json:"Title"`
}

// order matters?!?
type BadResponse struct {
	Response string `json:"Response"`
	Error    string `json:"Error"`
}

// faveCmd represents the fave command
var faveCmd = &cobra.Command{
	Use:   "fave",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		numArgs := len(args)
		if numArgs != 1 {
			fmt.Println("Usage: omdb fave \"<title of film>\"")
			return
		}

		searchKey := args[0]

		client := &http.Client{}
		params := &TitleParams{Title: searchKey, Plot: "short", ResponseType: "json"}
		req, err := sling.New().Get("https://www.omdbapi.com/").QueryStruct(params).Request()

		if err != nil { //TODO: use juju errors
			fmt.Println("error: " + err.Error())
		}
		log.WithFields(log.Fields{
			"REQUEST": req,
		}).Info("MAKING REQUEST")

		resp, err := client.Do(req)

		if err != nil { //TODO: use juju errors
			fmt.Println("error: " + err.Error())
		}

		defer resp.Body.Close() //I dont like defer

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil { //TODO: use juju errors
			fmt.Println("error: " + err.Error())
		}

		jsonString := string(body)

		log.WithFields(log.Fields{
			"API RESPONSE": jsonString,
		}).Info("RESPONSE RECEIVED")

		fmt.Println("\n****\nAPI response: \n" + jsonString)
		movie := Movie{}

		json.Unmarshal([]byte(jsonString), &movie)

		//FIXME: figure out why this wont work
		// if err != nil {
		// 	//TODO: use juju errors
		// 	fmt.Println("error: " + err.Error())
		//
		// }

		fmt.Println(movie.Title)

		title := movie.Title
		// if title == "" {
		// 	badResponse := BadResponse{}
		// 	err := json.Unmarshal([]byte(jsonString), &badResponse)
		// 	if err != nil {
		// 		//TODO: use juju errors
		// 		fmt.Println("error: " + err.Error())
		// 	}
		// 	fmt.Println("bad response!")
		// 	fmt.Println(badResponse)
		// }
		viper.SetDefault("favorite_movie", title)
		fmt.Println("Favorite Movie is now: " + title)
	},
}

func init() {
	RootCmd.AddCommand(faveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// faveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// faveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
