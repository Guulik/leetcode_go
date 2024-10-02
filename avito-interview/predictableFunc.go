package avito_interview

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.
func predictableFunc(ctx context.Context) (int64, error) {
	now := time.Now()

	defer func() {
		fmt.Println("Elapsed time:", time.Since(now))
	}()

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(1*time.Second))
	defer cancel()

	resultCh := make(chan int64, 1)
	go func() {
		res := unpredictableFunc()
		resultCh <- res
		close(resultCh)
	}()

	select {
	case res := <-resultCh:
		return res, nil
	case <-ctx.Done():
		return -1, ctx.Err()
	}

}

func RunPredicable() {
	fmt.Println("started")

	ctx := context.Background()
	result, err := predictableFunc(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
