package concurrency

//Реализовать контекст с таймаутом самостоятельно**
//
//Упрощённо реализовать:
//
//```go
//type MyContext struct {
//    Done <-chan struct{}
//}
//```
//
//С функцией:
//
//```go
//func WithTimeout(parent MyContext, d time.Duration) (MyContext, cancel)
//```
//
//Использовать только каналы.
