package concurrency

import (
	"fmt"
	"sync"
	"time"
)

//Канал отмены (or-channel)**
//
//Дано несколько `<-chan struct{}`.
//Написать функцию:
//
//```go
//func Or(chs ...<-chan struct{}) <-chan struct{}
//```
//
//Которая закрывается, когда **любой** из входных каналов закрылся.

func Task4() {
	stop1 := make(chan struct{})
	stop2 := make(chan struct{})
	stop3 := make(chan struct{})

	or := or(stop1, stop2, stop3)

	go func() {
		time.Sleep(time.Second)
		close(stop2) // закрыли второй → or тоже закроется
	}()

	<-or
	fmt.Println("or closed!")
}

func or(chs ...<-chan struct{}) <-chan struct{} {
	out := make(chan struct{})

	var once sync.Once

	for _, ch := range chs {
		c := ch
		go func() {
			<-c
			once.Do(func() {
				close(out)
			})
		}()
	}
	return out
}
