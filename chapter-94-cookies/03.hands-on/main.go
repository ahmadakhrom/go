package main

import (
	"io"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("_local")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:    "_local",
			Value:   "0",
			Path:    "/",
			Expires: time.Now().Add(time.Hour * 24),
		}
	}

	//conversion string to integer
	valCookies, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//conversion integer to string
	valCookies++

	//returns the original data type with an convert an integer to a string for the value of cookie
	c.Value = strconv.Itoa(valCookies)

	//set cookies
	http.SetCookie(w, c)
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, "<center> How many user access this page are: "+c.Value+"</center>")

}
