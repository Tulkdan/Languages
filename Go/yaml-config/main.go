package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Spec struct {
	Name        string             `yaml:"name"`
	Image       string             `yaml:"image"`
	Environment map[string]string  `yaml:"environment,omitempty"`
	Limits      *FunctionResources `yaml:"limits,omitempty"`
}

type FunctionResources struct {
	Memory string `yaml:"memory"`
	CPU    string `yaml:"cpu"`
}

func main() {
	spec := Spec{
		Image: "docker.io/functions/figlet:latest",
		Name:  "figlet",
	}

	bytesOut, err := yaml.Marshal(spec)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("figlet.yaml", bytesOut, os.ModePerm); err != nil {
		panic(err)
	}

	fmt.Printf("Wrote: figlet.yaml.. Ok.\n")
}
