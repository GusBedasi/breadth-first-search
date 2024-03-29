package main

import (
	"fmt"
	"strings"
)

func main() {
	graph := make(map[string][]string, 2)

	graph["gustavo"] = []string{"nicolle", "thales", "chancha"}

	graph["nicolle"] = []string{"dani"}
	graph["thales"] = []string{"renan"}
	graph["chancha"] = []string{"kevin"}

	queue := graph["gustavo"]

	for len(queue) > 0 {
		person, others := queue[0], queue[1:]
		if nameFinishWithIN(person) {
			fmt.Printf("%s finish with IN\n", person)
			return
		} else {
			fmt.Printf("%s DO NOT finishes with IN\n", person)
			others = append(others, graph[person]...)
			queue = others
		}
	}
}

func nameFinishWithIN(name string) bool {
	return strings.HasSuffix(name, "in")
}
