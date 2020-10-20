package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type users struct {
	username string
	password string
}

var dataUsrNm = map[string]string{}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/dashboard", dashboard)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

//func for directing to index page
func index(w http.ResponseWriter, r *http.Request) {
	//checkpoint cookie name
	_, err := r.Cookie("_sess_usrnm")
	if err != nil {
		tpl.ExecuteTemplate(w, "index.html", nil)
		return
	}

	_, err2 := r.Cookie("_sess_pwd")
	if err2 != nil {
		tpl.ExecuteTemplate(w, "index.html", nil)
		fmt.Println("or check here")
		return
	}

	//check value cookies if true pass, if false turn back page
	//if enable the database, check value was encrypted and compare.
	tpl.ExecuteTemplate(w, "dashboard.html", http.StatusOK)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	//if request is get check session
	//if request is post validating username & password

	//example data user like below
	dataUserDB := users{"andy", "123456"}

	//checkpoint #1
	if r.Method == http.MethodGet {
		fmt.Println("method get requested")

		c, err := r.Cookie("_sess_usrnm")
		if err != nil {
			http.Redirect(w, r, "/", 303) //http.StatusSeeOther
			return
		}

		//if user pwd already on DB
		//checking userName is already on DB (if use database)
		//c, _ := r.Cookie("_sess_usrnm")
		valUsrnm, _ := dataUsrNm[c.Value]
		if valUsrnm != dataUserDB.username {
			http.Redirect(w, r, "/", 303) //http.StatusSeeOther
			return
		}
		//check password
		cookiesValid(w, r, "_sess_pwd")
		c, err = r.Cookie("_sess_pwd")
		if err != nil {
			http.Redirect(w, r, "/", 303) //http.StatusSeeOther
			return
		}

		valPwd, _ := dataUsrNm[c.Value]
		if valPwd != dataUserDB.password {
			http.Redirect(w, r, "/", 303) //http.StatusSeeOther
			return
		}
		http.Redirect(w, r, "/dashboard", 200) //http.StatusOK
	}

	//end of checkpoint #2
	if r.Method == http.MethodPost {
		fmt.Println("method post requested")
		usrnmForm := r.FormValue("username")
		pwdForm := r.FormValue("password")

		//validating username and password
		//if accoout valid save data to dataFormUsr
		dataFormUsr := users{
			usrnmForm,
			pwdForm,
		}

		if usrnmForm != dataUserDB.username {
			http.Redirect(w, r, "/", 401) //http.StatusUnauthorized
			return
		}

		if pwdForm != dataUserDB.password {
			http.Redirect(w, r, "/", 401) //http.StatusUnauthorized
			return
		}

		cookiesUsrnm := &http.Cookie{
			Name:     "_sess_usrnm",
			Value:    dataFormUsr.username,
			HttpOnly: true,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 1), //1 year
		}

		cookiespwd := &http.Cookie{
			Name:     "_sess_pwd",
			Value:    dataFormUsr.password,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour * 1), //1 year
		}

		http.SetCookie(w, cookiesUsrnm)
		http.SetCookie(w, cookiespwd)

		tpl.ExecuteTemplate(w, "dashboard.html", nil)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	//cookiesValid(w, r, "_sess_usrnm")
	cookieUsrnm, err := r.Cookie("_sess_usrnm")
	if err != nil {
		http.Redirect(w, r, "/", 303) //http.StatusSeeOther
		return
	}

	//cookiesValid(w, r, "_sess_usrnm")
	cookiePwd, err := r.Cookie("_sess_pwd")
	if err != nil {
		http.Redirect(w, r, "/", 303) //http.StatusSeeOther
		return
	}

	//cookieUsrnm, _ = r.Cookie("_sess_usrnm")
	//cookiePwd, _ = r.Cookie("_sess_pwd")
	cookieUsrnm.MaxAge = -1
	cookiePwd.MaxAge = -1
	http.SetCookie(w, cookieUsrnm)
	http.SetCookie(w, cookiePwd)
	http.Redirect(w, r, "/", http.StatusSeeOther) //http.StatusSeeOther
	return

}

func cookiesValid(w http.ResponseWriter, r *http.Request, cookieName string) {
	_, err := r.Cookie(cookieName)
	if err != nil {
		http.Redirect(w, r, "/", 303) //http.StatusSeeOther
		return
	}
}

// I store username on value of first cookie
// and store password on value of second cookie
