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
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/spf13/cobra"

	log "github.com/Sirupsen/logrus"
)

// type Params struct {
// 	Count int `url:"count,omitempty"`
// }
type TitleParams struct {
	Title        string `url:"t,omitempty"`
	Plot         string `url:"plot,omitempty"`
	ResponseType string `url:"r,omitempty"`
}

// titleCmd represents the title command
var titleCmd = &cobra.Command{
	Use:   "title",
	Short: "Search OMDB by film title",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		numArgs := len(args)
		if numArgs != 1 {
			fmt.Println("Usage: omdb title \"<title of film>\"")
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

		log.WithFields(log.Fields{
			"API RESPONSE": string(body),
		}).Info("RESPONSE RECEIVED")

		fmt.Println("\n****\nAPI response: \n" + string(body))
	},
}

func init() {
	RootCmd.AddCommand(titleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// titleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// titleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
