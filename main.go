package main

import "github.com/gorilla/mux"

func main() {
	r := mux.NewRouter()
	r.Host(":8080")
}
