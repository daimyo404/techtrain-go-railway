package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/user/create", getUser)
	http.HandleFunc("/user/create", updateUser)

	http.ListenAndServe(":8080", nil)
}

func createUser(w http.ResponseWriter, r *http.Request) {

}

func getUser(w http.ResponseWriter, r *http.Request) {

}

func updateUser(w http.ResponseWriter, r *http.Request) {

}
