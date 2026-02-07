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

	mux.HandleFunc("GET /mark/add", Protected(InsertingPage))
	mux.HandleFunc("POST /mark/insert", Protected(InsertMarkHandler(db)))

	mux.HandleFunc("GET /mark/show", Protected(ShowMarkHandler(db)))

	mux.HandleFunc("GET /mark/login", LoginPage)
	mux.HandleFunc("POST /mark/auth", LoginHandler(db))

	mux.HandleFunc("POST /mark/logout", Protected(LogOutHandler))

	if err := http.ListenAndServe(":10055", mux); err != nil {
		panic(err)
	}
}
