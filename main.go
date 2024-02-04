package main

import (
	"log"
	"os"

	"github.com/nealwp/yaml-check/validator"
)

type Validation func(content []byte) error

type YamlFile struct {
	Path     string
	Validate Validation
}

func main() {

	files := []YamlFile{
		{Path: "./test_files/.deploy/base/kustomization.yaml", Validate: validator.ValidateKustomization},
		{Path: "./test_files/.deploy/base/service.yaml", Validate: validator.ValidateService},
	}

	for _, file := range files {
		content, err := os.ReadFile(file.Path)
		if err != nil {
			log.Fatal(err)
		}

		file.Validate(content)
	}

	return
}
