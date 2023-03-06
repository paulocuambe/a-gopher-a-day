package main

import (
	"fmt"
	"log"
	"os/exec"
	"os"
)

func main() {
	cmd := exec.Command("ls", "..")

	log.Println("Running: ", cmd.String())

	output, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ouptut bellow: ")
	fmt.Println(string(output))

	cmd2 := exec.Command("cat", "../README.md")
	
	log.Println("Running: ", cmd2.String())
	o2, err := cmd2.Output()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ouptut bellow: ")
	fmt.Println(string(o2))

	// I found it weird that I need to specify the name of the command as first argument
	// even after specifying the complete path. But its okay, it works.
	cmd3 := &exec.Cmd{
		Path: "/usr/local/go/bin/go",
		Args: []string{"go", "version"},
		Stderr: os.Stderr,
	}

	log.Println("Running: ", cmd3.String())

	o3, err := cmd3.Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ouput bellow: ")
	fmt.Println(string(o3))
}
