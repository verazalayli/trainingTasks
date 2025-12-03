package concurrency

//Реализовать семафор на каналах**
//
//**Условие:**
//Сделай структуру:
//
//```go
//type Semaphore struct {
//    ch chan struct{}
//}
//```
//
//Методы:
//
//- `Acquire()`
//
//- `Release()`
//
//
//Без `sync.Mutex`, только через каналы.

type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(limit int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, limit),
	}
}

func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ch
}
