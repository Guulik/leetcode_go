package avito_interview

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func GetPages() {
	var wg sync.WaitGroup
	addreses := []string{"https://avito.ru", "https://google.ru"}
	wg.Add(len(addreses))
	statusCh := make(chan int, len(addreses))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for _, addres := range addreses {
		go func(addres string) {
			defer wg.Done()
			resp, err := http.Get(addres)
			if err != nil {
				fmt.Println(err)
				return
			}
			statusCh <- resp.StatusCode
		}(addres)
	}

	go func() {
		wg.Wait()
		close(statusCh)
	}()
	for {
		select {
		case res, ok := <-statusCh:
			if !ok {
				fmt.Println("всё вывел")
				return
			}
			fmt.Println(res)

		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		}
	}
}
