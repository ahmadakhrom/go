package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"chapter-95-sessions/03.login-with-auth-and-sessions/modul"
	"chapter-95-sessions/03.login-with-auth-and-sessions/user"

	"golang.org/x/crypto/bcrypt"
)

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
	//checkpoint cookies
	//check _uid cookie
	modul.GetCookie(w, r, "_uid")

	//check if there is no "_sessID" cookie
	c, err := r.Cookie("_sessID")
	if err != nil {
		tpl.ExecuteTemplate(w, "index.html", nil)
		return
	}

	_, true := user.Cookie[c.Value]
	if !true {
		fmt.Println("no value on x line 51", true)
	}

	i, err := r.Cookie("_role")
	if err != nil {
		tpl.ExecuteTemplate(w, "index.html", i)
		fmt.Println("hold up")
		return
	}

	//role checker
	var data user.M
	data = user.M{
		"sessID": c.Name, "sessIdVal": c.Value, "role": i.Value,
	}
	switch i.Value {
	case "1":
		tpl.ExecuteTemplate(w, "dashboard.html", data)
		return
	case "2":
		fmt.Fprintln(w, "your entered admin page")
		return
	case "3":
		fmt.Fprintln(w, "your entered client page")
		return
	default:
		fmt.Fprintln(w, "forbidden to enter any page")
		return
	}

	//if enable the database, check value are encrypted and compare it. but it can be increase latency
	//how solve it? use a jwt
	//set uri to "domain/dashboard"

	//http.Redirect(w, r, "/dasboard", http.StatusOK)
	//tpl.ExecuteTemplate(w, "dashboard.html", data)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	//example data user
	userDB := user.Users{
		Username: "andy",
		Password: "$2a$04$GAgZBIywKpmH0jg7QTBsReALSk5C2au0PZj8p7mPqS3nsLtEogj8y",
		Role:     "1",
	}

	//checkpoint GET
	if r.Method == http.MethodGet {
		//check cookie "_uid" if nothing make it
		modul.GetCookie(w, r, "_uid")

		_, err := r.Cookie("_sessID")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		i, err := r.Cookie("_role")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		//role checker
		switch i.Value {
		case "1":
			tpl.ExecuteTemplate(w, "dashboard.html", nil)
			return
		case "2":
			fmt.Fprintln(w, "your entered admin page")
			return
		case "3":
			fmt.Fprintln(w, "your entered client page")
			return
		default:
			fmt.Fprintln(w, "forbidden to enter any page")
			return
		}

		//decrypt value cookie v and check result of decrypted with username on DB
		//check value cookie username are registered on DB or no
		//tpl.ExecuteTemplate(w, "dashboard.html", nil)
	}
	//end of checkpoint GET

	//checkpoint POST
	if r.Method == http.MethodPost {
		usrnmForm := r.FormValue("username")
		pwdForm := r.FormValue("password")
		role := r.FormValue("role")

		//validating username and password
		//if account valid save data to dataFormUsr
		FormUsrnm := user.Users{
			Username: usrnmForm,
			Password: pwdForm,
			Role:     role,
		}

		//matching username and password on db
		errCompr := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(FormUsrnm.Password))
		if errCompr != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		//generate with bycrypt the username and place on value cookie
		passRand, err := bcrypt.GenerateFromPassword([]byte(FormUsrnm.Password), bcrypt.MinCost)
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
		user.Cookie[c.Value] = c.Value

		//set cookie for role user
		d := &http.Cookie{
			Name:     "_role",
			Value:    role,
			HttpOnly: true,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 1), //1 year
		}
		http.SetCookie(w, d)

		//if you use database save cookiesUsrnm & cookiespwd store to database
		tpl.ExecuteTemplate(w, "dashboard.html", nil)
	}
	//end of checkpoint POST
}

//end of dashboard func

func logout(w http.ResponseWriter, r *http.Request) {
	//check cookie _uid
	c1, err := r.Cookie("_uid")
	if err != nil {
		http.Redirect(w, r, "/", 303) //http.StatusSeeOther
		return
	}

	//check cookie _sessID
	c2, err := r.Cookie("_sessID")
	if err != nil {
		http.Redirect(w, r, "/", 303) //http.StatusSeeOther
		return
	}

	//check cookie _role
	c3, err := r.Cookie("_role")
	if err != nil {
		http.Redirect(w, r, "/", 303) //http.StatusSeeOther
		return
	}

	//delete data on var DetailCookie
	delete(user.Cookie, c2.Value)

	//kill all cookies
	c1.MaxAge = -1
	c2.MaxAge = -1
	c3.MaxAge = -1

	http.SetCookie(w, c1)
	http.SetCookie(w, c2)
	http.SetCookie(w, c3)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
	//end of logout func
}
