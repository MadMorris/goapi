package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

var session = getTodoSession()
var collection = session.DB(getTodoDBName()).C("tasks")

// TaskList : route to task list
func TaskList(w http.ResponseWriter, r *http.Request) {
	var res Tasks
	err := collection.Find(nil).Sort("-_id").All(&res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// TaskAdd : route to task list
func TaskAdd(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var taskData Task
	err := decoder.Decode(&taskData)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(taskData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taskData)
}

// TaskByID : route to Task by ID
func TaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID := params["id"]

	if !bson.IsObjectIdHex(taskID) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(taskID)

	taskData := Task{}
	err := collection.FindId(oid).One(&taskData)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taskData)
}
