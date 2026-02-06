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
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	returnedError := req.FormValue("error")
	msg := req.FormValue("msg")
	tmpl := template.Must(template.ParseFiles("html/index.html", "html/add-mark.html"))
	data := map[string]any{
		"title":    "Inserting Marks Page",
		"username": "default",
		"msg":      msg,
		"error":    returnedError,
	}
	if err := tmpl.ExecuteTemplate(res, "index", data); err != nil {
		Check(err)
	}
}

func InsertMarkHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			log.Println(err)
			return
		}
		id := req.FormValue("ID")
		name := req.FormValue("Name")
		mark := req.FormValue("Mark")
		update := req.FormValue("forUpdate")

		id64, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Println(err)
			redirect := fmt.Sprint("/mark/add?error=", err)
			http.Redirect(res, req, redirect, http.StatusSeeOther)
			return
		}
		mark32, err := strconv.ParseFloat(mark, 32)
		if err != nil {
			log.Println(err)
			redirect := fmt.Sprint("/mark/add?error=", err)
			http.Redirect(res, req, redirect, http.StatusSeeOther)
			return
		}

		if id64 < 1 {
			redirect := fmt.Sprint("/mark/add?error=", "the id can't be less that 1")
			http.Redirect(res, req, redirect, http.StatusSeeOther)
			return
		}

		if mark32 < 0 || mark32 > 100.00 {
			redirect := fmt.Sprint("/mark/add?error=", "invalid mark value (0 <= mark <= 100.00 )")
			http.Redirect(res, req, redirect, http.StatusSeeOther)
			return
		}

		if update == "on" {
			err = UpdateStudent(db, id64, name, float32(mark32))
			if err != nil {
				log.Println(err)
				redirect := fmt.Sprint("/mark/add?error=", "update fail")
				http.Redirect(res, req, redirect, http.StatusSeeOther)
				return
			}
			http.Redirect(res, req, "/mark/add?msg= successful update", http.StatusSeeOther)
			return
		} else {
			err = InsertStudent(db, id64, name, float32(mark32))
			if err != nil {
				log.Println(err)
				redirect := fmt.Sprint("/mark/add?error=", "insertion fail")
				http.Redirect(res, req, redirect, http.StatusSeeOther)
				return
			}
			http.Redirect(res, req, "/mark/add?msg=a new record created", http.StatusSeeOther)
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
			"username": "default",
		}
		if err := tmpl.ExecuteTemplate(res, "index", data); err != nil {
			Check(err)
			http.Redirect(res, req, "/mark/show", http.StatusSeeOther)
		}
	}
}

func LoginPage(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	returnedError := req.FormValue("error")
	tmpl := template.Must(template.ParseFiles("html/index.html", "html/login.html"))
	data := map[string]any{
		"title":    "Login Page",
		"error":    returnedError,
		"username": "default",
	}
	if err := tmpl.ExecuteTemplate(res, "index", data); err != nil {
		Check(err)
	}
}
