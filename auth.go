package main

import "net/http"

func protected(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		next(res, req)
	}
}
