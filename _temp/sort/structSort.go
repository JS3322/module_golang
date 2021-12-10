package main

import (
	"fmt"
	"sort"
)

type Interface interface {
	Len() int
	// if i is small true
	Less(i, j int) bool
	Swap(i, j int)
}

type Student struct {
	name  string
	score float32
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].name < s[j].name
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// score sort struct init
type ByScore struct {
	Students
}

func (s ByScore) Less(i, j int) bool {
	return s.Students[i].score < s.Students[j].score
}

func main() {
	s := Students{
		{"JS1", 23.2},
		{"JS2", 51.2},
		{"JS3", 34.2},
	}

	sort.Sort(s)
	fmt.Println(s)

	sort.Sort(sort.Reverse(ByScore{s}))
	fmt.Println(s)
}
