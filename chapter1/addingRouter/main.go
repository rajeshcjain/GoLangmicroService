package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)


func main(){

	/*
	mux stands for "HTTP request multiplexer".
	So the important part here is we are declaring a mux and NewRouter()
	will return the new router and it will match the incoming requests
	against the list of registered routes and call the respective handler for
	a match of URL or a condition
	*/

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello",handleFunction)
	router.HandleFunc("/products/{category}/{id}",handleProducts)
	router.HandleFunc("/articals/{id}",handleArticals)
	fmt.Println("Listening on port 8080")

	http.ListenAndServe(":8080",router)

}

func handleArticals(w http.ResponseWriter, r *http.Request){

	//This is how we fetch data from the URI.
	//mux.Vars gives the route variables for the current request, if any
	Vars := mux.Vars(r)
	//gives the id variable.
	id := Vars["id"]
	fmt.Fprintf(w,"Articals , %q , Id : %q",r.URL.Path,id)
}

func handleProducts(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Products , %q",r.URL.Path)
}

func handleFunction(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello , %q",r.URL.Path)
}


