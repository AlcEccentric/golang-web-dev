package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func serve(conn net.Conn) {
	body := "Hello World! \nMethod: %s \nURI: %s \n"
	scanner := bufio.NewScanner(conn)

	var i = 0
	var method string
	var uri string
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
		if i == 0 {
			method = strings.Fields(ln)[0]
			uri = strings.Fields(ln)[1]
		}
		i++
	}
	defer conn.Close()
	// body := "Hello World"
	fmt.Println("Method: " + method)
	fmt.Println("URI: " + uri)
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)+len(uri)+len(method))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	fmt.Fprintf(conn, body, method, uri)

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
