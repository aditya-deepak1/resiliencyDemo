package hystrix

import (
	"context"
	"fmt"
	"net/http"
	"time"
	
	"github.com/afex/hystrix-go/hystrix"
)

const TestCommand = "test_command"

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()
	
	// make external call.
	resp := makeHTTPRequest(ctx)
	
	_, err := w.Write([]byte(resp))
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func makeHTTPRequest(ctx context.Context) string {
	
	resC := make(chan string, 1)
	fallbackC := make(chan string, 1)
	
	errC := hystrix.GoC(ctx, TestCommand, func(ctx context.Context) error {
		
		// Create http client.
		client := &http.Client{
			Timeout: time.Second * 5,
		}
		
		// Create request with parent context.
		req, err := http.NewRequest("GET", "https://aditya-hystrix-test.free.beeceptor.com/hello", nil)
		if err != nil {
			fmt.Println("Error forming HTTP request")
		}
		req.WithContext(ctx)
		
		start := time.Now()
		// Make http request.
		_, err = client.Do(req)
		
		if err != nil {
			fmt.Println("Error calling aditya-hystrix-test.free.beeceptor.com", err)
		}
		elapsedTime := time.Since(start).Milliseconds()
		
		fmt.Printf("Done calling aditya-hystrix-test.free.beeceptor.com in %d ms\n", elapsedTime)
		
		resC <- "success"
		
		return nil
	}, func(ctx context.Context, err error) error {
		fmt.Println("Running fallback due to : ", err)
		fallbackC <- err.Error()
		
		return nil
	})
	
	select {
	case res := <-resC:
		return res
	case fallbackRes := <-fallbackC:
		return fallbackRes
	case err := <-errC:
		return err.Error()
	}
}
