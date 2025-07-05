package main

import "fmt"

func main() {
	arr := []int{212, 5, 23, 199, 55, 777, 333, 64, 10, 55}
	fmt.Println("input", arr)

	l := len(arr) - 1
	fmt.Println("size", l)
	for i := 0; i < l; i++ {
		if arr[i] < arr[i+1] {
			arr[i+1], arr[i] = arr[i], arr[i+1]
		}
	}

	fmt.Println("smallest", arr[l])

}
