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

	body := "<h1>This is root  </h1>"
	scanner := bufio.NewScanner(conn)
	i := 0
	var m, u string
	for scanner.Scan() {

		ln := scanner.Text()
		if ln == "" {
			break
		}
		if i == 0 {
			u = strings.Fields(ln)[1]
			m = strings.Fields(ln)[0]
		}
		i++
	}
	defer conn.Close()
	// body := "Hello World"
	fmt.Printf("reach here %s %s \n", m, u)
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	// this just works in Firefox without <!doctype html> ......
	// ugly design just for ease
	if m == "GET" {
		if u == "/" {
			fmt.Fprintf(conn, "<h1>Get root</h1>")
		} else if u == "/apply" {
			fmt.Fprintf(conn, "<h1>Get apply</h1>")
		}

	} else if m == "POST" {
		if u == "/apply" {
			fmt.Fprintf(conn, "<h1>Post apply</h1>")
		}
	}

}
func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {

		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go serve(conn)

	}

}
