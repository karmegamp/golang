package main

import (
	"fmt"
	"time"
)

/*
ch := make(chan int, 3)

// critical section: only 3 client can access at a time
foo := func(r int) {
	ch<-1
	sleep(3 Sec)
	defer func(){ <-ch }()
}

for range []int{1..10} {
	go foo()
}

*/

func main() {

	ch := make(chan int, 3)

	// blocking: max 3 client can access same time
	foo := func(r int) {
		ch <- 1
		fmt.Printf("\n%d entry time: %v", r, time.Now())
		time.Sleep(time.Second * 3)
		defer func() {
			<-ch
			fmt.Printf("\n%d exit time: %v", r, time.Now())
		}()
	}

	for _, r := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		go foo(r)
	}

	time.Sleep(time.Second * 10)
}
