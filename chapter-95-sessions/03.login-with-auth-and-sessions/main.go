package main

import (
	"net/http"
	"text/template"
	"time"

	"chapter-95-sessions/03.login-with-auth-and-sessions/modul"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type users struct {
	username string
	password string
}

var detailCookie = map[string]string{}
var detailUser = map[string]users{}

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
	//checkpoint cookie name and password
	u := modul.GetCookie(w, r, "_uid")

	//u := (w, r, "_uid")
	//validation logged, direct to dashboard
	//if enable the database, check value was encrypted and compare it.
	tpl.ExecuteTemplate(w, "index.html", u)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	//checkpoint GET
	if r.Method == http.MethodGet {
		//check sessoin _uid if nothing make it
		c, err := r.Cookie("_uid")
		if err != nil {
			d := uuid.New()
			c = &http.Cookie{
				Name:     "_uid",
				Value:    d.String(),
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour * 24 * 365), //1 year
			}
			http.SetCookie(w, c)
		}

		_, err = r.Cookie("_sessID")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		//decrypt value cookie v and check result of decrypted with username on DB
		//check value cookie username are registered on DB or no
		tpl.ExecuteTemplate(w, "dashboard.html", nil)
	}
	//end of checkpoint GET

	//checkpoint POST
	//example data user
	userDB := users{"andy", "$2a$04$GAgZBIywKpmH0jg7QTBsReALSk5C2au0PZj8p7mPqS3nsLtEogj8y"}

	if r.Method == http.MethodPost {
		usrnmForm := r.FormValue("username")
		pwdForm := r.FormValue("password")

		//validating username and password
		//if account valid save data to dataFormUsr
		FormUsrnm := users{
			usrnmForm,
			pwdForm,
		}

		//matching username and password on db
		if usrnmForm != userDB.username {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		errCompr := bcrypt.CompareHashAndPassword([]byte(userDB.password), []byte(FormUsrnm.password))
		if errCompr != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		//genrate with bycrypt the username and place on value cookie
		passRand, err := bcrypt.GenerateFromPassword([]byte(FormUsrnm.password), bcrypt.MinCost)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusInternalServerError)
			return
		}
		//username and pasword are valid then make a session ID
		c := &http.Cookie{
			Name:     "_sessID",
			Value:    string(passRand),
			HttpOnly: true,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 1), //1 year
		}
		http.SetCookie(w, c)

		//if you use database save cookiesUsrnm & cookiespwd store to database
		tpl.ExecuteTemplate(w, "dashboard.html", nil)
	}
	//end of checkpoint POST
}

//end of dashboard func

func logout(w http.ResponseWriter, r *http.Request) {
	//check session _uid
	c1, err := r.Cookie("_uid")
	if err != nil {
		http.Redirect(w, r, "/", 303) //http.StatusSeeOther
		return
	}

	//check session _sessID
	c2, err := r.Cookie("_sessID")
	if err != nil {
		http.Redirect(w, r, "/", 303) //http.StatusSeeOther
		return
	}

	//kill all cookies
	c1.MaxAge = -1
	c2.MaxAge = -1
	http.SetCookie(w, c1)
	http.SetCookie(w, c2)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

//end of logout func
