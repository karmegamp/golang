package main

import "fmt"

func main() {

	// 1. infinity loop
	i := 10
	for {
		fmt.Println("Infinity", i)
		i--
		if i <= 0 {
			break
		}
	}

	// 2. while loop
	j := 1
	for j <= 10 {
		fmt.Println("upto 10", j)
		j++
	}

	// 3. conditional for loop
	for i := 0; i < 10; i++ {
		fmt.Println("counter:", i)
	}

	// 4. range loop (array, slice, map, channel)

	// slice loop
	s := []int{1, 2, 3, 4, 5}
	for _, v := range s { // index, value
		fmt.Println(v)
	}

	// map loop
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	for k, v := range m { // key, value
		fmt.Println(k, v)
	}

	// channel loop
	ch := make(chan int)
	go func() {
		i := 0
		for i < 10 {
			ch <- i
			i++
		}
		close(ch)
	}()
	for v := range ch { // channel value
		fmt.Println(v)
	}

	// other form of reading channel
	// for {
	// 	if v, ok := <-ch; ok {
	// 		fmt.Println(v)
	// 	} else {
	// 		break
	// 	}
	// }

}
