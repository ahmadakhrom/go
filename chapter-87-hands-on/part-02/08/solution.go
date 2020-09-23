package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	var i int
	var respMethod, respUri string

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if i == 0 {
			// we're in REQUEST LINE
			xs := strings.Fields(line)
			respMethod = xs[0]
			respUri = xs[1]
			fmt.Println("Method: ", respMethod)
			fmt.Println("Uri: ", respUri)
		}

		if line == "" {
			fmt.Println("end of http request header")
			break
		}
	i++

		//if line != "" {
		//	fmt.Println("this is end of the http request headers")
		//	break
		//}

	}
	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	//fmt.Println("Code got here.")
	//io.WriteString(conn, "in here responses written")
	io.WriteString(conn, "HTTP/1.1 200 OK \r\n")

	body := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Learn to GO</title>
</head>
<body>
	<center>HOLY COW THIS IS LOW LEVEL</center>
</body>
</html>
`
	fmt.Fprintf(conn, "Content-Length: %d \r\n", len(body)) //will prin on Header
	fmt.Fprintf(conn, "Content-Type: text/html \r\n")       //will print on Header
	fmt.Fprintf(conn, "Key-Pass \r\n")       //will print on Header
	io.WriteString(conn, "\r\n")
	//fmt.Fprintln(os.Stdout, "Hello world")
	io.WriteString(conn,body)

}
