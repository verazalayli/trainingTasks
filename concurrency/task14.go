package concurrency

import (
	"errors"
	"time"
)

//Таймаут для любой функции**
//
//Написать:
//
//```go
//func WithTimeout(fn func() error, d time.Duration) error
//```
//
//Функция должна запускать `fn()` в горутине, а снаружи ожидать либо результат, либо `time.After(d)`.

func WithTimeout(fn func() error, d time.Duration) error {
	ch := make(chan error, 1)

	go func() {
		ch <- fn()
	}()

	select {
	case err := <-ch:
		return err
	case <-time.After(d):
		return errors.New("timeout")
	}
}
