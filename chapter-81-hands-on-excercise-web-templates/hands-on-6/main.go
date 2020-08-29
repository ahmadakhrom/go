//package main
//
//import (
//	"html/template"
//	"log"
//	"os"
//)
//
//type DataCsv struct {
//	Date string
//	Open float64
//}
//
//type dataCsvOverall []DataCsv
//
//var tpl *template.Template
//
//func init(){
//	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
//}
//
//func main()  {
//	data := dataCsvOverall{
//	DataCsv{	Date: "2015-07-09",
//				Open: 523.119995,
//			},
//	DataCsv{	Date: "2015-07-08	",
//				Open: 521.049988,
//			},
//	DataCsv{	Date: "2015-07-07	",
//				Open: 523.130005,
//			},
//	}
//
//	err := tpl.Execute(os.Stdout, data)
//	if err != nil {
//		log.Fatalf("%s sdjhfgjhsfgjhds",err)
//	}
//}
package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type DataTable struct {
	Date time.Time
	Open float64
}


func main()  {
	http.HandleFunc("/",foo)
	http.ListenAndServe(":8080",nil)
}

func foo(res http.ResponseWriter, req *http.Request)  {
	//parse data.csv
	file := parsefiles("data.csv")

	//parse to tpl.gohtml
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalf("%s error occured",err)
	}

	//execute template
	err = tpl.Execute(res, file)
	if err != nil {
		log.Fatalf("%s error occured!",err)
	}
}


func parsefiles(filePath string) []DataTable  {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("%s error occured",err)
	}
	defer src.Close()

	dataCsv := csv.NewReader(src)
	rows, err := dataCsv.ReadAll()
	if err != nil {
		log.Fatalf("%s error occured", err)
	}

	DataTables := make([]DataTable, 0, len(rows))

	for i, row := range rows {
		if i == 0{
			continue
		}

		date, _ := time.Parse("2006-01-02",row[0])
		open, _ := strconv.ParseFloat(row[1],64)

		DataTables = append(DataTables, DataTable{
			Date: date,
			Open: open,
		})
	}
	return DataTables
}

































