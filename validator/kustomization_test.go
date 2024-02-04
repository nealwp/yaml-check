package validator

import (
	"testing"
)

type TestCase struct {
    Case []byte
    Want string 
}

func TestKustomizationValidator(t *testing.T) {

    goodYaml := `
resources:
  - deployment.yaml
  - service.yaml
  - virtual-service.yaml
  - authorization-policy.yaml
`
    badYaml := `
resources:
  - deployment.yaml
  - virtual-service.yaml
  - authorization-policy.yaml
`
    tests := []TestCase {
        { Case: []byte(goodYaml), Want: "" },
        { Case: []byte(badYaml), Want: "base/kustomization.yaml - \"resource\" missing item: service.yaml" },
    }

    for _, test := range tests {
        got, err := ValidateKustomization(test.Case)

        if err != nil {
            t.Fatal(err)
        }

        if got != test.Want {
            t.Fatalf("got %v, wanted %v", got, test.Want)
        }
    }
}
