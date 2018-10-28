package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()

}
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/login", a.login).Methods("POST")
	a.Router.HandleFunc("/signup", a.signup).Methods("POST")
	a.Router.HandleFunc("/services/{userid}", a.getService).Methods("GET")
	a.Router.HandleFunc("/services", a.putService).Methods("POST")
	a.Router.HandleFunc("/services", a.deleteService).Methods("DELETE")
}
func (a *App) login(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusNotFound, map[string]string{"error": "err"})
}
func (a *App) signup(w http.ResponseWriter, r *http.Request)        {}
func (a *App) getService(w http.ResponseWriter, r *http.Request)    {}
func (a *App) putService(w http.ResponseWriter, r *http.Request)    {}
func (a *App) deleteService(w http.ResponseWriter, r *http.Request) {}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
