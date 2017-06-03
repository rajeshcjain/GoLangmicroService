package main

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"fmt"
)


/*
               ALL REDIS RELATED FUNCTIONS

*/

//Id for Post and User...ID creation logic should be in this file...as
// other layer should not be aware of it.
var postId int
var userId int

func init(){
	postId = 0
	userId = 0
}



//Connection logic at one place to avoid little work :)
func Connect() redis.Conn {
	conn,err := redis.Dial("tcp",server + ":" + port)
	handlePanic(err)
	return conn
}


func findALL() []Post{
	conn := Connect()
	defer conn.Close()
	return nil
}


func createPost(post Post) int {
	postId++
	userId++
	post.Id = postId
	post.User.Id = userId
	fmt.Println("creating post ",post)
	conn := Connect()
	defer conn.Close()
	reply, err := conn.Do("SET",postId,post)

	handlePanic(err)
	if reply.(int) != 1{
		fmt.Println("No Post created")
		return 0
	}else{
		fmt.Println("Post created")
		return 1
	}

}

func updateData(postId int,post Post) int{
	conn := Connect()
	defer conn.Close()
	reply,err := conn.Do("SET",strconv.Itoa(postId),post)
        handlePanic(err)

	if reply.(int) != 1{
		fmt.Println(postId, "is not updated")
		return 0
	}else{
		fmt.Println(postId, "is updated")
		return 1

	}

}

func deletePost(key int) int {

	conn := Connect()
	defer conn.Close()
	reply,error := conn.Do("DEL",strconv.Itoa(key))
	handlePanic(error)

	if reply.(int) != 1 {
		fmt.Println("No Post removed")
		return 0
	}else{
		fmt.Println("Post removed")
             return 1
	}
}


func findPost(key int) Post{
	conn := Connect()
	defer conn.Close()
	reply,err := conn.Do("GET",strconv.Itoa(key))
	handlePanic(err)
	var post Post
	err = json.Unmarshal(reply.([]byte),&post)
	handlePanic(err)
	return post
}
