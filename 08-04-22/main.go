package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

var Limit int

func main() {
	// initialize the global limit
	dirDepth := os.Getenv("DIR_DEPTH")
	if l, err := strconv.Atoi(dirDepth); err != nil {
		Limit = 4
	} else {
		Limit = l
	}

	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatal("Too few arguments.")
	}

	if args[0] == "ls" {
		for _, v := range args[1:] {
			fmt.Println(printFile(v, 1))
		}

	} else {
		fmt.Println("Command not supported. Only [ls] is supported.")
	}
}

func printFile(name string, depth int) string {
	f, err := os.ReadDir(name)

	if err != nil {
		log.Fatal(err)
	}

	// path.Base only leaves the last part of the path
	str := fmt.Sprintf("📂 %s\n", path.Base(name))

	// stop recursion if depth is greater than or equal to the limit
	if depth >= Limit {
		return str
	}

	for _, file := range f {
		// add space
		space := strings.Repeat("    ", depth)

		inf, err := file.Info()

		if err != nil {
			log.Fatal(err)
		}

		if inf.IsDir() {
			dname := fmt.Sprintf("%s/%s", name, file.Name())

			// recursively print files in the directory
			str += fmt.Sprintf("%s%s", space, printFile(dname, depth+1))
		} else {
			str += fmt.Sprintf("%s📄 %s\n", space, file.Name())
		}
	}

	return str
}
