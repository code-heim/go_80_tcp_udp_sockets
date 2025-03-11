package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Create a TCP socket
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	// Send a message
	fmt.Fprintf(conn, "Hello!")

	// Read a response
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(buf[:n]))

	conn.Close()
}
