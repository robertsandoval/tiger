package main

import (
	"log"
	"os/exec"
	"os"
	"fmt"
)

func main() {
	path := os.Getenv("HOME") + "/bin/oc"
	fmt.Println(path)
	cmd := exec.Command("/home/sandoval/bin/oc", "version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
