package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Read from the client
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(buf[:n]))

	fmt.Fprintf(conn, "Echo - "+string(buf[:n]))
}

func handleConnectionNonBlocking(conn net.Conn) {
	defer conn.Close()

	for {
		// Set a deadline for reading data
		conn.SetReadDeadline(time.Now().Add(time.Second))

		// Read data from the client
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				// The read operation timed out, continue waiting for data
				continue
			} else {
				// Other errors (client disconnected, etc.)
				log.Println("Connection closed:", err)
				break
			}
		}

		// Print the received message
		fmt.Println("Received:", string(buf[:n]))

		// Set a deadline for writing data
		conn.SetWriteDeadline(time.Now().Add(time.Second))

		// Send a response to the client
		fmt.Fprintf(conn, "Echo!\n")
		if err != nil {
			log.Println("Error writing to client:", err)
			break
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer ln.Close()

	for {
		// Accept an incoming connection
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		// Handle the connection
		go handleConnectionNonBlocking(conn)
	}
}
