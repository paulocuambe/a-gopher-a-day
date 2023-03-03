package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) >= 1 {
		f, err := os.Open(arguments[0])

		if err != nil {
			log.Fatal(err)
		}

		b, err := io.ReadAll(f)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
	} else {
		log.Fatal("Please provide a filename or path to a file")
	}
}
