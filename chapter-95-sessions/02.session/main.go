package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
)

//Client are type for form field on frontend
type client struct {
	Username  string
	Firstname string
	Lastname  string
}

var tpl *template.Template

//DBsess for session user
var DBsess = map[string]string{} //session ID, User ID

//DBuser for user
var DBuser = map[string]client{} //user ID, user

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/doo", doo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	d := uuid.New()
	c, err := r.Cookie("_sess")

	//if client already has a cookie this code will be skipped and go to get a user
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:     "_sess",
			Value:    d.String(),
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour * 24 * 365), //1 year
		}
		http.SetCookie(w, c)
	}

	//get user
	var dataUser client
	username, status := DBsess[c.Value] // res contained status true or false  //xxxxxxxx
	if status == true {
		dataUser = DBuser[username] //zzzzzz
	}
	fmt.Println("line 56 --- Data user is :", dataUser)
	//process submission
	if r.Method == http.MethodPost {
		uForm := r.FormValue("username")
		fForm := r.FormValue("firstname")
		lForm := r.FormValue("lastname")

		dataUser = client{
			uForm,
			fForm,
			lForm,
		}
		DBsess[c.Value] = uForm  //map[xxxxxxxx-xxxxx-xxxxx-xxxxx-xxxxxxxxxxxxx]
		DBuser[uForm] = dataUser //map:{username: username firstname lastname}
	}
	fmt.Println("line 71 --- DBsess is :", DBsess)
	fmt.Println("line 72 --- DBuser is :", DBuser)
	tpl.ExecuteTemplate(w, "index.html", dataUser)
}

func bar(w http.ResponseWriter, r *http.Request) {
	//get cookie
	c, err := r.Cookie("_sess")
	if err != nil {
		http.Redirect(w, r, "/", 303) //http.StatusSeeOther
		return
	}

	username, status := DBsess[c.Value] //status contained status true or false
	if status != true {
		http.Redirect(w, r, "/", 303) //http.StatusSeeOther
		return
	}

	dataUser := DBuser[username]
	tpl.ExecuteTemplate(w, "bar.html", dataUser)
}

func doo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("_sess")
	if err != nil {
		http.Redirect(w, r, "/", 303) //http.StatusSeeOther
		return
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", 303) //http.StatusSeeOther
}
