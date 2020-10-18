package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("_sess")

	if err == http.ErrNoCookie {
		ID := uuid.New
		c = &http.Cookie{
			Name:  "_sess",
			Value: ID().String(),
			//Secure: true, just can  accessed with https
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour * 24 * 365), //1 year
		}
		http.SetCookie(w, c)
	}
	fmt.Fprintln(w, c)
}
