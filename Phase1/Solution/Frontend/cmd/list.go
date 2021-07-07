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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [all | sorted | inprogress | completed]",
	Short: "List all ToDos",
	Args: cobra.ExactValidArgs(1),
	ValidArgs: []string{"all", "sorted", "inprogress", "completed"},
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(("http://" + viper.GetString("api") + "/items/" + args[0]))
		if err != nil {
			fmt.Println("Problem retrieving ToDos! ", err)
			return
		}
		todos, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(todos))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
