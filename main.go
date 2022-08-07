package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// mount the handler.
	http.HandleFunc("/add", add)
	
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

func add(w http.ResponseWriter, r *http.Request) {
	
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()
	
	go func() {
		select {
		
		case <-time.After(time.Second * 20):
			fmt.Println("overslept")
		
		case <-ctx.Done():
			fmt.Println("Inside go routine")
			fmt.Println(ctx.Err())
		}
	}()
	
	// make external call.
	makeHTTPRequest(ctx)
	
	_, err := w.Write([]byte("Addition Done"))
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func makeHTTPRequest(ctx context.Context) {
	
	select {
	
	case <-ctx.Done():
		fmt.Println("Inside makeHTTRequest func")
		fmt.Println(ctx.Err())
	
	case <-time.After(time.Second * 5):
		fmt.Println("calling Google.com")
		
		client := &http.Client{
			Timeout: time.Second * 60,
		}
		
		req, err := http.NewRequest("GET", "https://google.com", nil)
		if err != nil {
			fmt.Println("Error forming HTTP request")
		}
		
		req.WithContext(ctx)
		start := time.Now()
		_, err = client.Do(req)
		
		if err != nil {
			fmt.Println("Error calling Google.com", err)
		}
		elapsedTime := time.Since(start).Milliseconds()
		
		fmt.Printf("Done calling Google.com in %d ms\n", elapsedTime)
	}
}
