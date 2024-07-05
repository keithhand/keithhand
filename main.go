package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

const (
	outputFile   = "README.md"
	templateFile = "readme.tmpl"
	configFile   = "config.yaml"

	iconSize = 40
)

type Config struct {
	Name  string
	Desc  string
	About []string
	Langs Languages
}

type Languages struct {
	Active   []Language
	Previous []Language
}

type Language struct {
	Name string
	Href string
	Type string
}

func (Language) Size() int {
	return iconSize
}

func (l Language) Url() string {
	file := l.Name
	switch l.Type {
	default:
		file += "-original"
	case "wordmark":
		file += "-original-wordmark"
	case "plain":
		file += "-plain"
	}
	url := fmt.Sprintf("https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/%s/%s.svg", l.Name, file)
	return url
}

func main() {
	var config Config

	cFile, err := os.Open(configFile)
	if err != nil {
		log.Println(err.Error())
	}
	defer cFile.Close()

	if cFile != nil {
		decoder := yaml.NewDecoder(cFile)
		if err := decoder.Decode(&config); err != nil {
			log.Println(err.Error())
		}
	}

	tmpl, err := template.New(templateFile).ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	oFile, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(oFile, config)
	if err != nil {
		panic(err)
	}
}
