package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	connectionString := getConnectionString()

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Println("error: 123")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM test")
	//rows, err := db.Query("show databases")
	if err != nil {
		log.Println(err)
		log.Println("error")
	}
	log.Println(rows)

	columns, err := rows.Columns()
	if err != nil {
		log.Println("error")
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



	tmpl, err := template.ParseFiles("/go/src/github.com/hisamura333/books-management/view/index.html")
	if err != nil {
		log.Println(err)
	}

	err = tmpl.Execute(w, stringList)
	if err != nil {
		log.Println(err)
	}
}

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	if env != "" {
		return env
	}
	return defaultValue
}

func getConnectionString() string {
	host := getParamString("MYSQL_DB_HOST", "db")
	port := getParamString("MYSQL_PORT", "3306")
	user := getParamString("MYSQL_USER", "root")
	pass := getParamString("MYSQL_PASSWORD", "")
	dbname := getParamString("MYSQL_DB", "books_management")
	if os.Getenv("APP_ENV") == "test" {
		dbname = strings.Join([]string{dbname, "_test"}, "")
	}
	protocol := getParamString("MYSQL_PROTOCOL", "tcp")
	dbargs := getParamString("MYSQL_DBARGS", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}