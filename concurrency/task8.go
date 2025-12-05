package concurrency

import (
	"sync"
	"sync/atomic"
)

//Реализовать `sync.Once` вручную**
//
//Сделать структуру:
//
//```go
//type MyOnce struct {
//    done uint32
//    mu   sync.Mutex
//}
//```
//
//Метод `Do(func())`.
//Использовать `atomic` для чтения состояния.

type customOnce struct {
	done uint32
	mu   sync.Mutex
}

func (o *customOnce) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
