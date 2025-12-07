package concurrency

import "sync"

//Объединить несколько каналов ошибок**
//
//Дано:
//
//```go
//errs := []<-chan error{ch1, ch2, ch3}
//```
//
//Нужно сделать функцию:
//
//```go
//func MergeErrors(chs ...<-chan error) <-chan error
//```
//
//Чтобы она вернула один канал, который получает все ошибки и закрывается после завершения всех каналов.

func mergeCh(chs ...<-chan error) <-chan error {
	fnlCh := make(chan error, len(chs))
	wg := sync.WaitGroup{}

	for _, ch := range chs {
		wg.Add(1)
		go func(<-chan error) {
			defer wg.Done()
			for err := range ch {
				fnlCh <- err
			}
		}(ch)
	}

	wg.Wait()
	close(fnlCh)
	return fnlCh
}

func Task15() {
	ch1 := make(chan error)
	ch2 := make(chan error)
	ch3 := make(chan error)

	ch1 <- nil
	ch2 <- nil
	ch3 <- nil
	errs := []<-chan error{ch1, ch2, ch3}

	err := mergeCh(errs...)
	if err != nil {
		panic(err)
	}
}
