package main

import (
	"fmt"
)

func main() {

	// 1. int and string
	var y int // seperate declare and assignment
	y = 1
	var q string
	q = "hello"
	fmt.Println("seperate declare and assignment", y, q)

	var x int = 1 // merge declare and init
	var s string = "hello"
	fmt.Println("merge declare and init", x, s)

	x1 := 2 // shorthand notion do not need var keyword and type
	s1 := "hello shorthand"
	fmt.Println("shorthand", x1, s1)

	// 2. Array
	var arr [3]int // array must has size/length
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println("array", arr)

	// 3. Slice (prefer shorthand)
	var sl []int        // default slice length is zero, so it is uninitialized
	sl = make([]int, 2) // here it initialize with zero, use make for init
	sl[0] = 1
	sl[1] = 2
	fmt.Println("declare slice", sl)

	sli := make([]int, 10) // declare and initialize, use shorthand, all non-custom init required make call
	fmt.Println("shorthand slice default init", sli)

	slie := []int{1, 2} // declare and custom init, use shorthand and curly brace
	fmt.Println("shortand slice manual init", slie)

	// 4. map
	var mymap map[string]string
	mymap = make(map[string]string) // all non-custom init required make call
	mymap["one"] = "1"
	mymap["two"] = "2"
	fmt.Println("declare map and assign value", mymap)

	var m1 = map[string]string{"one": "1", "two": "2"}
	fmt.Println("decalre and custom init", m1)

	m2 := map[string]string{"one": "1", "two": "2"}
	fmt.Println("declare and custom init shorthand", m2)

	m3 := make(map[string]string)
	m3["one"] = "1"
	m3["two"] = "2"
	fmt.Println("declare shorthand and default init map", m3)

	// 5. struct
	type emp struct {
		name string
		des  string
	}

	var e emp
	e.name = "Karmegam"
	e.des = "Architect"
	fmt.Println("struct declare and assignment", e)

	var e1 emp = emp{name: "Kar", des: "Architect"}
	fmt.Println("struct declare and custom init", e1)

	e2 := emp{des: "Architect", name: "Kar"}
	fmt.Println("struct declare shorthand init details attr and value", e2)

	e3 := emp{"Kar", "Architect"}
	fmt.Println("struct declare shorthand init details only value", e3)

}
