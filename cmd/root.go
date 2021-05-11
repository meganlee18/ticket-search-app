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
	"os"

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

			//TODO: refactor this to switch statement for readability
			//TODO: Also pull this out to separate function
			if searchOptions == 1 {
				fmt.Println("Select 1) Users or 2) Tickets or 3) Organizations")
				var options int
				fmt.Scanln(&options)

				if options == 1 {
					path := "./tickets/users.json"
					//TODO: depending on search options, return search results either user, ticket or org
					fmt.Println("Enter search term e.g _id or name")

					//check search term and see if it matches field
					var searchTerm string
					fmt.Scanln(&searchTerm)

					fmt.Println("Enter search value e.g 71")
					var searchValue int
					fmt.Scanln(&searchValue)

					app.DisplayUsersBasedOnSearchOptions(path, searchValue)
				}
			}

			if searchOptions == 2 {
				fmt.Println("----------------------\nYou can search Users with:")
				app.DisplayTicketFields("./tickets/users.json")

				fmt.Println("----------------------\nYou can search Tickets with:")
				app.DisplayTicketFields("./tickets/tickets.json")

				fmt.Println("----------------------\nYou can search Organizations with:")
				app.DisplayTicketFields("./tickets/organizations.json")
			}
		},
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
