/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	installDir string
	versionDir string
	config     string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init sets up the configuration for the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		setUpConfig()
	},
}

func init() {
	var installs = os.Getenv("HOME") + "/.ocp/clusters"
	var versions = os.Getenv("HOME") + "/.ocp/versions"
	initCmd.Flags().StringVarP(&installDir, "install-dir", "i", installs, "Set directory where installs go")
	initCmd.Flags().StringVarP(&versionDir, "version-dir", "v", versions, "Set directory where oc and openshift-install versions go")
	initCmd.Flags().StringVar(&config, "config", "", "Config to use if not taking defaults")
	rootCmd.AddCommand(initCmd)
}
func setUpConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".tiger" (without extension).
	tigerConfig.AddConfigPath(home)
	tigerConfig.SetConfigName(".tiger")
	tigerConfig.SetConfigType("yaml")
	tigerConfig.Set("install-dir", installDir)
	tigerConfig.Set("version-dir", versionDir)

	tigerConfig.AutomaticEnv() // read in environment variables that match
	fmt.Println("here")
	// If a config file is found, read it in.
	if err := tigerConfig.ReadInConfig(); err != nil {
		//we got an error trying to read the config
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//our error was a ConfigFileNotFound error. Lets try to create it
			if emptyFile, err := os.Create(home + "/.tiger.yaml"); err != nil {
				//we could not create the file. Lets log that but keep going using defaults
				fmt.Println("Could not create $HOME/.tiger.yaml file. Using defaults")
			} else {
				fmt.Println("Created file $HOME/.tiger.yaml")
				emptyFile.Close()
			}
		} else {
			//we could not read the file but it wasn't a ConfigNotFound error
			fmt.Println("We got the following error trying to read in file %s", err)
		}
	} else {
		//we read in the file...do we need to do anything else?
		fmt.Println("Using file located at $HOME/.tiger.yaml")
	}

	if err := tigerConfig.WriteConfig(); err != nil {
		fmt.Println("Error writing config file")
	} else {
		fmt.Println("Writing to: $HOME/.tiger.yaml")
	}
}
