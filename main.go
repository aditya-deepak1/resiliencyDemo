package main

import (
	"net/http"
	
	"resiliencyDemo/context"
	"resiliencyDemo/hystrix"
)

func main() {
	// mount the handler.
	http.HandleFunc("/context", context.Handler)
	http.HandleFunc("/hystrix", hystrix.Handler)
	
	// Configure hystrix commands.
	hystrix.ConfigureHystrix()
	
	// ask http server to start at 8080 port.
	server := http.Server{
		Addr: ":8080",
	}
	
	// starting server.
	err := server.ListenAndServe()
	if err != nil {
		panic("server crashed")
	}
}
