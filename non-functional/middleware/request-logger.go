package main

import (
	"fmt"
	"net/http"
)

/*

	Middleware is a wrapper function, which accept and returns http handler function
	Basically we can add pre and post processing logic for our request!

*/

func main() {
	http.HandleFunc("/", checkMethod(requestlog(handler)))
	http.ListenAndServe("localhost:9090", nil)
}

func requestlog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Host :", r.Host)
		fmt.Println("Method : ", r.Method)

		next(w, r)

		fmt.Println("Post processing log")
	}
}

func checkMethod(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			fmt.Println("Valid method")
			next(w, r)
		} else {
			fmt.Println("Invalid method")
		}
	}
}

func handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello dear!"))
}
