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
	"bufio"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"runtime"
)

var operatingsystem string
var tigerConfig = viper.New()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tiger",
	Short: "Utility command to manage OCP installs and binaries",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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

	operatingsystem = runtime.GOOS

	switch operatingsystem {
	case "windows":
		operatingsystem = "windows"
	case "darwin":
		operatingsystem = "mac"
	case "linux":
		operatingsystem = "linux"
	default:
		fmt.Printf("%s.\n", operatingsystem)
	}
	if operatingsystem == "windows" {
		fmt.Println("Command is only supported on Mac and Linux at this time")
		os.Exit(8008)
	}
	cobra.OnInitialize(initConfig)

}

// initConfig reads in config file and ENV variables if set.

//TODO Need set setup some variables from viper/env variables.
// Specifically need HOME, DownloadDir....etc
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".tiger" (without extension).
	tigerConfig.AddConfigPath(home)
	tigerConfig.SetConfigName(".tiger")
	tigerConfig.SetConfigType("yaml")
	tigerConfig.SetDefault("install-dir", os.Getenv("HOME")+"/.ocp/clusters")

	tigerConfig.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	_ = tigerConfig.ReadInConfig()

	//TODO add a variable to instruct how many versions back to query.
	// This will shorten up the back and forth requests to mirror.openshift.com

}

func runCmd(cmd *exec.Cmd) {
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {

		fmt.Println(err)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
}
