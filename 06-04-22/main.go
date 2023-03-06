package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "..")

	log.Println("Running: ", cmd.String())

	output, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))

	cmd2 := exec.Command("cat", "../README.md")
	
	log.Println("Running: ", cmd2.String())
	o2, err := cmd2.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(o2))
}
