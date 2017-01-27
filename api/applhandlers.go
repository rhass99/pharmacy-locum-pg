package api

import (
	//"fmt"
	"github.com/gorilla/sessions"
	"github.com/rhass99/pl-pq/storage"
	"log"
	"net/http"
	"text/template"
)

var cookieStore = sessions.NewCookieStore([]byte("Secret"))
var signupTmpl = template.Must(template.New("signup.html").ParseFiles("tmpl/signup.html"))
var loginTmpl = template.Must(template.New("login.html").ParseFiles("tmpl/login.html"))
var profileTmpl = template.Must(template.New("profile.html").ParseFiles("tmpl/profile.html"))

//Handler function for the Sign up page
func SignupAppl(w http.ResponseWriter, r *http.Request) {
	// Open pqsql from package storage
	switch {
	case r.Method == "POST":
		db, err := storage.DBConnect()
		if err != nil {
			log.Println("No DB connection")
		}
		// Define a new Applicant Sign up object
		var a *storage.Applog
		a = ProcessApplForm(r)
		err = StoreAppl(a, db)
		if err != nil {
			log.Println("Cannot store applicant")
		}
	case r.Method == "GET":
		err := signupTmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	}
}

// GET function for the login page
func LoginAppl(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		err := loginTmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	case r.Method == "POST":
		db, err := storage.DBConnect()
		if err != nil {
			log.Println("No DB connection")
		}
		var b *storage.Applog
		b = ProcessApplForm(r)
		signed := AuthAppl(b, db)
		switch {
		case signed == "userdoesntexist":
			http.Redirect(w, r, "/signup", 302)
		case signed == "wrongpassword":
			http.Redirect(w, r, "/login", 302)
		case signed == "userauthenticated":
			session, _ := cookieStore.Get(r, "user-session")
			session.Values["usermail"] = b.Email
			session.Save(r, w)
			http.Redirect(w, r, "/profile", 302)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("401 HTTP status code returned!"))
		}
		// if signed {
		// 	http.Redirect(w, r, "/profile", 302)
		// } else {
		// 	fmt.Fprintln(w, "Not Valid")
		// }
	}
}

// GET function for the Profile page
// Should be called after signup - login - profile
func ProfileApplGet(w http.ResponseWriter, r *http.Request) {
	db, err := storage.DBConnect()
	session, _ := cookieStore.Get(r, "user-session")
	var email interface{}
	email = session.Values["usermail"]
	s, _ := email.(string)
	a, err := storage.GetApplAuth(s, db)
	if err != nil {
		log.Println(err)
	}
	err = profileTmpl.Execute(w, a)
	if err != nil {
		log.Println(err)
	}
}
