package main

import "github.com/julienschmidt/httprouter"

type Route struct {
	Method string             // HTTP Method
	Path string               // Path of the end point
	Handle httprouter.Handle  //defining the type here.We are telling that
	// this is a function type.
}


var routers = []Route{
	{
		"GET",   //Home Page
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
	{
		"DELETE",  //Deleting the Post
		"/posts/:id",
		PostDel,
	},

}

