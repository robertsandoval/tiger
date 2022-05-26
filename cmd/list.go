/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stuff, right now OCP versions. versions|version",
}

func init() {
	rootCmd.AddCommand(listCmd)

}
