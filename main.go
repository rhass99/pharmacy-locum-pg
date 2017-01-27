package main

import (
	"github.com/gorilla/mux"
	"github.com/rhass99/pl-pq/api"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", api.SignupAppl)
	r.HandleFunc("/login", api.LoginAppl)
	r.HandleFunc("/profile", api.ProfileApplGet)
	http.ListenAndServe(":8080", r)
}
