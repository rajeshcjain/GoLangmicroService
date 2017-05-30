package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"
)

var processTheData = Data{
	FirstName: "Raj",
	LastName : "Jain",
	Address : "MyAddress",
	Sal : 10000000,
}

type Data struct {
	FirstName string
	LastName string
	Address string
	Sal int
}

func main(){

	router := httprouter.New()
	router.GET("/v1/processedData/:id",getProcessedData)
	router.POST("/v1/processedData/:id",createProcessedData)
	router.GET("/v1/processData/:category/:post",getCategoryData)
	router.GET("/v1/processData1/files/*filepath",getProcessFiles)
	fmt.Println("Listening on 8080")
	http.ListenAndServe(":8080",router)
}

func getProcessFiles(rw http.ResponseWriter,req *http.Request,params httprouter.Params ){
	/*
	Path: /files/*filepath

	Requests:
 		/files/                             match: filepath="/"
		/files/LICENSE                      match: filepath="/LICENSE"
 		/files/templates/article.html       match: filepath="/templates/article.html"
 		/files                              no match, but the router would redirect
	*/

	path := params.ByName("filepath")
	fmt.Fprintf(rw,string("path " + path))
}

func getCategoryData(rw http.ResponseWriter,req *http.Request,params httprouter.Params){

	/*
	Here we are processing two data items.
	Named parameters are dynamic path segments.
	They match anything until the next '/' or the path end:

	Path: /blog/:category/:post

Requests:
 /blog/go/request-routers            match: category="go", post="request-routers"
 /blog/go/request-routers/           no match, but the router would redirect
 /blog/go/                           no match
 /blog/go/request-routers/comments   no match

	*/
	cat := params.ByName("category")
	post := params.ByName("post")
	fmt.Println("category: ",cat,"\npost: ",post)
}

func getProcessedData(rw http.ResponseWriter,req *http.Request,params httprouter.Params){
	id := params.ByName("id")
	fmt.Println("The GET Request I got is ",id)
	data,err := json.Marshal(processTheData)
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)

	if err != nil{
		log.Fatalln(err)
	}
	fmt.Fprintf(rw,string(data))
}


func createProcessedData(rw http.ResponseWriter,req *http.Request,params httprouter.Params){
	id := params.ByName("id")
	fmt.Println("The POST Request I got is ", id)

	data, err := ioutil.ReadAll(req.Body)

	if err != nil{
		log.Fatalln(err)
	}

	var dataReceived Data

	//We do the Unmarshal to decode the data and to put it in to struct
	err = json.Unmarshal(data,&dataReceived)

	if err != nil{
		log.Fatalln(err)
	}

	DataSent := Data{
		FirstName:dataReceived.FirstName,
		LastName: dataReceived.LastName,
		Address:dataReceived.Address,
		Sal:dataReceived.Sal,
	}

	encodedData,_ := json.Marshal(DataSent)
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)

	fmt.Println("The data received is ",DataSent)
	fmt.Fprintf(rw, string(encodedData))

}
