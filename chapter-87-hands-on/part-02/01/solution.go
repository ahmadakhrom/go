package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalf("%s", err)
			continue
		}

		//write to connection
		io.WriteString(conn, " I see you connected")
		conn.Close()
	}
}