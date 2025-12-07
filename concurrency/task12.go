package concurrency

import "sync"

//WaitGroup своими руками**
//
//Реализовать:
//
//```go
//type MyWaitGroup struct {
//    ch chan struct{}
//}
//```
//
//Требования:
//
//- метод Add(n)
//
//- метод Done()
//
//- метод Wait() блокирует до нуля
//
//
//Может быть реализовано на каналах или атомиках.

type MyWaitGroup struct {
	mu      sync.Mutex
	ch      chan struct{}
	counter int
}

func NewMyWaitGroup() *MyWaitGroup {
	return &MyWaitGroup{
		ch: make(chan struct{}),
	}
}

func (mwg *MyWaitGroup) add(count int) {
	mwg.mu.Lock()
	defer mwg.mu.Unlock()

	if count < 0 {
		panic("negative Add")
	}

	if mwg.counter == 0 && count > 0 {
		mwg.ch = make(chan struct{})
	}

	mwg.counter += count
}

func (mwg *MyWaitGroup) done() {
	mwg.mu.Lock()
	defer mwg.mu.Unlock()

	if (mwg.counter) <= 0 {
		panic("Panic done when counter is 0 or less")
	} else {
		mwg.counter--
	}

	if mwg.counter == 0 {
		close(mwg.ch)
	}
}

func (mwg *MyWaitGroup) wait() {
	<-mwg.ch
}
