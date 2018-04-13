package main

import (
	"log"
	"net/http"
)

func main() {

	router := Router()
	log.Fatal(http.ListenAndServe(":8888", router))
}
