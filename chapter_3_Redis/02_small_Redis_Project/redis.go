package main

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)


/*
               ALL REDIS RELATED FUNCTIONS

*/


func findALL() []Post{
	conn,err := redis.Dial("tcp",server + ":" + port)
	handlePanic(err)
	defer conn.Close()
	return nil
}

func updateData(postId int,post Post){
	conn,err := redis.Dial("tcp",server + ":" + port)
	handlePanic(err)
	defer conn.Close()
	conn.Do("SET",postId,post)
}


func findPost(key int) Post{
	conn, error := redis.Dial("tcp",server + ":" + port)
	handlePanic(error)
	defer conn.Close()
	reply,err := conn.Do("GET",key)
	handlePanic(err)
	var post Post
	err = json.Unmarshal(reply.([]byte),&post)
	handlePanic(err)
	return post
}
