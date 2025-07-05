package main

import (
	"flag"
	"fmt"
)

func main() {

	max := flag.Int("max", 6, "the max value")
	min := flag.Int("min", 99, "the min value")
	str := flag.String("arg", "hello", "default string")
	fl := flag.Float64("temp", 38, "default temprature")

	flag.Parse()

	fmt.Println(*max) // read value by indirection (*)
	fmt.Println(*min)
	fmt.Println(*str)
	fmt.Println(*fl)
}
