package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "-c", "mkdir", "/data")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
