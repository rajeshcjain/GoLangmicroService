package main

import "time"

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


