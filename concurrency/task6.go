package concurrency

import (
	"context"
	"fmt"
)

//Pipeline из 3 стадий**
//
//Написать 3 стадии:
//
//1. чтение чисел из входного канала
//
//2. умножение на 2
//
//3. фильтрация только чётных
//
//
//Каждая стадия — своя горутина.
//Контролировать завершение через `context`.

func task6() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	s1 := numReader(ctx, ch)
	s2 := numMultiplaying(ctx, s1)
	s3 := filtr(ctx, s2)

	for v := range s3 {
		fmt.Println(v)
	}
}

func numReader(ctx context.Context, in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()
	return out
}

func numMultiplaying(ctx context.Context, in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				i := v * 2
				out <- i
			}
		}
	}()
	return out
}

func filtr(ctx context.Context, in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				if (v % 2) == 0 {
					out <- v
				}
			}
		}
	}()
	return out
}
