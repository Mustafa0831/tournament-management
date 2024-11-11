package main

import (
	"log"
	"net/http"
)

func main() {
	// TODO: Initialize router and handlers
	log.Println("Database Service is running on port 8086")
	http.ListenAndServe(":8086", nil)
}
