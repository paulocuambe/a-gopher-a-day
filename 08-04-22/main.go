package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatal("Too few arguments.")
	}

	
	if args[0] == "ls" {
		fmt.Println(printFile(args[1], 1))
	}
}

func printFile(name string, depth int) string {
	f, err := os.ReadDir(name)

	if err != nil {
		log.Fatal(err)
	}

	str := fmt.Sprintf("📂 %s\n", path.Base(name))
	for _, file := range f {
		space := strings.Repeat("    ", depth)
		inf, err := file.Info()

		if err != nil {
			log.Fatal(err)
		}

		if inf.IsDir() {
			n := fmt.Sprintf("%s/%s", name, file.Name())
			str += fmt.Sprintf("%s%s", space, printFile(n, depth + 1))
		} else {
			str += fmt.Sprintf("%s📄 %s\n", space, file.Name())
		}
	}

	return str
}