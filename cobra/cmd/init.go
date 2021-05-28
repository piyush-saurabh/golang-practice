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
	"log"
	"os"
	"os/user" // get user's home directory

	"github.com/cockroachdb/pebble" // simple key-value database
	"github.com/spf13/cobra"        // CLI
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "[CHANGE THIS] A brief description of your command",
	Long: `[CHANGE THIS] A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

[CHANGE THIS] Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This is called when user starts using the cli
		fmt.Println("init called")
		flag1, _ := cmd.Flags().GetString("flag1")
		flag2, _ := cmd.Flags().GetString("flag2")
		configDir, _ := cmd.Flags().GetString("db")

		// Create the local directory for the cli $HOME/.mycli and store the db
		// Check if the directory exists
		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			err := os.Mkdir(configDir, os.ModePerm)
			if err != nil {
				log.Fatal("Cannot create the configuation directory for the CLI")
			}

			fmt.Println("Configuration directory created successfully !!")

			// Location of local database (key value pair)
			dbLocation := fmt.Sprintf("%s/localcli.db", configDir)

			// Open the database
			db, err := pebble.Open(dbLocation, &pebble.Options{})
			if err != nil {
				log.Fatal("Unable to create the local database")
			}

			// Store some data in database (e.g. info passed as CLI input)
			sampleKey1 := []byte("key1")
			if err := db.Set(sampleKey1, []byte(flag1), pebble.Sync); err != nil {
				log.Fatal("Failed while writing flag 1")
			}

			sampleKey2 := []byte("key2")
			if err := db.Set(sampleKey2, []byte(flag2), pebble.Sync); err != nil {
				log.Fatal("Failed while writing flag 2")
			}

			fmt.Println("CLI initialized successfully !!")

			// Close the db
			if err := db.Close(); err != nil {
				log.Fatal("Unable to close database")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Create a base directory for this CLI. e.g $HOME/.mycli
	// Get the user's home directory
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)

	}

	rootDir := fmt.Sprintf("%s/.mycli", usr.HomeDir) // $HOME

	// Create flags for CLI
	// The shorthand of the flag should be only 1 character
	initCmd.Flags().StringP("flag1", "f", rootDir, "Help text for Flag 1") // mycli -f1/--flag1

	initCmd.Flags().StringP("flag2", "g", rootDir, "Help text for Flag 2") // mycli -f2/--flag2

	initCmd.Flags().StringP("db", "d", rootDir, "Help text for Database") // mycli -d/--db

}
