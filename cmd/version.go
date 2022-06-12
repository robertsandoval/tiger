/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	channel                             string
	latestFlag                          bool
	showurls                            bool
	version                             string
	devpreview                          bool
	oc_tar_filename                     string
	openshift_install_tar_filename      string
	oc_tar_filename_full                string
	openshift_install_tar_filename_full string
	oc_binary_source_location           string
	openshift_install_source_location   string
	oc_binary_target_location           string
	openshift_install_target_location   string
	ocp_version_directory               string
	oc_url                              string
	openshift_install_url               string
)

// versionCmd represents the version command
var setVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.CalledAs())
		if !validateChannel(channel) {
			fmt.Println(cmd.Flag("channel").Usage)
		} else {
			fmt.Println("In Set version")
			fmt.Println(cmd.Parent().Use)
		}

	},
}

//Get Version - Used to list available variants of a version in each channel
// if --download will require 2 flags --ocp-version && --channel OR --latest
var getVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		//	ocp_version_directory = os.Getenv("HOME") + "/" + OCP_VERSION_DOWNLOAD_DIR
		ocp_version_directory = tigerConfig.GetString("version-dir")
		fmt.Printf("version dir %s\n", ocp_version_directory)
		if !validateChannel(channel) {
			fmt.Println(cmd.Flag("channel").Usage)
			fmt.Printf("channel: %s\nversion: %s\n", channel, version)
			return
		}
		//Build download location string ($HOME/.ocp/versions/4.10/...
		if cmd.Flag("latest").Changed {
			version = "latest"
		} else if cmd.Flag("devpreview").Changed {
			version = "dev-preview"
		} else {
			version = channel + "-" + version
		}

		//This sets the version director to a specific version dir
		ocp_version_directory = ocp_version_directory + "/" + version
		cleanDir(ocp_version_directory)
		//TODO lets see if we can do this without passing version possibly
		buildFilenames(version)
		downloadBinaries(ocp_version_directory, version)
		createSymLinks(oc_binary_source_location, oc_binary_target_location)
		createSymLinks(openshift_install_source_location, openshift_install_target_location)

		//Delete whats there
		//if latest if false we need channel and version
		// if version is defined we default to stable
		// if no flags change we default to to latest stable
		// if dev preview do dev preview
	},
}

var listVersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List versions available to download for OCP",
	Run: func(cmd *cobra.Command, args []string) {
		var versions []OCPversion
		if !validateChannel(channel) {
			fmt.Println(cmd.Flag("channel").Usage)
			fmt.Printf("channel: %s\nversion: %s\n", channel, version)
			return
		}
		if cmd.Flag("version").Changed {
			versions = getVersion(version)
		} else {
			versions = getVersions(channel)
		}
		for _, ocp := range versions {
			fmt.Printf("%s-%d.%d -->  %d.%d.%d\n", ocp.channel, ocp.majorVersion, ocp.minorVersion, ocp.majorVersion, ocp.minorVersion, ocp.patchVersion)
		}

	},
}

func init() {
	//GET
	getCmd.AddCommand(getVersionCmd)
	getVersionCmd.Flags().StringVarP(&channel, "channel", "c", "stable", "Specify release channel. ")
	getVersionCmd.Flags().BoolVar(&latestFlag, "latest", false, "Get Latest Version (same as --channel=latest --version=<latest version>)")
	getVersionCmd.Flags().BoolVar(&showurls, "showurls", false, "Show oc and openshift-install download URLs")
	getVersionCmd.Flags().StringVarP(&version, "version", "v", "", "Specificy OCP version")
	getVersionCmd.Flags().BoolVar(&devpreview, "devpreview", false, "Get Dev Preview version")

	//Set
	setCmd.AddCommand(setVersionCmd)
	setVersionCmd.Flags().StringVarP(&version, "ocpversion", "v", "", "Set OCP Version to Use")

	//List
	listVersionsCmd.Flags().StringVarP(&channel, "channel", "c", "stable", "Specify release channel: stable|fast|candidate|latest ")
	listVersionsCmd.Flags().StringVarP(&version, "version", "v", "", "List stable, fast, latest and candidate versions for a specific version  ")
	listVersionsCmd.Flags().BoolVar(&showurls, "showurls", false, "List oc and openshift-install download URLs")
	listCmd.AddCommand(listVersionsCmd)
}

func validateChannel(channel string) bool {
	if channel != "stable" && channel != "candidate" && channel != "fast" && channel != "latest" {
		return false
	} else {
		return true
	}
}

//Build filenames based on flags
func buildFilenames(version string) {
	//oc tar_filname
	oc_tar_filename = "openshift-client-" + operatingsystem + ".tar.gz"
	openshift_install_tar_filename = "openshift-install-" + operatingsystem + ".tar.gz"
	//dir + tar_filename
	oc_tar_filename_full = ocp_version_directory + "/" + oc_tar_filename
	openshift_install_tar_filename_full = ocp_version_directory + "/" + openshift_install_tar_filename

	oc_binary_source_location = ocp_version_directory + "/oc"
	openshift_install_source_location = ocp_version_directory + "/openshift-install"

	oc_binary_target_location = os.Getenv("HOME") + "/bin/oc"
	openshift_install_target_location = os.Getenv("HOME") + "/bin/openshift-install"

	//	oc_download_file_location := dir + "openshift-client-" + operatingsystem + ".tar.gz"
	oc_url = OCP_MIRROR_URL + version + "/" + oc_tar_filename
	openshift_install_url = OCP_MIRROR_URL + version + "/" + openshift_install_tar_filename
	fmt.Println("-------------")
	fmt.Printf("ocp_version_directory:%s\n", ocp_version_directory)
	fmt.Printf("oc_tar_filename: %s\n", oc_tar_filename)
	fmt.Printf("oc_tar_filename_full: %s\n", oc_tar_filename_full)
	fmt.Printf("openshift_install_tar_filename: %s\n", openshift_install_tar_filename)
	fmt.Printf("openshift_install_tar_filename_full: %s\n", openshift_install_tar_filename_full)

	fmt.Printf("oc_url: %s\n", oc_url)
	fmt.Printf("openshift_install_url: %s\n", openshift_install_url)
	fmt.Println("-------------")
}

// download
//TODO  download from list instead, maybe pull from viper config
func downloadBinaries(dir, version string) {

	downloadFile(oc_tar_filename_full, oc_url)
	expandFile(dir, oc_tar_filename_full)
	downloadFile(openshift_install_tar_filename_full, openshift_install_url)
	fmt.Println("----")
	expandFile(dir, openshift_install_tar_filename_full)
	fmt.Println("----")
}

//TODO for now do this...
func downloadDevPreview() {

}
