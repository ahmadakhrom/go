package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"chapter-96-mysql/02.simple-app/modul"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var tpl *template.Template

//M type
type M map[string]string

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	//this example from Mr. todd McLeod while connecting the app with AWS
	//db, err = sql.Open("mysql", "awsuser:mypassword@tcp(mydbinstance.cakwl95bxza0.us-west-1.rds.amazonaws.com:3306)/test02?charset=utf8")
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/mydb?charset=utf8mb4,utf8")
	modul.CheckErr(err)
	defer db.Close()

	err = db.Ping()
	modul.CheckErr(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/read", read)
	http.HandleFunc("/find", findUser)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update-action", updateAction)
	http.HandleFunc("/update", update)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err = http.ListenAndServe(":8080", nil)
	modul.CheckErr(err)
}

//index handler
func index(w http.ResponseWriter, r *http.Request) {
	s := "Successfully directed to index."
	_, err := io.WriteString(w, s)
	modul.CheckErr(err)
}

//read user record
func read(w http.ResponseWriter, r *http.Request) {

	//select all rows on one table
	//trying to checking records
	rows, err := db.Query("select * from users")
	modul.CheckErr(err)
	defer rows.Close()

	//make variable to store a records result that scanned
	var userid, username, password, logid string

	//scanning records
	for rows.Next() {
		err = rows.Scan(&userid, &username, &password, &logid)
		modul.CheckErr(err)

		//call struct user
		dataRow := modul.Users{}
		//res := []modul.Users{}

		dataRow.UserID = userid
		dataRow.Username = username
		dataRow.Password = password
		dataRow.LogID = logid

		//append users struct with result record
		//data := append(res, dataRow)
		fmt.Fprintln(w, userid, "", username, "", password, "", logid)
	}
}

//find data user
func findUser(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("Name")
	w.Header().Set("Content-Type", "text/html;charset=utf8")

	io.WriteString(w, `
		<form action="/find" method="GET" >
		<input type="text" name="Name">
		<input type="submit"/> <br>
		<small>example type like: 10001 1002 1003</small>
	  </form>	 
		`)

	//make variable to store the result of scanning
	var userid, username, password, logid string

	//check the data were looking for
	row, err := db.Query("SELECT * from users where userid = ?", v)
	modul.CheckErr(err)
	defer row.Close()

	for row.Next() {
		err = row.Scan(&userid, &username, &password, &logid)
		modul.CheckErr(err)
	}

	//call struct user
	dataRow := modul.Users{}

	dataRow.UserID = userid
	dataRow.Username = username
	dataRow.Password = password
	dataRow.LogID = logid

	if username == "" {
		io.WriteString(w, "not found! ")
	} else {
		fmt.Fprintln(w, "found!")
		fmt.Fprintln(w, userid, username, password, logid)
	}
	return
}

func insert(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		userid := r.FormValue("userid")
		username := r.FormValue("username")
		password := r.FormValue("password")
		logid := r.FormValue("logid")

		stmt, err := db.Prepare("INSERT INTO users (userid, username, password, logid) VALUES (?,?,?,?)")
		modul.CheckErr(err)
		defer stmt.Close()

		res, err := stmt.Exec(userid, username, password, logid)
		modul.CheckErr(err)

		_, err = res.RowsAffected()
		modul.CheckErr(err)

		fmt.Fprint(w, "data successfully recorded!")

	case http.MethodGet:
		w.Header().Set("Content-Type", "text/html;charset=utf8")
		io.WriteString(w, `
	<form action="/insert" method="POST" >
	userid<input type="text" name="userid"> <br>
	username<input type="text" name="username"><br>
	password<input type="text" name="password"><br>
	logid<input type="text" name="logid"><br><br>
		<input type="submit"> <br>
	</form>	
	`)
	default:
		log.Fatal(http.StatusNotImplemented)
		return
	}
}

func updateAction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.Redirect(w, r, "/", http.StatusOK)
	case http.MethodPost:
		a := r.FormValue("userid")
		b := r.FormValue("username")
		c := r.FormValue("password")
		//d := r.FormValue("logid")

		//save update
		stmt, err := db.Prepare("UPDATE users SET username = ?, password= ? WHERE userid = ?")
		modul.CheckErr(err)
		defer stmt.Close()

		r, err := stmt.Exec(b, c, a)
		modul.CheckErr(err)

		_, err := r.RowsAffected()
		modul.CheckErr(err)

		fmt.Fprint(w, "record updated")

	default:
		http.Redirect(w, r, "/", http.StatusOK)
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "text/html;charset=utf8")
		io.WriteString(w, `
	<form action="/update" method="POST" >
	userid<input type="text" name="userid"> <br>
		<input type="submit"> <br>
	</form>	
		`)
	case http.MethodPost:
		//make variable to store the result of scanning
		var userid, username, password, logid string

		//select record
		a := r.FormValue("userid")
		err := db.QueryRow("SELECT * FROM users where userid = ?", a).Scan(&userid, &username, &password, &logid)
		modul.CheckErr(err)

		//call struct user
		dataRow := modul.Users{}

		dataRow.UserID = userid
		dataRow.Username = username
		dataRow.Password = password
		dataRow.LogID = logid

		data := M{
			"UserID":   userid,
			"Username": username,
			"Password": password,
			"LogID":    logid,
		}
		tpl.ExecuteTemplate(w, "edit.html", data)
	}
}
