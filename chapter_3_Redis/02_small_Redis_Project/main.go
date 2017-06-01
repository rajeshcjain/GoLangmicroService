package main

import (
	"time"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"github.com/garyburd/redigo/redis"
)

/*
We will create a simple blog service that writes JSON data as strings to the
browser on different url routes.

For instance,

browsing to
http://localhost/ will display a welcome message,
http://localhost/posts will display all posts in JSON and
http://localhost/posts/id will get you the post with the given id.
*/

type User struct{
	Id int
	UserName string
	Email string
}

type Comment struct{
	Id int
	User User
	Text string
	Timestamp time.Time

}

type Post struct{
	Id int
	User User
	Topic string
	Text string
	Timestamp time.Time
}

type Route struct {
	Method string             // HTTP Method
	Path string               // Path of the end point
	Handle httprouter.Handle  //defining the type here.We are telling that
	                          // this is a function type.
}


/*
We can not have a single post,User and comment on the blog..So defining the
Data structure for it...i.e multiple user,posts and comments.
*/

//var users = []User{}
//var posts = []Post{}
//var comments = []Comment{}

var users = make(map[int]User)
var posts = make(map[int]Post)
var comments = make(map[int]Comment)

var routers = []Route{
	{
		"GET",
		"/",
		Index,
	},
	{
		"GET",    //Here Get me all the posts
		"/posts",
		PostIndex,
	},
	{
		"GET",   //Get me particular post
		"/posts/:id",
		PostShow,
	},
	{
		"POST",  //Add in to existing posts
		"/posts",
		PostCreate,
	},

}



func handlePanic(err error){
	if err != nil{
		log.Panic(err)
	}
}

// Get all the posts
func PostIndex(rw http.ResponseWriter,r *http.Request,params httprouter.Params){

}

//Get Particular Post based on the parameter passed
func PostShow(rw http.ResponseWriter,r *http.Request,params httprouter.Params){

	//params.ByName() func returns the string...so need to convert it in to
	// integer
	id,err := strconv.Atoi(params.ByName("id"))
	fmt.Println("The id received as a parameter is ",id)
	handlePanic(err)
	post := findPost(id)
	ret,_ := json.Marshal(post)
	fmt.Fprintf(rw,string(ret))
}

//Welcome on Home page
func Index(rw http.ResponseWriter,r *http.Request,params httprouter.Params){
	fmt.Fprintf(rw,"Welcome to Blog Post.Please write Your Blogs here.")

}


//Add in to existing posts
func PostCreate(rw http.ResponseWriter,r *http.Request,params httprouter.Params){
	id,err := strconv.Atoi(params.ByName("id"))
	handlePanic(err)
	fmt.Println("The id received as a parameter is ",id)

	//First get the data in to byte stream..so that we could create structure out of
	//it.
	data,err := ioutil.ReadAll(r.Body)
        handlePanic(err)

	//Always remember to close the reader stream...so closing here
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	//Now time to get What we got as a Body from the POST request

	var newPost Post
	err = json.Unmarshal(data,&newPost)
	handlePanic(err)
	updateData(newPost)
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprintf(rw,strconv.Itoa(id))
}

/*
               ALL REDIS RELATED FUNCTIONS

*/


func findALL() []Post{
	conn,err := redis.Dial("tcp",":6379")
	handlePanic(err)
	defer conn.Close()


}

func updateData(post Post){
	conn,err := redis.Dial("tcp",":6379")
	handlePanic(err)
	defer conn.Close()
	conn.Do("SET",post.Id,post)

	conn.Do("SET","KEYS",)
}


func findPost(key int) Post{
	conn, error := redis.Dial("tcp",":6379")
	handlePanic(error)
	defer conn.Close()
	reply,err := conn.Do("GET",key)
	handlePanic(err)
	var post Post
	err = json.Unmarshal(reply.([]byte),&post)
	handlePanic(err)
	return post
}

func main(){

}
