package main

import (
	//"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../testing/"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.ListenAndServe(":3000", nil)
}
