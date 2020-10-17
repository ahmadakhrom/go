package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/doo", readCookies)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "_cookies_1",
		Value:   "my cookies value_1",
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24),
	})
	fmt.Fprintln(w, "error read _cookies_1")
}

func bar(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "_cookies_2",
		Value:   "my_cookies_value_2",
		//Path:    "/bar",
		Expires: time.Now().Add(time.Hour * 24),
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "_cookies_3",
		Value:   "my_cookies_value_3",
		//Path:    "/doo",
		Expires: time.Now().Add(time.Hour * 24),
	})
	fmt.Fprintln(w, "error read _cookies_2")
	fmt.Fprintln(w, "error read _cookies_3")
}

func readCookies(w http.ResponseWriter, r *http.Request) {

	cookies1, err := r.Cookie("_cookies_1")
	if err != nil {
		fmt.Fprintln(w, "error read cookies")
	} else {
		fmt.Fprintln(w, "cookies is: ", cookies1)
	}

	cookies2, err := r.Cookie("_cookies_2")
	if err != nil {
		fmt.Fprintln(w, "error read cookies")
	} else {
		fmt.Fprintln(w, "cookies is: ", cookies2)
	}

	cookies3, err := r.Cookie("_cookies_3")
	if err != nil {
		fmt.Fprintln(w, "error read cookies")
	} else {
		fmt.Fprintln(w, "cookies is: ", cookies3)
	}
}
