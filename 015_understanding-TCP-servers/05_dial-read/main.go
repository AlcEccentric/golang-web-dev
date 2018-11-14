package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	// net.Dial combines net.Listen and li.Accept
	// it directly listen to a port and return a net.Conn
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}
