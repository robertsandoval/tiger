/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionsCmd represents the versions command
var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var category string
		if stableFlag {
			category = "stable"
		} else if latestFlag {
			category = "latest"
		} else if fastFlag {
			category = "fast"
		} else if candidateFlag {
			category = "candidate"
		} else {
			category = "all"
		}
		versions := getVersions(category)
		for _, ocp := range versions {
			fmt.Println(ocp.version)
		}
	},
}

func init() {
	//	getCmd.AddCommand(versionsCmd)
	versionsCmd.Flags().BoolVarP(&stableFlag, "stable", "s", false, "Get Stable Versions")
	versionsCmd.Flags().BoolVarP(&fastFlag, "fast", "f", false, "Get Fast Versions")
	versionsCmd.Flags().BoolVarP(&candidateFlag, "candidate", "c", false, "Get Candidate Versions")
	versionsCmd.Flags().BoolVarP(&latestFlag, "latest", "l", false, "Get Latest Versions")
}
