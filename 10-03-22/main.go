package main

import "net/http"

func main() {
}

func ServeHTPP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("YOY", "Cool")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is up."))
}