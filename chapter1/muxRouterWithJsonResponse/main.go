package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"fmt"
)

type helloWorldResponse struct {
Message string `json:"message"`
}

/*
type product struct{
productId string `json:"message"`
productName string`json:"message"`
productCost int `json:"message"`
}
*/



func main(){

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello",handleHelloFunc)
//	router.HandleFunc("/product",handleProductFunc)
	http.ListenAndServe(":8080",router)
}


/*
func handleProductFunc(rw http.ResponseWriter,r *http.Request){
	response := product{productId : "1",
		            productName : "MacBook",
		            productCost : 100000,}
	data,_ := json.Marshal(response)
	fmt.Fprint(rw,string(data))
}
*/


func handleHelloFunc(rw http.ResponseWriter,r *http.Request ){
	response := helloWorldResponse{Message : "Hello World..I am here"}
	data,_ := json.Marshal(response)
	fmt.Fprintf(rw,string(data))
}