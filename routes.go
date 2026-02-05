package main

import (
	"fmt"
	"html/template"
	"net/http"
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

func InsertMarkHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	id := req.FormValue("ID")
	name := req.FormValue("Name")
	mark := req.FormValue("Mark")
	update := req.FormValue("forUpdate")

	fmt.Println("-----------------")
	fmt.Println(id)
	fmt.Println(name)
	fmt.Println(mark)
	fmt.Println(update)
	if update == "on" {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	fmt.Println("-----------------")
}

func ShowMarkHandler(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html", "html/show-mark.html"))
	data := map[string]any{
		"title": "Show Results",
		"students": []Student{
			{ID: 938, Name: "Mohmamad", Mark: 2.3},
			{ID: 2938, Name: "ali", Mark: 5.3},
			{ID: 298, Name: "omer", Mark: 9.3},
			{ID: 293, Name: "hassan", Mark: 2.3},
			{ID: 238, Name: "mark", Mark: 7.3},
		},
	}
	if err := tmpl.ExecuteTemplate(res, "index", data); err != nil {
		Check(err)
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

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	email := req.FormValue("email")
	password := req.FormValue("password")

	fmt.Println("-----------------")
	fmt.Println(email)
	fmt.Println(password)
	fmt.Println("-----------------")
}
