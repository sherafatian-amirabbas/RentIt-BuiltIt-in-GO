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
	"strconv"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [id]",
	Short: "Delete ToDo item",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := strconv.Atoi(args[0])
       
		if err != nil {
			fmt.Println("Error: First argument must be a integer")
		}

		// The code below is mostly taken from https://stackoverflow.com/questions/46310113/consume-a-delete-endpoint-from-golang

		req, err := http.NewRequest("DELETE", ("http://" + viper.GetString("api") + "/items/delete/" + args[0]), nil)
		if err != nil {
			fmt.Println("Problem deleting ToDo via REST: ", err)
			return
		}

		// Create client
		client := &http.Client{}

		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Problem deleting ToDo via REST: ", err)
			return
		}
		defer resp.Body.Close()

		// Read Response Body
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Problem deleting ToDo via REST: ", err)
			return
		}

		fmt.Println("ToDo item with ID", args[0], "was deleted if it existed")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
