package validator

import (
	"gopkg.in/yaml.v3"
)

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

func ValidateService(content []byte) error {

	s := Service{}

	err := yaml.Unmarshal(content, &s)

	if err != nil {
		return err
	}

	return nil
}
