package main

import (
	"nama_npm_pert4/handler" // adjust according to folder name (case sensitive)
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/", handler.API)
	// Replace the last 2 digits of the port with the last 2 digits of your NPM
	log.Println("localhost : 8029")
	http.ListenAndServe(":8029", nil)
}
