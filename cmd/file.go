package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

//TODO catch the error
func downloadFile(filepath string, url string) error {
	fmt.Printf("Writing %s \nTo %s", url, filepath)
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func cleanDir(dir string) {

	fmt.Printf("Deleting:  %s\n", dir)
	if err := os.RemoveAll(dir); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Creating:  %s\n", dir)
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		fmt.Println(err)
	}

}
func expandFile(dir, file string) {
	//	fmt.Printf("\ndir: %s\nfile: %s\n", dir, file)
	//TODO fix the tar output and log it or something
	cmd := exec.Command("tar", "xvfz", file, "-C", dir)
	runCmd(cmd)
}
func createSymLinks(source, target string) {
	fmt.Printf("Creating symlink: %s -> %s\n", source, target)
	os.Remove(target)
	os.Symlink(source, target)
}
