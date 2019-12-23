package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Name  string
		Phone int
	}{
		{"Feniks", 9999999999},
		{"Alice", 8888888888},
		{"Vary", 7777777777},
		{"Bob", 8888888888},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)
}
