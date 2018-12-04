package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/notes", getNoteHandler).Methods("GET")
	r.HandleFunc("/notes", createNoteHandler).Methods("POST")
	return r
}

func main() {
	fmt.Println("Starting server...")
	s := os.Getenv("JAWSDB_URL")
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	user := u.User.Username()
	password, _ := u.User.Password()
	host, port, _ := net.SplitHostPort(u.Host)
	dbname := strings.Replace(u.Path, "/", "", -1)

	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)

	fmt.Println("connString: ", connString)

	db, err := sql.Open(u.Scheme, connString)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	r := newRouter()
	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
