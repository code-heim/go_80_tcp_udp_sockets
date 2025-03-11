package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	for {
		// Read a message from the client
		buf := make([]byte, 1024)
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error reading:", err)
			continue
		}

		// Print received message
		fmt.Printf("Received from %s: %s\n", clientAddr, string(buf[:n]))

		// Send response to the client
		_, err = conn.WriteToUDP([]byte("Hey client!"), clientAddr)
		if err != nil {
			log.Println("Error writing:", err)
		}
	}
}
