package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {

	// NewScanner takes an io reader
	// return *Scanner type
	scanner := bufio.NewScanner(conn)

	// .Scan each time read a line (until EOF or \n )
	// return a bool whether there is something
	for scanner.Scan() {
		// .Text get the content read by .Scan
		/// bufio.ScanWords contains all the content of the message
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// ,which means .Scan would not detect an EOF in the connection that is still alive
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}
