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
)

// type Params struct {
// 	Count int `url:"count,omitempty"`
// }
type TitleParams struct {
	Title        string `url:"t,omitempty"`
	Plot         string `url:"plot,omitempty"`
	ResponseType string `url:"r,omitempty"`
}

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// params := &Params{Count: 5}
		client := &http.Client{}
		// req, err := sling.New().Get("https://www.omdbapi.com/?t=the+matrix&plot=short&r=json").Request()
		params := &TitleParams{Title: "the matrix", Plot: "short", ResponseType: "json"}
		req, err := sling.New().Get("https://www.omdbapi.com/").QueryStruct(params).Request()

		if err != nil { //TODO: use juju errors
			fmt.Println("error: " + err.Error())
		}

		resp, err := client.Do(req)

		if err != nil { //TODO: use juju errors
			fmt.Println("error: " + err.Error())
		}

		defer resp.Body.Close() //I dont like defer

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil { //TODO: use juju errors
			fmt.Println("error: " + err.Error())
		}
		fmt.Println("\n****\nAPI response: \n" + string(body))
	},
}

func init() {
	RootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
