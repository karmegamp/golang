package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/add", adder)
	http.ListenAndServe("localhost:8888", nil)
}

func adder(res http.ResponseWriter, req *http.Request) {

	// Read query param
	a, _ := strconv.Atoi(req.URL.Query().Get("a"))
	b, _ := strconv.Atoi(req.URL.Query().Get("b"))

	// Process query param
	c := a + b

	// Write response to client
	res.WriteHeader(200)
	fmt.Fprintf(res, strconv.Itoa(c))

	// Further way
	// io.WriteString(res, strconv.Itoa(c))

	// Other way to write response to client
	// s := strconv.Itoa(c)
	// var bu bytes.Buffer
	// bu.Write([]byte(s))
	// io.Copy(res, &bu)
}
