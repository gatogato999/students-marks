package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func InsertingPage(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html", "html/add-mark.html"))
	data := map[string]any{
		"title": "Inserting Marks Page",
	}
	if err := tmpl.ExecuteTemplate(res, "index", data); err != nil {
		Check(err)
	}
}

func InsertMarkHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id := req.FormValue("ID")
		name := req.FormValue("Name")
		mark := req.FormValue("Mark")
		update := req.FormValue("forUpdate")
		id64, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Println(err)
			return
		}
		mark32, err := strconv.ParseFloat(mark, 32)
		if err != nil {
			log.Println(err)
			return
		}
		if update == "on" {
			err = UpdateStudent(db, id64, name, float32(mark32))
			if err != nil {
				log.Println(err)
				return
			}
			return
		} else {
			err = InsertStudent(db, id64, name, float32(mark32))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}
	}
}

func ShowMarkHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		tmpl := template.Must(template.ParseFiles("html/index.html", "html/show-mark.html"))
		record, err := GetAllStudents(db)
		if err != nil {
			log.Println(err)
			return
		}
		data := map[string]any{
			"title":    "Show Results",
			"students": record,
		}
		if err := tmpl.ExecuteTemplate(res, "index", data); err != nil {
			Check(err)
		}
	}
}

func LoginPage(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html", "html/login.html"))
	data := map[string]any{
		"title": "Login Page",
	}
	if err := tmpl.ExecuteTemplate(res, "index", data); err != nil {
		Check(err)
	}
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		email := req.FormValue("email")
		password := req.FormValue("password")

		fmt.Println("-----------------")
		fmt.Println(email)
		fmt.Println(password)
		fmt.Println("-----------------")
	}
}
