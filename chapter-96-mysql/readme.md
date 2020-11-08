# Using MySQL

1. Install MySQL
-Download MySQL Community Server

1. We will need a MySQL driver
-go get github.com/go-sql-driver/mysql
-read the documentation on https://github.com/go-sql-driver/mysql#installation
-see all SQL drivers on https://github.com/golang/go/wiki/SQLDrivers
-Astaxie's book on https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.2.html 

1. Include the driver in your imports
-_ "github.com/go-sql-driver/mysql"
-Read the documentation
-Determine the Data Source Name
-user:password@tcp(localhost:5555)/dbname?charset=utf8
-Read the documentation on https://github.com/go-sql-driver/mysql#dsn-data-source-name

1. Open a connection
*db, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")
package sql golang on https://godoc.org/database/sql