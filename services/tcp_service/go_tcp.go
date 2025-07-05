package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	ln, _ := net.Listen("tcp", ":9999")
	for {
		c, _ := ln.Accept()
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	var msg string
	var i int = 1
	for i < 10 {
		gob.NewDecoder(c).Decode(&msg)
		fmt.Println("Received : ", msg)
		i++
	}
	defer c.Close()
}

func client1() {
	var i int = 1
	c1, _ := net.Dial("tcp", "localhost:9999") // using 127.0.0.1 give panic in gob.NewEncoder
	defer c1.Close()

	// panic recovery should be paired with defer
	// thatway it gaurantee to catch panic
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	for i < 10 {
		msg := "Hello dear"
		fmt.Println("Client1 Sending : ", msg)
		gob.NewEncoder(c1).Encode(&msg) // loopback IP for c1 will panic here
		i++
	}
}

func client2() {
	c2, _ := net.Dial("tcp", "localhost:9999")
	msg := "Hello dear"
	fmt.Println("Client2  Sending : ", msg)
	gob.NewEncoder(c2).Encode(&msg)
	c2.Close()
}

func main() {
	go server()
	go client1()
	go client2()

	var in string
	fmt.Scanln(&in)
}
