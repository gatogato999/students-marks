package main

import (
	"log"
	"net/http"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(" /mark/home", HomeHanlder)

	if err := http.ListenAndServe(":10055", mux); err != nil {
		panic(err)
	}
}
