package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main()  {
	//make a listen tcp
	listener, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer listener.Close()

	for  {
		conn, err := listener.Accept()
		if err != nil{
			log.Println(err.Error())
			continue
		}
		//use concurrency
		go handle(conn)
	}
}

//func handle to accomodate request and respond funcs
func handle(conn net.Conn)  {
	defer conn.Close()

	//read request
	request(conn)

	//response request
	respond(conn)
}

func request(conn net.Conn)  {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		lines := scanner.Text()
		fmt.Println(lines)

		if i == 0{
			//request line
			m := strings.Fields(lines)[0] //get method line
			u := strings.Fields(lines)[1] //get uri line
				fmt.Println("***Method",m)
				fmt.Println("***Method",u)
		}
		if lines == "" {
			//header are done
			break
		}
		i++
	}
}

func respond(conn net.Conn)  {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
