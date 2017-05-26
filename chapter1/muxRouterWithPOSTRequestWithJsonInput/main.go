package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type processInputData struct {
   Id string
   Name string
}

type inputResponse struct {
	Message string
}

func main(){

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/processInput",handler).Methods("POST")
	log.Println("listening on 8080")
	http.ListenAndServe(":8080",router)
}

func handler(w http.ResponseWriter,req *http.Request){

	/*
	 ioutil : This is a important util package for input output operations.

	 ReadAll(req.Body) : Read all the body and gives a byte stream.
	*/
	data,err := ioutil.ReadAll(req.Body)

	if err != nil{
		log.Fatalln(err)
		return
	}

	log.Println("data  ",data)

	/*
	  Now time to convert the byte stream in to struct.
	*/
	var inputData processInputData
	/*
	   json.unmarshal()--decode the data and put the stream in to the struct we
	   pass here.
	*/
	err = json.Unmarshal(data,&inputData)
	if err != nil{
		log.Fatalln(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := inputResponse{Message: "Got Your Input data ..Thanks"}

	/*Here we create the response and generate the json format..so that it could be
	sent over the network to client as a response*/
	data,_ = json.Marshal(response)

	//enconder := json.NewEncoder(w)
	//enconder.Encode(response)

	fmt.Fprintf(w,string(data))
}
