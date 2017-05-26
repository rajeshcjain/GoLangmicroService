package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

func main(){

	router := httprouter.New()
	router.GET("/",handleGetFunc)
	router.POST("/",handlePostFunc)
	fmt.Println("Listening on 8080")
	http.ListenAndServe(":8080",router)
}

func handleGetFunc(w http.ResponseWriter,req *http.Request,params httprouter.Params){

	fmt.Fprintf(w,"welcome from the GET Request")

}


func handlePostFunc(w http.ResponseWriter,req *http.Request,params httprouter.Params){
	fmt.Fprintf(w,"welcome from the POST request")
}