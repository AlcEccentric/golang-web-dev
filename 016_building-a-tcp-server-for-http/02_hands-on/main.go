package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func request(c net.Conn) string {
	rd := bufio.NewScanner(c)
	i := 0
	var uri string
	for rd.Scan() {
		line := rd.Text()
		fmt.Println(line)
		if i == 0 {
			uri = strings.Fields(line)[1]
			fmt.Println("***URI IS", uri)
		}
		if line == "" {
			break
		}
		i++
	}
	return uri
}

func response(c net.Conn, uri string) {
	body := "<!DOCTYPE html><html lang='en'><head><meta charet='UTF-8'><title></title></head><body><strong>" + uri + "</strong></body></html"
	fmt.Fprintf(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(c, "Content-Type: text/html\r\n")
	fmt.Fprintf(c, "\r\n")
	fmt.Fprintf(c, body)
}

func handle(c net.Conn) {

	// Notice to close conn after using
	defer c.Close()
	uri := request(c)
	response(c, uri)
}

func main() {
	// notice first parameter is protocal
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err.Error())
	}
	// Notice to close listenner after func
	defer li.Close()

	// Notice to use infinite loop to get connection forever
	// from specific port.
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err.Error())
		}
		go handle(conn)

	}

}
