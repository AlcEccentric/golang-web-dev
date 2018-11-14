package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func responseTCP(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Fprintln(os.Stdout, scanner.Text())
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
		go responseTCP(conn)

	}

}
