package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main()  {
	listener, err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalf("%s error connection",err)
	}
	defer listener.Close()

	for  {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn)  {
	defer conn.Close()

	//instruction
	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE: \n"+
		"SET Key Value \n"+
		"GET Key\n"+
		"Del Key\n\n"+
		"Example :\n"+
		"SET Fav Chococino\n" +
		"GET Fav\n\n\n")

	//read & write
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fs := strings.Fields(line)

		//logic
		if len(fs) < 1{
			continue
		}
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn,"%s\r\n",v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn,"Expected Value\r\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data [k] = v
		case "DEL":
		k := fs[1]
		delete(data, k)
		default:
			fmt.Fprintln(conn, "invalid command"+fs[0]+"\r\n")
			continue
		}
	}
}