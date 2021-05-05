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
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"../pkg/app"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ticket-search-app",
	Short: "This is a CLI app that searches and returns app",
	Long: `Ticket Search App is a simple CLI tool that searches app and returns the results in human readable format`,
	// Uncomment the following line if your bare app
	// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {

			//Ask user what they would like e.g users, tickets, org
			//Depend on input, call various functions
			fmt.Println("Welcome to Ticket Search App!\n\nSelect search options:\n - Press 1 to start searching\n - Press 2 to view a list of searchable fields")
			var searchOptions int

			fmt.Scanln(&searchOptions)
			if searchOptions == 1 {
				fmt.Println("Select 1) Users or 2) Tickets or 3) Organizations")
				var options int
				fmt.Scanln(&options)

				if options == 1 {
					////if user asks for tickets, call this
					//app.DisplayTickets()
					//
					////if user asks for users, call this
					////app.
					//app.DisplayUsers()
					//
					////if users asks for organizations, call this
					//app.DisplayOrganizations()
				}
			}

			if searchOptions == 2 {
				//return user fields e.g id, name etc
				fmt.Println("----------------------\nYou can search Users with:")
				readFields("./tickets/users.json")

				fmt.Println("----------------------\nYou can search Tickets with:")
				readFields("./tickets/tickets.json")

				fmt.Println("----------------------\nYou can search Organizations with:")
				readFields("./tickets/organizations.json")
			}
		},
}

func readFields(path string) {
	var result []map[string]interface{}

	data := app.ReadFile(path)
	unmarshalledResult, err := unmarshalData(data, result)
	if err != nil {
		fmt.Println("error: ", err)
	}

	sortedFields := displaySortedFields(unmarshalledResult)
	fmt.Println(strings.Join(removeDuplicateValues(sortedFields),  "\n"))
}

func unmarshalData(data []byte, result []map[string]interface{}) ([]map[string]interface{}, error) {
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	return result, err
}

func displaySortedFields(result []map[string]interface{}) []string {
	var fields []string

	for _, v := range result {
		for k := range v {
			fields = append(fields, k)
		}
	}

	sort.Strings(fields)
	return fields
}


func removeDuplicateValues(fields []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range fields {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your app.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ticket_search_app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ticket_search_app" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ticket_search_app")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
