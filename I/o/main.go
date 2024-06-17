package main

import (
	"net/http"

	"github.com/gorilla/mux"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey, This is Sanjana B- Bad lil Bitch's server"))
}
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/sanjana",Handler)

	http.ListenAndServe(":3000", r)

	


	
}