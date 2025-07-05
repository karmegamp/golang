package main

import "fmt"

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	// Keep read channels first
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Ch1:", msg1)
			case msg2 := <-c2:
				fmt.Println("Ch2 : ", msg2)
			}
		}
	}()

	// Write channels next
	go func() {
		for i := range 10 {
			c1 <- i
		}
	}()

	go func() {
		for i := range 10 {
			c2 <- i
		}
	}()

	var input string
	fmt.Scanln(&input)
}
