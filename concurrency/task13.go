package concurrency

import (
	"context"
	"fmt"
	"time"
)

//Воркер, который слушает 2 канала (select-case)**
//
//Условие:
//
//Есть 2 канала:
//
//- `jobs <-chan int`
//
//- `control <-chan string`
//
//
//Нужно написать воркера:
//
//- если пришёл `job` — обработать;
//
//- если `control` == `"stop"` — завершить;
//
//- если ничего не приходит 2 секунды — завершить по таймауту.

func Task13() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	jobs := make(chan int)
	control := make(chan string)

	// запускаем воркер в отдельной горутине
	go worker(ctx, jobs, control)

	// отправляем тестовые данные
	time.Sleep(1 * time.Second)
	jobs <- 42

	time.Sleep(1 * time.Second)
	control <- "stop"

	time.Sleep(1 * time.Second)
}

func worker(ctx context.Context, jobs chan int, control chan string) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("jobs closed, exiting")
				return
			}
			fmt.Println("received job:", job)
		case cmd, ok := <-control:
			if !ok || cmd == "stop" {
				fmt.Println("received stop, exiting")
				return
			}
		case <-ctx.Done():
			fmt.Println("timeout:", ctx.Err())
			return
		}
	}
}
