package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main()  {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("%s error occured",err)
	}

	defer listener.Close()

	for  {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("%s error occured", err)
			continue
		}

		io.WriteString(conn,"\n Hello from TCP Server \n")
		fmt.Fprintln(conn, "how is your day?")
		fmt.Fprintf(conn, "%v", "well, I hope")

		conn.Close()
	}
}
