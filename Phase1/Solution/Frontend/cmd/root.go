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
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ApiHost string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "frontend",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}


func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&ApiHost, "api", "", "API URL")
	viper.BindPFlag("api", rootCmd.PersistentFlags().Lookup("api"))
}

func initConfig() {
	if ApiHost == "" {
		value, success := os.LookupEnv("ESI_HOMEWORK1_API_HOST")
		if !success {
			fmt.Println("WARNING: No ESI_HOMEWORK1_API_HOST has been set in environment variables and API flag hasn't been passed, defaulting API host to localhost:90")
			ApiHost = "localhost:90"
		} else {
			ApiHost = value
		}
	}
}