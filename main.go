package main

import (
	"log"
	"net/http"
)

func Check(err error) {
	if err != nil {
		log.Println(err)
	}
}

func main() {
	db, err := GainAccessToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = CreateTables(db)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.Handle(
		"/marks/static/",
		http.StripPrefix("/marks/static/", http.FileServer(http.Dir("static"))),
	)

	mux.HandleFunc("GET /marks/add", Protected(InsertingPage))
	mux.HandleFunc("POST /marks/insert", Protected(InsertMarkHandler(db)))

	mux.HandleFunc("GET /marks/show", Protected(ShowMarkHandler(db)))

	mux.HandleFunc("GET /marks/login", LoginPage)
	mux.HandleFunc("POST /marks/auth", LoginHandler(db))

	mux.HandleFunc("POST /marks/logout", Protected(LogOutHandler))

	if err := http.ListenAndServe(":10055", mux); err != nil {
		panic(err)
	}
}
