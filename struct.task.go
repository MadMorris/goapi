package main

//Task data about task
type Task struct {
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Completed bool   `json:"completed"`
}

//Tasks is a collection of tasks
type Tasks []Task
