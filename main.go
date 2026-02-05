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
	mux := http.NewServeMux()
	mux.HandleFunc(" /mark/add", InsertMarkHandler)
	mux.HandleFunc(" /mark/show", ShowMarkHandler)
	mux.HandleFunc(" /mark/login", LoginPage)
	mux.HandleFunc(" /mark/auth", LoginHandler)
	mux.Handle("static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	if err := http.ListenAndServe(":10055", mux); err != nil {
		panic(err)
	}
}
