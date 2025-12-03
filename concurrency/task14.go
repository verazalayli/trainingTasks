package concurrency

//Таймаут для любой функции**
//
//Написать:
//
//```go
//func WithTimeout(fn func() error, d time.Duration) error
//```
//
//Функция должна запускать `fn()` в горутине, а снаружи ожидать либо результат, либо `time.After(d)`.
