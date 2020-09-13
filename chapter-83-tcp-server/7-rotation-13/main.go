package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main()  {
	listener, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalf("%s error connection", err)
	}
	defer listener.Close()

	for  {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handling(conn)
	}
}

func handling(conn net.Conn)  {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		line := strings.ToLower(scanner.Text())
		byteSLice := []byte(line)
		r := rotation13(byteSLice)

		fmt.Fprintf(conn, "%s - %s\n\n",line,r )
	}
}

func rotation13(byteString []byte) []byte  {
	var r13 = make([]byte, len(byteString))
	for i, v := range byteString{
		//ascii 97 - 122
		if v <= 109{
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return  r13
}
