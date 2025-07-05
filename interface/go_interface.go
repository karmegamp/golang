package main

import (
	"fmt"
	"sort"
)

type printer interface {
	print()
}

type mydetails struct {
	name string
}

func (m mydetails) print() {
	fmt.Println(m.name)
}

func main() {

	var p printer
	p = mydetails{"Karmegam"}
	p.print()

	// Sort interface usage
	var names byName
	names = append(names, person{"kar1", 43})
	names = append(names, person{"kar3", 43})
	names = append(names, person{"kar2", 43})

	fmt.Print(names)
	sort.Sort(names)
	fmt.Print(names)
}

type person struct {
	name string
	age  int
}

type byName []person

func (b byName) Len() int           { return len(b) }
func (b byName) Less(i, j int) bool { return b[i].name < b[j].name }
func (b byName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
