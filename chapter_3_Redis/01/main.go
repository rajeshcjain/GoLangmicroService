package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)


type Data struct {
	FirstName string
	LastName string
	Address string
	Sal int
}

func main(){
        var router = httprouter.New()
	router.POST("/v1/test/:key",handleFunc)
	router.GET("/v1/test/:key",handleGet)
	fmt.Println("Listening 8080")
        http.ListenAndServe(":8080",router)
	fmt.Println("Ending")
}

func handleGet(rw http.ResponseWriter,req *http.Request,params httprouter.Params){
	fmt.Println("GET Call")
	key := params.ByName("key")
	fmt.Println("Key",key)

	byteStream := make([]byte,0,50)
	/*
	This is a important concept here...We are getting the slice from the
	function as a returned value...but the append function expect the
	elements....so we are giving it a list of element rather then a slice.
	*/
	byteStream = append(byteStream, getFromRedis(key)...)

	fmt.Fprintf(rw,string(byteStream))
}

func handleFunc(rw http.ResponseWriter,req *http.Request,params httprouter.Params){
    	fmt.Println("POST Call")
	key := params.ByName("key")
	fmt.Println("Key",key)

	data,err := ioutil.ReadAll(req.Body)

	if err != nil{
		log.Fatalln(err)
	}

	var marshaledData Data
	errOfDecoding := json.Unmarshal(data,&marshaledData)


	if errOfDecoding != nil {
		log.Fatalln(errOfDecoding)
	}

	fmt.Println("Data received is ",marshaledData)
	updateRedis(data,key)
	fmt.Println("Updated the Redis..")
}


/*

				THIS IS IMPORTANT
Connect to the Redis server...As redis is on my local machine ..so its easy...
Need to see how to do it when it is on remote server ...or there is a cluster
of redis servers...There are other parameters which i need to consider...like
default value and how to set TTL.

*/

func connect() redis.Conn{
	connWithRedis,err := redis.Dial("tcp",":6379")

	if err != nil{
		log.Fatalln(err)
	}

	return connWithRedis

}

/*

 			SET Method

*/
func updateRedis(data []byte,key string){
	c := connect()
	defer c.Close()
	c.Do("SET",key,data)
}



/*YET TO SEE THIS PART,There is a problem here...Will check*/

func getFromRedis(key string) []byte{
	c := connect()
	defer c.Close()
	fmt.Println("Connected with Redis")

	reply,err := c.Do("GET",key)
	fmt.Println("reply from Redis ")
	if err != nil{
		log.Fatalln(err)
	}

	/*
	Here its important...to note that how to fetch data from redis...
	I am here converting the data received from redis to byte stream..
	and then using the Unmashal method json...convert it in to byte stream
	and then put in to struct and then marshal it..and send it as a response.
	*/
	var myData Data
	error := json.Unmarshal(reply.([]byte),&myData)

	if error != nil{
		log.Fatalln(error)
	}

	byteStream,err := json.Marshal(myData)
	if err != nil{
		log.Fatalln(err)
	}

	return byteStream
}




