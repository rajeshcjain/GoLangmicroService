package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"fmt"
	"log"
)

type helloWorldResponse struct {
Message string
Message1 string
}

/*
This is important.If you keep the first letter of the field of struct as small
then,json will not be marshaled properly.So the first letter should be capital
letter.

*/
type product struct{
ProductID string
ProductName string
}




func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello",handleHelloFunc)
	router.HandleFunc("/product",handleProductFunc)
	http.ListenAndServe(":8080",router)
}



func handleProductFunc(rw http.ResponseWriter,r *http.Request){
	log.Println("Giving output in JSON....")
	response := product{ProductID : "1",
		ProductName : "MacBook Pro",}
        log.Println("Response ",response)
	/*
	 Here first we need to convert the response in to json so Marshal it.
	*/
	data,_ := json.Marshal(response)
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	log.Println("Data ",data)
	fmt.Fprintf(rw,string(data))
}



func handleHelloFunc(rw http.ResponseWriter,r *http.Request ){
	response := helloWorldResponse{Message : "Hello World..I am here",
		Message1 : "Message1",}
	data,_ := json.Marshal(response)
	fmt.Fprintf(rw,string(data))
}