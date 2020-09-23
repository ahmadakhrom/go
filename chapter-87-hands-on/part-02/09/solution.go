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
			fmt.Println("Methooooooooooooooooood: ", respMethod)
			fmt.Println("Uriiiiiiiiiiiiiii: ", respUri)
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
	switch {
	case respMethod == "GET" && respUri == "/" :
		HandleIndex(conn)
	case respMethod == "GET" && respUri == "/apply" :
		HandleApply(conn)
	case respMethod == "POST" && respUri == "/apply" :
		HandlePost(conn)
	default:
		HandleDefault(conn)
	}
	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	//fmt.Println("Code got here.")
	//io.WriteString(conn, "in here responses written")
	//io.WriteString(conn, "HTTP/1.1 200 OK \r\n")

//	body := `<!DOCTYPE html>
//<html lang="en">
//<head>
//   <meta charset="UTF-8">
//   <title>Learn to GO</title>
//</head>
//<body>
//	<center>GET Index <br>
//   <a href="/route1">route1</a> <br>
//   <a href="/route2">route2</a> <br>
//</center>
//</body>
//</html>
//`
	//fmt.Fprintf(conn, "Content-Length: %d \r\n", len(body)) //will prin on Header
	//fmt.Fprintf(conn, "Content-Type: text/html \r\n")       //will print on Header
	//fmt.Fprintf(conn, "Key-Pass \r\n")       //will print on Header
	//io.WriteString(conn, "\r\n")
	////fmt.Fprintln(os.Stdout, "Hello world")
	//io.WriteString(conn,body)
}

func HandleIndex(c net.Conn)  {
	body :=`
			<html lang="en">
			<head>
  			<meta charset="UTF-8">
  					<title>Learn to GO</title>
			</head>
			<body>
					<center>GET Index <br>
  					<a href="/index">index</a> <br>
					<a href="/apply">apply</a> <br>
			</center>
			</body>
			</html>		
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	//fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func HandleApply(c net.Conn)  {
	body :=`
			<html lang="en">
			<head>
  			<meta charset="UTF-8">
  					<title>Learn to GO</title>
			</head>
			<body>
					<center>POS Apply <br>
  					<a href="/index">index</a> <br>
					<a href="/apply">apply</a> <br>
					<form method="POST" action="/apply">
						<input type="hidden" value="value text">	
						<input type="submit" value="submit value">	
					</form>
			</center>
			</body>
			</html>		
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	//fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func HandlePost(c net.Conn)  {
	body :=`
			<html lang="en">
			<head>
  			<meta charset="UTF-8">
  					<title>Learn to GO</title>
			</head>
			<body>
					<center>POS Apply <br>
  					<a href="/">index</a> <br>
  					<a href="/apply">apply</a> <br>
			</center>
			</body>
			</html>		
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func HandleDefault(c net.Conn)  {
	body :=`
			<html lang="en">
			<head>
  			<meta charset="UTF-8">
  					<title>Learn to GO</title>
			</head>
			<body>
					<center>Default
			</center>
			</body>
			</html>		
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
