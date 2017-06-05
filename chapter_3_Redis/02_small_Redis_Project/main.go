package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
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


/*
We can not have a single post,User and comment on the blog..So defining the
Data structure for it...i.e multiple user,posts and comments.
*/

//var users = []User{}
//var posts = []Post{}
//var comments = []Comment{}

//var users = make(map[int]User)
//var posts = make(map[int]Post)
//var comments = make(map[int]Comment)


/*For reading the configuration file from property file*/
var server string
var port string




func handlePanic(err error){
	if err != nil{
		log.Panic(err)
	}
}

//Welcome on Home page
func Index(rw http.ResponseWriter,r *http.Request,params httprouter.Params){
	content,err := ioutil.ReadFile("welcomepage.html")
	handlePanic(err)
	fmt.Fprintf(rw,string(content))
}

//Deleting the Post
func PostDel(rw http.ResponseWriter,w *http.Request, params httprouter.Params){
	postId := params.ByName("id")
	id,err := strconv.Atoi(postId)
	handlePanic(err)
	ret := deletePost(id)

	if ret == 1{
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type","application/json")
		fmt.Fprintf(rw,"Post number " + postId)
	}else{
		rw.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(rw,"Post Does not exists")
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
	if id == 0 {
		rw.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(rw,"No Content found")
	}
	post := findPost(id)

	/*
	              IMPORTANT
	This is interesting here,We are not allowed to compare the Post struct with
	 nil as they are of different types."nil" can be used for comparision on with
	 functions,channels,interface,slices,maps and pointers.So we got the post from
	 the findPost and then we are taking the pointer of that so that we could compare
	  it
	*/
	if &post == nil {
		rw.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(rw,"No Content found")
	}

	ret,_ := json.Marshal(post)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type","application/json")
	fmt.Fprintf(rw,string(ret))
}


//Add in to existing posts
func PostCreate(rw http.ResponseWriter,r *http.Request,params httprouter.Params){
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
	//handlePanic(err)
	postId := createPost(newPost)
	fmt.Println("one post submitted with id ",postId)

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type","application/json")
	fmt.Fprintf(rw,"blog id " + strconv.Itoa(postId))
}


func main(){
	router := registerRouters()
	/*So finally we register here*/
	fmt.Println("Listening 8080")
	http.ListenAndServe(":8080",router)

}
