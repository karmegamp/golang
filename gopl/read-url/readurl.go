package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	resp, err := http.Get("http://www.amazon.in")
	if err != nil {
		fmt.Println(err)
	}

	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(len(string(data)))
	for v := range strings.SplitSeq(string(data), "\n") {
		fmt.Println(v)
	}
}
