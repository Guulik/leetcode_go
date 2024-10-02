package avito_interview

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const numRequests = 10000

var count int64

func networkRequest() {
	time.Sleep(time.Millisecond)
	atomic.AddInt64(&count, 1)
}

func RunNetworkRequest() {
	var wg sync.WaitGroup
	for range numRequests {
		wg.Add(1)
		go func() {
			defer wg.Done()
			networkRequest()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
