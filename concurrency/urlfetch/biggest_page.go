package main

import (
	"fmt"
	"io"
	"net/http"
)

type page struct {
	url  string
	size int
}

func main() {

	results := make(chan page)

	urls := []string{"http://www.amazon.com",
		"http://www.google.com",
		"http://www.wikipedia.com",
	}

	// Read channel
	go func() {
		var biggest page
		for range urls {
			res := <-results // blocking call
			fmt.Println(res)
			if biggest.size < res.size {
				biggest = res
			}
		}
		fmt.Println("Biggest page", biggest.url, biggest.size)
	}()

	// Write channel
	for _, url := range urls {
		go func() {

			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, _ := io.ReadAll(res.Body)
			results <- page{url, len(bs)}
		}()
	}

	// defer func() { close(results) }()

	var input string
	fmt.Scanln(&input)
}
