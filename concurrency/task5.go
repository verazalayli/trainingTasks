package concurrency

import (
	"fmt"
	"sync"
	"time"
)

//Fan-out + Fan-in**
//
//Есть канал строк.
//Нужно:
//
//1. Фан-аут — запустить N горутин, каждая обрабатывает строку.
//
//2. Фан-ин — собирать результаты в один канал.
//
//3. Закрывать каналы корректно.
//
//
//Проверка: не потерять данные, не сделать дедлок.

func Task5() {
	urls := []string{"http://a.com/a", "http://b.com/b", "http://c.com/c"}

	ch1 := make(chan string, 3)
	ch2 := make(chan string, 3)
	ch3 := make(chan string, 3)

	ch1 <- urls[0]
	ch1 <- urls[0]
	ch1 <- urls[0]
	ch2 <- urls[1]
	ch2 <- urls[1]
	ch2 <- urls[1]
	ch3 <- urls[2]
	ch3 <- urls[2]
	ch3 <- urls[2]
	close(ch1)
	close(ch2)
	close(ch3)

	n := len(urls)
	out := fanIn(ch1, ch2, ch3)
	fanOut(out, n)
}

func fanOut(ch <-chan string, n int) {
	wg := sync.WaitGroup{}

	for range n {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for c := range ch {
				time.Sleep(200 * time.Millisecond)
				fmt.Println(c)
			}
		}()
	}
	wg.Wait()

}

func fanIn(chs ...<-chan string) <-chan string {
	output := make(chan string)
	wg := sync.WaitGroup{}

	for _, ch := range chs {
		wg.Add(1)

		go func(c <-chan string) {
			defer wg.Done()
			for v := range c {
				output <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
