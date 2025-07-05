package main

import "fmt"

func main() {

	// create 10 goroutines
	for i := 1; i <= 10; i++ {
		go fmt.Println("Hello")
	}

	// without below statement, we will not see the above output, why?
	// because the main is a parent goroutine which exit immediately.
	// All goroutine terminated when main goroutine end. But if we raise
	// the loop number to 100 then we may get a chance to execute few
	// goroutine before main terminates!
	var input string
	fmt.Scanln(&input)
}
