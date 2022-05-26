/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:           "get",
	Short:         "Display one or more resources. [versions|clusters]",
	SilenceUsage:  true,
	SilenceErrors: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return errors.New("Please use version|clusters")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println(len(args))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
