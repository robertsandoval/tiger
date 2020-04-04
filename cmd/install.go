/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/AlecAivazis/survey"
	"github.com/spf13/cobra"
	"os"
)

var qs = []*survey.Question{
	{
		Name:   "clusterName",
		Prompt: &survey.Input{Message: "Cluster Name?"},
	},
}

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			ClusterName   string    // if the types don't match, survey will convert it
		}{}

		// perform the questions
		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		clusterPath := ("/usr/local/ocp/clusters/" + os.Getenv("USER") + "/" +  answers.ClusterName)
		if _, err := os.Stat(clusterPath); ! os.IsNotExist(err) {
			fmt.Println("Cluster with this name already exists")
			return
		}
		os.MkdirAll(clusterPath, os.ModePerm)

	},
}

func init() {
	createCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
