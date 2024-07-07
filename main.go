package main

import (
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

const (
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

func (l Language) Svg() string {
	var sufx string
	switch l.Type {
	default:
		sufx = "original"
	case "wordmark":
		sufx = "original-wordmark"
	case "plain":
		sufx = "plain"
	}
	url := fmt.Sprintf("https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/%[1]s/%[1]s-%[2]s.svg", l.Name, sufx)
	return url
}

func readConfig(cfg string) (*Config, error) {
	var config *Config
	file, err := os.Open(cfg)
	if err != nil {
		return nil, fmt.Errorf("opening file %s: %w", cfg, err)
	}
	defer file.Close()
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			return nil, fmt.Errorf("decoding file %s: %w", cfg, err)
		}
	}
	return config, nil
}

func readTemplate(tmpl string) (*template.Template, error) {
	file, err := template.New(tmpl).ParseFiles(tmpl)
	if err != nil {
		return nil, fmt.Errorf("parsing template %s: %w", tmpl, err)
	}
	return file, nil
}

func writeOutput(out string, tmpl *template.Template, cfg *Config) (*os.File, error) {
	oFile, err := os.Create(out)
	if err != nil {
		return nil, fmt.Errorf("creating output %s: %w", out, err)
	}
	err = tmpl.Execute(oFile, cfg)
	if err != nil {
		return nil, fmt.Errorf("applying template to %s: %w", out, err)
	}
	return oFile, nil
}

var (
	outputFile   = "README.md"
	templateFile = "readme.tmpl"
	configFile   = "config.yaml"
)

func main() {
	cfg, err := readConfig(configFile)
	if err != nil {
		panic(err)
	}
	tpl, err := readTemplate(templateFile)
	if err != nil {
		panic(err)
	}
	_, err = writeOutput(outputFile, tpl, cfg)
	if err != nil {
		panic(err)
	}
}
