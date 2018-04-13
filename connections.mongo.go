package main

import (
	"gopkg.in/mgo.v2"
)

// getTodoSession: Return mongo cnx
func getTodoSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

// getTodoDBName: Return mongo database name
func getTodoDBName() string {
	return "todo_db"
}
