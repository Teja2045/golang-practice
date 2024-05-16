package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://api.sampleapis.com/beers/ale",
		"https://api.sampleapis.com/coffee/hot",
		//"https://api.example.com/orders",
	}

	results := make(chan string)

	for _, url := range urls {
		go fetchAPI(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}
}

func fetchAPI(ctx context.Context, url string, results chan<- string) {

	// stupid select to test
	// select {
	// case <-ctx.Done():
	// 	fmt.Println("Context canceled or deadline exceeded:", ctx.Err())
	// }
	now := time.Now()
	defer func() {
		end := time.Now()
		fmt.Println("time took to complete the request (", url, "): ", end.Sub(now).Milliseconds(), "ms")
	}()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	//body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		results <- "something bad..."
		return
	}
	if err != nil {
		results <- fmt.Sprintf("Error making request to %s: %s", url, err.Error())
		return
	}
	defer resp.Body.Close()

	results <- fmt.Sprintf("Response from %s: %d", url, resp.StatusCode)
}
