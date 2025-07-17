package task7_requestCaching

import (
	"fmt"
	"time"
)

type Cache struct {
	m map[int]int
}

func NewCache() *Cache {
	return &Cache{m: make(map[int]int)}
}

func (c *Cache) Call(n int) int {
	if val, ok := c.m[n]; ok {
		fmt.Printf("Cache hit for %d\n", n)
		return val
	}
	fmt.Printf("Cache miss for %d\n", n)
	result := slowSquare(n)
	c.m[n] = result
	return result
}

func slowSquare(n int) int {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("slowSquare executed:", n)
	return n * n
}

func Task7() {
	cache := NewCache()

	fmt.Println("First call (slow):", cache.Call(5))
	fmt.Println("Second call (fast):", cache.Call(5))
	fmt.Println("Call with new arg:", cache.Call(8))
	fmt.Println("Second call (fast):", cache.Call(8))
}
