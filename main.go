package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Users struct {
	ID       string `json:"id,omitempty"`
	Password string `json:"password,omitempty"`
}
type Services struct {
	NAME string `json:"name,omitempty"`
	ID   string `json:"id,omitempty`
}

var users []Users
var services []Services

func GetServices(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var result []Services
	for _, item := range services {
		if item.ID == params["userid"] {
			result = append(result, item)
		}
	}
	json.NewEncoder(w).Encode(result)
}

func PostServices(w http.ResponseWriter, r *http.Request) {
	var service Services
	_ = json.NewDecoder(r.Body).Decode(&service)
	services = append(services, service)
	json.NewEncoder(w).Encode("{'status':'ok'}")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user Users
	_ = json.NewDecoder(r.Body).Decode(&user)
	for _, item := range users {
		if item.ID == user.ID {
			if item.Password == user.Password {
				json.NewEncoder(w).Encode("{'status':'ok'}")
				return
			}
			json.NewEncoder(w).Encode("{'status':'not ok'}")
		}
	}
}
func Signup(w http.ResponseWriter, r *http.Request) {
	var user Users
	_ = json.NewDecoder(r.Body).Decode(&user)
	//TODO check dup
	users = append(users, user)
	json.NewEncoder(w).Encode("{'status':'ok'}")
}

func main() {
	router := mux.NewRouter()
	users = append(users, Users{ID: "oyt", Password: "1234"})
	services = append(services, Services{NAME: "service1", ID: "oyt"})
	router.HandleFunc("/services/{userid}", GetServices).Methods("GET")
	router.HandleFunc("/services", PostServices).Methods("POST")
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/signup", Signup).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
