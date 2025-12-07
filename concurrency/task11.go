package concurrency

import (
	"fmt"
	"sync/atomic"
)

//Concurrent counter на атомиках**
//
//Сделать структуру:
//
//```go
//type Counter struct {
//    v int64
//}
//```
//
//Методы:
//
//- `Inc()`
//
//- `Load()`
//
//
//Сравнить с `Mutex` по поведению.

type Counter struct {
	v uint64
}

func (c *Counter) Inc() {
	atomic.AddUint64(&c.v, 1)
}

func (c *Counter) Load() uint64 {
	return atomic.LoadUint64(&c.v)
}

func Task11() {
	c := new(Counter)
	for i := 1; i < 100000; i++ {
		go func() {
			c.Inc()
		}()
		go func() {
			fmt.Println(c.Load())
		}()
	}
}
