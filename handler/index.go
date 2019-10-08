package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	db, err := sql.Open("mysql", "root:@/books_management")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		log.Println("error: ${err}")
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Println("error: ${err}")
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var stringList []string
	for rows.Next() {
		err = rows.Scan(scanArgs...)

		if err != nil {
			panic(err.Error())
		}

		var value string


		for i, col := range values {

			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
				stringList = append(stringList, value)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}



	tmpl, err := template.ParseFiles("view/index.html")
	if err != nil {
		log.Println("error: ${err}")
	}

	err = tmpl.Execute(w, stringList)
	if err != nil {
		log.Println("error: ${err}")
	}
}
