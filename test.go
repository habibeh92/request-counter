package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func fetch(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error reading response body for %s: %v", url, err)
		return
	}

	ch <- fmt.Sprintf("URL: %s\nResponse: %s", url, body)
}

func main() {
	url := "http://localhost:8080/count"
	numRequests,_ := strconv.Atoi(os.Args[1])
	var wg sync.WaitGroup
	ch := make(chan string, numRequests)

	startTime := time.Now()

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fetch(url, ch)
		}()
	}


	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Elapsed time: %s\n", elapsedTime)
}