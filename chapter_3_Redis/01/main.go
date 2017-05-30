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
        http.ListenAndServe(":8080",router)
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

func connect() redis.Conn{
	connWithRedis,err := redis.Dial("tcp",":6379")
	if err != nil{
		log.Fatalln(err)
	}

	return connWithRedis

}

func updateRedis(data []byte,key string){
	c := connect()
	c.Do("SET",key,data)
	c.Close()
}



/*YET TO SEE THIS PART,There is a problem here...Will check*/
/*
func getFromRedis(key string) string{
	c := connect()
	reply,err := c.Do("GET",key)

	if err != nil{
		log.Fatalln(err)
	}
	c.Close()
	return reply

}
*/

