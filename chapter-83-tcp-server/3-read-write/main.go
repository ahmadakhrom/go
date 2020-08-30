package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main()  {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("%s error occured", err)
	}
	defer listener.Close()

	for{
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn)  {
	//adding timeout connection
	//...
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "I heard that you say : %s\n",line)
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?

	fmt.Println("Code got here.")
}
