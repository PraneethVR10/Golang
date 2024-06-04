package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hey this is Praneeth and welcome to my server"))
}

func HandleProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to the Prodcuts section, Please choose a product of your choice"))
}

func HandleArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pick an Article of yoru choice and read it"))
}



func main() {
r := mux.NewRouter()	
r.HandleFunc("/hello",handler)
r.HandleFunc("/products",HandleProducts)
r.HandleFunc("/article",HandleArticle)
http.ListenAndServe(":3030",r)
}