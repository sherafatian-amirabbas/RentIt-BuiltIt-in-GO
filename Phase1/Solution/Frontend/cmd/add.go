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
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [id] [title] [description] [priority]",
	Short: "Add a ToDo",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("Error: First argument must be an integer")
		}

		priority, err := strconv.Atoi(args[3])

		if err != nil {
			fmt.Println("Error: Fourth argument must be an integer")
		}

		todo := model.TodoItem{
			ID:         id,
			Title:      args[1],
			Desciption: args[2],
			Priority:   priority,
		}

		todoJSON, _ := json.Marshal(todo)
		_, err = http.Post("http://" + viper.GetString("api") + "/items/create", "", bytes.NewBuffer(todoJSON))
		if err != nil {
			fmt.Println("Problem adding new ToDo via REST: ", err)
			return
		}

		fmt.Println("New ToDo item with ID: ", id, " was added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
