package concurrency

import (
	"fmt"
	"sync"
	"time"
)

//Мини-кэш с TTL**
//
//Сделать структуру Cache:
//
//- map + Mutex
//
//- каждый ключ имеет TTL
//
//- каждые 500ms должна очищаться просроченная запись (через `Ticker`).
//
//
//Задача: не создать утечку горутин.

type myCache struct {
	mu   sync.Mutex
	data map[string]time.Time
	ttl  time.Duration
	t    *time.Ticker
	stop chan struct{}
}

func newCache(ttl time.Duration) *myCache {
	c := &myCache{
		data: make(map[string]time.Time),
		ttl:  ttl,
		stop: make(chan struct{}),
	}

	c.startGC()

	return c
}

func (c *myCache) addToCache(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = time.Now().Add(c.ttl)
}

func (c *myCache) getFromCache(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	exp, ok := c.data[key]
	if !ok {
		return false
	}

	if time.Now().After(exp) {
		delete(c.data, key)
		return false
	}

	return true
}

func (c *myCache) startGC() {
	c.t = time.NewTicker(c.ttl * time.Millisecond)

	go func() {
		for {
			select {
			case <-c.t.C:
				c.cleanupExpired()
			case <-c.stop:
				c.t.Stop()
				return
			}
		}
	}()
}

func (c *myCache) cleanupExpired() {
	now := time.Now()

	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.data {
		if now.After(v) {
			delete(c.data, k)
		}
	}
}

func (c *myCache) Stop() {
	close(c.stop)
}

func Task9() {
	cache := newCache(500 * time.Millisecond)
	defer cache.Stop()
	cache.addToCache("a")

	time.Sleep(3 * time.Nanosecond)
	fmt.Println("alive:", cache.getFromCache("a"))

	time.Sleep(3 * time.Second)
	fmt.Println("alive:", cache.getFromCache("a"))
}
