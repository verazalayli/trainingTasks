package concurrency

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
