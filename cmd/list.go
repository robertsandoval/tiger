/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	stableFlag    bool
	chanel        string
	latestFlag    bool
	fastFlag      bool
	candidateFlag bool
	devFlag       bool
	showurls      bool
	ocpversion    string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:       "list",
	Short:     "A brief description of your command",
	Long:      `Long description`,
	ValidArgs: []string{"versions", "version"},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var category string

		if args[0] == "versions" {
			if stableFlag {
				category = "stable"
			} else if latestFlag {
				category = "latest"
			} else if fastFlag {
				category = "fast"
			} else if candidateFlag {
				category = "candidate"
			} else if devFlag {
				category = "dev"
			} else {
				category = "all"
			}
			versions := getVersions(category)
			for _, ocp := range versions {
				fmt.Println(ocp.version)
			}
		} else if args[0] == "version" {
			getVersion(ocpversion)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&chanel, "chanel", "c", "", "List specific OCP version")
	listCmd.Flags().BoolVar(&stableFlag, "stable", false, "Get Stable Version")
	listCmd.Flags().BoolVar(&fastFlag, "fast", false, "Get Fast Version")
	listCmd.Flags().BoolVar(&candidateFlag, "candidate", false, "Get Candidate Version")
	listCmd.Flags().BoolVar(&latestFlag, "latest", false, "Get Latest Version")
	listCmd.Flags().BoolVar(&devFlag, "dev", false, "Get Dev Preview Version")
	listCmd.Flags().BoolVar(&showurls, "showurls", false, "List oc and openshift-install download URLs")
	listCmd.Flags().StringVar(&ocpversion, "ocpverstion", "", "List specific OCP version")
	//TODO add this once we update to cobra version 1.5
	//	listCmd.MarkFlagsMutuallyExclusive("stable", "fast", "candidate", "latest", "dev")
}
