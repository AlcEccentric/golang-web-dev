package main

import (
	"io"
	"log"
	"net"
)

func responseTCP(conn net.Conn) {
	io.WriteString(conn, "I see you connected\n")
}
func main() {
	li, err := net.Listen("tcp", ":8080")
	defer li.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {

		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		responseTCP(conn)
		conn.Close()

	}

}
