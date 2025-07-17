package main

import "fmt"

func main() {
	v := 10
	p := &v
	fmt.Println("address:", p, " value:", *p)
	*p = 11
	fmt.Println("address:", p, " value:", *p)

	// pointer value access
	u := 10
	fmt.Println(&v == &v, &v == &u, &u == nil)
	fmt.Println(*&u == 10, *p == 11, *&u == 0)

	// extending variable visibility
	ff := f()
	*ff = 2
	fmt.Printf("%d", *ff)
	incr(ff)
	fmt.Printf("\n%d\n", *ff)

	// default pointer value is nil
	var q *int
	fmt.Println(q == nil)

	// unnamed variable, uses pointer indirection
	r := new(int)
	fmt.Println(*r)
	*r = 5
	fmt.Println(*r)

	// pointer to unname int factory
	s := newInt()
	t := newInt()
	*s = 6
	fmt.Println(s == t, *s, *t)
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p = *p + 2
	return *p
}

// factory, generates unnamed variable
func newInt() *int {
	return new(int)
}
