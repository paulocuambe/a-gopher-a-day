package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Api struct {
	Name        string  `json:"name" yaml:"name"`
	Version     string  `json:"version" yaml:"version"`
	Description string  `json:"description" yaml:"description"`
	Author      *Author `json"author" yaml:"author"`
}

type Author struct {
	Name       string `json:"name" yaml:"name"`
	Twitter    string `json:"twitter" yaml:"twitter"`
	Github     string `json:"github" yaml:"github"`
	Profession string `json:"profession" yaml:"profession"`
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

	// Reading yaml file
	log.Printf("\nYaml")

	f2, err := os.Open("api.yml")

	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(f2)
	if err != nil {
		log.Fatal(err)
	}

	var conf2 Api

	err = yaml.Unmarshal(b, &conf2)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v", conf2)
	log.Printf("%#v", conf2.Author)

}
