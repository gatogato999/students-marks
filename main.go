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

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("GET /mark/add", InsertingPage)
	mux.HandleFunc("POST /mark/insert", InsertMarkHandler)

	mux.HandleFunc("GET /mark/show", ShowMarkHandler)

	mux.HandleFunc("GET /mark/login", LoginPage)
	mux.HandleFunc("POST /mark/auth", LoginHandler)

	if err := http.ListenAndServe(":10055", mux); err != nil {
		panic(err)
	}
}
