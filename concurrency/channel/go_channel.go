package main

import "fmt"

func main() {
	ch := make(chan int)

	// Read channel must come first
	// if the channel is from local scope then no need to pass it as func argument
	// we can create function as an object
	print := func() {
		for v := range ch {
			fmt.Println(v)
		}
	}
	go print() // read channel must be used in goroutine

	// Write channel must start after listerner
	// write context can run in goroutine if we add some wait before main terminate
	func() {
		for i := range 10 {
			ch <- i
		}
	}()

	// var input string
	// fmt.Scanln(&input)
}
