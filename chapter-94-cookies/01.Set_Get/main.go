package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", readCookie)
	http.ListenAndServe(":8080", nil)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "_cookies",
		Value: "10",
		Path:  "/",
		//Domain:     "",
		Expires: time.Now().Add(time.Hour * 24),
		//RawExpires: "",
		//MaxAge:     0,
		//Secure:     false,
		//HttpOnly:   false,
		//SameSite:   0,
		//Raw:        "",
		//Unparsed:   nil,
	})
	fmt.Fprint(w, "cookies")
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	dataCookie, err := r.Cookie("_cookies")
	if err != nil {
		http.Error(w, err.Error(), 400) //http.StatusBadRequest
		return
	}
	fmt.Fprintln(w, "data cookies is:", dataCookie)
}
