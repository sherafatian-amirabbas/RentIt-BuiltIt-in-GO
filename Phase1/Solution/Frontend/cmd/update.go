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
	model "ESI-Homework1/Frontend/DataModel"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id] [title] [description] [priority]",
	Short: "Update a ToDo",
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("Error: First argument must be a integer")
		}

		priority, err := strconv.Atoi(args[3])

		if err != nil {
			fmt.Println("Error: Fourth argument must be a integer")
		}

		todo := model.TodoItem{
			ID:         id,
			Title:      args[1],
			Desciption: args[2],
			Priority:   priority,
		}

		todoJSON, _ := json.Marshal(todo)
		// The code below is mostly taken from https://stackoverflow.com/questions/46310113/consume-a-delete-endpoint-from-golang

		req, err := http.NewRequest("PUT", ("http://" + viper.GetString("api") + "/items/update/" + args[0]), bytes.NewBuffer(todoJSON))
		if err != nil {
			fmt.Println("Problem updating ToDo via REST: ", err)
			return
		}

		// Create client
		client := &http.Client{}

		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Problem updating ToDo via REST: ", err)
			return
		}
		defer resp.Body.Close()

		// Read Response Body
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Problem updating ToDo via REST: ", err)
			return
		}

		fmt.Println("ToDo item with ID", args[0], "was updating if it existed")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
