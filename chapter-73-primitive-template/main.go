package main

import "fmt"

func main()  {
	say := "hello word i'm here.."
	string := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>
Hello world
</title>
</head>
<body>
<h2>`+say+`</h2>
</body>
</html>`

fmt.Println(string)
}

//to build and run program : "go run 02.newServerMux.go > index.html"