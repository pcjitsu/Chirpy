package main

import (
	"net/http"
)

func main() {
	//Creating a new request router that matches paths to handlers
	mux := http.NewServeMux()

	//configuring the server struct
	server := &http.Server{
		Addr:    ":8080", //port
		Handler: mux,     //send all incoming requests to our router
	}

	//for any path strating iwth "/" serve files from the directory
	//request for /index.html returns ./index.htnml
	mux.Handle("/", http.FileServer(http.Dir(".")))
	//Starts the server
	server.ListenAndServe()
}

//go build -o out && ./out
