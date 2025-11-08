package main

import (
	"encoding/json"
	"io"
	"os"

	stolos_yoke "github.com/stolos-cloud/stolos/yoke-base/pkg/stolos-yoke"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/yokecd/yoke/pkg/flight"
)

func main() {
	// TODO: Change all these parameters
	airway := stolos_yoke.AirwayInputs{
		NamePlural:   "base",
		NameSingular: "bases",
		Kind:         "Base",
		Version:      "v1apha1",
		DisplayName:  "Base Scaffold (Change me!)",
	}

	stolos_yoke.Run[Base](airway, run)
}

func run() ([]byte, error) {
	// When this flight is invoked, the atc will pass the JSON representation of the Base instance to this program via standard input.
	// We can use the yaml to json decoder so that we can pass yaml definitions manually when testing for convenience.
	var base Base
	if err := yaml.NewYAMLToJSONDecoder(os.Stdin).Decode(&base); err != nil && err != io.EOF {
		return nil, err
	}

	if err := validateSpec(&base); err != nil && err != io.EOF {
		return nil, err
	}

	// Create the k8s resources
	return json.Marshal([]flight.Resource{
		createResources(base),
	})
}

func validateSpec(base *Base) error {
	// TODO : Validate the spec and set sane defaults

	if base != nil {
		return nil
	}

	return nil
}

// TODO : Implement functions which return standard k8s resources to create.
func createResources(base Base) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		Data: map[string]string{
			"HelloWorld": base.Spec.SomeProperty,
		},
	}
}
