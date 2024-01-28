package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Kustomization struct {
	Resources []string `yaml:"resources"`
}

type Service struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name   string `yaml:"name"`
		Labels struct {
			App string `yaml:"app"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
	Spec struct {
		Selector struct {
			App string `yaml:"app"`
		} `yaml:"selector"`
		Ports []struct {
			Name string `yaml:"name"`
			Port int    `yaml:"port"`
		} `yaml:"ports"`
	} `yaml:"spec"`
}

func main() {
	dat, err := os.ReadFile("./test_files/.deploy/base/kustomization.yaml")
	if err != nil {
		log.Fatal(err)
	}

	k := Kustomization{}

	err = yaml.Unmarshal(dat, &k)

	if err != nil {
		log.Fatal(err)
	}

	// we should have 4 elements in this list
	fmt.Print(k)

	dat, err = os.ReadFile("./test_files/.deploy/base/service.yaml")

	if err != nil {
		log.Fatal(err)
	}
	s := Service{}

	err = yaml.Unmarshal(dat, &s)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(s)
}
