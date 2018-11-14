package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		fmt.Fprintln(os.Stdout, scanner.Text())
	}
	defer conn.Close()
	body := "Hello World"
	fmt.Println("Code got here.")
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	fmt.Fprint(conn, body)

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
