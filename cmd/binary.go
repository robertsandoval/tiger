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
	"os"
	"path"
	"path/filepath"

	"github.com/AlecAivazis/survey"
	"github.com/spf13/cobra"
)

var binaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "Sets up openshift-install and oc command versions",
	Long: `\nSets up openshift-install and oc command veresions
Lists a set of version options for the binaries and 
creates symlinks into $USER/bin `,
	Run: func(cmd *cobra.Command, args []string) {
		install_clients := "/usr/local/ocp/install-clients"
		version_choices := make(map[string]string)

		if install_clients != "" {
			versions, err := filepath.Glob(filepath.Join(install_clients, "oc*", "*-4*"))
			keys := make([]string, 0, len(versions))

			if err != nil {
				fmt.Println("error finding binary version directories")
			}

			for _, binary_path := range versions {
				version := path.Base(binary_path)
				version_choices[version] = binary_path
				keys = append(keys, version)

			}

			var qs = []*survey.Question{
				{
					Name: "Version",
					Prompt: &survey.Select{
						Message: "Choose a version:",
						Options: keys,
					},
				},
			}

			// the answers will be written to this struct
			answers := struct {
				Version string // survey will match the question and field names
			}{}

			// perform the questions
			err = survey.Ask(qs, &answers)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			var link_openshift_install = os.Getenv("HOME") + "/bin/openshift-install"
			var link_oc = os.Getenv("HOME") + "/bin/oc"
			var target_openshift_install = version_choices[answers.Version] + "/openshift-install"
			var target_oc = version_choices[answers.Version] + "/oc"

			os.Remove(link_openshift_install)
			os.Remove(link_oc)
			os.Symlink(target_openshift_install, link_openshift_install)
			os.Symlink(target_oc, link_oc)

		}
	},
}

func init() {
	setCmd.AddCommand(binaryCmd)
}
