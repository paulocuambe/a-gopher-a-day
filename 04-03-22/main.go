package main

import (
	"encoding/json"
	"log"
	"os"
)

type Api struct {
	Version     string  `json:"version"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Author      *Author `json"author"`
}

func (a *Api) String() string {
	return a.Name
}

type Author struct {
	Name       string `json:"name"`
	Twitter    string `json:"twitter"`
	Github     string `json:"github"`
	Profession string `json:"profession"`
}

func main() {
	f, err := os.Open("api.json")

	if err != nil {
		log.Fatal(err)
	}

	var conf Api

	json.NewDecoder(f).Decode(&conf)

	log.Printf("%#v", conf)
	log.Printf("%#v", conf.Author)
}
