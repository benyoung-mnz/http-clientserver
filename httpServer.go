package main

import (
	"fmt"
	"net"
)

func handleConnetction(conn net.Conn) {
	// Read client message
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	// Print client message
	fmt.Println("Received message:", string(buf[:n]))

	// Close the connection
	conn.Close()

}

func main() {
	// listen for incoming connections
	ln, err := net.Listen("tcp", ":4000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer ln.Close()

	fmt.Println("Listening on port 4000...")

	for {
		// Accept new connections
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			continue
		}

		// Handle the connection in a new go routine
		go handleConnetction(conn)
	}
}
