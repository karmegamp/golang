package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// response message
type Response struct{}

func (s *Response) Negate(i int, reply *int) error {
	*reply = -i
	return nil
}

func (s *Response) Echo(msg string, reply *string) error {
	*reply = "echo " + msg
	return nil
}

func server() {
	rpc.Register(new(Response)) // note 1, rpc.Register for message structure
	ln, _ := net.Listen("tcp", ":9999")
	for {
		c, _ := ln.Accept()
		go rpc.ServeConn(c) // note 2, rpc.ServeConn
	}
}

func client() {
	c, _ := rpc.Dial("tcp", "localhost:9999") // note 3, rpc.Dial
	var res int
	c.Call("Response.Negate", int(2), &res) // Remote procedure to invoke, input, output
	fmt.Println(res)
	var echo string
	c.Call("Response.Echo", "hello", &echo)
	fmt.Print(echo)
}

func main() {

	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
