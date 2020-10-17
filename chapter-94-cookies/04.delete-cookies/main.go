package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/doo", doo)

	http.ListenAndServe(":8080", nil)
}

//index page
func index(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `<center> <a href="/foo">write cookies>></center>`)
}

//setup cookies
func foo(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "_local",
		Value:  "0",
		Domain: "",
		Path:   "/",
	})
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `<center> cookies is set</center>`)
	io.WriteString(w, `<center> <a href="/bar">list all cookies>></center>`)
	io.WriteString(w, `<center> <a href="/">index>></center>`)
}

//list of cookies
func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("_local")
	if err != nil {
		http.Redirect(w,r,"/foo", 303) //see other
	}
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, "<center> value cookies: "+c.Value+"</center>")
	io.WriteString(w, `<center> <a href="/doo">delete cookies</center>`)
	io.WriteString(w, `<center> <a href="/">index>></center>`)
}

//delete cookies
func doo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("_local")
	if err != nil {
		http.Redirect(w,r,"/foo", 303) //see other
	}
	//delete cookie
	c.MaxAge = -1

	http.SetCookie(w, c)
	http.Redirect(w, r, "/",303)
}
//index -> set cookies -> cookie list -> delete cookies -> cookies list