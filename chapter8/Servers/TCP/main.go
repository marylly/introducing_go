package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Listener interface {
	// Accepts waits for and returns the next connection to the listener.
	Accept() (c net.Conn, err error)

	// Close closes the listener
	// Any blocked Accept operations will be unblocked and return errors.
	Close() error

	// Addr returns the listener's network address.
	Addr() net.Addr
}

func server() {
	// listen on a port
	ln, err := net.Listen("tcp",":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}

func client() {
	// connect to the server
	c, err := net.Dial("tcp","127.0.0.1:9999")
	if err != nil  {
		fmt.Println(err)
		return
	}

	// send the message
	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err1 := gob.NewEncoder(c).Encode(msg)
	if err1 != nil {
		fmt.Println(err1)
	}
	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}