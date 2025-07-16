package task7_requestCaching

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Cache[T any] struct {
	items T
}

func Task7() {
	cache := Cache[int]{items: slowSquare(rand.IntN(100))}
	fmt.Println(callForInt(cache))
}

func slowSquare(n int) int {
	time.Sleep(100 * time.Millisecond)
	defer fmt.Println("slowSquare: ", n)
	return n * n
}

func callForInt(cache Cache[int]) any {
	return cache.items
}
