package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"sync"
	"time"
)

// Construct the reusable object generator
func objectFactory() any {
	return struct{}{}
}

// Create pool of reusable object collection cache
func objectPool() *sync.Pool {
	pool := &sync.Pool{
		New: objectFactory,
	}

	// pre create 10 object in the pool
	for i := range 10 {
		d := pool.New()
		d = i
		pool.Put(d)
	}

	// test for not returning the object to pool
	fmt.Println("Getting random object from pool : ", pool.Get())
	fmt.Println("Getting random object from pool : ", pool.Get())
	fmt.Println("===============================================")

	return pool
}

// TCP Server
func server() {
	// get the pool objects a cache
	cache := objectPool()

	// server listen port
	ln, _ := net.Listen("tcp", ":9999")

	// serve connection
	for {
		c, _ := ln.Accept()

		// get & return pool object for low memory footprint
		obj := cache.Get()
		go handleConnection(c, obj)
		cache.Put(obj)
	}
}

// Handle request
func handleConnection(c net.Conn, obj any) {
	var msg string
	gob.NewDecoder(c).Decode(&msg)
	fmt.Println("\nReceived : ", msg, "\nUsing pool object: ", obj)
	fmt.Println("===============================================")

	// type assertion, and simulate delay processing client
	if v, ok := obj.(int); ok {
		time.Sleep(time.Second << v)
	}

	defer c.Close()
}

// Client connect
func client(i int) {

	// using 127.0.0.1 give panic in gob.NewEncoder
	c1, _ := net.Dial("tcp", "localhost:9999")
	defer c1.Close()

	msg := fmt.Sprintf("Hello dear from client %d", i)
	fmt.Println("Client ", i, "Sending : ", msg)

	// loopback IP for c1 will panic here
	gob.NewEncoder(c1).Encode(&msg)
}

// Another client for random connect
func clientRandom() {
	c2, _ := net.Dial("tcp", "localhost:9999")
	msg := "Hello dear (Random)"
	fmt.Println("Client (Random) sending : ", msg)
	gob.NewEncoder(c2).Encode(&msg)
	c2.Close()
}

func main() {

	// start server
	go server()

	// simulate client connection
	for i := range 4 {
		go client(i)
		time.Sleep(time.Second << i)

		if i%2 == 0 {
			go clientRandom()
		}
	}

	var in string
	fmt.Scanln(&in)
}
