package main

import (
	"fmt"
	"time"
)

func main() {

	simpleBackoff := func() {
		fmt.Println("Constant delay retry (15 sec)")
		duration := time.Now().Add(1 * time.Minute)
		fmt.Println("Sleeping until ", duration)
		for time.Now().Before(duration) {
			time.Sleep(15 * time.Second)
			fmt.Printf("timenow         %v\n", time.Now())
		}
		fmt.Println("Sleeping done ")

	}

	incrDelay := 1
	expBackoff := func() {
		fmt.Println("Exponential backoff: incremental delay retry")
		duration := time.Now().Add(1 * time.Minute)
		fmt.Println("Sleeping until ", duration)
		for time.Now().Before(duration) {
			sleeptime := time.Second << incrDelay
			fmt.Println("Sleeping for ", sleeptime)
			time.Sleep(sleeptime)
			fmt.Printf("timenow         %v\n", time.Now())
			incrDelay++
		}
		fmt.Println("Sleeping done ")

	}

	simpleBackoff()

	expBackoff()
}
