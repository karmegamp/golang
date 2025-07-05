package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	ch := make(chan string)

	if len(os.Args) <= 1 {
		go fetch("https://google.com", ch)
		data := <-ch
		fmt.Println(data)
	} else {
		for _, v := range os.Args[1:] {
			go fetch(v, ch)
		}
		for range os.Args[1:] {
			data := <-ch
			fmt.Println(data)
		}
	}
}

func fetch(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	ch <- strconv.Itoa(len(data)) + " : " + url
}
