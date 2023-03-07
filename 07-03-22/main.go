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
func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What's username you want information about? ")
	n, err := reader.ReadString('\n')
	// the returned string will end with \n
	n = strings.Trim(n, "\n")

	if err != nil {
		log.Fatal(err)
	}

	u, err := user.Lookup(n)

	if err != nil {
		log.Fatal(err)
	}

	tmp := "User\n\tUid: \t%s \n\tGid: \t%s \n\tName: \t%s \n\tUsername: \t%s \n\tHomeDir: \t%s\n"
	fmt.Printf(tmp, u.Uid, u.Gid, u.Username, u.Name, u.HomeDir)
}