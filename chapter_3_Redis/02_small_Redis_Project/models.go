package main

import (
	"time"
)

type User struct{
	Id int           `json:"id"`
	UserName string   `json:"username"`
	Email string      `json:"email"`
}

type Comment struct{
	Id int            `json:"id"`
	User User         `json:"user"`
	Text string        `json:"text"`
	Timestamp time.Time  `json:"timestamp"`

}

type Post struct{
	Id int     `json:"id"`
	User User   `json:"user"`
	Topic string  `json:"topic"`
	Text string   `json:"comment"`
	Timestamp time.Time  `json:"timestamp"`
}
/*
{
	"id" : "1",
	"user" : {
		"id" : 1,
		"username" : "rajesh jain",
		"email" : "rajeshcjain81@gmail.com"
	},
	"topic" : "blog1",
	"comment" : "comment",
	"timestamp" : "2017-05-06 23:00:00 +0000 UTC"
}
*/
