//psql -U rami -f setup.sql -d phlocum

package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	//"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rami"
	password = "phlocum"
	dbname   = "phlocum"
)

type Applog struct {
	Id           int
	Pid          string `schema:"pid"`
	Firstname    string `schema:"firstname"`
	Lastname     string `schema:"lastname"`
	Email        string `schema:"email"`
	Password     string `schema:"password"`
	Password2    string `schema:"password2"`
	PasswordHash string
}

// type ApplogSignin struct {
// 	Email        string `schema:"email"`
// 	Password     string `schema:"password"`
// 	PasswordHash string
// }

//func Authenticate(email string)

func DBConnect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (a *Applog) CreateAppl(db *sql.DB) error {
	var qemail string
	err := db.QueryRow("SELECT email from applog where email = $1", a.Email).Scan(&qemail)
	switch {
	case err == sql.ErrNoRows:
		statement := "insert into applog (pid, firstname, lastname, email, passwordhash) values ($1, $2, $3, $4, $5) returning id;"
		stmt, err := db.Prepare(statement)
		if err != nil {
			return err
		}
		defer stmt.Close()
		err = stmt.QueryRow(a.Pid, a.Firstname, a.Lastname, a.Email, a.PasswordHash).Scan(&a.Id)
		if err != nil {
			return err
		}
		return nil
	case err == nil:
		fmt.Println("User Exist")
	default:
		return err
	}
	return nil
}

// func GetApplByEmail(email string, db *sql.DB) (Applog, error) {
// 	a := Applog{}
// 	err := db.QueryRow("SELECT email, firstname, lastname, FROM applog WHERE email = $1", email).Scan(&a.Email, &a.Firstname, &a.Lastname)
// 	return a, err
// }

func GetApplByID(pid string, db *sql.DB) (Applog, error) {
	a := Applog{}
	err := db.QueryRow("SELECT pid, firstname, lastname, email FROM applog WHERE pid = $1", pid).Scan(&a.Pid, &a.Firstname, &a.Lastname, &a.Email)
	return a, err
}

func GetApplAuth(email string, db *sql.DB) (Applog, error) {
	a := Applog{}
	err := db.QueryRow("SELECT email, firstname, lastname, passwordhash FROM applog WHERE email = $1", email).Scan(&a.Email, &a.Firstname, &a.Lastname, &a.PasswordHash)
	return a, err
}
