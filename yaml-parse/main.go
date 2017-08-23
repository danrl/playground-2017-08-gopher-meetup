package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// type definitions

type yamlWalker struct {
	Name string
	Type string
}
type yamlFile struct {
	Walkers []yamlWalker
}

func main() {

	yamlFileContent := []byte(`---
walkers:
  - name: Night King
    type: scary ice dude
  - name: Cordell Walker
    type: texas ranger
    info: Do not play games with this one!
`)

	myData := yamlFile{}

	err := yaml.Unmarshal(yamlFileContent, &myData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	for _, walker := range myData.Walkers {
		fmt.Printf("I know about a walker named %v and it is a %v.\n",
			walker.Name,
			walker.Type)
	}
}
