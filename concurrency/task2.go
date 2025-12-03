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
