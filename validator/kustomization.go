package validator

import (
	"log"

	"gopkg.in/yaml.v3"
)

type Kustomization struct {
	Resources []string `yaml:"resources"`
}

func ValidateKustomization(content []byte) {
	k := Kustomization{}

    err := yaml.Unmarshal(content, &k)

	if err != nil {
		log.Fatal(err)
	}

    return
}
