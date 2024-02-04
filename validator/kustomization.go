package validator

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Kustomization struct {
	Resources []string `yaml:"resources"`
}

func ValidateKustomization(content []byte) (string, error) {

	k := Kustomization{}

	err := yaml.Unmarshal(content, &k)

	if err != nil {
		return "", err
	}

	expected := []string{
		"deployment.yaml",
		"service.yaml",
		"virtual-service.yaml",
		"authorization-policy.yaml",
	}

	for i, e := range expected {
		if k.Resources[i] != e {
			return fmt.Sprintf("base/kustomization.yaml - \"resource\" missing item: %s", e), nil
		}
	}

	return "", nil
}
