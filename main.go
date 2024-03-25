package main

import (
	"fmt"
	"github.com/codebeginer110/containers/sorted_map"
)

func main() {
	m := sorted_map.NewSortedMap[string]()
	m.Insert("a", 1)
	m.Insert("b", 2)
	m.Insert("c", 3)
	m.Insert("d", 4)
	m.Search("a")
	m.Range(func(v any) {
		fmt.Println(v)
	})
}
