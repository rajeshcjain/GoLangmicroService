package main

import (
	"github.com/julienschmidt/httprouter"
)


/*
Registering the routers
Reading it from routers slice and range it and register with the httprouter.
*/
func registerRouters() *httprouter.Router{

	r := httprouter.New()

	for _,val := range routers{
		r.Handle(val.Method,val.Path,val.Handle)
	}
	return r
}