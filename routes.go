package main

import (
	"html/template"
	"net/http"
)

type Student struct {
	ID   int64
	Name string
	Mark float32
}

func HomeHanlder(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html"))

	data := map[string]any{
		"title":      "this my title",
		"authorized": true,
		"students": []Student{
			{ID: 938, Name: "Mohmamad", Mark: 2.3},
			{ID: 2938, Name: "ali", Mark: 5.3},
			{ID: 298, Name: "omer", Mark: 9.3},
			{ID: 293, Name: "hassan", Mark: 2.3},
			{ID: 238, Name: "mark", Mark: 7.3},
		},
	}

	if err := tmpl.Execute(res, data); err != nil {
		Check(err)
	}
}
