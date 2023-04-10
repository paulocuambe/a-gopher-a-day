package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.OpenFile("index.html", os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get("http://gosenx.com")

	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	n, err := f.Write(b)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Wrote %d into index.html", n)
	}

	fmt.Println(string(b))
}