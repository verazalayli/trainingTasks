package concurrency

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
