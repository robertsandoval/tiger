/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	channel    string
	latestFlag bool
	showurls   bool
	ocpversion string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:       "list",
	Short:     "List stuff, right now OCP versions. versions|version",
	ValidArgs: []string{"versions", "version"},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if args[0] == "versions" && validateChannel(channel) {
			versions := getVersions(channel)
			for _, ocp := range versions {
				fmt.Printf("%s-%d.%d -->  %d.%d.%d\n", ocp.channel, ocp.majorVersion, ocp.minorVersion, ocp.majorVersion, ocp.minorVersion, ocp.patchVersion)
			}
		} else if args[0] == "version" {
			getVersion(ocpversion)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&channel, "channel", "c", "stable", "Specify release channel. ")
	listCmd.Flags().BoolVar(&latestFlag, "latest", false, "Get Latest Version")
	listCmd.Flags().BoolVar(&showurls, "showurls", false, "List oc and openshift-install download URLs")
	listCmd.Flags().StringVar(&ocpversion, "ocpverstion", "", "List specific OCP version")
	//	listCmd.MarkFlagRequired("channel")

}
func validateChannel(channel string) bool {
	if channel != "stable" && channel != "candidate" && channel != "fast" {
		return false
	} else {
		return true
	}
}
