package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main()  {
	listener, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalf("%s error occured",err)
	}
	defer listener.Close()

	for{
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("%s error occured", err)
			continue
		}
		go handling(conn)
	}
}

func handling(conn net.Conn)  {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		line := scanner.Text()
		fmt.Println(line)
	}
	defer conn.Close()
	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?

	fmt.Println("Code got here.")
}
