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
		fmt.Fprintf(conn, "I heard you say: %s\n", scanner.Text())
	}
	defer conn.Close()
	fmt.Println("Code got here.")
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
		go serve(conn)

	}

}
