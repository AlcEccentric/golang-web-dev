package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	// Get listener for tcp
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	// Close the listener after main
	defer li.Close()

	for {
		// Get connection from TCP listener
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// Write tcp message
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}
