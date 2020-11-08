package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mydb?charset=utf8mb4,utf8")
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/checkpoint", checkpoint)
	http.HandleFunc("/dashboard", dashboard)
	err = http.ListenAndServe(":8080", nil)
	checkErr(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to index")
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("error : %s", err)
		return
	}
}

func checkpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//validation
		//code here
		//check session, jwt

		//default route
		tpl.ExecuteTemplate(w, "login.html", nil)
	case http.MethodPost:
		//validation etc
		fmt.Fprint(w, "welcome to dashboard user")
	default:
		log.Fatal(http.StatusNotImplemented)
		return
	}
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to dashboard!") // write data to response
}
