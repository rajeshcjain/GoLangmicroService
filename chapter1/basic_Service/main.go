package main

import (
	"net/http"
	"fmt"
	"log"
)


func main(){

	port := 8080
	http.HandleFunc("/hello",helloFunc)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port),nil))

}

func helloFunc( res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res,"Hello World!!!!")
}


