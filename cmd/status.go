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
	"os"
	"strings"

	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

type Cluster struct {
	clusterName     string
	api_endpoint    string
	cluster_console string
}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Displays currently created cluster directories and additional info",
	Long: `Displays currently created cluster directories and additional info

Currently lists ClusterName, OWNER, and API-Endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		clusters := []*Cluster{}
		fmt.Printf("%-24s%-16s%-24s\n", "ClusterName", "OWNER", "API Endpoint")
		path := os.Getenv("HOME") + "/.ocp/clusters"

		//get users cluster directories
		files, err := ioutil.ReadDir(path)

		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			temp := new(Cluster)
			temp.clusterName = file.Name()
			temp.api_endpoint = getAPIEndpoint(path, temp.clusterName)
			clusters = append(clusters, temp)
		}

		for i := range clusters {
			cluster := clusters[i]
			fmt.Printf("%-24s%-24s\n", cluster.clusterName, cluster.api_endpoint)
		}

	},
}

func getAPIEndpoint(path string, clustername string) string {

	f, err := os.Open(path + "/" + clustername + "/auth/kubeconfig")
	if err != nil {
		return "Kubeconfig file not found"
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "server") {
			//			fmt.Println(strings.TrimLeft(scanner.Text(), "    server:"))
			return strings.TrimLeft(scanner.Text(), "    server:")
		}

		line++
	}

	if err := scanner.Err(); err != nil {
		return "error"
	}
	return "No API Endpoint Found"
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	statusCmd.Flags().BoolP("all", "a", false, "print all clusters")
}
