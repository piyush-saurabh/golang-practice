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
	"strings"

	utils "github.com/piyush-saurabh/golang-practice/cobra/utils"
	"github.com/spf13/cobra"
)

// podCmd represents the pod command
var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pod called")

		// Arguments are recorded here
		fmt.Println("Here are the arguments of card command : " + strings.Join(args, ","))

		// Read the flags
		name, _ := cmd.Flags().GetString("name")
		namespace, _ := cmd.Flags().GetString("namespace")
		output, _ := cmd.Flags().GetString("output")

		// fmt.Println("value of name flag is : " + name)
		// fmt.Println("value of namespace flag is : " + namespace)
		// fmt.Println("value of output flag is : " + output)

		utils.Test(name, namespace, output)
	},
}

func init() {
	createCmd.AddCommand(podCmd)

	podCmd.PersistentFlags().StringP("namespace", "n", "default", "Possible Values: kube-system, gate-keeper")

	podCmd.PersistentFlags().StringP("output", "o", "", "Possible Values: yaml, json")

	podCmd.PersistentFlags().StringP("name", "a", "", "Name of the pod")

	// Make name of the pod mandatory. if not provided, return error
	podCmd.MarkPersistentFlagRequired("name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
