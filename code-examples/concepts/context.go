package concepts

import (
	"context"
	"fmt"
	"log"
	"time"
)

// https://medium.com/@matryer/context-keys-in-go-5312346a868d
func thirPartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 80)

	return "some Response", nil
}

func fetchUserById(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)

	defer cancel()

	val := ctx.Value("username")
	fmt.Println("username:", val)

	type result struct {
		userId string
		err    error
	}

	resultChan := make(chan result, 1)

	go func() {
		user, err := thirPartyHTTPCall()
		resultChan <- result{userId: user, err: err}
	}()

	select {
	// done()
	// 1. the context timeout is exceended
	// 2. the context has been manually cancelled
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultChan:
		return res.userId, res.err
	}
}

func PackageContext() {
	ctx := context.WithValue(context.Background(), "username", "jbarron")
	start := time.Now()
	result, err := fetchUserById(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the response tokk %v:%v\n", time.Since(start), result)
}
