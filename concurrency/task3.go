package concurrency

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать rate limiter (N запросов в секунду)**
//
//Использовать `time.Ticker`.
//Функция принимает список задач и запускает их с ограничением 10 событий/сек.

func rateLimiter(n int, tasks []func()) {
	interval := time.Second / time.Duration(n)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, task := range tasks {
		<-ticker.C
		go func(t func()) {
			defer wg.Done()
			task()
		}(task)
	}

	wg.Wait()
}

func Task3() {
	tasks := []func(){
		func() { fmt.Println("task 1", time.Now()) },
		func() { fmt.Println("task 2", time.Now()) },
		func() { fmt.Println("task 3", time.Now()) },
		func() { fmt.Println("task 4", time.Now()) },
		func() { fmt.Println("task 5", time.Now()) },
		func() { fmt.Println("task 6", time.Now()) },
	}

	rateLimiter(2, tasks) // 2 задачи в секунду
}
