package main

import (
	"fmt"
	"goapi/logger"
	"net/http"
	"time"
)

// Ver : route to Ver
func Ver(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("log-2006-01-02.15:04:05.test")
	go logger.Log("internal", "INFO", "Test log").ToFile()
	fmt.Fprintf(w, now)
}
