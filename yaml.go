package main

import (
		"fmt"
		"log"

		"gopkg.in/yaml.v2"
)

var data = `
tasks:
  deploy:
    - command: git clone
      desc: Fetching new commits
      dir: /boo/boo
    - command: bundle install
      desc: Bundle install
      sudo: yes
      dir: /boo/boo
`

type T struct {
	Tasks struct{
		Deploy []struct{",flow"
	}
}

func main() {
		t := T{}

		err := yaml.Unmarshal([]byte(data), &t)
		if err != nil {
				log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- t:\n%v\n\n", t)

		d, err := yaml.Marshal(&t)
		if err != nil {
				log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- t dump:\n%s\n\n", string(d))

		m := make(map[interface{}]interface{})

		err = yaml.Unmarshal([]byte(data), &m)
		if err != nil {
				log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- m:\n%v\n\n", m)

		d, err = yaml.Marshal(&m)
		if err != nil {
				log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- m dump:\n%s\n\n", string(d))
}