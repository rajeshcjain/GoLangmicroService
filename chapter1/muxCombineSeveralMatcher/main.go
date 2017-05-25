package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main(){

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",func(rw http.ResponseWriter,r *http.Request){
		fmt.Fprintf(rw,"HelloWorld")
	}).Methods("GET")//.Schemes("http", "https").

         http.ListenAndServe(":8080",router)
}



