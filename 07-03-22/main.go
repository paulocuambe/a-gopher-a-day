package main

import (
	"os"
	"os/user"
	"fmt"
	"bufio"
	"log"
	"strings"
)

// go run . 
// what's the user name you want to look up?
// gosenx
// Display all the information about the user

// what's the group name you want to look up?
// root
// Display all th information about the group
func main() {
	// args := os.Args[1:]

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What's username you want information about? ")
	name, err := reader.ReadString('\n')

	name = strings.Trim(name, "\n")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", name)

	u, err := user.Lookup(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", u)
}