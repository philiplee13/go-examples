package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	result, err := fetchDataFromThirdParty(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result:", result)
	fmt.Println("Total time:", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchDataFromThirdParty(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	resp := make(chan Response)
	go func() {
		data, err := getData()
		resp <- Response{value: data, err: err}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("Timeout happened when getting data")
		case resp := <-resp:
			return resp.value, resp.err
		}
	}
}

func getData() (int, error) {
	time.Sleep(5 * time.Second)
	return 10, nil
}
