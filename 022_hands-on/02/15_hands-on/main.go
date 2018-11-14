package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func serve(conn net.Conn) {
	// this just works in Firefox without <!doctype html> ......
	body := "<h1>Hello World!</h1>"
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
	}
	defer conn.Close()
	// body := "Hello World"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	fmt.Fprintf(conn, body)

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
		go serve(conn)

	}

}
